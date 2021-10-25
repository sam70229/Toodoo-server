// package controller

// import (
// 	"Toodoo/database"
// 	"context"

// 	"github.com/google/uuid"
// )

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func InsertTodo(todo Todo) interface{} {
// 	sqlStatment := `
// 	INSERT INTO todo (uid, title, category, createddate, priority, completed, recentlyDelete)
// 	VALUES ($1, $2, $3, $4, $5, $6, $7)
// 	RETURNING uid`
// 	var uid uuid.UUID
// 	row, err := database.QueryOne(sqlStatment, todo.Uid, todo.Title, todo.Category, todo.CreatedDate, todo.Priority, todo.Completed, todo.RecentlyDelete)
// 	row.Scan(&uid)
// 	checkError(err)
// 	return (map[string]string{"uid": uid.String()})
// }

// func FetchTodo(crx context.Context, id string) Todo {
// 	sqlStatment := `
// 	SELECT * FROM todo
// 	WHERE uid=$1
// 	`
// 	row, err := database.QueryOne(sqlStatment, id)
// 	// row := DB.QueryRow(sqlStatment, id)
// 	var todo Todo
// 	err = row.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.CreatedDate, &todo.ExpiredDate, &todo.Priority, &todo.RemindDate, &todo.Desc, &todo.Completed, &todo.RecentlyDelete)
// 	checkError(err)

// 	return todo
// }

// func FetchTodos() []Todo  {
// 	sqlStatment := `
// 	SELECT * FROM todo
// 	ORDER BY priority
// 	`
// 	rows, err := database.QueryMany(sqlStatment)
// 	checkError(err)
// 	var todos []Todo

// 	for rows.Next() {
// 		var todo Todo
// 		err = rows.Scan(&todo.Uid, &todo.Title, &todo.Category, &todo.CreatedDate, &todo.ExpiredDate, &todo.Priority, &todo.RemindDate, &todo.Desc, &todo.Completed, &todo.RecentlyDelete)
// 		todos = append(todos, todo)
// 	}

// 	return todos
// }

// // func updateTodo(id string) {
// // 	sqlStatement := `
	
// // 	`
// // }

// // func disconnect() error {
// // 	return DB.Close()
// // }