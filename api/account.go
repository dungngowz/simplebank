package api

import (
	"net/http"

	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/dungngowz/simple_bank/util"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID   int64  `json:"userID" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ErrorBadRequest(ctx, err)
	}

	args := db.CreateAccountParams{
		UserID:   req.UserID,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		util.ErrorBadRequest(ctx, err)
	}

	ctx.JSON(http.StatusOK, account)
}
