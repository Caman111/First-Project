package api

import (
	"Bins/Files"
	"Bins/models"
	"Bins/storage"
)

type Service struct {
	DB     storage.Database
	FM     *Files.FileManager
	Config *models.Config
}

func NewService(db storage.Database, fm *Files.FileManager, cfg *models.Config) *Service {
	return &Service{
		DB:     db,
		FM:     fm,
		Config: cfg,
	}
}
