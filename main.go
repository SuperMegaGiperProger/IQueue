package main

import (
	"net/http"
//	"regexp"
)

import (
	"fmt"
	"./models"
	//"./models/queue_item"
	//"./models/queue"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	models.InitDB()
	//http.HandleFunc("/list/", makeHandler(viewHandler))
	//http.HandleFunc("/", handler)

	//http.ListenAndServe(":8080", nil)
}
