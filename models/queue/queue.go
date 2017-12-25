package queue

import (
	"../../models"
	"../queue_item"
	"database/sql"
	"strconv"
	"fmt"
)

type Queue struct {
	Id          int32
	name        string
	firstItemId int32
	lastItemId  int32
}

func (q *Queue) get(row *sql.Row) (err error) {
	err = row.Scan(&(q.Id), &(q.name), &(q.firstItemId), &(q.lastItemId))
	return
}

func (q Queue) set(fieldName, value string) {
	models.SetFieldById("queues", fieldName, value, q.Id)
}

func New(name string) (q Queue, err error) {
	fakeItem, _ := queue_item.New(0)
	q = Queue{0, name, fakeItem.Id, fakeItem.Id}
	row := models.DB.QueryRow(`INSERT INTO queues (name, first_item_id, last_item_id)
							  			VALUES ($1, $2, $3)
							  			RETURNING id`, name, fakeItem.Id, fakeItem.Id)
	err = row.Scan(&(q.Id))
	fmt.Println(err)
	return
}

func Find(id int32) (q Queue) {
	q.get(models.DB.QueryRow(`SELECT * FROM queues WHERE id=$1`, id))
	return
}

func (q Queue) Items(c chan queue_item.QueueItem) {
	for currItem := queue_item.Find(q.firstItemId).Next();; currItem = currItem.Next() {
		c <- currItem
		if currItem.NextItemId == 0 { break }
	}
	close(c)
}

func (q Queue) firstItem() queue_item.QueueItem {
	return queue_item.Find(q.firstItemId)
}

func (q Queue) RemoveRow() {
	models.DB.Exec(`DELETE FROM queues WHERE id=$1`, q.Id)
}

func (q Queue) Remove() {
	c := make(chan queue_item.QueueItem)
	go q.Items(c)
	for item := range c {
		item.RemoveRow()
	}
	queue_item.Find(q.firstItemId).Remove()
	q.RemoveRow()
	return
}

func (q *Queue) Push(newItem queue_item.QueueItem) {
	queue_item.Push(q.lastItemId, &newItem)
	q.lastItemId = newItem.Id
	q.set("last_item_id", strconv.Itoa(int(newItem.Id)))
}

func (q *Queue) Erase(itemId int32) {
	eraseItem := queue_item.Find(itemId)
	if eraseItem.NextItemId == 0 {
		q.lastItemId = eraseItem.PrevItemId
		q.set("last_item_id", strconv.Itoa(int(eraseItem.PrevItemId)))
	}
	eraseItem.Remove()
}

func (q *Queue) Pop() {
	if q.lastItemId == q.firstItemId { return }
	q.Erase(queue_item.Find(q.firstItemId).Next().Id)
}

func (q Queue) ToSlice() (items []queue_item.QueueItem) {
	items = nil
	c := make(chan queue_item.QueueItem)
	go q.Items(c)
	for item := range c {
		items = append(items, item)
	}
	return
}