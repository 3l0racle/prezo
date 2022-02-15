package components

/*
  * Defines an agent and functions for creating one

*/

import (
  "fmt"
  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)

type Agent struct{
  FirstName string
  SecondName string
  Phone string
  Email string
  IdNumber string
  AgentId string
  PaId string //create users
  PollingStationName string
  WardName string
  Constituency string
  County string
  CreatedAt string
  UpdatedAt string
}

//a transaction to users to create login credentials and to votes to set initial vote count to zero
func CreateAgent(a Agent) error{
  tx,err := db.Begin()
  if err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Srintf("EPIA: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  var result sql.Result
  //create the agent
  result,err = tx.Exec("INSERT INTO `prezo`.`agents` (`firstname`,`secondname`,`phoneno`,`email`,`idnumber`,`agentid`,`pollingstationname`,`wardname`,`constituency`,`county`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);",a.FirstName,a.SecondName,a.Phone,a.Email,a.IdNumber,a.AgentId,a.PollingStationName,a.WardName,a.Constituency,a.County,a.CreatedAt,a.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("CAEWIP: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agen")
  }//&v.Number,&v.PollingStation,&v.WardName,&V.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt
  //set the total number of votes in thtat polling station to zero
  result,err = tx.Exec("INSERT INTO `prezo`.`votes` (`number`,`pollingstationname`,`wardname`,`constituency`,`county`,`agentid`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?);",0,a.PollingStation,a.WardName,a.Constituency,a.County,a.AgentId,a.CreatedAt,a.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("CAEIIVC: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent's polling station initial vote count")
  }
  /*&c.CandidateId,&c.Votes,&c.PollingStationName,&c.AgentId,&c.CreatedAt,&c.UpdatedAt
  //create the initial candidate vote to zero
  result,err = tx.Exec("INSERT INTO `prezo`.`runningmates` (`candidateid`,`votes`,`pollingstationname`,`agentid`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?,?,?,?);",c.CandidateId,c.Votes,c.PollingStationName,c.AgentId,c.CreatedAt,c.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECPRM: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president's running mate")
  }*/
  var passwordHash []byte
  passwordHash,err = bcrypt.GenerateFromPassword([]byte(a.AgentId),bcrypt.DefaultCost)
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("passgen",fmt.Sprintf("CAEGPH: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  result,err = tx.Exec("INSERT INTO `prezo`.`users` (`email`,`userid`,`creatorid`,`password`,`active`,`updated`) VALUES(?,?,?,?,?,?);",a.Email,a.AgentId,a.PaId,passwordHash,true,false)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("CAEIU: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  if err = tx.Commit(); err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("CACE: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  return nil
}

func ListAgents()([]Agent,error){
  stmt := "SELECT * FROM `prezo`.`agents`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql","ELA: ",err)
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all agents.")
  }
  defer rows.Close()
  var agents []Agent
  for rows.Next(){
    var a Agent
    err = rows.Scan(&a.FirstName,&a.SecondName,&a.Phone,&a.Email,&a.IdNumber,&a.AgentId,&a.PollingStationName,&a.WardName,&a.Constituency,&a.County,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all agents.")
    }
    agents = append(agents,a)
  }
  return agents,nil
}

//a transaction (f*n big one)
func ListAgentsByPaid(paid string)([]Agent,error){
  return nil,nil
}

func ViewAgent(agentIdNumber string)(*Agent,error){
  var a Agent
  row := db.QueryRow("SELECT * FROM `prezo`.`agents` WHERE agentid	 = ?;",agentIdNumber)
  err := row.Scan(&a.FirstName,&a.SecondName,&a.Phone,&a.Email,&a.IdNumber,&a.AgentId,&a.PollingStationName,&a.WardName,&a.Constituency,&a.County,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESA id: %s, %s",agentIdNumber,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing agent with id of %s",agentIdNumber))
  }
  return &a,nil
}
