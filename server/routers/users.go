package routers

import (
	"net/http"
	"srv/db"
	"srv/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRouterUsers(router *gin.RouterGroup) {
	router.GET("/v1/users", getUsers)
	router.GET("/v1/users/:id", getUser)
	router.POST("/v1/users", createUser)
}

func getUser(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := uint(u64)
	user, err := db.GetUser(id)
	transactions, err := db.GetTransactionsUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"user":         user,
		"transactions": db.NormalizeTransactions(transactions),
	})

}

func getUsers(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var user = structs.CreateUserRequest{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
	}
	if user.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: "doesn't have username",
		})
		return
	}
	id, err := db.CreateUser(user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user_id": id,
	})
}
