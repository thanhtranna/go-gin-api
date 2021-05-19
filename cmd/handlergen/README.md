## Excuting an order

```$xslt
// test_handler is the package name in ./internal/api/controller/
./scripts/handlergen.sh test_handler
```

## Template file reference

```go
package test_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i in order to avoid being implemented by other packages
	i()

	// Create create user
	// @Tags Test
	// @Router /test/create [post]
	Create() core.HandlerFunc

	// Update edit user
	// @Tags Test
	// @Router /test/update [post]
	Update() core.HandlerFunc

	// Delete delete user
	// @Tags Test
	// @Router /test/delete [post]
	Delete() core.HandlerFunc

	// Detail user detail
	// @Tags Test
	// @Router /test/detail [post]
	Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	userService user_service.UserService
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		userService: user_service.NewUserService(db, cache),
	}
}

func (h *handler) i() {}

```

The above will generate 4 files
- func_create.go
- func_update.go
- func_delete.go
- func_detail.go

## func_create.go reference

```go
package test_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

// Create Create user
// @Summary create user
// @Description creates a user
// @Tags Test
// @Accept json
// @Produce json
// @Param Request body createRequest true "request information"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /test/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {

	}
}

```