//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ras96/gcg/internal/handler"
	"github.com/Ras96/gcg/internal/repository"
	"github.com/google/wire"
	"golang.org/x/tools/imports"
)

var (
	handlerSet = wire.NewSet(
		handler.NewHandlers,
	)

	repositorySet = wire.NewSet(
		repository.NewRepositories,
		repository.NewParserRepository,
		repository.NewGeneratorRepository,
	)
)

func NewHandlers(opts *imports.Options) handler.Handlers {
	wire.Build(
		handlerSet,
		repositorySet,
	)

	return nil
}
