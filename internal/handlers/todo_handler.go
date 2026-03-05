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

		userIDInterface, exists := c.Get("user_id")

		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "user_id not found in the context."})
			return
		}

		userID := userIDInterface.(string)


		var input CreateTodoInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodo(c.Request.Context(), pool, input.Title, input.Completed, userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, todo)
	}
}

func GetAllTodohandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		userIDInterface , exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "user_id not found in context."})
			return
		}

		//why are we asserting a type?
		//cause we specifically want string

		userID := userIDInterface.(string)

		todos, err := repository.GetAllTodos(pool, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)

	}
}

func GetTodoByIDHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		userIDInterface , exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "user_id not found in context."})
			return
		}

		//why are we asserting a type?
		//cause we specifically want string

		userID := userIDInterface.(string)


		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TODO id."})
			return
		}

		todo, err := repository.GetTodoByID(pool, id, userID)
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

		userIDInterface , exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "user_id not found in context."})
			return
		}

		//why are we asserting a type?
		//cause we specifically want string

		userID := userIDInterface.(string)

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

		existing, err := repository.GetTodoByID(pool, id, userID)
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

		todo, err := repository.UpdateTodo(pool, id, title, completed, userID)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todo)

	}

}

func DeleteToDoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		userIDInterface , exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "user_id not found in context."})
			return
		}

		//why are we asserting a type?
		//cause we specifically want string

		userID := userIDInterface.(string)

			idStr := c.Param("id")
			id, err :=	strconv.Atoi(idStr)
			if err != nil {

				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo id."})
			}

			err = repository.DeleteTodo(pool, id, userID)
			if err != nil {

					if err.Error() == "todo with id"+idStr+"not found!" {
					c.JSON(http.StatusNotFound, gin.H{"error" : "Todo not found."})
					return
			}	

			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message" : "Todo deleted successfully."})
}
}