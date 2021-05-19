package handler

import (
	"context"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/graph/generated"
	"github.com/xinliangnote/go-gin-api/internal/graph/resolvers"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/zap"
)

var _ Gql = (*gql)(nil)

type Gql interface {
	i()
	Playground() core.HandlerFunc
	Query() core.HandlerFunc
}

type gql struct {
	logger *zap.Logger
	db     db.Repo
	cache  cache.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Gql {
	return &gql{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (g *gql) i() {}

func (g *gql) Query() core.HandlerFunc {

	// Define extension fields
	extensions := make(map[string]interface{})

	h := handler.New(generated.NewExecutableSchema(
		resolvers.NewRootResolvers(g.logger, g.db, g.cache)),
	)

	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	// Set transport
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New(1000))

	// Enable sidebar documents
	h.Use(extension.Introspection{})

	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(c core.Context) {
		var responses interface{}

		defer func() {
			// Setup core log
			c.GraphPayload(responses)
		}()

		// Setup core trace_id
		extensions["trace_id"] = c.Trace().ID()

		h.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
			resp := next(ctx)
			resp.Extensions = extensions
			responses = resp
			return resp
		})

		// Setup core context
		coreContext := context.WithValue(c.Request().Context(), resolvers.CoreContextKey, c)
		h.ServeHTTP(c.ResponseWriter(), c.Request().WithContext(coreContext))
	}
}

func (g *gql) Playground() core.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")
	return func(c core.Context) {
		h.ServeHTTP(c.ResponseWriter(), c.Request())
	}
}
