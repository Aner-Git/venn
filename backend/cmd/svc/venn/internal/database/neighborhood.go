package database

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
)

func (d *GormDatabase) GetNeighborhoods() ([]*model.Neighborhood, error) {
	var hoods []*model.Neighborhood
	var err error
	err = d.DB.Joins("PublicTransportAvailablity").Order("name").Find(&hoods).Error

	return hoods, err
}
