package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// this is not json
var todos = []todo{
	{ID: "1", Item: "Code", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Sleep", Completed: false},
	{ID: "4", Item: "Eat", Completed: false},
}

// information about incoming request
//
//	|			|Type
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) //conversts structure to json
}

func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
		//puts return in todo. if structure of todo not match input throws error
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// used by handler
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id) //gets id
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id) //gets id
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

// gets structure from on id
func getTodosById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default() //server
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo) //:id indicates dynamic. This is a path parameter
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090") //run server on port 9090
}
