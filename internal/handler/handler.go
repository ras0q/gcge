package handler

import "github.com/Ras96/gcg/internal/repository"

type Handlers struct {
	Gen  genHandler
	Root rootHandler

	Repo *repository.Repositories
}

func NewHandlers(repo *repository.Repositories) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) SetupGen(output *string) {
	h.Gen = genHandler{
		repo: h.Repo,
		opts: &genOpts{
			output: output,
		},
	}
}

func (h *Handlers) SetupRoot() {
	h.Root = rootHandler{}
}
