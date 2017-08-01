package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	mux *mux.Router
}

var (
	// Router is exported so any content type file can just call frontend.Router.HandleFunc(....
	// otherwise, a New() func might be needed, but it is not clear where it should be called from
	Router *router
)

func (r *router) HandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(path, fn)
}

func (r *router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(res, req)
}

func init() {
	Router = &router{
		mux: mux.NewRouter(),
	}

	http.Handle("/", Router)
}
