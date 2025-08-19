package handlers

import (
	"gin-auth-sql/database"
	"gin-auth-sql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	username := c.GetString("username")
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.QueryRow("INSERT INTO todos(username,title,completed) VALUES($1,$2,$3) RETURNING id,created_at", username, todo.Title, todo.Completed).Scan(&todo.ID, &todo.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	todo.Username = username
	c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	username := c.GetString("username")
	rows, err := database.DB.Query("SELECT id,username,title,completed,created_at FROM todos WHERE username=$1", username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Username, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning todo"})
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	username := c.GetString("username")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id"})
		return
	}
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := database.DB.Exec("UPDATE todos SET title=$1,completed=$2 WHERE id=$3 AND username=$4", todo.Title, todo.Completed, id, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func DeleteTodo(c *gin.Context) {
	username := c.GetString("username")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id"})
		return
	}
	res, err := database.DB.Exec("DELETE FROM todos WHERE id=$1 AND username=$2", id, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
