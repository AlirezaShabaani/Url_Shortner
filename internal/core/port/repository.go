package port

import "url_shortner/internal/core/domain"

type DBRepository interface {
	Save(data domain.Data) error
	ReadDb(surl string) (domain.Data, error)
}

type CacheRepository interface {
	ReadCache(surl string) (string, error)
	Cache(data domain.Data) error
}
