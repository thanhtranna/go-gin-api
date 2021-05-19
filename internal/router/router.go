package router

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"
	"github.com/xinliangnote/go-gin-api/internal/pkg/metrics"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/internal/router/middleware"
	"github.com/xinliangnote/go-gin-api/pkg/file"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	grpConn grpc.ClientConn
	middles middleware.Middleware
}

type Server struct {
	Mux       core.Mux
	Db        db.Repo
	Cache     cache.Repo
	GrpClient grpc.ClientConn
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	openBrowserUri := "http://127.0.0.1" + configs.ProjectPort()

	_, ok := file.IsExists(configs.ProjectInstallFile())
	if !ok { // Not Installed
		openBrowserUri += "/install"
	} else { // It has been installed
		// Initialize DB
		dbRepo, err := db.New()
		if err != nil {
			logger.Fatal("new db err", zap.Error(err))
		}
		r.db = dbRepo

		// Initialize Cache
		cacheRepo, err := cache.New()
		if err != nil {
			logger.Fatal("new cache err", zap.Error(err))
		}
		r.cache = cacheRepo

		// Initialize gRPC client
		gRPCRepo, err := grpc.New()
		if err != nil {
			logger.Fatal("new grpc err", zap.Error(err))
		}
		r.grpConn = gRPCRepo
	}

	mux, err := core.New(logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithPanicNotify(notify.OnPanicNotify),
		core.WithRecordMetrics(metrics.RecordMetrics),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.middles = middleware.New(logger, r.cache, r.db)

	// Setup WEB routing
	setWebRouter(r)

	// Setup API routing
	setApiRouter(r)

	// Setup GraphQL routing
	setGraphQLRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache
	s.GrpClient = r.grpConn

	return s, nil
}
