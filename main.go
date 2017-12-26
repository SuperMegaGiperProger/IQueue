package main

import (
	"net/http"
//	"regexp"
)

import (
	"./models"
	"./controllers/queue"
)

func main() {
	models.InitDB()

	http.HandleFunc("/", queue.Create)
	http.HandleFunc("/create", queue.Create)
	http.HandleFunc("/push", queue.Push)
	http.HandleFunc("/show", queue.Show)
	http.HandleFunc("/save", queue.Save)
	http.HandleFunc("/remove", queue.Remove)
	http.HandleFunc("/pop", queue.Pop)
	http.ListenAndServe(":8080", nil)
}
