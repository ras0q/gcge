package service

import "github.com/Ras96/gcg/internal/model"

type GeneratorService interface {
	GenerateConstructors(file *model.File, output string, isPrivate bool) ([]byte, error)
}
