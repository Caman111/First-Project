package main

import (
	"Bins/Files"
	"Bins/api"
	"Bins/storage"
	"fmt"
)

func main() {

	db := storage.NewMemoryDB()
	fm := Files.NewMemoryFiles()

	service := api.NewService(db, fm)
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
		fmt.Println("Это не JSON файл.")
	}
}
