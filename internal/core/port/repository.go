package port

import "url_shortner/internal/core/domain"

type DBRepository interface {
	Save(data domain.Data) (surl string,err error)
	ReadDb(surl string) (data domain.Data,err error)
}

type CacheRepository interface {
	ReadCache(surl string) (ourl string,err error)
	Cache(data domain.Data) (err error)
}
