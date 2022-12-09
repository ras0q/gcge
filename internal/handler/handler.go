package handler

import "github.com/ras0q/gcg/internal/service"

type Handlers struct {
	Srv *service.Services
}

func NewHandlers(srv *service.Services) *Handlers {
	return &Handlers{
		Srv: srv,
	}
}
