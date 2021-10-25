package api
import (
	"context"

	"Toodoo/model"
)

type APIStore interface {
	//Todo
	GetTodoById(ctx context.Context, id string) (model.Todo, error)
	AddTodo(ctx context.Context, todo model.Todo) (string, error)
	GetTodos(ctx context.Context, args ...interface{}) ([]model.Todo, error)

	//Category
	GetCategories(ctx context.Context) ([]model.Category, error)
	AddCategory(ctx context.Context, category model.Category) (string, error)
}
