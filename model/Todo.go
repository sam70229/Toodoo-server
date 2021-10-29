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
	Create_at int64 `json:"create_at"`
    Expire_at null.Int `json:"expire_at,omitempty"`
	Priority int `json:"priority"`
	Remind_at null.String `json:"remind_at,omitempty"`
	Comment null.String `json:"comment,omitempty"`
	Completed bool `json:"completed"`
	RecentlyDelete bool `json:"recentlyDelete"`
}

func (t *Todo) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}
