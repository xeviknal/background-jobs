package models

import (
	"time"
)

type Job struct {
	Id        int64     `db:"id, primarykey, autoincrement" json:"id"`
	ObjectId  int64     `db:"object_id" json:"object_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func NewJob(ObjectId int64) Job {
	return Job{
		ObjectId:  ObjectId,
		CreatedAt: time.Now(),
	}
}
