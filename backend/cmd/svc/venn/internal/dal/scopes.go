package dal

import (
	"gorm.io/gorm"
)

const (
	defaultMaxPageSize = 100
	deafultMinPageSize = 10
)

func PaginateDefault(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return Paginate(page, pageSize, defaultMaxPageSize, deafultMinPageSize)
}

func Paginate(page, pageSize, maxPageSize, minPageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > maxPageSize:
			pageSize = maxPageSize
		case pageSize <= 0:
			pageSize = minPageSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
