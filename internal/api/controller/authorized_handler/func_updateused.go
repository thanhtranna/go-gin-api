package authorized_handler

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

// UpdateUsed update the caller to enable/disable
// @Summary Update the caller to enable/disable
// @Description updates the caller to enable/disable
// @Tags API.authorized
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "Hashid"
// @Param used formData int true "Is it enabled 1: Yes -1: No"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized/used [patch]
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

		err = h.authorizedService.UpdateUsed(c, id, req.Used)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedUpdateError,
				code.Text(code.AuthorizedUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
