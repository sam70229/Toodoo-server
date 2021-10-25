package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoryString(t *testing.T) {
	category := &Category{
		Uid: uuid.New(),
		Title: "Test category",
		CreatedDate: time.Now().Unix(),

	}
	b, err := json.Marshal(category)
	assert.Nil(t, err)
	assert.Equal(t, string(b), category.String())
}