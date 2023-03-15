package api

import (
	"net/http"
	"simplebank/sqlc"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required, currency"`
}

func (server *Server) transferRequest(ctx *gin.Context) {
	var body transferRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !server.validAccount(ctx, body.FromAccountID, body.Currency) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "currency mismatch"})
		return
	}

	if !server.validAccount(ctx, body.ToAccountID, body.Currency) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "currency mismatch"})
		return
	}

	arg := sqlc.TransferTxParams{
		FromAccountID: body.FromAccountID,
		ToAccountID:   body.ToAccountID,
		Amount:        body.Amount,
	}
	account, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccountByAccountId(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}

	if account.Currency != currency {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "currency mismatch"})
		return false
	}

	return true
}
