package repository

import (
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/jinzhu/gorm"
)

type PodcastRepository struct {
	DB *gorm.DB
}

func NewPodcastRepository(db *gorm.DB) domain.PodcastRepository {
	return &PodcastRepository{db}
}

func (repo *PodcastRepository) Get() (res []domain.Podcast, err error) {

	var podcasts []domain.Podcast
	errRow := repo.DB.Find(&podcasts).Error

	return podcasts, errRow
}
