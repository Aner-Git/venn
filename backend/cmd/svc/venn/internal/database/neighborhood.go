package database

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal"
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
)

func (d *GormDatabase) GetNeighborhoods(orderBy string, page, pageSize int, condition ...any) ([]*model.Neighborhood, error) {
	var hoods []*model.Neighborhood
	var err error
	handle := d.DB.Joins("PublicTransportAvailablity").Order(orderBy).Scopes(dal.PaginateDefault(page, pageSize))
	if len(condition) == 1 {
		handle = handle.Where(condition[0])
	} else if len(condition) > 1 {
		handle = handle.Where(condition[0], condition[1:]...)
	}
	err = handle.Find(&hoods).Error
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
