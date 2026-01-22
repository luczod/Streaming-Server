// Package repository is to foo
package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"authServer/internal/model"

	"github.com/labstack/gommon/log"
)

var errQuery = errors.New("error query key")

type IKeyRepository interface {
	FindStreamKey(name, key string) (*model.Keys, error)
}

type keysRepository struct {
	*sql.DB
}

func NewKeysReposiroy(db *sql.DB) IKeyRepository {
	return &keysRepository{
		db,
	}
}

func (r *keysRepository) FindStreamKey(name, key string) (*model.Keys, error) {
	fmt.Println("========= Loooking for:", name, key)
	keys := &model.Keys{}
	row := r.QueryRow(`SELECT * FROM "Lives" WHERE "name"=$1 AND "stream_key"=$2`, name, key)

	err := row.Scan(&keys.Name, &keys.Key)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Keys{}, nil
		}

		return &model.Keys{}, errQuery
	}

	return keys, nil
}
