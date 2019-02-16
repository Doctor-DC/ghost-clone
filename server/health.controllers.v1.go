package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/macaron.v1"
	"net/http"
)

func (s *Server) HealthEndpoint(c *macaron.Context) {
	fmt.Println(s.Cache.Get(c.Req.Request.RequestURI))
	c.JSON(http.StatusOK, gin.H{"msg": "Hi, I'm OK. No Worry :)"})
}
