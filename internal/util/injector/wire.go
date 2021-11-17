//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ras96/gcg/internal/handler"
	"github.com/Ras96/gcg/internal/service"
	"github.com/google/wire"
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
