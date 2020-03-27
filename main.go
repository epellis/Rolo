package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Login information for a user
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginUser(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "email": json.Email})
}

func main() {
	r := gin.Default()
	r.POST("/auth/login", loginUser)
	r.Use(static.Serve("/", static.LocalFile("./client/public", true)))
	panic(r.Run())
}
