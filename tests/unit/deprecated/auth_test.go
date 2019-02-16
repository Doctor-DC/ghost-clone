package deprecated

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSuccessLogin(t *testing.T) {
	payload := `{
	"username": "dev",
	"password": "dev"
}`
	res, _ := http.Post(httpSrv.URL+"/auth/login", "application/json", bytes.NewBuffer([]byte(payload)))
	if res.StatusCode != http.StatusOK {
		t.Fatal(res.StatusCode)
	}

	defer res.Body.Close()
}

func TestFailedLogin(t *testing.T) {
	payload := `{
	"username": "dev",
	"password": "de"
}`
	res, _ := http.Post(httpSrv.URL+"/auth/login", "application/json", bytes.NewBuffer([]byte(payload)))
	if res.StatusCode != 500 {
		t.Fatal(res.StatusCode)
	}

	defer res.Body.Close()
}

//func TestUnAuthorized(t *testing.T) {
//	res, _ := http.Post(httpSrv.URL+"/posts", "", nil)
//	if res.StatusCode != http.StatusUnauthorized {
//		t.Fatal(res.StatusCode)
//	}
//
//	defer res.Body.Close()
//}

func loginReq() string {
	payload := `{
	"username": "dev",
	"password": "dev"
}`
	res, _ := http.Post(httpSrv.URL+"/auth/login", "application/json", bytes.NewBuffer([]byte(payload)))

	defer res.Body.Close()

	login := struct {
		Token string `json:"token"`
	}{}

	json.NewDecoder(res.Body).Decode(&login)
	return login.Token
}

func TestAuthorized(t *testing.T) {
	token := loginReq()
	req, _ := http.NewRequest("GET", httpSrv.URL+"/posts", nil)
	req.Header.Set("Authorization", "Bearer " + token)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		t.Fatal(res.StatusCode)
	}
}

