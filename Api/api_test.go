package api

import (
	"encoding/json"
	"os"
	"testing"
)

const testAPIURL = "https://api.jsonbin.io/v3/b"

func TestCreateBin(t *testing.T) {
	key := os.Getenv("X_MASTER_KEY")
	if key == "" {
		t.Skip("Нет API ключа")
	}

	data := []byte(`{"test":"create test","number":123}`)

	resp, err := CreateBin(testAPIURL, data)
	if err != nil {
		t.Fatal("Ошибка создания:", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		t.Fatal("Ошибка парсинга ответа:", err)
	}

	metadata, ok := result["metadata"].(map[string]interface{})
	if !ok {
		t.Fatal("Нет metadata в ответе")
	}

	id, ok := metadata["id"].(string)
	if !ok || id == "" {
		t.Fatal("Не получили ID бина")
	}

	DeleteBin(testAPIURL + "/" + id)
}

func TestGetBin(t *testing.T) {
	key := os.Getenv("X_MASTER_KEY")
	if key == "" {
		t.Skip("Нет API ключа")
	}
	testData := []byte(`{"test":"get test","value":42}`)
	resp, err := CreateBin(testAPIURL, testData)
	if err != nil {
		t.Fatal("Ошибка создания тестового бина:", err)
	}

	var createResult map[string]interface{}
	json.Unmarshal(resp, &createResult)
	binID := createResult["metadata"].(map[string]interface{})["id"].(string)

	url := testAPIURL + "/" + binID
	getResp, err := GetBin(url)
	if err != nil {
		t.Fatal("Ошибка получения бина:", err)
	}

	var getResult map[string]interface{}
	err = json.Unmarshal(getResp, &getResult)
	if err != nil {
		t.Fatal("Ошибка парсинга полученных данных:", err)
	}

	record, ok := getResult["record"].(map[string]interface{})
	if !ok {

		record = getResult
	}

	if record["test"] != "get test" {
		t.Errorf("Ожидалось 'get test', получено %v", record["test"])
	}
	if record["value"] != float64(42) {
		t.Errorf("Ожидалось 42, получено %v", record["value"])
	}

	DeleteBin(url)
}

func TestUpdateBin(t *testing.T) {
	key := os.Getenv("X_MASTER_KEY")
	if key == "" {
		t.Skip("Нет API ключа")
	}

	initialData := []byte(`{"test":"initial","count":1}`)
	resp, err := CreateBin(testAPIURL, initialData)
	if err != nil {
		t.Fatal("Ошибка создания:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(resp, &result)
	binID := result["metadata"].(map[string]interface{})["id"].(string)
	url := testAPIURL + "/" + binID

	updatedData := []byte(`{"test":"updated","count":99,"new":"field"}`)
	_, err = UpdateBin(url, updatedData)
	if err != nil {
		t.Fatal("Ошибка обновления:", err)
	}

	getResp, err := GetBin(url)
	if err != nil {
		t.Fatal("Ошибка получения после обновления:", err)
	}

	var getResult map[string]interface{}
	json.Unmarshal(getResp, &getResult)

	record, ok := getResult["record"].(map[string]interface{})
	if !ok {
		record = getResult
	}

	if record["test"] != "updated" {
		t.Errorf("Ожидалось 'updated', получено %v", record["test"])
	}
	if record["count"] != float64(99) {
		t.Errorf("Ожидалось 99, получено %v", record["count"])
	}
	if record["new"] != "field" {
		t.Errorf("Ожидалось 'field', получено %v", record["new"])
	}

	DeleteBin(url)
}

func TestDeleteBin(t *testing.T) {
	key := os.Getenv("X_MASTER_KEY")
	if key == "" {
		t.Skip("Нет API ключа")
	}

	data := []byte(`{"test":"to be deleted"}`)
	resp, err := CreateBin(testAPIURL, data)
	if err != nil {
		t.Fatal("Ошибка создания:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(resp, &result)
	binID := result["metadata"].(map[string]interface{})["id"].(string)
	url := testAPIURL + "/" + binID

	_, err = DeleteBin(url)
	if err != nil {
		t.Fatal("Ошибка удаления:", err)
	}
	_, err = GetBin(url)
	if err == nil {
		t.Error("Бин все еще доступен после удаления")
	}
}
