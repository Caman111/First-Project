package models

type LocalBin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LocalData struct {
	Bins []LocalBin `json:"bins"`
}

type Config struct {
	Key string
}
