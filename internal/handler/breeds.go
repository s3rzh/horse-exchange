package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/spf13/viper"
)

type getBreedsResponse struct {
	Data []entity.Breed `json:"data"`
}

func (h *Handler) getBreeds(c *gin.Context) {
	breads, err := h.service.Breed.GetBreeds()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, viper.GetString("messages.errors.retry_later"))
		return
	}

	c.JSON(http.StatusOK, getBreedsResponse{
		Data: breads,
	})
}
