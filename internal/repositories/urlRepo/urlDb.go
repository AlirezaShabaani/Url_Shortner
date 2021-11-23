package urlRepo

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

func (ur *urlRepositories) Save(data domain.Data) (surl string,err error) {
	resp := ur.db.Where(&domain.Data{
		Ourl: data.Ourl,
	}).FirstOrCreate(&data)
	if resp.Error != nil {
		return data.Ourl,resp.Error
	}
	surl = data.Surl
	return
}

func (ur *urlRepositories) ReadDb(surl string) (data domain.Data, err error) {
	ur.db.Table("url_shortener.urls").Where(&domain.Data{
		Surl: surl,
	}).Scan(&data)
	return
}
