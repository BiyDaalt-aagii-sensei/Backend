package server

import (
	db "bd/db/sqlc"
	"bd/utils"
	"database/sql"
	"log"
	"log/slog"
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

type UserResponse struct {
	Id               int64     `json:"id"`
	Username         string    `json:"username"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	PasswordAtChange time.Time `json:"password_at_change"`
	CreatedAt        time.Time `json:"created_at"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:         user.Username,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		PasswordAtChange: user.PasswordAtChange,
		CreatedAt:        user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

type updatePasswordRequest struct {
	Id       int64  `json:"id" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=10"`
}

type updatePasswordResponse struct {
	Username         string    `json:"username"`
	FirstName        string    `json:"firstname"`
	LastName         string    `json:"lastname"`
	Email            string    `json:"email"`
	PasswordAtChange time.Time `json:"password_at_change"`
	CreatedAt        time.Time `json:"created_at"`
}

func (server *Server) updateuserPassword(ctx *gin.Context) {
	var req updatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	hashPassowrd, err := utils.HashedPassword(req.Password)
	if err != nil {
		slog.Error("unable to hash password", slog.Any("err", err))
		return
	}
	arg := db.UpdatePasswordUserParams{
		Id:       req.Id,
		Password: hashPassowrd,
	}

	updatePassword, err := server.store.UpdatePasswordUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := updatePasswordResponse{
		Username:         updatePassword.Username,
		FirstName:        updatePassword.FirstName,
		LastName:         updatePassword.LastName,
		Email:            updatePassword.Email,
		PasswordAtChange: updatePassword.PasswordAtChange,
		CreatedAt:        updatePassword.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password"binding:"required,min=10"`
}

type loginResponse struct {
	AccessToken string       `json:"access_token"`
	user        UserResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, errResponse(err))
		return
	}

	accesToken, err := server.token.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := loginResponse{
		AccessToken: accesToken,
		user:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)

}
