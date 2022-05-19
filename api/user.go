package api

import (
	"net/http"

	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/dungngowz/simple_bank/util"
	"github.com/gin-gonic/gin"
)

type createUserParams struct {
	Username       string `json:"username" binding:"required,alphanum"`
	HashedPassword string `json:"hashedPassword" binding:"required,min=6"`
	Fullname       string `json:"fullname" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserParams

	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorBadRequest(ctx, err)
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		Fullname:       req.Fullname,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		util.ErrorBadRequest(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
