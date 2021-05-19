package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Page              int    `json:"page"`               // which page
	PageSize          int    `json:"page_size"`          // Display the number of items per page
	BusinessKey       string `json:"business_key"`       // caller key
	BusinessSecret    string `json:"business_secret"`    // caller secret
	BusinessDeveloper string `json:"business_developer"` // caller docking person
	Remark            string `json:"remark"`             // Remark
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := authorized_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.BusinessKey != "" {
		qb.WhereBusinessKey(db_repo.EqualPredicate, searchData.BusinessKey)
	}

	if searchData.BusinessSecret != "" {
		qb.WhereBusinessSecret(db_repo.EqualPredicate, searchData.BusinessSecret)
	}

	if searchData.BusinessDeveloper != "" {
		qb.WhereBusinessDeveloper(db_repo.EqualPredicate, searchData.BusinessDeveloper)
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
