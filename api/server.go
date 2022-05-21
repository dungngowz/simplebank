package api

import (
	"fmt"

	db "github.com/dungngowz/simple_bank/db/sqlc"
	"github.com/dungngowz/simple_bank/token"
	"github.com/dungngowz/simple_bank/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests and responses for our bamking service.
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	// Add routes to the router
	router.POST("/login-user", server.loginUser)
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
