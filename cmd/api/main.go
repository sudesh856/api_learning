package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"todo_api/internal/handlers"
	"todo_api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Configuration was not correct!",err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database.",err)

	}

	defer pool.Close()


	var router *gin.Engine = gin.Default()

	router.GET("/", func(c *gin.Context){



		c.JSON(200, gin.H{

			"message": "todo api",
			"status": "success",
			"database": "connected",
		})
	})


	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.POST("/auth/login", handlers.LoginHandler(pool, cfg))

	protected := router.Group("/todos")
	protected.Use(middleware.AuthMiddleware(cfg))
	{

	protected.POST("", handlers.CreateTodoHandler(pool))	
	protected.GET("/", handlers.GetAllTodohandler(pool))
	protected.GET("/:id", handlers.GetTodoByIDHandler(pool))
	protected.PUT("/:id", handlers.UpdateToDoHandler(pool))
	protected.DELETE("/:id", handlers.DeleteToDoHandler(pool))
	}


	
	

	//test route for middleware.
	router.GET("/test_for_protected", middleware.AuthMiddleware(cfg), handlers.TestProtectedHandler())

	router.Run(":" + cfg.Port)


}

	