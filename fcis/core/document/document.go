package document

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Content   []byte    `json:"context"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
