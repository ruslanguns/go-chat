package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ruslanguns/go-chat/internal/database"
	"github.com/ruslanguns/go-chat/internal/handler"
	"github.com/ruslanguns/go-chat/internal/repository"
	"github.com/ruslanguns/go-chat/internal/service"
)

type Server struct {
	port int

	db database.Service

	userHandler    *handler.UserHandler
	channelHandler *handler.ChannelHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()

	// Run Migrations
	err := db.Migrate()
	if err != nil {
		panic(fmt.Sprintf("failed to run migrations %v", err))
	}

	gormDB := db.GetDB()
	userRepo := repository.NewUserRepository(gormDB)
	channelRepo := repository.NewChannelRepository(gormDB)

	userService := service.NewUserService(userRepo)
	channelService := service.NewChannelService(channelRepo, userRepo)

	newServer := &Server{
		port:           port,
		db:             db,
		userHandler:    handler.NewUserHandler(userService),
		channelHandler: handler.NewChannelHandler(channelService),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
