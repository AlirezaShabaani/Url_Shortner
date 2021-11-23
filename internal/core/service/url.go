package service

import (
	"errors"
	"url_shortner/internal/core/domain"
	"url_shortner/internal/core/port"
	"url_shortner/pkg/random"
	"url_shortner/pkg/uidgen"
)

type service struct {
	DBRepository    port.DBRepository
	CacheRepository port.CacheRepository
	uidgen          uidgen.UIDGen
	random          random.ShortenerService
}

func New(repositoryServices port.DBRepository, uidgen uidgen.UIDGen, CacheRepository port.CacheRepository,random random.ShortenerService) *service {
	return &service{DBRepository: repositoryServices, uidgen: uidgen, CacheRepository: CacheRepository,random: random}
}

func (srvs *service) Save(ourl string) (surl string, err error) {
	surl = srvs.random.New(4)
	surl, err = srvs.DBRepository.Save(domain.Data{Ourl: ourl, Surl: surl})
	if err != nil {
		return "", errors.New("trouble when saving shortened url in database")
	}
	err = srvs.CacheRepository.Cache(domain.Data{Ourl: ourl, Surl: surl})
	if err != nil {
		return "", errors.New("trouble when caching shortened url")
	}
	return surl, nil
}


func (srvs *service) Read(surl string) (ourl string, err error) {
	ourl, err = srvs.CacheRepository.ReadCache(surl)
	if err != nil || ourl == "" {
		data, err := srvs.DBRepository.ReadDb(surl)
		if err != nil {
			return "", errors.New("nothing found")
		}
		return data.Ourl, nil
	}
	return ourl, nil
}
