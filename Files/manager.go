package Files

import (
	"os"
	"path/filepath"
	"strings"
)

type FileManager interface {
	Save(name, data string)
	Load(name string) string
}

type MemoryFiles struct {
	files map[string]string
}

func NewMemoryFiles() *MemoryFiles {
	return &MemoryFiles{files: make(map[string]string)}
}

func (m *MemoryFiles) Save(n, d string) {
	m.files[n] = d
}

func (m *MemoryFiles) Load(n string) string {
	return m.files[n]
}

func ReadFileReadll(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
func IsJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}
