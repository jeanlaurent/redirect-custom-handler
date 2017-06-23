package main

import (
	"fmt"
	"net/http"
)

// Server servers the redirect
type Server struct {
	customHandler string
}

// Start the server on the specified port
func (server *Server) Start() error {

	http.HandleFunc("/", server.redirect)

	fmt.Println("Server rerouting to custom handler", server.customHandler)

	return http.ListenAndServe(":8080", nil)
}

func (server *Server) redirect(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, fmt.Sprintf("%s://%s", server.customHandler, request.URL.RequestURI()), http.StatusSeeOther)
}
