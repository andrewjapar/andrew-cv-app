package handler

import (
	"net/http"
	"strconv"

	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type WeddingHandler struct {
	Repository domain.WeddingRepository
}

func NewWeddingHandler(e *echo.Echo, repo domain.WeddingRepository) {
	handler := &WeddingHandler{
		Repository: repo,
	}
	e.GET("/weddings/:id", handler.GetByUserID)
}

func (handler *WeddingHandler) GetByUserID(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			ResponseError{
				Message: err.Error(),
				Code:    500,
			},
		)
	}

	listAr, err := handler.Repository.GetByUserID(int64(userID))
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
