package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
)

func (s *service) ModifyPassword(ctx core.Context, id int32, newPassword string) (err error) {
	model := admin_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"password":     password.GeneratePassword(newPassword),
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+password.GenerateLoginToken(id), cache.WithTrace(ctx.Trace()))
	return
}
