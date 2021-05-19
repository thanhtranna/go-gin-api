package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type UpdateMenuData struct {
	Name string // Menu name
	Link string // link address
	Icon string // icon
}

func (s *service) Modify(ctx core.Context, id int32, menuData *UpdateMenuData) (err error) {
	model := menu_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"name":         menuData.Name,
		"link":         menuData.Link,
		"icon":         menuData.Icon,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
