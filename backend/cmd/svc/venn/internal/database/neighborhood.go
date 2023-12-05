package database

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
)

func (d *GormDatabase) GetNeighborhoods(offset, limit int) ([]*model.Neighborhood, error) {
	var hoods []*model.Neighborhood
	var err error
	err = d.DB.Joins("PublicTransportAvailablity").Order("name").Offset(offset).Limit(limit).Find(&hoods).Error

	return hoods, err
}

func (d *GormDatabase) CountNeighborhood(condition ...any) (int64, error) {
	c := int64(-1)
	handle := d.DB.Model(new(model.Neighborhood))
	if len(condition) == 1 {
		handle = handle.Where(condition[0])
	} else if len(condition) > 1 {
		handle = handle.Where(condition[0], condition[1:]...)
	}
	err := handle.Count(&c).Error
	return c, err
}
