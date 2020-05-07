package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	podcastHandler "github.com/andrewjapar/andrew-cv-app/app/podcast/handler"
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/andrewjapar/andrew-cv-app/domain/mocks"
)

func TestGet(t *testing.T) {
	// Given
	var mockPodcast domain.Podcast
	err := faker.FakeData(&mockPodcast)
	assert.NoError(t, err)

	mockPodcastList := make([]domain.Podcast, 0)
	mockPodcastList = append(mockPodcastList, mockPodcast)

	mockRepo := new(mocks.PodcastRepository)
	mockRepo.On("Get").Return(mockPodcastList, nil)

	// When
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/podcasts", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := podcastHandler.PodcastHandler{
		Repository: mockRepo,
	}
	err = handler.Get(c)
	require.NoError(t, err)

	// Then
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestGet_GivenError_ShouldReturn500(t *testing.T) {
	// Given
	var mockPodcast []domain.Podcast

	mockRepo := new(mocks.PodcastRepository)
	mockRepo.On("Get").Return(mockPodcast, errors.New("Unknown Error"))

	// When
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/podcasts", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := podcastHandler.PodcastHandler{
		Repository: mockRepo,
	}
	err = handler.Get(c)
	require.NoError(t, err)

	// Then
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
