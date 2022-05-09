package api

import (
	"net/http"

	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserOwner string `json:"userOwner" binding:"required"`
	Currency  string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	args := db.CreateAccountParams{
		UserOwner: req.UserOwner,
		Currency:  req.Currency,
		Balance:   0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
