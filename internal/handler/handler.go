package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/s3rzh/horse-exchange/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		horses := api.Group("/horses")
		{
			horses.GET("/", h.getAllHorses)
			horses.PATCH("/:id/buy", h.buyHorseByID)
			horses.POST("/rent", h.rentHorseByIDWithPeriod)
			horses.PATCH("/rent/:id/pay", h.payHorseByRentID)
		}

		catalog := api.Group("/catalog")
		{
			breeds := catalog.Group("/breeds")
			{
				breeds.GET("/", h.getBreeds)
				breeds.GET("/:id", h.getHorsesByBreedID)
			}

			tasks := catalog.Group("/tasks")
			{
				tasks.GET("/", h.getTasks)
				tasks.GET("/:id", h.getHorsesByTaskID)
			}
		}
	}

	return router
}
