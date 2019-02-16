package deprecated

import (
	"github.com/cyantarek/ghost-clone-project/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpSrv *httptest.Server

func TestHealthRoute(t *testing.T) {
	res, err := http.Get(httpSrv.URL+"/health")
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(res.StatusCode)
	}
	//defer httpSrv.Close()
}

func init() {
	router := server.GetRouter()
	httpSrv = httptest.NewServer(router)
}
