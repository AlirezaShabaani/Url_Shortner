package urlservices

import (
	"gorm.io/gorm"
	"url_shortner/internal/core/domain"
)

type urlRepositories struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *urlRepositories {
	return &urlRepositories{db: db}
}

func (ur *urlRepositories) Save(data domain.Data) (err error) {
	resp := ur.db.FirstOrCreate(data)
	if resp.Error != nil {
		return resp.Error
	}
	return
}

func (ur *urlRepositories) ReadDb(surl string) (data domain.Data, err error) {
	ur.db.Table("url_shortener.urls").Where(&domain.Data{
		Surl: surl,
	}).Scan(&data)
	return
}
