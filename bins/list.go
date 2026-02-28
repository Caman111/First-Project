package bins

import "time"

type BinList struct {
	ID        string    `jsom:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
