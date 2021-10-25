package database

import (
	"context"

	"Toodoo/model"

	"github.com/google/uuid"
)

func (c *Client) AddTodo(ctx context.Context, todo model.Todo) (string, error) {
	sqlStatment := `
	INSERT INTO todo (uid, title, category, createddate, expiredDate, priority, completed, recentlyDelete)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING uid`
	var uid uuid.UUID
	row := c.db.QueryRowContext(ctx, sqlStatment, todo.Uid, todo.Title, todo.Category, todo.CreatedDate, todo.ExpiredDate, todo.Priority, todo.Completed, todo.RecentlyDelete)
	row.Scan(&uid)

	return uid.String(), nil
}

func (c *Client) GetTodos(ctx context.Context, args ...interface{}) ([]model.Todo, error) {
	sqlStatment := `
	SELECT * FROM todo
	WHERE expiredDate >= $1
	AND expiredDate <= $2
	ORDER BY priority
	`
	c.logger.Infow("sqlstatment", "string", sqlStatment, "args", args)
	rows, err := c.db.QueryContext(ctx, sqlStatment, args...)
	checkError(err)
	var todos []model.Todo

	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.CreatedDate, &todo.ExpiredDate, &todo.Priority, &todo.RemindDate, &todo.Desc, &todo.Completed, &todo.RecentlyDelete)
		todos = append(todos, todo)
	}
	c.logger.Infow("GetTodos", "todos", todos)
	return todos, nil
}

func (c *Client) GetTodoById(ctx context.Context, id string) (model.Todo, error) {
	sqlStatment := `
	SELECT * FROM todo
	WHERE uid=$1
	`
	row := c.db.QueryRowContext(ctx, sqlStatment, id)

	var todo model.Todo
	err := row.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.CreatedDate, &todo.ExpiredDate, &todo.Priority, &todo.RemindDate, &todo.Desc, &todo.Completed, &todo.RecentlyDelete)
	checkError(err)
	
	c.logger.Infow("GetTodoById", "todo", todo)
	return todo, nil
}