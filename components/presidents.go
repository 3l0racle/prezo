package components


import (
  "fmt"
  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)

type President struct{
  NickName string
  FirstName string
  SecondName string
  LastName string
  Phone string
  Email string
  PartyName string
  PresidentId string
  RunningMateId string
  VoteCount int
  CreatedAt string
  UpdatedAt string
}

//a transaction to users to create login creds for him n his/her running mate and create running mate also
//vote count is initially set to zero
func CreatePresident(p President,r RunningMate)error{
  tx,err := db.Begin()
  if err != nil{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EPIP: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president")
  }
  var result sql.Result

  //create the president
  result,err = tx.Exec("INSERT INTO `prezo`.`presidents` (`nickname`,`firstname`,`secondname`,lastname,`phoneno`,`email`,`partyname`,`seniorid`,`runmateid`,`votecount`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);",p.NickName,p.FirstName,p.SecondName,p.LastName,p.Phone,p.Email,p.PartyName,p.PresidentId,p.RunningMateId,p.VoteCount,p.CreatedAt,p.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECP: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president")
  }
  //create the running mate
  result,err = tx.Exec("INSERT INTO `prezo`.`runningmates` (`nickname`,`firstname`,`secondname`,`lastname`,`phoneno`,`email`,`partyname`,`seniorid`,`runmateid`,`created_at`,`updated_at`) VALUES(?,?,?,?,?,?,?,?,?,?,?);",r.NickName,r.FirstName,r.SecondName,r.LastName,r.Phone,r.Email,p.PartyName,p.PresidentId,p.RunningMateId,r.CreatedAt,r.UpdatedAt)
  rowsAffec,_ = result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECPRM: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president's running mate")
  }
  //create their login credentials
  var passwordHash []byte
  var rmpassHash []byte
  //create their hashes
  passwordHash,err = bcrypt.GenerateFromPassword([]byte(p.PresidentId),bcrypt.DefaultCost)
  if err != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("passgen",fmt.Sprintf("ECPH: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president")
  }
  rmpassHash,err = bcrypt.GenerateFromPassword([]byte(r.RunningMateId),bcrypt.DefaultCost)
  if err != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("passgen",fmt.Sprintf("ECRH: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president's running mate")
  }
  //create their login creds
  result,err = tx.Exec("INSERT INTO `prezo`.`users` (`email`,`userid`,`creatorid`,`password`,`active`,`updated`) VALUES(?,?,?,?,?,?);",p.Email,p.PresidentId,"SAM DID THIS!",passwordHash,true,false)
  rowsAffec,_ = result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECPU: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president")
  }
  result,err = tx.Exec("INSERT INTO `prezo`.`users` (`email`,`userid`,`creatorid`,`password`,`active`,`updated`) VALUES(?,?,?,?,?,?);",r.Email,r.RunningMateId,"SAM DID THIS!",rmpassHash,true,false)
  rowsAffec,_ = result.RowsAffected()
  if err != nil || rowsAffec != 1{
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECRU: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president's running mate username")
  }
  //commit their login creds
  if err = tx.Commit(); err != nil {
    _ = tx.Rollback()
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ECCP: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating president")
  }
  return nil
}

func GetPresident(presidentId string) (*President,error){
  var p President
  row := db.QueryRow("SELECT * FROM `prezo`.`presidents` WHERE (`presidentid` = ?);",presidentId)
  err := row.Scan(&p.NickName,&p.FirstName,&p.SecondName,&p.LastName,&p.Phone,&p.Email,&p.PartyName,&p.PresidentId,&p.RunningMateId,&p.VoteCount,&p.CreatedAt,&p.UpdatedAt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EVP id: %s  %s",presidentId,err))
    helpers.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing president with the id of %s",presidentId))
  }
  return &p,nil
}

func ShowPresidents()([]President,error){
  stmt := "SELECT * FROM `prezo`.`presidents` ORDER BY `votecount` DESC;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELP: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all presidential candidates.")
  }
  defer rows.Close()
  var presidents []President
  for rows.Next(){
    var p President
    err = rows.Scan(&p.NickName,&p.FirstName,&p.SecondName,&p.LastName,&p.Phone,&p.Email,&p.PartyName,&p.PresidentId,&p.RunningMateId,&p.VoteCount,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESPR: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all presidents.")
    }
    presidents = append(presidents,p)
  }
  return presidents,nil
}
