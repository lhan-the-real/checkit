package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckCategory string

const (
	GYM    CheckCategory = "Gym"
	CODING CheckCategory = "Coding"
)

type CheckMark struct {
	ID       string        `json:"id"`
	UserID   string        `json:"user_id"`
	Time     string        `json:"time"`
	Content  string        `json:"content"`
	Category CheckCategory `json:"category"`
}

type Server struct {
	router *gin.Engine
}

func NewServer() Server {
	router := gin.Default()
	srv := Server{
		router: router,
	}
	srv.router.GET("/checkmarks", srv.getCheckMarks)
	return srv
}

func (s Server) Start(address string) error {
	return s.router.Run(address)
}

func (s Server) getCheckMarks(c *gin.Context) {
	mycheckmarks := []CheckMark{
		{
			ID:       "1",
			UserID:   "2",
			Time:     time.Now().GoString(),
			Content:  "Run",
			Category: GYM,
		},
		{
			ID:       "2",
			UserID:   "2",
			Time:     time.Now().GoString(),
			Content:  "Run",
			Category: GYM,
		},
	}
	c.IndentedJSON(http.StatusOK, mycheckmarks)
}
