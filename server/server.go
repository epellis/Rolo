package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type userModel struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

// Server stores all state and associated handlers
type Server struct {
	router *gin.Engine
	db     *gorm.DB
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

		var user userModel
		if err := s.db.First(&user, "email = ? AND password = ?", json.Email, json.Password).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "email": json.Email})
	}
}

func (s *Server) handleSignup() func(*gin.Context) {
	type signup struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var json signup
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("Username:", json.Username)
		s.db.Create(&userModel{Username: json.Username, Email: json.Email, Password: json.Password})

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func (s *Server) Run() error {
	defer s.db.Close()
	return s.router.Run()
}

func Default() (*Server, error) {
	s := &Server{}
	s.router = gin.Default()

	var err error
	s.db, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		return nil, fmt.Errorf("Gorm Open Issue: %v", err)
	}
	s.db.AutoMigrate(&userModel{})

	s.router.Use(gzip.Gzip(gzip.DefaultCompression))
	s.router.Use(static.Serve("/", static.LocalFile("./client/public", true)))
	s.router.POST("/auth/signup", s.handleSignup())
	s.router.POST("/auth/login", s.handleLogin())
	return s, nil
}
