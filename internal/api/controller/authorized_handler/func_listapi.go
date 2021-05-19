package authorized_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type listAPIRequest struct {
	Id string `form:"id"` // hashID
}

type listAPIData struct {
	HashId      string `json:"hash_id"`      // hashID
	BusinessKey string `json:"business_key"` // caller key
	Method      string `json:"method"`       // caller secret
	API         string `json:"api"`          // caller docking person
}

type listAPIResponse struct {
	List []listAPIData `json:"list"`
}

// ListAPI caller interface address list
// @Summary caller interface address list
// @Description caller interface address list
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param business_key query string false "caller key"
// @Success 200 {object} listAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [get]
func (h *handler) ListAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listAPIRequest)
		res := new(listAPIResponse)
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

		// Query business_key by id
		authorizedInfo, err := h.authorizedService.Detail(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedDetailError,
				code.Text(code.AuthorizedDetailError)).WithErr(err),
			)
			return
		}

		searchAPIData := new(authorized_service.SearchAPIData)
		searchAPIData.BusinessKey = authorizedInfo.BusinessKey

		resListData, err := h.authorizedService.ListAPI(c, searchAPIData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedListAPIError,
				code.Text(code.AuthorizedListAPIError)).WithErr(err),
			)
			return
		}

		res.List = make([]listAPIData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				h.logger.Info("hashids err", zap.Error(err))
			}

			data := listAPIData{
				HashId:      hashId,
				BusinessKey: v.BusinessKey,
				Method:      v.Method,
				API:         v.Api,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
