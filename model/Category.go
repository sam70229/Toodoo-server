package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Category struct {
	Uid uuid.UUID `json:"uid"`
	Title string `json:"title"`
	CreatedDate int64 `json:"createddate"`
}

func (c *Category) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}