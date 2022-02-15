package components

/*
  * defines work and structure for the presidents' PA's
*/

import (
  "fmt"
  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)


type PresPA struct{
  FirstName string
  SecondName string
  Phone string
  Email string
  IdNumber string
  PasId string
  PresId string
  CreatedAt string
  UpdatedAt string
}

//transaction to users
func CreatePA(pa PresPA)error{
  tx,err := db.Begin()
  if err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Srintf("EPCPPA: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  var result sql.Result
  result,err = tx.Exec("INSERT INTO `prezo`.`pa` (`firstname`,`secondname`,`phoneno`,`email`,`paid`,`presid`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?);",pa.FirstName,pa.SecondName,pa.Phone,pa.Email,pa.IdNumber,pa.PasId,pa.PresId,pa.CreatedAt,pa.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECPA: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while inserting into PA")
  }
  var passwordHash []byte
  passwordHash,err = bcrypt.GenerateFromPassword([]byte(pa.AgentId),bcrypt.DefaultCost)
  if err != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("passgen",fmt.Sprintf("CPAEGPH: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating agent")
  }
  result,err = tx.Exec("INSERT INTO `prezo`.`users` (`email`,`userid`,`creatorid`,`password`,`active`,`updated`) VALUES(?,?,?,?,?,?);",pa.Email,pa.PasId,pa.PresId,passwordHash,true,false)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("PAEIU: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating PA")
  }
  if err = tx.Commit(); err != nil {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECCPA: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating PA")
  }
  return nil
}

func ListPAs()([]PresPA,error){
  stmt := "SELECT * FROM `prezo`.`pa`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EPLAPAs: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all PA's.")
  }
  defer rows.Close()
  var pas []PresPA
  for rows.Next(){
    var p PresPA
    err = rows.Scan(&p.FirstName,&p.SecondName,&p.Phone,&p.Email,&p.IdNumber,&p.PasId,&p.PresId,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAPAs: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all PA's.")
    }
    pas = append(pas,p)
  }
  return pas,nil
}

func ViewPA(paid string)(*PresPA,error){
  var pa PresPA
  row := db.QueryRow("SELECT * FROM `prezo`.`pa` WHERE paid	 = ?;",paid)
  err := row.Scan(&pa.FirstName,&pa.SecondName,&pa.Phone,&pa.Email,&pa.IdNumber,&pa.PasId,&pa.PresId,&pa.CreatedAt,&pa.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EVPA with id %s %s",agentId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing pa with id of %s",paid))
  }
  return &v,ni
}

func ListPAByCreator(creatorsId string)([]PresPA,error){
  stmt := "SELECT * FROM `prezo`.`pa` WHERE (`presid` = ?);"
  rows,err := db.Query(stmt,creatorsId)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EPLPABC: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all PA's.")
  }
  defer rows.Close()
  var pas []PresPA
  for rows.Next(){
    var  p PresPA
    err = rows.Scan(&pa.FirstName,&pa.SecondName,&pa.Phone,&pa.Email,&pa.IdNumber,&pa.PasId,&pa.PresId,&pa.CreatedAt,&pa.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESPPA: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all PA's.")
    }
    pas = append(pas,p)
  }
  return votes,nil
}
