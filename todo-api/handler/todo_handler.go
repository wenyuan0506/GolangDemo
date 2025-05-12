package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo-api/config"
	"todo-api/service"
	"todo-api/util"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := service.GetAllTodos()
	util.JSON(w, http.StatusOK, todos)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, _ := strconv.Atoi(idStr)
	todo, found := service.GetTodoByID(id)
	if !found {
		http.NotFound(w, r)
		return
	}
	util.JSON(w, http.StatusOK, todo)
}

func ConnString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable",
		config.GetEnv("MSSQL_USER", "sa"),
		config.GetEnv("MSSQL_PASSWORD", "password"),
		config.GetEnv("MSSQL_SERVER", "localhost"),
		config.GetEnv("MSSQL_DATABASE", "todo"),
	)
}

func MssqlTest(w http.ResponseWriter, r *http.Request) {
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
	util.JSON(w, http.StatusOK, map[string]string{"message": message})
}

func GetConnString(w http.ResponseWriter, r *http.Request) {
	// 取得 MSSQL 連線字串
	connString := ConnString()
	util.JSON(w, http.StatusOK, map[string]string{"connString": connString})
}
