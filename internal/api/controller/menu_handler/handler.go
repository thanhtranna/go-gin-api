package menu_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/menu_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create create/edit menu
	// @Tags API.menu
	// @Router /api/menu [post]
	Create() core.HandlerFunc

	// Detail menu details
	// @Tags API.menu
	// @Router /api/menu/{id} [get]
	Detail() core.HandlerFunc

	// Delete delete menu
	// @Tags API.menu
	// @Router /api/menu/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed Update menu is enabled/disabled
	// @Tags API.menu
	// @Router /api/menu/used [patch]
	UpdateUsed() core.HandlerFunc

	// List menu list
	// @Tags API.menu
	// @Router /api/menu [get]
	List() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	hashids     hash.Hash
	menuService menu_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu_service.New(db, cache),
	}
}

func (h *handler) i() {}
