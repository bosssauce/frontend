package frontend_test

import (
	"io/ioutil"
	"net/http"
)

// this is only needed for the go test.
func testServer(url string) {
	go http.ListenAndServe(url, nil)
}

func testGet(url, page string) string {
	resp, err := http.Get("http://" + url + page)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(text)
}
