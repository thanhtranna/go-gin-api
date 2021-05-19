package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // Caller key
	Method      string `json:"method"`       // Request method
	API         string `json:"api"`          // Request address
}

func (s *service) CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error) {
	model := authorized_api_repo.NewModel()
	model.BusinessKey = authorizedAPIData.BusinessKey
	model.Method = authorizedAPIData.Method
	model.Api = authorizedAPIData.API
	model.CreatedUser = ctx.UserName()
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	s.cache.Del(cacheKeyPrefix+authorizedAPIData.BusinessKey, cache.WithTrace(ctx.Trace()))
	return
}
