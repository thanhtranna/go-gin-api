package menu_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type updateUsedRequest struct {
	Id   string `form:"id"`   // Primary key ID
	Used int32  `form:"used"` // Whether to enable 1: Yes -1: No
}

type updateUsedResponse struct {
	Id int32 `json:"id"` // Primary Key ID
}

// UpdateUsed Update menu is enabled/disabled
// @Summary Update the menu to enable/disable
// @Description Update the menu to enable/disable
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id formData string true "Hashid"
// @Param used formData int true "Is it enabled 1: Yes -1: No"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/used [patch]
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.menuService.UpdateUsed(c, id, req.Used)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MenuUpdateError,
				code.Text(code.MenuUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
