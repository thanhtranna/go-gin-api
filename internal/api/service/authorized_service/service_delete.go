package authorized_service

import (
	"gorm.io/gorm"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	// First check if id exists
	authorizedInfo, err := authorized_repo.NewQueryBuilder().
		WhereIsDeleted(db_repo.EqualPredicate, -1).
		WhereId(db_repo.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	model := authorized_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+authorizedInfo.BusinessKey, cache.WithTrace(ctx.Trace()))
	return
}
