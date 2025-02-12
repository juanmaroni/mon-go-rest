package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConnection struct {
	Client *mongo.Client
	Db *mongo.Database
	Collection *mongo.Collection
	Connected bool
}

// TODO: Pass context as parameter
// TODO: Pass Host URI as parameter
// TODO: Logs
func NewConnection(dbName, collectionName string) MongoConnection {
	const LocalhostUri string = "mongodb://localhost:27017"

	client, err := mongo.Connect(options.Client().ApplyURI(LocalhostUri))

	if err != nil {
		panic(fmt.Sprintf("Error: couldn't connect to '%s' found.", LocalhostUri)) // Log
	}

	db, err := getDatabase(*client, dbName)

	if err != nil {
		panic(err) // Log
	}

	collection, err := getCollection(*db, collectionName)

	if err != nil {
		panic(err) // Log
	}

	fmt.Println("Connected") // Log

	return MongoConnection {
		Client: client,
		Db: db,
		Collection: collection,
		Connected: true,
	}
}

func (conn *MongoConnection) CloseConnection() {
	if err := conn.Client.Disconnect(context.TODO()); err != nil {
		panic(err) // Log
	}

	conn.Connected = false
	fmt.Println(conn.Connected) // Log
}

func getDatabase(client mongo.Client, dbName string) (*mongo.Database, error) {
	dbList, err := client.ListDatabaseNames(context.TODO(), bson.M{"name": dbName})

	if err != nil || len(dbList) != 1 || dbName != dbList[0] {
		return nil, errors.New(fmt.Sprintf("Error: couldn't find database '%s'.", dbName))
	}

	return client.Database(dbName), nil
}

func getCollection(db mongo.Database, collectionName string) (*mongo.Collection, error) {
	collectionList, err := db.ListCollectionNames(context.TODO(), bson.M{"name": collectionName})

	if err != nil || len(collectionList) != 1 || collectionName != collectionList[0] {
		return nil, errors.New(fmt.Sprintf("Error: couldn't find collection '%s'.", collectionName))
	}

	return db.Collection(collectionName), nil
}
