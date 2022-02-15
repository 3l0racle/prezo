package components

/*
  * Contains struct defining news types and alert level
*/

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/3l0racle/prezo/helpers"
)

type News struct{
  Title string
  Description string
  NewsId string
  Handled bool
  Level string
  CreatorId string
  ForEveryone bool
  CreatedAt string
  UpdatedAt string
}

func CreateNews(n News)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `prezo`.`news` (title,description,newsid,handled,level,creatorid,foreveryone,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EPCN: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while creating news, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(n.Title,n.Description,n.NewsId,n.Handled,n.Level,n.CreatorId,n.ForEveryone,n.CreatedAt,n.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("EZRAWCN: %s",err))
    logError(e)
    return errors.New("Server encountered an error while creating deal.")
  }
  return nil
}

func DeleteNews(newsId string)error{
  del,err := db.Prepare("DELETE FROM `prezo`.`news` WHERE (`newsid` = ?);")
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EDN with id: %s : %s",newsId,err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while deleting news.")
  }
  defer del.Close()
  var res sql.Result
  res,err = del.Exec(newsId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EDNZRA with id: %s : %s",newsId,err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while deleting news.")
  }
  return nil
}

func MarkNewAsHandled(newsId string)error{
  upStmt := "UPDATE `prezo`.`news` SET `handled` = ? WHERE (`newsid` = ?);";
  stmt,er := db.Prepare(upStmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EMNH: %s",err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while marking news as handled.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(true,newsId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EMNH: news id: %s  %s",newsId,err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while marking news as handled.")
  }
  return nil
}

func ListAllUnhandledNews()([]News,error){
  stmt := "SELECT * FROM `prezo`.`news` WHERE (`handled` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,false)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELUHN: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all unhandled news.")
  }
  defer rows.Close()
  var news []News
  for rows.Next(){
    var n News
    err = rows.Scan(&n.Title,&n.Description,&n.NewsId,&n.Handled,&n.Level,&n.CreatorId,&n.ForEveryone,&n.CreatedAt,&n.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAUHN: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all unhandled news.")
    }
    news = append(news,n)
  }
  return news,nil
}

func ListAllHandledNews()([]News,error){
  stmt := "SELECT * FROM `prezo`.`news` WHERE (`handled` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,true)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELHN: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all handled news.")
  }
  defer rows.Close()
  var news []News
  for rows.Next(){
    var n News
    err = rows.Scan(&n.Title,&n.Description,&n.NewsId,&n.Handled,&n.Level,&n.CreatorId,&n.ForEveryone,&n.CreatedAt,&n.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAHN: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all handled news.")
    }
    news = append(news,n)
  }
  return news,nil
}

//lists all news from the latest
func ListAllNews()([]News,error){
  stmt := "SELECT * FROM `prezo`.`news` ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("ELN: %s",err))
    helpers.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all news.")
  }
  defer rows.Close()
  var news []News
  for rows.Next(){
    var n News
    err = rows.Scan(&n.Title,&n.Description,&n.NewsId,&n.Handled,&n.Level,&n.CreatorId,&n.ForEveryone,&n.CreatedAt,&n.UpdatedAt)
    if err != nil{
      e := helpers.LogErrorToFile("sql",fmt.Sprintf("ESAN: %s",err))
      helpers.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all news.")
    }
    news = append(news,n)
  }
  return news,nil
}
