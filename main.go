package main

import (
	"Bins/Files"
	"Bins/api"
	"Bins/config"
	"Bins/storage"
	"fmt"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.Key)

	db := storage.NewMemoryDB()
	fm := Files.NewMemoryFiles()

	service := api.NewService(db, fm, &cfg)
	service.Run()

}
