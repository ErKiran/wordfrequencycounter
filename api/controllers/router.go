package controllers

import (
	"log"
	"net/http"

	"wordcount/api/middlewares"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.Use(middlewares.CORS)
	server.Router.HandleFunc(path, middlewares.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) InitRoutes() {
	server.Router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080"+"/docs/swagger.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	server.Router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	server.setJSON("/api/v1/health", server.Health, "GET")
	server.setJSON("/api/v1/wordcount", server.WordCountText, "POST")
	server.setJSON("/api/v1/wordcount/file", server.WordCountFile, "POST")

}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
