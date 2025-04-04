package upload

import (
	"encoding/json"
	"vertical/shared/db"
	"vertical/shared/domain"
)

type Repo struct {
	db *db.Conn
}

func (r *Repo) Save(doc domain.Document) error {
	if bytes, err := json.Marshal(doc); err != nil {
		return err
	} else {
		r.db.Set(doc.ID, bytes)
	}
	return nil
}
