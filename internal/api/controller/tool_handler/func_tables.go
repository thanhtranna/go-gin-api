package tool_handler

import (
	"fmt"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type tablesRequest struct {
	DbName string `form:"db_name"` // database name
}

type tablesResponse struct {
	List []tableData `json:"list"` // Data sheet list
}

type tableData struct {
	Name    string `json:"table_name"`    // data table name
	Comment string `json:"table_comment"` // Data table comment
}

// Tables query Table
// @Summary query Table
// @Description query Table
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param db_name formData string true "database name"
// @Success 200 {object} tablesResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/tables [post]
func (h *handler) Tables() core.HandlerFunc {
	return func(c core.Context) {
		req := new(tablesRequest)
		res := new(tablesResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", req.DbName)

		// TODO Support for querying multiple databases later
		rows, err := h.db.GetDbR().Raw(sqlTables).Rows()
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SearchMySQLError,
				code.Text(code.SearchMySQLError)).WithErr(err),
			)
			return
		}

		defer rows.Close()

		for rows.Next() {
			var info tableData
			err = rows.Scan(&info.Name, &info.Comment)
			if err != nil {
				fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
				continue
			}

			res.List = append(res.List, info)
		}

		c.Payload(res)
	}
}
