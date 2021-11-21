package port

import "url_shortner/internal/core/domain"

type RepositoryServices interface {
	Save(data domain.Data)error
	Cache(data domain.Data)error
	ReadCache(surl string)(string,error)
	ReadDb(surl string)(domain.Data,error)
}
