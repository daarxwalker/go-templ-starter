package facade

import (
	"context"
	
	"common/pkg/service/cache_service"
)

func Cache(c context.Context) *cache_service.CacheService {
	service, ok := c.Value(cache_service.Token).(*cache_service.CacheService)
	if !ok {
		panic(cache_service.Token + " not found in context")
	}
	return service
}
