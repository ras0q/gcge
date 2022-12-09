package service

import (
	"github.com/ras0q/gcg/internal/model"
)

type AnalyzerService interface {
	AnalyzeFile(filename string) (*model.File, error)
}
