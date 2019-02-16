package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/google/uuid"
	"gopkg.in/macaron.v1"
	"net/http"

	oauth2Models "gopkg.in/oauth2.v3/models"
)

func (s *Server) userAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error){
	storeP, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid, ok := storeP.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		storeP.Set("ReturnUri", r.Form)
		_ = storeP.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	storeP.Delete("LoggedInUserID")
	_ = storeP.Save()
	return
}

func (s *Server) GetCredentials(c *macaron.Context) {
	newClientId := uuid.New().String()[:8]
	newClientSecret := uuid.New().String()[:8]
	err := s.ClientStore.Set(newClientId, &oauth2Models.Client{
		ID: newClientId,
		Secret: newClientSecret,
		Domain: "http://localhost",
	})
	if err != nil {

	}
	c.JSON(200, gin.H{"client_id": newClientId, "client_secret": newClientSecret})
}

func (s *Server) GetToken(c *macaron.Context) {
	err := s.Oauth2.HandleTokenRequest(c.Resp, c.Req.Request)
	if err != nil {
		fmt.Println(err)
	}
}
