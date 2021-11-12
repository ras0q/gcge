package handler

import (
	"github.com/Ras96/gcg/internal/repository"
)

type Handlers struct {
	Repo *repository.Repositories
}

func NewHandlers(r *repository.Repositories) *Handlers {
	return &Handlers{r}
}
