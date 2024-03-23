package server

import (
	db "bd/db/sqlc"
	"bd/token"
	"bd/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store  *db.Store
	router *gin.Engine
	token  token.Maker
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Токен үүсгэж чадсангүй")
	}

	server := &Server{
		config: config,
		store:  store,
		token:  token,
	}

	server.routes()
	return server, nil
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
