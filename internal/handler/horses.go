package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/spf13/viper"
)

type getAllHorsesResponse struct {
	Data []entity.Horse `json:"data"`
}

func (h *Handler) getAllHorses(c *gin.Context) {
	horses, err := h.service.Horse.GetAllHorses()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, viper.GetString("messages.errors.retry_later"))
		return
	}

	c.JSON(http.StatusOK, getAllHorsesResponse{
		Data: horses,
	})
}

func (h *Handler) buyHorseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.invalid_id"))
		return
	}

	rows, err := h.service.Horse.BuyHorseByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, viper.GetString("messages.errors.retry_later"))
		return
	}

	if rows == 0 {
		newErrorResponse(c, http.StatusNotFound, viper.GetString("messages.errors.rented_or_sold"))
		return
	}

	newSuccessResponse(c, http.StatusOK, viper.GetString("messages.responses.successful_buy"))
}

func (h *Handler) getHorsesByBreedID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.invalid_id"))
		return
	}

	horses, err := h.service.Horse.GetHorsesByBreedID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, viper.GetString("messages.errors.retry_later"))
		return
	}

	if len(horses) == 0 {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.not_found"))
		return
	}

	c.JSON(http.StatusOK, getAllHorsesResponse{
		Data: horses,
	})

}

func (h *Handler) getHorsesByTaskID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.invalid_id"))
		return
	}

	horses, err := h.service.Horse.GetHorsesByTaskID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, viper.GetString("messages.errors.retry_later"))
		return
	}

	if len(horses) == 0 {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.not_found"))
		return
	}

	c.JSON(http.StatusOK, getAllHorsesResponse{
		Data: horses,
	})
}

func (h *Handler) rentHorseByIDWithPeriod(c *gin.Context) {
	var input entity.Rent
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.invalid_format"))
		return
	}

	var dt string = "2006-01-02 15:04:05"
	nowTime := time.Now()

	sinceTime, err := time.Parse(dt, input.Since)
	if err != nil || nowTime.After(sinceTime) {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.wrong_since_time_format"))
		return
	}

	tillTime, err := time.Parse(dt, input.Till)
	if err != nil || nowTime.After(tillTime) || sinceTime.After(tillTime) {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.wrong_till_time_format"))
		return
	}

	id, err := h.service.Horse.RentHorseByIDWithPeriod(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, viper.GetString("messages.errors.retry_later"))
		return
	}

	if id == 0 {
		newErrorResponse(c, http.StatusOK, viper.GetString("messages.errors.busy_period"))
		return
	}

	newSuccessResponse(c, http.StatusOK, fmt.Sprintf(viper.GetString("messages.responses.buy_url"), id))
}

func (h *Handler) payHorseByRentID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("messages.errors.invalid_id"))
		return
	}

	err = h.service.Horse.PayHorseByRentID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, http.StatusOK, viper.GetString("messages.responses.successful_pay"))
}
