package repository

import (
	"ecommerce/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProvincesRepository struct {
	Repository[entity.Provinces]
	Log *logrus.Logger
}

func NewProvincesRepositories(log *logrus.Logger) *ProvincesRepository {
	return &ProvincesRepository{
		Log: log,
	}
}

func (r *ProvincesRepository) FindAll(tx *gorm.DB) ([]entity.Provinces, error) {
	var provinces []entity.Provinces
	err := tx.Find(&provinces).Error
	if err != nil {
		r.Log.WithError(err).Error("Failed to find Provinces")
		return nil, err
	}
	return provinces, nil
}
