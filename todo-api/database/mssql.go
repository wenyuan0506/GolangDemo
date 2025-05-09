package database

import (
	"database/sql"
	"fmt"
	"log"
	"todo-api/config"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() string {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable",
		config.GetEnv("MSSQL_USER", "sa"),
		config.GetEnv("MSSQL_PASSWORD", "password"),
		config.GetEnv("MSSQL_SERVER", "localhost"),
		config.GetEnv("MSSQL_DATABASE", "todo"),
	)

	var err error
	var message string
	DB, err = sql.Open("sqlserver", connString)

	if err != nil {
		message = "無法連線到資料庫，請檢查連線字串或資料庫狀態"
		log.Fatal(message, err)
	}

	if err = DB.Ping(); err != nil {
		message = "無法 ping 資料庫，請檢查資料庫狀態"
		log.Fatal(message, err)
	}
	message = "✅ 資料庫連線成功"
	log.Println(message)

	// 回傳資料庫連線狀態
	return message
}
