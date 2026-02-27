package api

import (
	"Bins/Files"
	"Bins/config"
	"Bins/storage"
	"fmt"
)

type Service struct {
	DB     storage.Database
	FM     Files.FileManager
	Config *config.Config
}

func NewService(db storage.Database, fm Files.FileManager, cfg *config.Config) *Service {
	return &Service{
		DB:     db,
		FM:     fm,
		Config: cfg,
	}
}

func (s *Service) Run() {
	s.DB.Set("user", "Kama")
	s.FM.Save("log.txt", "started")

	fmt.Println("API ключ:", s.Config.Key)
}
