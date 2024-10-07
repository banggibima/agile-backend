package mongo

import (
	"github.com/banggibima/agile-backend/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(config *config.Config) (*mongo.Client, error) {
	options := options.Client().ApplyURI(config.Mongo.URI)

	client, err := mongo.Connect(options)
	if err != nil {
		return nil, err
	}

	return client, nil
}
