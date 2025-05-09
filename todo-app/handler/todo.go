// handler/todo.go
package handler

import (
	"fmt"
	"todo-app/model"
)

func PrintSampleTodo() {
	todo := model.Todo{ID: 1, Title: "Learn Go", Done: false}
	fmt.Printf("Todo: %+v\n", todo)
}
