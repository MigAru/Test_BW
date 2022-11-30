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
    router.POST("/v1/transactions", createTransaction)
}

func getTransaction(c *gin.Context) {
    u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
            c.JSON(http.StatusBadRequest, gin.H{
                "message":"does't find opiration_type",
            }) 
            return
    }

    if ok := db.ValidateBalace(req.UserID, req.Amount, operationType); !ok {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "breake reuest because operation not valid",
        })
        return
    }

    transactionID, err := db.CreateTransaction(req.Amount, operationType, req.UserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "breake request because operation not valid",
        })
        return
    }
    c.AbortWithStatusJSON(http.StatusCreated, structs.MessageResponse{
        Message: fmt.Sprint(transactionID),
    })

}
