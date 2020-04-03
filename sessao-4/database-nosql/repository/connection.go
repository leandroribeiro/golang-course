package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Db *mongo.Database

const (
	// Name of the database.
	DBName = "GoCourse"
	URI    = "mongodb://mestre:siga-o-mestre@127.0.0.1:27017"
)

/*func GetClient() *mongo.Client  {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	return client
}*/

func Open() (err error) {
	// Base context.
	ctx := context.Background() // Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return
	}

	Db = client.Database(DBName)
	fmt.Println(Db.Name())

	return
}
