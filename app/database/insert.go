package database

import (
	"Biocad/app/model"
	"context"
)

func InsertFile(file model.File) error {
	_, err := fileCollection.InsertOne(context.Background(), file)
	if err != nil {
		return err
	}
	return nil
}