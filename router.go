package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// export the Router so any content type file can just call frontend.Router.HandleFunc(....
	// otherwise, a New() func might be needed, but it is not clear where it should be called from
	Router *mux.Router
)

func init() {
	Router = mux.NewRouter()

	http.Handle("/", Router)
}
