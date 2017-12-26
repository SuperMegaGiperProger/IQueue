package main

import (
	"net/http"
//	"regexp"
)

import (
	//"fmt"
	"./models"
	//"./models/queue_item"
	//"./models/queue"
	"html/template"
	"./controllers/queue"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/queue/create.html")
	t.Execute(w, nil)
}

func main() {
	models.InitDB()

	http.HandleFunc("/", handler)
	http.HandleFunc("/save", queue.Save)
	http.ListenAndServe(":8080", nil)
}
