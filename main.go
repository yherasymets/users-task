package main

import (
	"log"
	"net/http"

	"github.com/yherasymets/users-task/app"
)

func main() {
	app := app.NewApp()
	handler := app.Router()
	log.Println("server running on :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}
}
