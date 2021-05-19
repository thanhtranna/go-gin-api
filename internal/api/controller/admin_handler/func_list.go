package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/time_parse"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type listRequest struct {
	Page     int    `form:"page"`      // which page
	PageSize int    `form:"page_size"` // The number of items displayed per page
	Username string `form:"username"`  // username
	Nickname string `form:"nickname"`  // Nickname
	Mobile   string `form:"mobile"`    // mobile phone number
}

type listData struct {
	Id          int    `json:"id"`           // ID
	HashID      string `json:"hashid"`       // hashid
	Username    string `json:"username"`     // username
	Nickname    string `json:"nickname"`     // Nickname
	Mobile      string `json:"mobile"`       // mobile phone number
	IsUsed      int    `json:"is_used"`      // Whether to enable 1: Yes -1: No
	CreatedAt   string `json:"created_at"`   // created time
	CreatedUser string `json:"created_user"` // created by
	UpdatedAt   string `json:"updated_at"`   // updated time
	UpdatedUser string `json:"updated_user"` // Updated by
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PrePageCount int `json:"pre_page_count"`
	} `json:"pagination"`
}

// List admin list
// @Summary administrator list
// @Description admin list
// @Tags API.admin
// @Accept json
// @Produce json
// @Param page query int false "page number"
// @Param page_size query string false "Number of items displayed per page"
// @Param username query string false "Username"
// @Param nickname query string false "nickname"
// @Param mobile query string false "mobile phone number"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(admin_service.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Username = req.Username
		searchData.Nickname = req.Nickname
		searchData.Mobile = req.Mobile

		resListData, err := h.adminService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithErr(err),
			)
			return
		}

		resCountData, err := h.adminService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithErr(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PrePageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				h.logger.Info("hashids err", zap.Error(err))
			}

			data := listData{
				Id:          cast.ToInt(v.Id),
				HashID:      hashId,
				Username:    v.Username,
				Nickname:    v.Nickname,
				Mobile:      v.Mobile,
				IsUsed:      cast.ToInt(v.IsUsed),
				CreatedAt:   v.CreatedAt.Format(time_parse.CSTLayout),
				CreatedUser: v.CreatedUser,
				UpdatedAt:   v.UpdatedAt.Format(time_parse.CSTLayout),
				UpdatedUser: v.UpdatedUser,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
