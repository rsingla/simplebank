package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "simplebank/sqlc"
)

type createAccount struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR INR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByAccountId(ctx, id)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, account)
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	account, err := server.store.DeleteAccount(ctx, id)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, account)
}
