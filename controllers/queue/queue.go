package queue

import (
	"net/http"
	"net/url"
	queue_model "../../models/queue"
	"../../models/queue_item"
)

func Save(w http.ResponseWriter, r *http.Request) {
	body, _ := url.ParseQuery(r.URL.RawQuery)
	q, _ := queue_model.New(body["queue_name"][0])
	for _, username := range body["item[]"] {
		q.Push(queue_item.New(username))
	}
	http.Redirect(w, r, "/", http.StatusFound)
}