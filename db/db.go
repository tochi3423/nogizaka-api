package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/kelseyhightower/envconfig"
)

type (
  Conn struct {
    User     string `required:"true"`
    Password string `required:"true"`
    Protocol string `required:"true"`
    Address  string `required:"true"`
    Schema   string `required:"true"`
  }
)

var db *gorm.DB

func New() *gorm.DB {
  var conn Conn

  DBMS := "mysql"
  USER := os.Getenv("DB_USER")
  PASS := os.Getenv("DB_PASSWORD")
  PROTOCOL := os.Getenv("DB_PROTOCOL")
  ADDRESS := os.Getenv("DB_ADDRESS")
  DBNAME := "nogizaka"

  err := envconfig.Process("db", &conn)
  if err == nil {
    USER = conn.User
    PASS = conn.Password
    PROTOCOL = conn.Protocol
    ADDRESS = conn.Address
    DBNAME = conn.Schema
  }

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "(" + ADDRESS + "" + ")" + "/" + DBNAME

  db, err := gorm.Open(DBMS, CONNECT+"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
  if err != nil {
    panic("failed to connect database!!")
  }

  db.LogMode(true)
  return db
}