package handlers

import (
	"net/http"
	"strconv"
	"todo_api/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoInput struct {
	Title     *string `json:"title"`
	Completed *bool   `jsson:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input CreateTodoInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodo(c.Request.Context(), pool, input.Title, input.Completed)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, todo)
	}
}

func GetAllTodohandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		todos, err := repository.GetAllTodos(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)

	}
}

func GetTodoByHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TODO id."})
			return
		}

		todo, err := repository.GetTodoByID(pool, id)
		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found!"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, todo)
	}

}

func UpdateToDoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID."})
			return
		}

		var input UpdateTodoInput

		if err = c.ShouldBindJSON(&input); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.Title == nil && input.Completed == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "at least one field (title or completed) must be provided."})
			return
		}

		existing, err := repository.GetTodoByID(pool, id)
		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Todo NOT found!"})
				return
			}
		}

		title := existing.Title
		if input.Title != nil {
			title = *input.Title
		}

		completed := existing.Completed
		if input.Completed != nil {

			completed = *input.Completed
		}

		todo, err := repository.UpdateTodo(pool, id, title, completed)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todo)

	}

}

func DeleteToDoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Param("id")
	id, err :=	strconv.Atoi(idStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid todo ID."})
	}

	err  = repository.DeleteTodo(pool, id)
	if err != nil { 
		if err.Error() == "todo with id" +idStr+ "was not found." {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Todo deleted successfully!"})


	}
}
