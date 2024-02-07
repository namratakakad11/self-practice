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

var todos = []todo{
	{ID: "1", Item: "cleanroom", Completed: false},
	{ID: "2", Item: "read book", Completed: false},
	{ID: "3", Item: "task 2", Completed: false},
}
	
// get all data
// incoming http request is store in context eg data inside of request body
func gettodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// adding data
func addtodos(c *gin.Context) {
	var newtodo todo

	if err := c.ShouldBindJSON(&newtodo); err != nil {
		return
	}
	todos = append(todos, newtodo)

	c.IndentedJSON(http.StatusCreated, newtodo)
}


func deletebyid(c *gin.Context){
	id := c.Param("id")
	todo, err := getById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found by id"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, todo)

}
// get data by id logical part
func getById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("id not found")
}

// handler  for get data by id
func gettodohandler(c *gin.Context) {
	id := c.Param("id")
	todo, err := getById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found by id"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

//  patch function
func updatebyid(c *gin.Context) {
	id := c.Param("id")
	todo, err := getById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found by id"})
		return
	}
	// it will make true false and false true
	todo.Completed = !todo.Completed

	c.IndentedJSON(http.StatusOK, todo)
}

// main function to call all function
func main() {
	rout := gin.Default()
	rout.GET("/todos", gettodos)
	rout.GET("/todos/:id", gettodohandler)
	rout.PATCH("/todos/:id", updatebyid)
	rout.POST("/todos", addtodos)
	rout.Run("localhost:8080")
}



//pointers reference to memory address -