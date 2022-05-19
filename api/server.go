package api

import (
	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests and responses for our bamking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add routes to the router
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
