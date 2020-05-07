package repository_test

import (
	"testing"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	repository "github.com/andrewjapar/andrew-cv-app/app/podcast/repository"
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	db, mock, err := sqlmock.New()
	gormDb, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("stub db connection error: '%s'", err)
	}

	mockPodcasts := []domain.Podcast{
		domain.Podcast{
			ID: 1, Title: "First Podcast by Andrew", Description: "NoDesc",
			AudioURL: "https://test.com", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		domain.Podcast{
			ID: 2, Title: "Second Podcast by Andrew", Description: "NoDesc",
			AudioURL: "https://test.com", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "audio_url", "created_at", "updated_at"}).
		AddRow(mockPodcasts[0].ID, mockPodcasts[0].Title, mockPodcasts[0].Description, mockPodcasts[0].AudioURL, mockPodcasts[0].CreatedAt, mockPodcasts[0].UpdatedAt).
		AddRow(mockPodcasts[1].ID, mockPodcasts[1].Title, mockPodcasts[1].Description, mockPodcasts[1].AudioURL, mockPodcasts[1].CreatedAt, mockPodcasts[1].UpdatedAt)

	query := "SELECT \\* FROM \"podcasts\""

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := repository.NewPodcastRepository(gormDb)
	podcasts, err := repo.Get()

	assert.NoError(t, err)
	assert.Equal(t, mockPodcasts, podcasts)
}
