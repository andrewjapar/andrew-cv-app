package repository_test

import (
	"testing"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	repository "github.com/andrewjapar/andrew-cv-app/app/wedding/repository"
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetByUserID(t *testing.T) {

	db, mock, err := sqlmock.New()
	gormDb, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("stub db connection error: '%s'", err)
	}

	mockWeddings := []domain.Wedding{
		domain.Wedding{
			ID: 1, Title: "Wedding Andrew", Description: "NoDesc",
			WeddingDate: time.Now(), Code: "test", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		domain.Wedding{
			ID: 2, Title: "Wedding Ivana", Description: "NoDesc",
			WeddingDate: time.Now(), Code: "test", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "wedding_date", "code", "created_at", "updated_at"}).
		AddRow(mockWeddings[0].ID, mockWeddings[0].Title, mockWeddings[0].Description, mockWeddings[0].WeddingDate, mockWeddings[0].Code, mockWeddings[0].CreatedAt, mockWeddings[0].UpdatedAt).
		AddRow(mockWeddings[1].ID, mockWeddings[1].Title, mockWeddings[1].Description, mockWeddings[1].WeddingDate, mockWeddings[1].Code, mockWeddings[1].CreatedAt, mockWeddings[1].UpdatedAt)

	query := "SELECT \"weddings\"\\.\\* FROM \"weddings\" left join wedding_organizers on wedding_organizers\\.wedding_id = weddings\\.id left join users on wedding_organizers\\.user_id = users\\.id WHERE \\(users\\.id = \\$1\\)"

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := repository.NewWeddingRepository(gormDb)
	Weddings, err := repo.GetByUserID(2)

	assert.NoError(t, err)
	assert.Equal(t, mockWeddings, Weddings)
}
