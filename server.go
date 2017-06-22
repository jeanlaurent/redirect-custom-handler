package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Server servers the redirect
type Server struct {
	port          int
	customHandler string
}

// Start the server on the specified port
func (server *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/{path:.*}", handleError(server.redirect)).Methods(http.MethodGet)

	fmt.Println("Server rerouting to custom handler", server.customHandler)
	fmt.Println("listening on", server.port)

	return http.ListenAndServe(fmt.Sprintf(":%d", server.port), router)
}

func (server *Server) redirect(response http.ResponseWriter, request *http.Request) error {
	url := request.URL.RequestURI()
	response.Header().Set("Location", fmt.Sprintf("%s://%s", server.customHandler, url))
	response.WriteHeader(http.StatusSeeOther)
	return nil
}

func handleError(handler func(response http.ResponseWriter, request *http.Request) error) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if err := handler(response, request); err != nil {
			http.Error(response, err.Error(), 500)
		}
	}
}
