package install_handler

import (
	"go.uber.org/zap"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	View() core.HandlerFunc
	Execute() core.HandlerFunc
	Restart() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) i() {}
