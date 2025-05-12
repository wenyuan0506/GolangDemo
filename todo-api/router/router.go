package router

import (
	"net/http"
	"todo-api/handler"
	"todo-api/middleware"
)

func SetupRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/todos", middleware.Logger(http.HandlerFunc(handler.GetTodos)))
	r.Handle("/todos/", middleware.Logger(http.HandlerFunc(handler.GetTodoByID)))
	r.Handle("/mssqlTest/", middleware.Logger(http.HandlerFunc(handler.MssqlTest)))
	r.Handle("/connString/", middleware.Logger(http.HandlerFunc(handler.GetConnString)))
	r.Handle("/allTableNames/", middleware.Logger(http.HandlerFunc(handler.GetAllTableNames)))
	r.Handle("/table", middleware.Logger(http.HandlerFunc(handler.GetTableData)))
	return r
}
