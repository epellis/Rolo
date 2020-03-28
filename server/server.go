package server

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Server stores all state and associated handlers
type Server struct {
	router *gin.Engine
}

func (s *Server) handleLogin() func(*gin.Context) {
	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var json login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "email": json.Email})
	}
}

func (s *Server) handleSignup() func(*gin.Context) {
	type signup struct {
		User     string `json:"user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var json signup
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func (s *Server) Run() error {
	return s.router.Run()
}

func Default() *Server {
	s := &Server{}
	s.router = gin.Default()

	s.router.Use(gzip.Gzip(gzip.DefaultCompression))
	s.router.Use(static.Serve("/", static.LocalFile("./client/public", true)))
	s.router.POST("/auth/signup", s.handleSignup())
	s.router.POST("/auth/login", s.handleLogin())
	return s
}
