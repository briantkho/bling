package server

import (
	"database/sql"
	"github.com/briantkho/bling/internal/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db:     db,
		router: gin.Default(),
	}
}

// SetupRoutes Add Routes Here/*
func (s *Server) SetupRoutes() {
	s.router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "we are so live rn") })

	userRoutes := s.router.Group("/api/user")
	{
		userHandler := user.InitUserDependencies(s.db)
		userRoutes.GET("/:user_id", userHandler.GetUserByID)
		userRoutes.POST("", userHandler.CreateUser)
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
