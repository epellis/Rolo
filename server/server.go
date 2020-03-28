package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Posts    []Post
}

type Post struct {
	gorm.Model
	URL    string
	Notes  string
	UserID int
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

		var user User
		if s.db.First(&user, "email = ?", json.Email).RecordNotFound() {
			c.JSON(http.StatusOK, gin.H{"success": false})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)) != nil {
			c.JSON(http.StatusOK, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "userid": user.ID, "email": user.Email, "username": user.Username})
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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		s.db.Create(&User{Username: json.Username, Email: json.Email, Password: string(hashedPassword)})

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func (s *Server) handleNewPost() func(*gin.Context) {
	type post struct {
		URL    string `json:"url"`
		Notes  string `json:"notes"`
		UserID int    `json:"userid"`
	}
	return func(c *gin.Context) {
		var json post
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if UserID exists in database
		var user User
		if s.db.Find(&user, json.UserID).RecordNotFound() {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("Create Post Failed")})
			return
		}

		s.db.Create(&Post{URL: json.URL, Notes: json.Notes, UserID: json.UserID})

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
	s.db.AutoMigrate(&User{}, &Post{})

	s.router.Use(gzip.Gzip(gzip.DefaultCompression))

	identityKey := "id"
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(os.Getenv("SECRET_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			type login struct {
				Username string `form:"username" json:"username" binding:"required"`
				Password string `form:"password" json:"password" binding:"required"`
			}
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "user" && password == "password") {
				return &User{
					Username: "User",
					Email:    "Email",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Username == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	s.router.Use(static.Serve("/", static.LocalFile("./client/public", true)))
	s.router.POST("/auth/signup", s.handleSignup())
	// s.router.POST("/auth/login", s.handleLogin())
	s.router.POST("/auth/login", authMiddleware.LoginHandler)
	auth := s.router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	// s.router.POST("/create", s.handleNewPost())
	// s.router.NoRoute()  - TODO
	return s, nil
}
