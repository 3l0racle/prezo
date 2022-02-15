package components

/*
  * defines how votes are reflected into the DB
*/

import (
  "fmt"
  "errors"
  //"database/sql"
//  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)

type Vote struct{
  Number int //this are the total number of votes in that perticular polling station
  PollingStation string
  WardName string
  Constituency string
  County string
  AgentId string
  CreatedAt string
  UpdatedAt string
}

func ShowAllVotes()([]Vote,error){
  stmt := "SELECT * FROM `prezo`.`votes`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAV: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all votes.")
  }
  defer rows.Close()
  var votes []Vote
  for rows.Next(){
    var v Vote
    err = rows.Scan(&v.Number,&v.PollingStation,&v.WardName,&v.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESCAV: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all votes.")
    }
    votes = append(votes,v)
  }
  return votes,nil
}

//shows the votes of a perticular polling station
func ShowVote(agentId string)(*Vote,error){
  var v Vote
  row := db.QueryRow("SELECT * FROM `prezo`.`votes` WHERE agentid	 = ?;",agentId)
  err := row.Scan(&v.Number,&v.PollingStation,&v.WardName,&v.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EVPV of %s",agentId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing votes for agent with id of %s",agentId))
  }
  return &v,nil
}

func ShowCountyVotes(countyName string)([]Vote,error){
  stmt := "SELECT * FROM `prezo`.`votes` WHERE (`county` = ?);"
  rows,err := db.Query(stmt,countyName)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELCNV: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all county votes.")
  }
  defer rows.Close()
  var votes []Vote
  for rows.Next(){
    var v Vote
    err = rows.Scan(&v.Number,&v.PollingStation,&v.WardName,&v.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESCNV: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all county votes.")
    }
    votes = append(votes,v)
  }
  return votes,nil
}

func ShowConstituencyVotes(consName string)([]Vote,error){
  stmt := "SELECT * FROM `prezo`.`votes` WHERE (`constituency` = ?);"
  rows,err := db.Query(stmt,consName)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELCV: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all constituency votes.")
  }
  defer rows.Close()
  var votes []Vote
  for rows.Next(){
    var v Vote
    err = rows.Scan(&v.Number,&v.PollingStation,&v.WardName,&v.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESCV: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all constituency votes.")
    }
    votes = append(votes,v)
  }
  return votes,nil
}

func ShowWardVotes(wardName string)([]Vote,error){
  stmt := "SELECT * FROM `prezo`.`votes` WHERE (`wardname` = ?);"
  rows,err := db.Query(stmt,wardName)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EQWV: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all ward votes.")
  }
  defer rows.Close()
  var votes []Vote
  for rows.Next(){
    var v Vote
    err = rows.Scan(&v.Number,&v.PollingStation,&v.WardName,&v.Constituency,&v.County,&v.AgentId,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESWR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all ward votes.")
    }
    votes = append(votes,v)
  }
  return votes,nil
}

func CountVote(votes []Vote)(int,error){
  var count int
  for _,vote := range votes{
    count += vote.Number
  }
  return count,nil
}
