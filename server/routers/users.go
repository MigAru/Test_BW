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

//	@BasePath	/api/v1
//	@Summary	users
//	@Schemes
//	@Param			user_id	path	int	true	"User ID"
//	@Description	gives user transactions and user data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.UserTransactions
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/users/{user_id} [get]
func getUser(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.MessageResponse{
            Message: err.Error(),
        })
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
	c.AbortWithStatusJSON(http.StatusOK, db.UserTransactions{
		User:         user,
		Transactions: db.NormalizeTransactions(transactions),
	})

}

//	@BasePath	/api/v1
//	@Summary	users
//	@Schemes
//	@Description	get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		db.User
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/users [get]
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

//	@BasePath	/api/v1
//	@Summary	users
//	@Schemes
//	@Param			request	body	structs.CreateUserRequest	true	"Create Param"
//	@Description	creating new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Create			201	{object}	structs.MessageResponse
//	@Router			/users [post]
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
