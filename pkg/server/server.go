package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/config/database"
	v1routes "github.com/ja1felipe/user-service/pkg/server/routes/v1"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() *Server {
	return &Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	dbConnection := database.NewDB()
	dbConnection.Connect()
	db := dbConnection.GetDB()

	router := s.server
	v1routes.ConfigRoutes(router, db)

	log.Printf("Server running at port: %s", s.port)
	log.Fatal(router.Run(":" + s.port))
}
