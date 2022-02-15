package components
/*
  * Acts as a sanitizer for components functions
*/

import (
  "fmt"
  "log"
  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  "github.com/3l0racle/prezo/helpers"
  _ "github.com/go-sql-driver/mysql"
)

/* Not exactly used just an initalization of what everyones day to day activities mlooks like */
type Users struct{
  Email string
  UserId string
  CreatorId string
  Password string
  Active bool
  Updated bool
}

//this prevents them from loging in only
func MarkUserAsInactive(uid string)error{
  return nil
}


func GetPAsTotalNumberrOfAgents(paid string) (int64,error){
  var count int64
  return count,nil
}

type User struct {
  FirstName string
  SecondName string
  Data string
}

type RecentNews struct{
  Title string
  NewsId string
  CreatedAt string
  UpdatedAt string
}

func GetTopFiveRecentNews()([]News,error){
  var trn []News
  //var rn RecentNews
  news,err := ListAllNews()
  if err != nil{
    log.Println(err)
    return nil,errors.New("Server encountered an error while listing news")
  }
  var count int
  for _,rn := range news{
    if count <= 4{
      trn = append(trn,rn)
      count += 1
    }
  }
  return trn,nil
}

func Authenticate(password,email string)(bool,string){
  var userEmail,hash,userId string
  stmt := "SELECT email,userid,password FROM `prezo`.`users` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&userEmail,&userId,&hash)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    helpers.Logerror(e)
    return false,userId
  }
  err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  if err != nil{
    e := helpers.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    helpers.Logerror(e)
    return false,userId
  }
  return true,userId
}


func IsPresident(email string) (bool,error){
  var usrMail string
  stmt := "SELECT email FROM `prezo`.`presidents` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&usrMail)
  if err != nil{
    if err == sql.ErrNoRows {
      return false,errors.New("Email is not for any president")
    }
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for president verification %s",err))
    helpers.Logerror(e)
    return false,errors.New("Server encountered an error verifying president")
  }
  if usrMail == email{
    return true,nil
  }
  return true,nil
}

func IsPA(email string) (bool,error){
  var usrMail string
  stmt := "SELECT email FROM `prezo`.`pa` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&usrMail)
  if err != nil{
    if err == sql.ErrNoRows {
      return false,errors.New("Email is not for any pa")
    }
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for pa verification %s",err))
    helpers.Logerror(e)
    return false,errors.New("Server encountered an error verifying pa")
  }
  if usrMail == email{
    return true,nil
  }
  return true,nil
}


func IsAgent(email string) (bool,error){
  var usrMail string
  stmt := "SELECT email FROM `prezo`.`agents` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&usrMail)
  if err != nil{
    if err == sql.ErrNoRows {
      return false,errors.New("Email is not for any agent")
    }
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for agent verification %s",err))
    helpers.Logerror(e)
    return false,errors.New("Server encountered an error verifying agent")
  }
  if usrMail == email{
    return true,nil
  }
  return true,nil
}
