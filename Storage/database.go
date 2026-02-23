package storage

import (
	"encoding/json"
	"os"
)

type BinData struct {
	Data []byte `json:"data"`
}

type Database interface {
	Set(key, value string)
	Get(key string) string
}

type MemoryDB struct {
	data map[string]string
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{data: make(map[string]string)}
}

func (m *MemoryDB) Set(k, v string) {
	m.data[k] = v
}

func (m *MemoryDB) Get(k string) string {
	return m.data[k]
}

func SaveBin(bin []byte, filename string) error {
	data := BinData{Data: bin}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func SaveBinList(bins [][]byte, filename string) error {
	dataList := make([]BinData, len(bins))
	for i, b := range bins {
		dataList[i] = BinData{Data: b}
	}
	jsonData, err := json.MarshalIndent(dataList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)

}

func LoadBinListFromFile(filename string) ([][]byte, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dataList []BinData
	err = json.Unmarshal(jsonData, &dataList)
	if err != nil {
		return nil, err
	}
	bins := make([][]byte, len(dataList))
	for i, item := range dataList {
		bins[i] = item.Data
	}
	return bins, nil
}
