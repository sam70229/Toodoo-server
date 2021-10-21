package model

import (
	// "database/sql"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

)

type Todo struct {
	Uid uuid.UUID `json:"uid"`
	Title string `json:"title"`
	Category string `json:"category"`
	CreatedDate int64 `json:"createdDate"`
	ExpiredDate null.Int `json:"expiredDate,omitempty"`
	Priority int `json:"priority"`
	RemindDate null.String `json:"remindDate"`
	Desc null.String `json:"desc"`
	Completed bool `json:"completed"`
	RecentlyDelete bool `json:"recentlyDelete"`
}

func NewTodo() Todo {
	var todo Todo
	todo.Uid = uuid.New()
	todo.Title = "Test title"
	todo.CreatedDate = time.Now().Unix()
	todo.Completed = false
	todo.RecentlyDelete = false
	return todo
}

// let uid: UUID
// var title: String
// var category: String
// var createdDate: Date
// var expiredDate: Date?
// var priority: Int
// var remind: String?
// var desc: String?
// var subtasks: [TDSubtask]?
// var completed: Bool
// var delete: Bool