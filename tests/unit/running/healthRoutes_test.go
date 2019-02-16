package running

import (
	"encoding/json"
	"fmt"
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/cyantarek/ghost-clone-project/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpSrv *httptest.Server

var url = ""

func TestHealthRoute(t *testing.T) {
	res, err := http.Get(httpSrv.URL + "/health")
	if err != nil {
		t.Fatal(err, res.StatusCode)
	} else {
		fmt.Println(res.StatusCode)
	}
	//defer httpSrv.Close()
}

func init() {
	router, _ := server.GetRouter()
	httpSrv = httptest.NewServer(router)
	url = httpSrv.URL + "/api/"+ config.VERSION
	setupOAuthTest()
}

func setupOAuthTest() {
	res, _ := http.Get(url + fmt.Sprintf("/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=all", "default", "default"))
	if res.StatusCode != http.StatusOK {
		log.Fatal()
	}

	_ = json.NewDecoder(res.Body).Decode(&tokenRespD)
	if tokenRespD.AccessToken == "" {
		log.Fatal()
	}
	defer res.Body.Close()
}
