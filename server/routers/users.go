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
    }
    id := uint(u64)
    user, err := db.GetUser(id)
    c.AbortWithStatusJSON(http.StatusOK, user)

}

func getUsers(c *gin.Context) {
    users, err := db.GetUsers()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    }
    c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
    var user = structs.CreateUserRequest{}
    if err := c.BindJSON(&user); err != nil{
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    }
    if user.Username == "" {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "doesn't have username"})
    }
    id, err := db.CreateUser(user.Username)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    }
    c.JSON(http.StatusOK, gin.H{
        "user_id": id,
    })
}
