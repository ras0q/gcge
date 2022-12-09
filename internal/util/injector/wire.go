//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/ras0q/gcg/internal/handler"
	"github.com/ras0q/gcg/internal/service"
)

func NewHandlers() *handler.Handlers {
	wire.Build(
		handler.NewHandlers,
		service.NewServices,
		service.NewAnalyzerService,
		service.NewGeneratorService,
	)

	return nil
}
