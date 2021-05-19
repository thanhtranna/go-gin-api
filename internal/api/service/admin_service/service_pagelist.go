package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Page     int    // Which page
	PageSize int    // The number of items displayed per page
	Username string // Username
	Nickname string // Nickname
	Mobile   string // phone number
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*admin_repo.Admin, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := admin_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.Username != "" {
		qb.WhereUsername(db_repo.EqualPredicate, searchData.Username)
	}

	if searchData.Nickname != "" {
		qb.WhereNickname(db_repo.EqualPredicate, searchData.Nickname)
	}

	if searchData.Mobile != "" {
		qb.WhereMobile(db_repo.EqualPredicate, searchData.Mobile)
	}

	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
