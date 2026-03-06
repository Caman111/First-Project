package main

import (
	"Bins/Files"
	"Bins/api"
	"Bins/models"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const apiURL = "https://api.jsonbin.io/v3/b"
const localFile = "my.json"

func loadLocal() models.LocalData {

	var data models.LocalData

	content, err := os.ReadFile(localFile)
	if err != nil {
		return data
	}

	json.Unmarshal(content, &data)

	return data
}

func saveLocal(data models.LocalData) {

	content, _ := json.MarshalIndent(data, "", "  ")

	os.WriteFile(localFile, content, 0644)

}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не найден .env файл")
	}

	create := flag.Bool("create", false, "")
	get := flag.Bool("get", false, "")
	update := flag.Bool("update", false, "")
	deleteF := flag.Bool("delete", false, "")
	list := flag.Bool("list", false, "")

	file := flag.String("file", "", "")
	name := flag.String("name", "", "")
	id := flag.String("id", "", "")

	flag.Parse()

	local := loadLocal()

	if *create {

		data, err := Files.ReadFile(*file)
		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := api.CreateBin(apiURL, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		var r map[string]interface{}
		json.Unmarshal(res, &r)

		meta := r["metadata"].(map[string]interface{})
		binID := fmt.Sprintf("%v", meta["id"])

		local.Bins = append(local.Bins, models.LocalBin{
			ID:   binID,
			Name: *name,
		})

		saveLocal(local)

		fmt.Println("created:", binID)
	}

	
	if *get {

		url := apiURL + "/" + *id

		res, err := api.GetBin(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(res))
	}
	
	if *update {

		data, err := Files.ReadFile(*file)
		if err != nil {
			fmt.Println(err)
			return
		}

		url := apiURL + "/" + *id

		_, err = api.UpdateBin(url, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("updated:", *id)
	}

	if *deleteF {

		url := apiURL + "/" + *id

		_, err := api.DeleteBin(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, b := range local.Bins {

			if b.ID == *id {

				local.Bins = append(local.Bins[:i], local.Bins[i+1:]...)
				break

			}

		}

		saveLocal(local)

		fmt.Println("deleted:", *id)
	}

	if *list {

		for _, b := range local.Bins {

			fmt.Println(b.ID, "-", b.Name)

		}

	}

}
