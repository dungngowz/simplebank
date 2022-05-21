package api

import (
	"net/http"

	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/dungngowz/simple_bank/util"
	"github.com/gin-gonic/gin"
)

//////////////////////////////////// CREATE USER ////////////////////////////////////////////////////////////
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ErrorBadRequest(ctx, err)
		return
	}

	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		util.ErrorInternalServer(ctx, err)
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		util.ErrorBadRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
