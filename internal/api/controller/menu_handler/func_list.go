package menu_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/menu_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type listData struct {
	Id     int32  `json:"id"`      // ID
	HashID string `json:"hashid"`  // hashid
	Pid    int32  `json:"pid"`     // Parent class ID
	Name   string `json:"name"`    // Menu name
	Link   string `json:"link"`    // link address
	Icon   string `json:"icon"`    // icon
	IsUsed int32  `json:"is_used"` // Whether to enable 1=Enable -1=Disable
}

type listResponse struct {
	List []listData `json:"list"`
}

// List menu list
// @Summary menu list
// @Description menu list
// @Tags API.menu
// @Accept json
// @Produce json
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		res := new(listResponse)
		resListData, err := h.menuService.List(c, new(menu_service.SearchData))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithErr(err),
			)
			return
		}

		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				h.logger.Info("hashids err", zap.Error(err))
			}

			data := listData{
				Id:     v.Id,
				HashID: hashId,
				Pid:    v.Pid,
				Name:   v.Name,
				Link:   v.Link,
				Icon:   v.Icon,
				IsUsed: v.IsUsed,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
