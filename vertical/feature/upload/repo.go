package upload

import (
	"fmt"
	"vert/shared/db"
	"vert/shared/domain"
)

type Repo struct {
	db *db.Conn
}

func (r *Repo) Save(_ domain.Document) error {
	fmt.Println("Saving document ...")
	return nil
}
