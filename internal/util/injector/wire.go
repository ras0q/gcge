//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ras96/gcg/internal/handler"
	"github.com/Ras96/gcg/internal/repository"
	"github.com/google/wire"
)

var (
	mainSet = wire.NewSet(
		repository.NewParserRepository,
		repository.NewRepositories,
		handler.NewHandlers,
	)
)

func NewHandlers() handler.Handlers {
	wire.Build(
		mainSet,
	)

	return nil
}
