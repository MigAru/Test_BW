package routers

import (
	"fmt"
	"net/http"
	"srv/db"
	"srv/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRouterTransactions(router *gin.RouterGroup) {
	router.GET("/v1/transactions/:id", getTransaction)
	router.GET("/v1/transactions/:id/pop", popTransactionsQueue)
	router.POST("/v1/transactions", createTransaction)
}

//	@BasePath	/api/v1
//	@Summary	transactions
//	@Schemes
//	@Param			user_id	path	int	true	"Transaction ID"
//	@Description	gives active from queue transactions user
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.TransactionResponse
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/transactions/{user_id}/pop [get]
func popTransactionsQueue(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	id := uint(u64)
	transaction, err := db.GetFirstActiveTransaction(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	transactionResp := db.NormalizeTransactions([]db.Transaction{transaction})
	c.AbortWithStatusJSON(http.StatusOK, transactionResp[0])
}

//	@BasePath	/api/v1
//	@Summary	transactions
//	@Schemes
//	@Param			transaction_id	path	int	true	"Transaction ID"
//	@Description	gives user transactions and user data
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.TransactionResponse
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/transactions/{transaction_id} [get]
func getTransaction(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	id := uint(u64)
	transaction, err := db.GetTransaction(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	transactionResp := db.NormalizeTransactions([]db.Transaction{transaction})
	c.AbortWithStatusJSON(http.StatusOK, transactionResp[0])

}

//	@BasePath	/api/v1
//	@Summary	transactions
//	@Schemes
//	@Param			request	body	structs.CreateTransactionRequest	true	"Create Param"
//	@Description	creating new transaction and return id string
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	structs.MessageResponse
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/transactions [post]
func createTransaction(c *gin.Context) {
	req := structs.CreateTransactionRequest{}
	operationType := -1

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	switch req.OperationType {
	case "add":
		operationType = db.AddPrice
	case "reduce":
		operationType = db.ReducePrice
	default:
		c.JSON(http.StatusBadRequest, structs.MessageResponse{
			Message: "does't find opiration_type",
		})
		return
	}

	if ok := db.ValidateBalace(req.UserID, req.Amount, operationType); !ok {
		c.JSON(http.StatusBadRequest, structs.MessageResponse{
			Message: "breake reuest because operation not valid",
		})
		return
	}

	transactionID, err := db.CreateTransaction(req.Amount, operationType, req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.MessageResponse{
			Message: "breake request because operation not valid",
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusCreated, structs.MessageResponse{
		Message: fmt.Sprint(transactionID),
	})

}
