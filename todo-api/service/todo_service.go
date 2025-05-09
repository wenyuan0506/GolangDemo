package service

import (
	"todo-api/database"
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

func MssqlTest(todo model.Todo) {
	database.InitDB()
}
