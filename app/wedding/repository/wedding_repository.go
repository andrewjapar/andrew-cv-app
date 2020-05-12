package repository

import (
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/jinzhu/gorm"
)

type WeddingRepository struct {
	DB *gorm.DB
}

func NewWeddingRepository(db *gorm.DB) domain.WeddingRepository {
	return &WeddingRepository{db}
}

func (repo *WeddingRepository) GetByUserID(id int64) (res []domain.Wedding, err error) {

	var weddings []domain.Wedding
	errRow := repo.DB.Table("weddings").
		Joins("left join wedding_organizers on wedding_organizers.wedding_id = weddings.id left join users on wedding_organizers.user_id = users.id").
		Where("users.id = ?", id).
		Scan(&weddings).Error

	return weddings, errRow
}
