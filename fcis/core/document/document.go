package document

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Content   []byte    `json:"context"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(name string) Document {
	return Document{
		ID:        time.Now().UnixNano(),
		Name:      name,
		Content:   []byte{},
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	}
}
