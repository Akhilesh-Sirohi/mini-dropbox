package router

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mini-dropbox/app/config"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var commonMiddlewareForApiGroup []mux.MiddlewareFunc

func Initialize(config *config.Config) {
	router := NewServiceRouter(config.App.GitCommitHash)

	serveRoutes(router, config)
}

// initalize the service route for all the groups
func NewServiceRouter(gitCommitHash string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/commit.txt", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, gitCommitHash) // nosemgrep go.lang.security.audit.xss.no-fprintf-to-responsewriter.no-fprintf-to-responsewritera
	})

	for _, routeGroup := range routeList {
		registerRouteGroup(router, routeGroup)
	}
	return router
}

func registerRouteGroup(router *mux.Router, routeGroup route) {
	var middlewareList []mux.MiddlewareFunc

	if len(routeGroup.middleware) != 0 {
		middlewareList = routeGroup.middleware
	}

	route := router.PathPrefix(routeGroup.group).Subrouter()
	if routeGroup.group == "/api" {
		route.Use(commonMiddlewareForApiGroup...)
	}

	route.Use(middlewareList...)

	registerRoutes(route, routeGroup.endpoints)
}

func registerRoutes(router *mux.Router, endpoints []endpoint) {
	for _, endpoint := range endpoints {
		router.HandleFunc(endpoint.path, endpoint.handler).Methods(endpoint.method, http.MethodOptions)
	}
}

// serveRoutes will serves the both the service Route and Prometheus route
func serveRoutes(router *mux.Router, config *config.Config) {
	srv := &http.Server{
		Addr:    getListenAddress(&config.App),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Info("Shutting down Server ...")

	time.Sleep(time.Duration(config.App.ShutdownDelay) * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.App.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown with error : ", err)
	}

	log.Info("Server exiting")
}

func getListenAddress(app *config.App) string {
	return net.JoinHostPort(app.Hostname, app.Port)
}
