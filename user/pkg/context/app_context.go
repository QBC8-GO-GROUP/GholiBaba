package context

import (
	"context"
	"log/slog"
	"os"

	"gorm.io/gorm"
)

var defaultLogger *slog.Logger

func init() {
	defaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

type appContext struct {
	context.Context
	db           *gorm.DB
	shouldCommit bool
	logger       *slog.Logger
}

func SetLogger(ctx context.Context, logger *slog.Logger) {
	if appCtx, ok := ctx.(*appContext); ok {
		appCtx.logger = logger
	}
}

func GetLogger(ctx context.Context) *slog.Logger {
	appCtx, ok := ctx.(*appContext)
	if !ok || appCtx.logger == nil {
		return defaultLogger
	}

	return appCtx.logger
}
