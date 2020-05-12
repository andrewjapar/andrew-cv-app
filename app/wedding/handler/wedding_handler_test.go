package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	WeddingHandler "github.com/andrewjapar/andrew-cv-app/app/wedding/handler"
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/andrewjapar/andrew-cv-app/domain/mocks"
)

func TestGetByUserID(t *testing.T) {
	// Given
	var mockWedding domain.Wedding
	err := faker.FakeData(&mockWedding)
	assert.NoError(t, err)

	mockWeddingList := make([]domain.Wedding, 0)
	mockWeddingList = append(mockWeddingList, mockWedding)
	userID := 2

	mockRepo := new(mocks.WeddingRepository)
	mockRepo.On("GetByUserID", int64(userID)).Return(mockWeddingList, nil)

	// When
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/weddings/"+strconv.Itoa(userID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("weddings/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(userID))
	handler := WeddingHandler.WeddingHandler{
		Repository: mockRepo,
	}
	err = handler.GetByUserID(c)
	require.NoError(t, err)

	// Then
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestGet_GivenError_ShouldReturn500(t *testing.T) {
	// Given
	var mockWeddingList []domain.Wedding
	userID := 2

	mockRepo := new(mocks.WeddingRepository)
	mockRepo.On("GetByUserID", int64(userID)).Return(mockWeddingList, errors.New("Unknown Error"))

	// When
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/weddings/"+strconv.Itoa(userID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("weddings/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(userID))
	handler := WeddingHandler.WeddingHandler{
		Repository: mockRepo,
	}
	err = handler.GetByUserID(c)
	require.NoError(t, err)

	// Then
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
