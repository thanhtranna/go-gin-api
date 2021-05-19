package middleware

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"go.uber.org/zap"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// i In order to avoid being implemented by other packages
	i()

	// JWT Middleware
	Jwt(ctx core.Context) (userId int64, userName string, err errno.Error)

	// Resubmit Middleware
	Resubmit() core.HandlerFunc

	// DisableLog Don't use log
	DisableLog() core.HandlerFunc

	// Signature Signature verification, use signature algorithm pkg/signature
	Signature() core.HandlerFunc

	// Token Signature verification, verification of logged-in users
	Token(ctx core.Context) (userId int64, userName string, err errno.Error)
}

type middleware struct {
	logger            *zap.Logger
	cache             cache.Repo
	db                db.Repo
	authorizedService authorized_service.Service
	adminService      admin_service.Service
}

func New(logger *zap.Logger, cache cache.Repo, db db.Repo) Middleware {
	return &middleware{
		logger:            logger,
		cache:             cache,
		db:                db,
		authorizedService: authorized_service.New(db, cache),
		adminService:      admin_service.New(db, cache),
	}
}

func (m *middleware) i() {}
