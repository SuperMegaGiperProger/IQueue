package queue

import (
	"net/http"
	"net/url"
	queue_model "../../models/queue"
	"../../models/queue_item"
	"strconv"
	"html/template"
	"fmt"
)

func getQ(r *http.Request) queue_model.Queue {
	body, _ := url.ParseQuery(r.URL.RawQuery)
	q_id, _ := strconv.Atoi(body["queue_id"][0])
	return queue_model.Find(int32(q_id))
}

func Save(w http.ResponseWriter, r *http.Request) {
	body, _ := url.ParseQuery(r.URL.RawQuery)
	q, _ := queue_model.New(body["queue_name"][0])
	for _, username := range body["item[]"] {
		q.Push(queue_item.New(username))
	}
	http.Redirect(w, r, "/", http.StatusFound)
}


type Items struct {
	Queue queue_model.Queue
	Items []queue_item.QueueItem
}

func Show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/queue/show.html")
	q := getQ(r)
	fmt.Println(t.Execute(w, Items{q, q.ToSlice()}))
}

func Create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/queue/create.html")
	t.Execute(w, nil)
}

func Push(w http.ResponseWriter, r *http.Request) {
	body, _ := url.ParseQuery(r.URL.RawQuery)
	q := getQ(r)
	q.Push(queue_item.New(body["username"][0]))
	http.Redirect(w, r, "/show?queue_id=" + body["queue_id"][0], http.StatusFound)
}

func Pop(w http.ResponseWriter, r *http.Request) {body, _ := url.ParseQuery(r.URL.RawQuery)
	q := getQ(r)
	q.Pop()
	http.Redirect(w, r, "/show?queue_id=" + body["queue_id"][0], http.StatusFound)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	getQ(r).Remove()
	http.Redirect(w, r, "/", http.StatusFound)
}