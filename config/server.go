package config

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func NewServer(host string, port string, router *mux.Router) (*Server, error) { 
	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://" + host + ":" + port,
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE",
		},
	})
	handler := corsOptions.Handler(router)

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("🚀 Server running on port http://localhost%s\n", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
