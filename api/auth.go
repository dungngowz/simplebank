package api

import (
	"database/sql"
	"net/http"

	"github.com/dungngowz/simple_bank/util"
	"github.com/gin-gonic/gin"
)

//////////////////////////////////// LOGIN USER ////////////////////////////////////////////////////////////
type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ErrorBadRequest(ctx, err)
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.ErrorNotFound(ctx, err)
			return
		}
		util.ErrorInternalServer(ctx, err)
		return
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		util.ErrorUnauthorized(ctx, err)
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		util.ErrorInternalServer(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"user":        user,
	})
}
