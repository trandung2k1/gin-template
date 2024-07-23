package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)
	// r.GET("/", s.HelloWorldHandler)
	// r.GET("/health", s.healthHandler)
	getRoutes(r)
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func getRoutes(r *gin.Engine) {
	api := r.Group("/api")
	AuthRoutes(api)
}
