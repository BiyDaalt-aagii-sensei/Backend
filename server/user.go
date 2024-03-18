package server

import (
	db "bd/db/sqlc"
	"bd/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username  string `json:"username" binding:"required,alphanum"`
	Password  string `json:"password" binding:"required,min=10"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Id               int64     `json:"id"`
	Username         string    `json:"username"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	PasswordAtChange time.Time `json:"password_at_change"`
	CreatedAt        time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errResponse(err))
		return
	}
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		log.Fatalf("cant hashed password %v", err)
	}
	arg := db.CreateUserParams{
		Username:  req.Username,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errResponse(err))
		return
	}

	rsp := createUserResponse{
		Id:               user.Id,
		Username:         user.Username,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		PasswordAtChange: user.PasswordAtChange,
		CreatedAt:        user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
