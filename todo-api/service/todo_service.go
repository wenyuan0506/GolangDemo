package service

import (
	"database/sql"
	"fmt"
	"log"
	"todo-api/config"
	"todo-api/model"
)

var todos = []model.Todo{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: true},
}

func GetAllTodos() []model.Todo {
	return todos
}

func GetTodoByID(id int) (model.Todo, bool) {
	for _, t := range todos {
		if t.ID == id {
			return t, true
		}
	}
	return model.Todo{}, false
}

func ConnString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable",
		config.GetEnv("MSSQL_USER", "sa"),
		config.GetEnv("MSSQL_PASSWORD", "password"),
		config.GetEnv("MSSQL_SERVER", "localhost"),
		config.GetEnv("MSSQL_DATABASE", "todo"),
	)
}

func MssqlTest() string {
	var DB *sql.DB
	// 取得連線字串
	connString := ConnString()
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
	return message
}

func GetConnString() string {
	connString := ConnString()
	return connString
}
