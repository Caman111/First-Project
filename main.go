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

	filename := "data.json"

	content, err := Files.ReadFileReadll(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println("Содержимое файла:", string(content))

	if Files.IsJSONFile(filename) {
		fmt.Println("Это JSON файл!")
	} else {
		fmt.Println("Это не JSON файл!")
	}
}
