package running

//func TestCreatePostEmptyFields(t *testing.T) {
//	payload := `{
//	"title": "",
//	"slug": "aaa",
//	"markdown": "aaa",
//	"html": "aaa",
//	"is_featured": false,
//	"is_page": false,
//	"is_published": false,
//	"image": "aaaa",
//	"meta_description": "meta",
//	"tags": ["aaa", "bbb", "cccc"],
//	"author_id": 1
//}`
//
//	token := loginReq()
//	req, _ := http.NewRequest("POST", httpSrv.URL+"/posts", bytes.NewBuffer([]byte(payload)))
//	req.Header.Set("Authorization", "Bearer " + token)
//	req.Header.Set("Content-Type", "application/json")
//
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode == 200 {
//		t.Fatal(res.StatusCode)
//	}
//	defer res.Body.Close()
//}

//func TestCreatePostNew(t *testing.T) {
//	payload := `{
//	"title": "DDD",
//	"slug": "aaa",
//	"markdown": "aaa",
//	"html": "aaa",
//	"is_featured": false,
//	"is_page": false,
//	"is_published": false,
//	"image": "aaaa",
//	"meta_description": "meta",
//	"tags": ["aaa", "bbb", "cccc"],
//	"author_id": 1
//}`
//	token := loginReq()
//	req, _ := http.NewRequest("POST", httpSrv.URL+"/posts", bytes.NewBuffer([]byte(payload)))
//	req.Header.Set("Authorization", "Bearer " + token)
//	req.Header.Set("Content-Type", "application/json")
//
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode != 200 {
//		t.Fatal(res.StatusCode)
//	}
//	defer res.Body.Close()
//}

//func TestCreatePostExists(t *testing.T) {
//	payload := `{
//	"title": "DDD",
//	"slug": "aaa",
//	"markdown": "aaa",
//	"html": "aaa",
//	"is_featured": false,
//	"is_page": false,
//	"is_published": false,
//	"image": "aaaa",
//	"meta_description": "meta",
//	"tags": ["aaa", "bbb", "cccc"],
//	"author_id": 1
//}`
//	token := loginReq()
//	req, _ := http.NewRequest("POST", httpSrv.URL+"/posts", bytes.NewBuffer([]byte(payload)))
//	req.Header.Set("Authorization", "Bearer " + token)
//	req.Header.Set("Content-Type", "application/json")
//
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode == 200 {
//		t.Fatal(res.StatusCode)
//	}
//	defer res.Body.Close()
//}

//func TestUpdatePost(t *testing.T) {
//	payload := `{
//	"title": "EEE",
//	"slug": "aaa",
//	"markdown": "aaa",
//	"html": "aaa",
//	"is_featured": false,
//	"is_page": false,
//	"is_published": false,
//	"image": "aaaa",
//	"meta_description": "meta",
//	"tags": ["aad", "bbb", "cccc"],
//	"author_id": 1
//}`
//
//	token := loginReq()
//	req, _ := http.NewRequest("PUT", httpSrv.URL+"/posts/aaa", bytes.NewBuffer([]byte(payload)))
//	req.Header.Set("Authorization", "Bearer " + token)
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode != 200 {
//		t.Fatal(res.StatusCode)
//		return
//	}
//
//	defer res.Body.Close()
//}

//func TestDeletePostNonExistent(t *testing.T) {
//	token := loginReq()
//
//	req, _ := http.NewRequest("DELETE", httpSrv.URL+"/posts/aaad", nil)
//	req.Header.Set("Authorization", "Bearer " + token)
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode != 404 {
//		t.Fatal(res.StatusCode)
//	}
//
//	defer res.Body.Close()
//}
//
//func TestDeletePostExistent(t *testing.T) {
//	token := loginReq()
//
//	req, _ := http.NewRequest("DELETE", httpSrv.URL+"/posts/aaa", nil)
//	req.Header.Set("Authorization", "Bearer " + token)
//	res, _ := http.DefaultClient.Do(req)
//	if res.StatusCode != 204 {
//		t.Fatal(res.StatusCode)
//	}
//
//	defer res.Body.Close()
//	defer httpSrv.Close()
//}
