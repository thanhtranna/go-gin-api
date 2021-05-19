package tool_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type hashIdsDecodeRequest struct {
	Id string `uri:"id"` // Ciphertext to be decrypted
}

type hashIdsDecodeResponse struct {
	Val int `json:"val"` // Decrypted value
}

// HashIdsDecode HashIds decrypt
// @Summary HashIds decrypt
// @Description HashIds decrypt
// @Tags API.tool
// @Accept json
// @Produce json
// @Param id path string true "ciphertext to be decrypted"
// @Success 200 {object} hashIdsDecodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/decode/{id} [get]
func (h *handler) HashIdsDecode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsDecodeRequest)
		res := new(hashIdsDecodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		res.Val = hashId[0]

		c.Payload(res)
	}
}
