package database

import "github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"

type Loader struct {
	db *GormDatabase
}

func NewLoader(DB *GormDatabase) *Loader {
	return &Loader{db: DB}
}

func (l *Loader) Migrate() error {
	return l.db.DB.AutoMigrate(&model.Neighborhood{})
}
