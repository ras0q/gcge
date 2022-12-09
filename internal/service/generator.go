package service

import "github.com/ras0q/gcg/internal/model"

type GeneratorService interface {
	GenerateConstructors(file *model.File, output string, isPrivate bool) ([]byte, error)
}
