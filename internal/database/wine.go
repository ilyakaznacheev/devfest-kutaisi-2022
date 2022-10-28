package database

import (
	"context"

	"github.com/ilyakaznacheev/devfest-kutaisi-2022/internal/model"
)

func (db *Database) AddWine(ctx context.Context, id string, w model.Wine) error {
	data, err := db.marshal(&w)
	if err != nil {
		return err
	}

	return db.client.HSet(ctx, db.collection, id, data).Err()
}

func (db *Database) GetWine(ctx context.Context, id string) (*model.Wine, error) {
	data, err := db.client.HGet(ctx, db.collection, id).Result()
	if err != nil {
		return nil, err
	}

	var w model.Wine
	err = db.unmarshal(data, &w)
	return &w, err
}

func (db *Database) GetWineList(ctx context.Context) (map[string]model.Wine, error) {
	data, err := db.client.HGetAll(ctx, db.collection).Result()
	if err != nil {
		return nil, err
	}

	res := make(map[string]model.Wine, len(data))

	for k, v := range data {
		var w model.Wine
		if err := db.unmarshal(v, &w); err != nil {
			return nil, err
		}
		res[k] = w
	}

	return res, nil
}
