package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, name string) Bin {
	return Bin{
		ID:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
func mBain() {
	Bin := newBin("1", true, "мой bin")
	fmt.Printf("%+v\n", Bin)
}
