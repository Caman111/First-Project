package main

import (
	files "Bins/Files"
	"fmt"
)

func main() {
	filename := "data.json"

	content, err := files.ReadFileReadll(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println("Содержимое файла:", string(content))

	if files.IsJSONFile(filename) {
		fmt.Println("Это JSON файл!")
	} else {
		fmt.Println("Это не JSON файл.")
	}
}
