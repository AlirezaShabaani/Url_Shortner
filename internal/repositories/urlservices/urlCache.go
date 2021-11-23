package urlservices

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"url_shortner/internal/core/domain"
)

type redisService struct {
	rdis redis.Client
}

func NewCache(rdis *redis.Client) *redisService {
	return &redisService{rdis: *rdis}
}

func (r *redisService) Cache(data domain.Data) (err error) {
	status := r.rdis.Set(context.TODO(),data.Surl,data.Ourl,time.Hour * 999999)
	if status.Err() != nil{
		return status.Err()
	}
	return
}


func (r *redisService) ReadCache(surl string) (ourl string, err error) {
	ourl, err = r.rdis.Get(context.TODO(),surl).Result()
	if err != nil {
		return
	}
	return
}
