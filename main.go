package main

import (
	"net/http"
	//	"regexp"
)

import (
	"./models"
	"./controllers/queue"
)

var routes = map[string]func(http.ResponseWriter, *http.Request) {
	"/": queue.Create,
	"/create": queue.Create,
	"/push": queue.Push,
	"/show": queue.Show,
	"/save": queue.Save,
	"/remove": queue.Remove,
	"/pop": queue.Pop,
}

func initHandlers() {
	for pattern, handler := range routes {
		http.HandleFunc(pattern, handler)
	}
}

func main() {
	models.InitDB()
	initHandlers()
	http.ListenAndServe(":8080", nil)
}
