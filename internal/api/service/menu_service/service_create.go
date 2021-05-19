package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateMenuData struct {
	Pid   int32  // Parent class ID
	Name  string // Menu name
	Link  string // link address
	Icon  string // icon
	Level int32  // Menu type 1: Level 1 menu 2: Level 2 menu
}

func (s *service) Create(ctx core.Context, menuData *CreateMenuData) (id int32, err error) {
	model := menu_repo.NewModel()
	model.Pid = menuData.Pid
	model.Name = menuData.Name
	model.Link = menuData.Link
	model.Icon = menuData.Icon
	model.Level = menuData.Level
	model.CreatedUser = ctx.UserName()
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
