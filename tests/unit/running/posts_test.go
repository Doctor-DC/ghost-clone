package running

import (
	"encoding/json"
	"fmt"
	"github.com/cyantarek/ghost-clone-project/models"
	"net/http"
	"testing"
)

func TestGetAllPosts(t *testing.T) {
	res, _ := http.Get(url + "/posts?access_token=" + tokenRespD.AccessToken)
	if res.StatusCode != http.StatusOK {
		t.Fatal()
	}

	var data struct {
		Posts           []models.Post `json:"posts"`
		Meta models.PostMeta `json:"meta"`
	}
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		t.Fatalf("error sending request: %v", err)
	}
	if len(data.Posts) == 0 {
		t.Fatalf("empty data returned")
	}

	if data.Meta.Total != len(data.Posts) {
		t.Fatalf("meta total is not correct")
	}

	defer res.Body.Close()
}

func TestGetNonExistentPost(t *testing.T) {
	t.Run("get post by slug", func(r *testing.T) {
		res, _ := http.Get(url + "/posts/slug/111?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})

	t.Run("get post by id", func(r *testing.T) {
		res, _ := http.Get(url + "/posts/111?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode == http.StatusOK {
			r.Fatal(res.StatusCode)
		} else {
			fmt.Println(res.StatusCode)
		}
		defer res.Body.Close()
	})
}

func TestGetExistentPost(t *testing.T) {
	t.Run("get post by slug", func(r *testing.T) {
		res, _ := http.Get(url + "/posts/slug/welcome-to-ghost?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})

	t.Run("get post by id", func(r *testing.T) {
		res, _ := http.Get(url + "/posts/5c5e8db93443be02f7e1173f?access_token=" + tokenRespD.AccessToken)
		if res.StatusCode != http.StatusOK {
			t.Fatal()
		}
		defer res.Body.Close()
	})
}
