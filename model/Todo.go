package model

import (
	"encoding/json"

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

func (t *Todo) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}
