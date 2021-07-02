package service

import "github.com/bluele/gcache"

// Service - Service struct
type Service struct {
	//redisClient *redis.Client
	cacheClient gcache.Cache
}

// NewCacheClient - Creates a new Cache client
func NewCacheClient() *Service {
	client := gcache.New(20).
		LRU().
		Build()
	return &Service{cacheClient: client}
}

// SetToken -
func (s *Service) SetToken(key, token string) error {
	return s.cacheClient.Set(key, token)
}

// GetToken -
func (s *Service) GetToken(key string) (interface{}, error) {
	return s.cacheClient.Get(key)
}
