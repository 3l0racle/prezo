package components

import (
  "time"
  "database/sql"
)

var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
