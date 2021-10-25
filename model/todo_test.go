package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTodoString(t *testing.T) {
	todo := &Todo{
		Uid: uuid.New(),
		Title: "Test title",
		Category: "Test Category",
		CreatedDate: time.Now().Unix(),

	}
	b, err := json.Marshal(todo)
	assert.Nil(t, err)
	assert.Equal(t, string(b), todo.String())
}