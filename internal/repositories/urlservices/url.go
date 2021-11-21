package urlservices

import (
	"gorm.io/gorm"
	"url_shortner/internal/core/domain"
)

type urlRepositories struct {
	db *gorm.DB
}

func New(db *gorm.DB) *urlRepositories {
	return &urlRepositories{db: db}
}

func (ur *urlRepositories) Save(data domain.Data) (err error) {
	return
}

func (ur *urlRepositories) ReadDb(surl string) (data domain.Data, err error) {
	return
}
