package main

import (
	"fmt"
	"log"
	"net/http"

	"todo-api/config"
	"todo-api/router"
)

func main() {
	config.LoadEnv()
	r := router.SetupRouter()
	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
