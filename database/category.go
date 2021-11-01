package database

import (
	"Toodoo/model"
	"context"

	"github.com/google/uuid"
)

func (c *Client) GetCategories(ctx context.Context) ([]model.Category, error) {
	sqlStatement :=`
	SELECT * FROM category
	ORDER BY createddate
	`
	rows, err := c.db.QueryContext(ctx, sqlStatement)
	checkError(err)

	var categories []model.Category

	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.Uid, &category.Title, &category.Create_at)
		categories = append(categories, category)
	}
	c.logger.Infow("GetCategories", "categories", categories)
	return categories, nil
}

func (c *Client) AddCategory(ctx context.Context, category model.Category) (string, error) {
	sqlStatment := `
	INSERT INTO category (uid, title, createddate)
	VALUES ($1, $2, $3)
	RETURNING uid`
	var uid uuid.UUID
	row := c.db.QueryRowContext(ctx, sqlStatment, category.Uid, category.Title, category.Create_at)
	row.Scan(&uid)

	return uid.String(), nil
}
