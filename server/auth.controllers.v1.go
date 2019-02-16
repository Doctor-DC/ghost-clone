package server

import (
	"github.com/cyantarek/ghost-clone-project/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Server) Login(c *gin.Context) {
	var loginPayload structs.LoginPayload

	_ = c.BindJSON(&loginPayload)

	// Database Logic
	if loginPayload.Username == "dev" && loginPayload.Password == "dev" {
		claims := &jwt.MapClaims{"exp": time.Now().Add(time.Minute*1).Unix(), "username": loginPayload.Username}

		tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		tokString, err := tok.SignedString(s.SignKey)
		if err != nil {
			c.JSON(401, APIErrors{"ErrTokenValidation", ErrTokenValidation})
			return
		}
		c.JSON(200, gin.H{"token": tokString})
	} else {
		c.JSON(500, APIErrors{"ErrLogin", ErrLogin})
	}
}

func (s *Server) Signup(c *gin.Context) {

}