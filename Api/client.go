package api

import (
	"Bins/Files"
	"Bins/storage"
)

type Service struct {
	DB storage.Database
	FM Files.FileManager
}

func NewService(db storage.Database, fm Files.FileManager) *Service {
	return &Service{DB: db, FM: fm}
}

func (s *Service) Run() {
	s.DB.Set("user", "Kama")
	s.FM.Save("log.txt", "started")
}
