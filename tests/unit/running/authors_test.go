package running

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	res, _ := http.Get(url + "/authors?access_token=" + tokenRespD.AccessToken)
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	defer res.Body.Close()
}

func TestGetNonExistentUser(t *testing.T) {
	t.Run("get user by slug failed test", func(r *testing.T) {
		res, _ := http.Get(url + "/authors/slug/asasas?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})

	t.Run("get user by id failed test", func(r *testing.T) {
		res, _ := http.Get(url + "/authors/111?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})
}

func TestGetExistentUser(t *testing.T) {
	t.Run("get user by slug", func(r *testing.T) {
		res, _ := http.Get(url + "/authors/slug/john?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})

	t.Run("get user by id", func(r *testing.T) {
		res, _ := http.Get(url + "/authors/5c5eb87c3443be02f7e11ed5?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})
}
