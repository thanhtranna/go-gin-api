package authorized_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createAPIRequest struct {
	Id     string `form:"id"`     // HashID
	Method string `form:"method"` // request method
	API    string `form:"api"`    // Request address
}

type createAPIResponse struct {
	Id int32 `json:"id"` // Primary Key ID
}

// CreateAPI authorized caller interface address
// @Summary authorized caller interface address
// @Description authorized caller interface address
// @Tags API.authorized
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "HashID"
// @Param method formData string true "request method"
// @Param api formData string true "request address"
// @Success 200 {object} createAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [post]
func (h *handler) CreateAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAPIRequest)
		res := new(createAPIResponse)
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

		createAPIData := new(authorized_service.CreateAuthorizedAPIData)
		createAPIData.BusinessKey = authorizedInfo.BusinessKey
		createAPIData.Method = req.Method
		createAPIData.API = req.API

		createId, err := h.authorizedService.CreateAPI(c, createAPIData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedCreateAPIError,
				code.Text(code.AuthorizedCreateAPIError)).WithErr(err),
			)
			return
		}

		res.Id = createId
		c.Payload(res)
	}
}
