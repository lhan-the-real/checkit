package server

import (
	"checkit/mark"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) getCheckMarks(c *gin.Context) {
	mycheckMarks, err := s.controller.ListMarks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error)
	}
	c.IndentedJSON(http.StatusOK, mycheckMarks)
}

func (s Server) createMark(c *gin.Context) {
	var requestMark mark.CheckMark
	if err := c.ShouldBindJSON(&requestMark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.controller.CreateMark(requestMark)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error)
	}
	c.IndentedJSON(http.StatusOK, "")
}
