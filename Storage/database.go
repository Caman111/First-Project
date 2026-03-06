package storage

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

func (m *MemoryDB) Set(k, v string)     { m.data[k] = v }
func (m *MemoryDB) Get(k string) string { return m.data[k] }
