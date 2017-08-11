package frontend_test

import (
	"fmt"
	"github.com/bosssauce/frontend"
	"net/http"
)

func viewHomePage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("HomePage"))
}

func viewOther(res http.ResponseWriter, req *http.Request) {
	vars := frontend.Vars(req)
	res.Write([]byte("this is page " + vars["keyName"]))
}

func postOther(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("PostOther"))
}

func Example_register() { // Use in init()

	frontend.RegisterGet(frontend.ListHandle{{
		URL:     "/",
		Handler: viewHomePage,
	}, {
		URL:     "/other/{keyName}",
		Handler: viewOther,
	}}.ToHandlers()...)

	frontend.RegisterPost(frontend.ListHandle{{
		URL:     "/other",
		Handler: postOther,
	}}.ToHandlers()...)

	runTest() // Simple tests.
	// Output:
	// HomePage
	// this is page test
	// this is page another
}

func runTest() {
	url := "localhost:7357"
	testServer(url)
	fmt.Println(testGet(url, "/"))
	fmt.Println(testGet(url, "/other/test"))
	fmt.Println(testGet(url, "/other/another"))
}
