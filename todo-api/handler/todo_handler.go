package handler

import (
	"net/http"
	"strconv"
	"strings"

	"todo-api/database"
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

func MssqlTest(w http.ResponseWriter, r *http.Request) {
	// 測試 MSSQL 連線
	message := database.InitDB()
	util.JSON(w, http.StatusOK, map[string]string{"message": message})
}
