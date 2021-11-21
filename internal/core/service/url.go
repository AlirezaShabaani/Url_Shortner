package service

import (
	"errors"
	"url_shortner/internal/core/domain"
	"url_shortner/internal/core/port"
	"url_shortner/pkg/uidgen"
)

type service struct {
	repositoryServices port.RepositoryServices
	uidgen uidgen.UIDGen
}

func New(repositoryServices port.RepositoryServices,uidgen uidgen.UIDGen) *service {
	return &service{repositoryServices: repositoryServices,uidgen: uidgen}
}




func (srvs *service) Save(ourl string) (surl string,err error) {
	surl = srvs.uidgen.New()
	err = srvs.repositoryServices.Save(domain.Data{Ourl:  ourl, Surl:  surl})
	if err != nil {
		return "", errors.New("trouble when saving shortened url in database")
	}
	err = srvs.repositoryServices.Cache(domain.Data{Ourl:  ourl, Surl:  surl})
	if err != nil {
		return "", errors.New("trouble when caching shortened url")
	}
	return surl, nil
}




func (srvs *service) Read(surl string) (ourl string,err error) {
	ourl,err = srvs.repositoryServices.ReadCache(surl)
	if err != nil || ourl == ""{
		data, err := srvs.repositoryServices.ReadDb(surl)
		if err != nil {
			return "",errors.New("nothing found")
		}
		return data.Ourl,nil
	}
	return "", nil
}


