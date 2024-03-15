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

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
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

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, exists := todos[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yikes, Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo TodoItem
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, exists := todos[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yikes, Todo not found"})
		return
	}
	todos[id] = updatedTodo
	c.JSON(http.StatusOK, updatedTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, exists := todos[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yikes, Todo not found"})
		return
	}
	delete(
		todos,
		id,
	)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted 👍."})
}
