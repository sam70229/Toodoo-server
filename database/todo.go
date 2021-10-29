package database

import (
	"context"
	"fmt"

	"Toodoo/model"

	// "github.com/google/uuid"
)

func (c *Client) AddTodo(ctx context.Context, todo model.Todo) (model.Todo, error) {
	sqlStatment := `
    INSERT INTO todo (uid, title, category, create_at, expire_at, priority, remind_at, description, completed, recentlyDelete)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING uid, title, category, create_at, expire_at, priority, remind_at, description, completed, recentlyDelete`
	var insertTodo model.Todo
	row := c.db.QueryRowContext(ctx, sqlStatment, todo.Uid, todo.Title, todo.Category, todo.Create_at, todo.Expire_at, todo.Priority, todo.Remind_at, todo.Comment, todo.Completed, todo.RecentlyDelete)

	row.Scan(&insertTodo.Uid, &insertTodo.Title, &insertTodo.Category, &insertTodo.Create_at, &insertTodo.Expire_at, &insertTodo.Priority, &insertTodo.Remind_at, &insertTodo.Comment, &insertTodo.Completed, &insertTodo.RecentlyDelete)
	fmt.Println(insertTodo)
	
	return insertTodo, nil
}

func (c *Client) GetTodos(ctx context.Context, args ...interface{}) ([]model.Todo, error) {
	sqlStatment := `
	SELECT * FROM todo
	WHERE expire_at >= $1
	AND expire_at <= $2
	ORDER BY priority
	`
	c.logger.Infow("sqlstatment", "string", sqlStatment, "args", args)
	rows, err := c.db.QueryContext(ctx, sqlStatment, args...)
	checkError(err)
	var todos []model.Todo

	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.Create_at, &todo.Expire_at, &todo.Priority, &todo.Remind_at, &todo.Comment, &todo.Completed, &todo.RecentlyDelete)
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
	err := row.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.Create_at, &todo.Expire_at, &todo.Priority, &todo.Remind_at, &todo.Comment, &todo.Completed, &todo.RecentlyDelete)
	checkError(err)

	c.logger.Infow("GetTodoById", "todo", todo)
	return todo, nil
}

func (c *Client) GetTodosByCategory(ctx context.Context, category string) ([]model.Todo, error) {
	sqlStatement := `
	SELECT * FROM todo
	WHERE category=$1
	`
	rows, err := c.db.QueryContext(ctx, sqlStatement, category)

	if err != nil {
		return nil, err
	}
	var todos []model.Todo
	
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.Create_at, &todo.Expire_at, &todo.Priority, &todo.Remind_at, &todo.Comment, &todo.Completed, &todo.RecentlyDelete)
		todos = append(todos, todo)
	}

	return todos, nil
}
