package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteFile(filename string) error {
	filter := bson.M{"_id": filename}

	_, err := fileCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
