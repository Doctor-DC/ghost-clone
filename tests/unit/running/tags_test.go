package running

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllTags(t *testing.T) {
	res, _ := http.Get(url + "/tags?access_token=" + tokenRespD.AccessToken)
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	defer res.Body.Close()
}

func TestGetNonExistentTag(t *testing.T) {
	t.Run("get tag by slug failed test", func(r *testing.T) {
		res, _ := http.Get(url + "/tags/slug/111?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})

	t.Run("get tag by id failed test", func(r *testing.T) {
		res, _ := http.Get(url + "/tags/111?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})
}

func TestGetExistentTag(t *testing.T) {
	t.Run("get tag by slug", func(r *testing.T) {
		res, _ := http.Get(url + "/tags/slug/getting-started?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})

	t.Run("get tag by id", func(r *testing.T) {
		res, _ := http.Get(url + "/tags/5c5ea52c3443be02f7e11bac?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})
}
