package queue_item

import (
	"../../models"
	"database/sql"
	"strconv"
	"fmt"
)

type QueueItem struct {
	Id         int32
	UserName   string
	NextItemId int32
	PrevItemId int32
}

func get(row *sql.Row) (qItem QueueItem) {
	var id, nextItemId, prevItemId sql.NullInt64
	var userName sql.NullString
	row.Scan(&id, &userName, &nextItemId, &prevItemId)
	qItem.Id = int32(id.Int64)
	if userName.Valid {
		qItem.UserName = userName.String
	}
	if nextItemId.Valid {
		qItem.NextItemId = int32(nextItemId.Int64)
	}
	if prevItemId.Valid {
		qItem.PrevItemId = int32(prevItemId.Int64)
	}
	return
}

func Set(id int32, fieldName string, value string) {
	models.SetFieldById("queue_items", fieldName, value, id)
}

func New(userName string) (qItem QueueItem, err error) {
	qItem = QueueItem{0, userName, 0, 0}
	row := models.DB.QueryRow(`INSERT INTO queue_items (user_id)
							  			VALUES ($1)
							  			RETURNING id`, userName)
	err = row.Scan(&(qItem.Id))
	return
}

func Find(id int32) (qItem QueueItem) {
	qItem = get(models.DB.QueryRow(`SELECT * FROM queue_items WHERE id=$1`, id))
	return
}

func (qItem QueueItem) RemoveRow() {
	fmt.Println(models.DB.Exec(`DELETE FROM queue_items WHERE id=$1`, qItem.Id))
}

func (qItem QueueItem) Remove() {
	if qItem.PrevItemId != 0 {
		Set(qItem.PrevItemId, "next_item_id", strconv.Itoa(int(qItem.NextItemId)))
	}
	if qItem.NextItemId != 0 {
		Set(qItem.NextItemId, "prev_item_id", strconv.Itoa(int(qItem.PrevItemId)))
	}
	qItem.RemoveRow()
	return
}

func (qItem QueueItem) Next() (nextItem QueueItem) {
	 nextItem = Find(qItem.NextItemId)
	 return
}

func Push(lastId int32, nextQItem *QueueItem) {
	Set(lastId, "next_item_id", strconv.Itoa(int(nextQItem.Id)))
	nextQItem.PrevItemId = lastId
	Set(nextQItem.Id, "prev_item_id", strconv.Itoa(int(lastId)))
}