package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func mBain() {
	Bin := NewBin("1", true, "мой bin")
	fmt.Printf("%+v\n", Bin)
}
