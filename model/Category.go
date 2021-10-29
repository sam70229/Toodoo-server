package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Category struct {
	Uid uuid.UUID `json:"uid"`
	Title string `json:"title"`
	Create_at int64 `json:"create_at"`
}

func (c *Category) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
