package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var todos = make(map[string]TodoItem)

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.GET("/todos/:id", getTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	todoSlice := make([]TodoItem, 0, len(todos))
	for _, todo := range todos {
		todoSlice = append(todoSlice, todo)
	}
	c.JSON(http.StatusOK, todoSlice)
}

func createTodo(c *gin.Context) {
	var newTodo TodoItem
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos[newTodo.ID] = newTodo
	c.JSON(http.StatusCreated, newTodo)
}
