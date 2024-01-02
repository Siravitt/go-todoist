package router

import (
	"fmt"
	"net/http"

	"github.com/Siravitt/go-todoist/handler"
	"github.com/Siravitt/go-todoist/logs"
	"github.com/Siravitt/go-todoist/middlewares/auth_middleware"
	"github.com/Siravitt/go-todoist/repository/todo_repository"
	"github.com/Siravitt/go-todoist/repository/user_repository"
	"github.com/Siravitt/go-todoist/service/auth_service"
	"github.com/Siravitt/go-todoist/service/todo_service"
	"github.com/Siravitt/go-todoist/service/user_service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func RunServer(db *sqlx.DB) {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// User init
	userRepo := user_repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Auth init
	authService := auth_service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := auth_middleware.NewAuthMiddleware()

	// Todo init
	todoRepo := todo_repository.NewTodoRepositoryDB(db)
	todoService := todo_service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// Auth route
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)

	// User route
	e.GET("/users", userHandler.GetUsers)
	e.GET("/user/:id", userHandler.GetUser)

	protected := e.Group("/v1", authMiddleware.AuthorizationMiddleware)

	// ! Todo route
	protected.GET("/todos", todoHandler.GetTodos)
	protected.GET("/todo/:id", todoHandler.GetTodo)
	protected.POST("/todo", todoHandler.AddTodo)
	// protected.PATCH("/todo/:id", )
	// protected.DELETE("/todo/:id", )

	logs.Info("Server running at port: " + viper.GetString("app.port"))
	if err := e.Start(fmt.Sprintf(":%v", viper.GetInt("app.port"))); err != http.ErrServerClosed {
		logs.Error(err)
		panic(err)
	}
}
