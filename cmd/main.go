package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mariiasalikova/go-time-tracker/controllers"
	"github.com/mariiasalikova/go-time-tracker/models"
	"github.com/mariiasalikova/go-time-tracker/repositories"
	"github.com/mariiasalikova/go-time-tracker/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading .env: %s", err.Error())
	}

	r := gin.New()

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	// User components
	userRepo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userController := &controllers.UserController{Service: userService}

	// Task components
	taskRepo := &repositories.TaskRepository{DB: db}
	taskService := &services.TaskService{Repo: taskRepo}
	taskController := &controllers.TaskController{Service: taskService}

	// Time components
	timeRepo := &repositories.TimeRepository{DB: db}
	timeService := &services.TimeService{Repo: timeRepo}
	timeController := &controllers.TimeController{Service: timeService}

	// User routes
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.GET("/users", userController.ListUsers)

	// Task routes
	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks/:id", taskController.GetTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)
	r.GET("/tasks", taskController.ListTasks)

	// Time routes
	r.POST("/time/start", timeController.StartTime)
	r.PUT("/time/stop/:id", timeController.StopTime)
	r.GET("/time/entries/:userId", timeController.GetTimeEntries)

	srv := new(models.Server)

	//todo: groupize and log it

	logrus.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured while db closing: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()

}
