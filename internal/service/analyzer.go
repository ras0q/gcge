package service

import (
	"github.com/Ras96/gcg/internal/model"
)

type AnalyzerService interface {
	AnalyzeFile(filename string) (*model.File, error)
}
