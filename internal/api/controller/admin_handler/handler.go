package admin_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login administrator login
	// @Tags API.admin
	// @Router /api/admin/login [post]
	Login() core.HandlerFunc

	// Logout administrator logout
	// @Tags API.admin
	// @Router /api/admin/logout [post]
	Logout() core.HandlerFunc

	// ModifyPassword change password
	// @Tags API.admin
	// @Router /api/admin/modify_password [patch]
	ModifyPassword() core.HandlerFunc

	// Detail get personal information
	// @Tags API.admin
	// @Router /api/admin/info [get]
	Detail() core.HandlerFunc

	// ModifyPersonalInfo edit personal information
	// @Tags API.admin
	// @Router /api/admin/modify_personal_info [patch]
	ModifyPersonalInfo() core.HandlerFunc

	// Create add manager
	// @Tags API.admin
	// @Router /api/admin [post]
	Create() core.HandlerFunc

	// List administrator list
	// @Tags API.admin
	// @Router /api/admin [get]
	List() core.HandlerFunc

	// Delete delete administrator
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed update administrator to enable/disable
	// @Tags API.admin
	// @Router /api/admin/used [patch]
	UpdateUsed() core.HandlerFunc

	// ResetPassword reset password
	// @Tags API.admin
	// @Router /api/admin/reset_password/{id} [patch]
	ResetPassword() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	cache        cache.Repo
	hashids      hash.Hash
	adminService admin_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:       logger,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		adminService: admin_service.New(db, cache),
	}
}

func (h *handler) i() {}
