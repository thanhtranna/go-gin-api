package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// HashIdsEncode HashIds encryption
	// @Tags API.tool
	// @Router /api/tool/hashids/encode/{id} [get]
	HashIdsEncode() core.HandlerFunc

	// HashIdsDecode HashIds decrypt
	// @Tags API.tool
	// @Router /api/tool/hashids/decode/{id} [get]
	HashIdsDecode() core.HandlerFunc

	// SearchCache query cache
	// @Tags API.tool
	// @Router /api/tool/cache/search [post]
	SearchCache() core.HandlerFunc

	// ClearCache clears the cache
	// @Tags API.tool
	// @Router /api/tool/cache/clear [patch]
	ClearCache() core.HandlerFunc

	// Dbs query DB
	// @Tags API.tool
	// @Router /api/tool/data/dbs [get]
	Dbs() core.HandlerFunc

	// Tables query Table
	// @Tags API.tool
	// @Router /api/tool/data/tables [post]
	Tables() core.HandlerFunc

	// SearchMySQL executes SQL statements
	// @Tags API.tool
	// @Router /api/tool/data/mysql [post]
	SearchMySQL() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	hashids hash.Hash
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:  logger,
		db:      db,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
