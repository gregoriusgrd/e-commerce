package repository

import (
	"ecommerce/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CitiesRepository struct {
	Repository[entity.Cities]
	Log *logrus.Logger
}

func NewCitiesRepository(log *logrus.Logger) *CitiesRepository {
	return &CitiesRepository{
		Log: log,
	}
}

func (r *CitiesRepository) FindByProvinceId(tx *gorm.DB, id int) ([]entity.Cities, error) {
	var cities []entity.Cities
	err := tx.Where("province_id = ?", id).Find(&cities).Error
	if err != nil {
		r.Log.WithError(err).Error("failed to find cities")
		return nil, err
	}
	return cities, nil
}
