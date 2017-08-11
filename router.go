/*
Package frontend is a ponzu addon to handle page routing for your ponzu site.

Usage

See example.
*/
package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	*mux.Router
}

var (
	// Router is exported so any content type file can just call frontend.Router.HandleFunc(....
	// otherwise, a New() func might be needed, but it is not clear where it should be called from
	Router = &router{
		Router: mux.NewRouter(),
	}
)

func init() {
	http.Handle("/", Router)
}

// Vars returns the route variables for the current request, if any.
func Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

// RegisterGet registers GET handlers.
func RegisterGet(toHandle ...ToHandler) {
	for _, hand := range toHandle {
		Router.HandleFunc(hand.ToHandler()).Methods("GET")
	}
}

// RegisterPost registers POST handlers.
func RegisterPost(toHandle ...ToHandler) {
	for _, hand := range toHandle {
		Router.HandleFunc(hand.ToHandler()).Methods("POST")
	}
}

// ToHandler defines a registrable handler with routing path.
type ToHandler interface {
	ToHandler() (url string, handler http.HandlerFunc)
}

// ToHandle defines a route to register with an URL and Handler func.
type ToHandle struct {
	URL     string
	Handler http.HandlerFunc
}

// ToHandler implements ToHandler.
func (th ToHandle) ToHandler() (string, http.HandlerFunc) {
	return th.URL, th.Handler
}

// ListHandle defines a list of ToHandle.
type ListHandle []ToHandle

// ToHandlers converts the list to a list of interface ToHandler.
func (l ListHandle) ToHandlers() (handlers []ToHandler) {
	for _, handle := range l {
		handlers = append(handlers, handle)
	}
	return
}
