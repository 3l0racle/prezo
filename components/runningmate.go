package components

import (
  "fmt"
  "errors"
  "github.com/3l0racle/prezo/helpers"
)

//A running mate can only be created by the president and the governor and thats it
//created like a transaction
type RunningMate struct{
  NickName string
  FisrtName string
  SecondName string
  LastName string
  Phone string
  Email string
  PartyName string
  SeniorId string//ID of the senior be it president or governor
  RunningMateId string
  CreatedAt string
  UpdatedAt string
}


func ListRunningMates()([]RunningMate,error){
  stmt := "SELECT * FROM `prezo`.`runningmates`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELP: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all presidential candidates.")
  }
  defer rows.Close()
  var rmates []RunningMate
  for rows.Next(){
    var r RunningMate
    err = rows.Scan(&r.NickName,&r.FirstName,&r.SecondName,&r.LastName,&r.Phone,&r.Email,&r.PartyName,&r.SeniorId,&r.RunningMateId,&r.CreatedAt,&r.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESPR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all presidents.")
    }
    mates = append(mates,r)
  }
  return mates,nil
}

func GetSeniorsRunningMate(snmId string)(*RunningMate,error){
  var r RunningMate
  row := db.QueryRow("SELECT * FROM `prezo`.`runningmates` WHERE seniorid = ?;",snmId)
  err := row.Scan(&r.NickName,&r.FirstName,&r.SecondName,&r.LastName,&r.Phone,&r.Email,&r.PartyName,&r.SeniorId,&r.RunningMateId,&r.CreatedAt,&r.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESG id: %s, %s",snmId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing running mate with id of %s",snmId))
  }
  return &g,nil
}

func GetRunningMate(rnmId string)(*RunningMate,error){
  var r RunningMate
  row := db.QueryRow("SELECT * FROM `prezo`.`runningmates` WHERE runmateid = ?;",rnmId)
  err := row.Scan(&r.NickName,&r.FirstName,&r.SecondName,&r.LastName,&r.Phone,&r.Email,&r.PartyName,&r.SeniorId,&r.RunningMateId,&r.CreatedAt,&r.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESG id: %s, %s",rnmId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing running mate with id of %s",rnmId))
  }
  return &g,nil
}
