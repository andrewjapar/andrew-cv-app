package handler

import (
	"net/http"

	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type PodcastHandler struct {
	Repository domain.PodcastRepository
}

func NewPodcastHandler(e *echo.Echo, repo domain.PodcastRepository) {
	handler := &PodcastHandler{
		Repository: repo,
	}
	e.GET("/podcasts", handler.Get)
}

func (handler *PodcastHandler) Get(c echo.Context) error {

	listAr, err := handler.Repository.Get()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			ResponseError{
				Message: err.Error(),
				Code:    500,
			},
		)
	}

	return c.JSON(http.StatusOK, listAr)
}
