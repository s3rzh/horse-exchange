package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/spf13/viper"
)

type getTasksResponse struct {
	Data []entity.Task `json:"data"`
}

func (h *Handler) getTasks(c *gin.Context) {
	tasks, err := h.service.Task.GetTasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, viper.GetString("messages.errors.retry_later"))
		return
	}

	c.JSON(http.StatusOK, getTasksResponse{
		Data: tasks,
	})
}
