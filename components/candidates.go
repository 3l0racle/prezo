package components

/*
  *Defines how the votes are cast
  * Transacts all votes
*/

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)

type CandidatesVotes struct {
  CandidateId string
  Votes int
  PollingStationName string
  AgentId string
  CreatedAt string
  UpdatedAt string
}

//takes in a list of the candidates votes
func UpdateCandidateVotes(cvs CandidatesVotes)error{
  tx,err := db.Begin()
  if err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EPUCV: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  //Geters
  var presidentVoteCount,pollingStationVoteCount,candidateVoteCount int
  row := tx.QueryRow("SELECT votes FROM `prezo`.`candidates` WHERE (`candidateid` = ?);",cvs.CandidateId)
  err = row.Scan(&candidateVoteCount)
  if err != nil {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUCV: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  row = tx.QueryRow("SELECT number FROM `prezo`.`votes` WHERE (`agentid` = ?);",cvs.AgentId)
  err = row.Scan(&pollingStationVoteCount)
  if err != nil {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUCV: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  row = tx.QueryRow("SELECT votecount FROM `prezo`.`presidents` WHERE (`seniorid ` = ?);",cvs.CandidateId)
  err = row.Scan(&presidentVoteCount)
  if err != nil {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUCV: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  presidentVoteCount += cvs.Votes
  pollingStationVoteCount += cvs.Votes
  candidateVoteCount += cvs.Votes
  fmt.Println("[+] Preparing to update president with id %s vote count to %d at polling station %s",cvs.CandidateId,presidentVoteCount,cvs.PollingStationName)
  var result sql.Result
  //update presidents vote count
  result,execErr := tx.Exec(`UPDATE presidents SET votecount = ? WHERE seniorid = ?`,presidentVoteCount,cvs.CandidateId)
  rowsAffec,_ := result.RowsAffected()
  if execErr != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUPVC: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  //update polling station vote count
  result,execErr = tx.Exec(`UPDATE votes SET number = ? WHERE agentid = ?`,pollingStationVoteCount,cvs.CandidateId)
  rowsAffec,_ = result.RowsAffected()
  if execErr != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUPSVC: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  //update the candidates votes
  result,execErr = tx.Exec(`UPDATE candidates SET votes = ? WHERE candidateid = ?`,candidateVoteCount,cvs.CandidateId)
  rowsAffec,_ = result.RowsAffected()
  if execErr != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EUPVC: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  if err = tx.Commit(); err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECPUVC: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while updating candidates vote count")
  }
  return nil
}

func GetAllCandidatesVotes(cid string)([]CandidatesVotes,error){
  stmt := "SELECT * FROM `prezo`.`candidates` WHERE (`candidateid` = ?);"
  rows,err := db.Query(stmt,cid)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EGCV: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all candidates votes.")
  }
  defer rows.Close()
  var cvs []CandidatesVotes
  for rows.Next(){
    var c CandidatesVotes
    err = rows.Scan(&c.CandidateId,&c.Votes,&c.PollingStationName,&c.AgentId,&c.CreatedAt,&c.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESCVR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all candidate votes.")
    }
    cvs = append(cvs,c)
  }
  return cvs,nil
}

func GetAllCandidatesVotesFrmAPerticularPollingStation(agentId,psName string)(*CandidatesVotes,error){
  var c CandidatesVotes
  row := db.QueryRow("SELECT * FROM `prezo`.`candidates` WHERE (`pollingstationame` = ? AND `agentid` = ?);",psName,agentId)
  err := row.Scan(&c.CandidateId,&c.Votes,&c.PollingStationName,&c.AgentId,&c.CreatedAt,&c.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EVPSA name: %s of agent %s  %s",psName,agentId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing votes for polling station: %s with agent id of %s",psName,agentId))
  }
  return &c,nil
}

func CountCandidateVotes(cvs []CandidatesVotes) (int,error){
  var count int
  for _,cv := range CandidatesVotes {
    count += cv.Votes
  }
  return count,nil
}


func AuditVotes()bool{
  return true
}









/*
Something to be tried later on
type Votting interface{
  CreateVote() error
}*/
