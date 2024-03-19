package server

import (
	db "bd/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// User
	router.POST("/api/createuser", server.createUser)
	router.POST("/api/update/password", server.updateuserPassword)

	server.router = router
	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func (server *Server) errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
