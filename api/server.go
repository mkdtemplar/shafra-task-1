package api

import (
	"shafra-task-1/internal/database/interfaces"
	"shafra-task-1/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	Router *gin.Engine
	store  interfaces.IUserInterface
}

func NewServer(config utils.Config, store interfaces.IUserInterface) (*Server, error) {

	server := &Server{config: config, store: store}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	gin.ForceConsoleColor()

	router.POST("/create-user", s.CreateUser)
	router.GET("/get-user/:id", s.GetUserByID)
	router.PATCH("/update-user/:id", s.UpdateUser)
	router.DELETE("/delete-user/:id", s.DeleteUser)

	s.Router = router
}

func (s *Server) Run(address string) error {
	return s.Router.Run(address)
}
