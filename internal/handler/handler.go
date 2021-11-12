package handler

import (
	"github.com/Ras96/gcg/internal/repository"
	"github.com/spf13/cobra"
)

type Handlers interface {
	Root(cmd *cobra.Command, args []string)
	Gen(cmd *cobra.Command, args []string)
}

type handlers struct {
	repo *repository.Repositories
}

func NewHandlers(r *repository.Repositories) Handlers {
	return &handlers{r}
}
