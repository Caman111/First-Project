package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFileReadll(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
func IsJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}
