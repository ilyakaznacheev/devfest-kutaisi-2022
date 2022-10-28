package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ilyakaznacheev/devfest-kutaisi-2022/internal/model"
)

func (s *Server) AddWine(c *gin.Context) {
	var w model.Wine

	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.NewString()

	if err := s.repo.AddWine(c.Request.Context(), id, w); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusCreated, id)
}

func (s *Server) GetWine(c *gin.Context) {
	id := c.Param("id")

	w, err := s.repo.GetWine(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, w)
}

func (s *Server) ListWine(c *gin.Context) {
	w, err := s.repo.GetWineList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, w)
}
