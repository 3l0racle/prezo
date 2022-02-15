package components

import (
  "fmt"
  "errors"
  /*"database/sql"
  _ "github.com/go-sql-driver/mysql"*/
  "github.com/3l0racle/prezo/helpers"
)

type Governor struct{
  NickName string
  FisrtName string
  SecondName string
  LastName string
  Phone string
  Email string
  PartyName string
  GovernorId string
  VoteCount int
  RunningMateId string
  CreatedAt string
  UpdatedAt string
}

//transaction creating governor his/her running mate and their login creds
func CreateGovernor(g Governor,r RunningMate)error{
  return nil
}


func ListGovernors()([]Governor,error){
  stmt := "SELECT * FROM `prezo`.`governors` ORDER BY `votecount` DESC;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELP: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all governors")
  }
  defer rows.Close()
  var gvns Governor
  for rows.Next(){
    var g Governor
    err = rows.Scan(&g.NickName,&g.FirstName,&g.SecondName,&g.LastName,&g.Phone,&g.Email,&g.PartyName,&g.GovernorId,&g.VoteCount,&d.RunningMateId,&g.CreatedAt,&g.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESGR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all governors.")
    }
    gvns = append(gvns,g)
  }
  return presidents,nil
}

func ViewGovernor(gid string)(*Governor,error){
  var g Governor
  row := db.QueryRow("SELECT * FROM `prezo`.`governors` WHERE (`governorid` = ?);",gid)
  err := row.Scan(&g.NickName,&g.FirstName,&g.SecondName,&g.LastName,&g.Phone,&g.Email,&g.PartyName,&g.GovernorId,&g.VoteCount,&d.RunningMateId,&g.CreatedAt,&g.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESG id: %s, %s",gid,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing governor with id of %s",gid))
  }
  return &g,nil
}
