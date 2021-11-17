package service

import "github.com/Ras96/gcg/internal/model"

type GeneratorService interface {
	GenerateConstructors(file *model.File, filename string) ([]byte, error)
}
