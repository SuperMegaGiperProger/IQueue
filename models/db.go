package models

import (
	"database/sql"
	_ "github.com/lib/pq"
  "gopkg.in/yaml.v2"
	"fmt"
  "log"
  "io/ioutil"
)

const CONFIG_FILE_NAME = "models/db.yml"

var DB *sql.DB = nil

type DB_config struct {
  User string
  Password string
  DBName string
}

func DBConfigFileData() []byte {
  data, err := ioutil.ReadFile(CONFIG_FILE_NAME)
  if err != nil {
    log.Fatal(err)
  }
  return data
}

func DBConfig() string {
  config := DB_config{}
  err := yaml.Unmarshal(DBConfigFileData(), &config)
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("user=%s password=%s dbname=%s", config.User, config.Password, config.DBName)
}

func InitDB() (err error) {
	DB, err = sql.Open("postgres", DBConfig())
  return
}

func SetFieldById(tableName, fieldName, value string, id int32) {
	query := fmt.Sprintf("UPDATE %s SET %s=%s WHERE id=%d", tableName, fieldName, value, id)
	DB.Exec(query)
}
