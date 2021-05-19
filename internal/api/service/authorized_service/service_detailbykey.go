package authorized_service

import (
	"encoding/json"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

// Define the cache structure
type CacheAuthorizedData struct {
	Key    string         `json:"key"`     // caller key
	Secret string         `json:"secret"`  // caller secret
	IsUsed int32          `json:"is_used"` // Caller enabled state 1=enabled -1=disabled
	Apis   []cacheApiData `json:"apis"`    // Apis authorized by the caller
}

type cacheApiData struct {
	Method string `json:"method"` // request method
	Api    string `json:"api"`    // request address
}

func (s *service) DetailByKey(ctx core.Context, key string) (cacheData *CacheAuthorizedData, err error) {
	// Query cache
	cacheKey := cacheKeyPrefix + key
	value, err := s.cache.Get(cacheKey, cache.WithTrace(ctx.RequestContext().Trace))

	cacheData = new(CacheAuthorizedData)
	if err == nil && json.Unmarshal([]byte(value), cacheData) == nil {
		return
	}

	// Query caller information
	authorizedInfo, err := authorized_repo.NewQueryBuilder().
		WhereIsDeleted(db_repo.EqualPredicate, -1).
		WhereBusinessKey(db_repo.EqualPredicate, key).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		return nil, err
	}

	// Query caller authorized API information
	authorizedApiInfo, err := authorized_api_repo.NewQueryBuilder().
		WhereIsDeleted(db_repo.EqualPredicate, -1).
		WhereBusinessKey(db_repo.EqualPredicate, key).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		return nil, err
	}

	// Set cache data
	cacheData = new(CacheAuthorizedData)
	cacheData.Key = key
	cacheData.Secret = authorizedInfo.BusinessSecret
	cacheData.IsUsed = authorizedInfo.IsUsed
	cacheData.Apis = make([]cacheApiData, len(authorizedApiInfo))

	for k, v := range authorizedApiInfo {
		data := cacheApiData{
			Method: v.Method,
			Api:    v.Api,
		}
		cacheData.Apis[k] = data
	}

	cacheDataByte, _ := json.Marshal(cacheData)

	err = s.cache.Set(cacheKey, string(cacheDataByte), time.Hour*24, cache.WithTrace(ctx.Trace()))
	if err != nil {
		return nil, err
	}

	return
}
