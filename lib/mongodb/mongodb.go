package mongodb

import (
	"context"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Cli *mongo.Client

func Initialize(config config.Mongodb) {
	if config.Uri != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
		if err != nil {
			panic(err)
		}

		Cli = client

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = Cli.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(err)
		}
	}
}
