package main

import (
	"net/http"
)

import (
	"./models"
	"./controllers/queue"
	"flag"
	"fmt"
	"log"
)

var routes = map[string]func(http.ResponseWriter, *http.Request) {
	"/": queue.List,
	"/create": queue.Create,
	"/push": queue.Push,
	"/show": queue.Show,
	"/save": queue.Save,
	"/remove": queue.Remove,
	"/pop": queue.Pop,
	"/list": queue.List,
}

const DEFAULT_PORT = 3000

var port int

func init_flags() {
	flag.IntVar(&port, "port", DEFAULT_PORT, "server port")
	flag.Parse()
}

func handleStatic() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func initHandlers() {
	for pattern, handler := range routes {
		http.HandleFunc(pattern, handler)
	}
}

func port_to_s() string {
	return fmt.Sprintf(":%d", port)
}

func main() {
	init_flags()
	models.InitDB()
	handleStatic()
	initHandlers()
	err := http.ListenAndServe(port_to_s(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
