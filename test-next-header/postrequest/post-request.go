package postrequest

import (
	"log"
	"net/http"
)

func DoPostRequest(cookieHeader string) *http.Response {
	url := "http://127.0.0.1:3000/api/post"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("cookie", cookieHeader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
