package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoOpener defines what is needed to instantiate a new mongo.Client. It doesn't
// implement methods from stoppage.Stopper interface because they should be defined
// in a Entity Repository and implemented by its Data Sources.
type MongoOpener interface {
	Open(context.Context) (*mongo.Client, error)
}
