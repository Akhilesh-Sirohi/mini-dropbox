package router

import (
	"github.com/mini-dropbox/app/router/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

type route struct {
	group      string
	middleware []mux.MiddlewareFunc
	endpoints  []endpoint
}

type endpoint struct {
	method  string
	path    string
	handler handlerFunc
}

var routeList = []route{
	{
		group: "/",
		endpoints: []endpoint{
			{http.MethodGet, "/status", controllers.App.Status},
		},
	},
}
