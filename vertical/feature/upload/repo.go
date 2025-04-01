package upload

import (
	"fmt"
	"vert/shared/db"
	"vert/shared/domain"
)

type Repository struct {
	db *db.DB
}

func (r *Repository) Save(_ domain.Document) error {
	fmt.Println("Saving document ...")
	return nil
}
