package api

import (
	db "TailorShop/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/customers", server.createCustomer)
	router.GET("/customers/:id", server.getCustomer)
	router.GET("/customers", server.listCustomer)
	server.router = router
	return server
}
