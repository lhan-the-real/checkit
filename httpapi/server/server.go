package server

import (
	"checkit/mark"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	controller mark.Controller
}

func NewServer() Server {
	router := gin.Default()
	srv := Server{
		router:     router,
		controller: mark.NewController(),
	}
	v1 := srv.router.Group("/v1")
	{
		v1.GET("/checkMarks", srv.getCheckMarks)
		v1.POST("/checkMarks", srv.createMark)
	}
	return srv
}

func (s Server) Start(address string) error {
	return s.router.Run(address)
}
