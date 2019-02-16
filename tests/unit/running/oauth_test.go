package running

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type credentialsResp struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type tokenResp struct {
	AccessToken string `json:"access_token"`
}

var credentialsRespD credentialsResp
var tokenRespD tokenResp

func TestOAuth(t *testing.T) {
	t.Run("get credentials", func(r *testing.T) {
		res, _ := http.Get(url + "/oauth/credentials")
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		_ = json.NewDecoder(res.Body).Decode(&credentialsRespD)
		if credentialsRespD.ClientId == "" {
			t.Fatal()
		}
		defer res.Body.Close()
	})

	t.Run("get token", func(r *testing.T) {
		res, _ := http.Get(url + fmt.Sprintf("/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=all", credentialsRespD.ClientId, credentialsRespD.ClientSecret))
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}

		_ = json.NewDecoder(res.Body).Decode(&tokenRespD)
		if tokenRespD.AccessToken == "" {
			t.Fatal()
		}
		defer res.Body.Close()
	})

	t.Run("use token", func(r *testing.T) {
		res, _ := http.Get(url + fmt.Sprintf("/posts?access_token=%s",tokenRespD.AccessToken))
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})
}
