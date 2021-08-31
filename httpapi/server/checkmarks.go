package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
