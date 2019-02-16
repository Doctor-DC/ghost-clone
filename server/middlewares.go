package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
)

func (s *Server) ValidateJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader != "" {
			tok := strings.Split(authHeader, " ")

			if len(tok) == 2 {
				token, err := jwt.Parse(tok[1], func(token *jwt.Token) (interface{}, error) {
					return s.VerifyKey, nil
				})
				if err != nil {
					c.JSON(401, map[string]string{"message": err.Error()})
					c.Abort()
					return
				}
				if token.Valid {
					c.Next()
				} else {
					c.JSON(401, map[string]string{"message": "Invalid auth token"})
					c.Abort()
					return
				}
			}
		} else {
			c.JSON(401, map[string]string{"message": "Auth token required"})
			c.Abort()
		}
	}
}

func (s *Server) ValidateOauthToken(c *macaron.Context) {
	if viper.GetBool("server.oauth2") {
		_, err := s.Oauth2.ValidationBearerToken(c.Req.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, APIErrors{"ErrInvalidAccessToken", ErrInvalidAccessToken})
			return
		}
		c.Next()
	} else {
		c.Next()
	}
}

// deprecated
func (s *Server) CachedCheck(c *macaron.Context) {
	content, _ := s.Cache.Get(c.Req.RequestURI)
	if content == nil {
		fmt.Println("cache hit")
		c.JSON(200, string(content))
		return
	} else {
		fmt.Println("cache miss")
		c.Next()
	}
}
