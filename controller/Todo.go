package controller

import (
	"Toodoo/logger"
	_ "Toodoo/logger"

	Model "Toodoo/model"
	"fmt"

	"reflect"
	"strings"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var TodoList []Model.Todo

// type ApiResponse struct {
// 	ResultCode string
// 	Data interface{}
// }

type ApiResponse struct {
	Data interface{}
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("Received %s %s", r.URL, r.Body)
	decoder := json.NewDecoder(r.Body)
	logger.Info.Printf("Body %v", decoder)
	// decoder.More()
	var request map[string][]Model.Todo
	err := decoder.Decode(&request)
	logger.Info.Printf("request %v", request)
	checkError(err)

	uid := InsertTodo(request["todos"][0])

	response := ApiResponse{uid}
	d, _ := json.Marshal(response)

	w.Write(d)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	 todos := FetchTodos()
	 response := ApiResponse{todos}
	 d, _ := json.Marshal(response)

	 w.Write(d)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["uid"]
	logger.Info.Println("Received ", r.URL, "query for uid:", queryId)

	var targetTodo Model.Todo
	targetTodo = FetchTodo(queryId)

	response := ApiResponse{targetTodo}
	d, _ := json.Marshal(response)
	w.Write(d)
}

func diffTodo(t1 Model.Todo, t2 Model.Todo) map[string]interface{} {
	m := make(map[string]interface{})
	if t1 == t2 {
		return m
	}
	t1v, t1f := reflect.ValueOf(t1), reflect.TypeOf(t1)
	t2v := reflect.ValueOf(t2)

	for i := 0; i < t1f.NumField(); i++ {
		if t1v.Field(i).Interface() != t2v.Field(i) {
			m[strings.ToLower(t1f.Field(i).Name)] = t2v.Field(i).Interface()
		}
		// fmt.Println(strings.ToLower(v1.Field(i).Name), v.Field(i).Interface())
	}
	fmt.Println(m)
	return m
}
