package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type dbsResponse struct {
	List []dbData `json:"list"` // Database list
}

type dbData struct {
	DbName string `json:"db_name"` // Database name
}

// Dbs query DB
// @Summary query DB
// @Description query DB
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} dbsResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/dbs [get]
func (h *handler) Dbs() core.HandlerFunc {
	return func(c core.Context) {
		res := new(dbsResponse)

		// TODO supports querying multiple databases later
		data := dbData{
			DbName: configs.Get().MySQL.Read.Name,
		}

		res.List = append(res.List, data)
		c.Payload(res)
	}
}
