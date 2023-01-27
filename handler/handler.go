package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cache map[string]int
}

type album struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
}

func (h *Handler) GetHot100(c *gin.Context) {
	date := c.Param("date")
	c.JSON(http.StatusOK, date)
}

func (h *Handler) PostAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
