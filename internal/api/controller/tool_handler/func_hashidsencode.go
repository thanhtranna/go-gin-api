package tool_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type hashIdsEncodeRequest struct {
	Id int32 `uri:"id"` // Number to be encrypted
}

type hashIdsEncodeResponse struct {
	Val string `json:"val"` // Encrypted value
}

// HashIdsEncode HashIds encryption
// @Summary HashIds encryption
// @Description HashIds encryption
// @Tags API.tool
// @Accept json
// @Produce json
// @Param id path string true "number to be encrypted"
// @Success 200 {object} hashIdsEncodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/encode/{id} [get]
func (h *handler) HashIdsEncode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsEncodeRequest)
		res := new(hashIdsEncodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(req.Id)})
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		res.Val = hashId

		c.Payload(res)
	}
}
