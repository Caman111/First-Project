package Files

import "os"

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (f *FileManager) Save(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}