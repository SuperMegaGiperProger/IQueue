package models

import (
	"database/sql"
	//"fmt"
	_ "github.com/lib/pq"
	//"fmt"
	"fmt"
)

var DB *sql.DB = nil

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=123 dbname=iqueue")
	return err
}

func SetFieldById(tableName, fieldName, value string, id int32) {
	query := fmt.Sprintf("UPDATE %s SET %s=%s WHERE id=%d", tableName, fieldName, value, id)
	DB.Exec(query)
}