package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoConnection struct {
	Client *mongo.Client
	Db *mongo.Database
	Collection *mongo.Collection
	Connected bool
}

// TODO: Pass context as parameter
// TODO: Logs
func NewConnection(dbHostUri, dbName, collectionName string) MongoConnection {
	// Set client options
	opts := options.Client()
	opts.ApplyURI(dbHostUri)
	opts.SetServerSelectionTimeout(1 * time.Second)
    opts.SetConnectTimeout(1 * time.Second)
	
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(fmt.Sprintf("Error: couldn't connect to host '%s'.", dbHostUri)) // Log
	}
    
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(fmt.Sprintf("Error: unreachable host '%s'.", dbHostUri)) // Log
    }

	fmt.Println("Connected to MongoDB.") // Log

	db, err := getDatabase(*client, dbName)

	if err != nil {
		panic(err) // Log
	}

	fmt.Printf("Connected to database '%s'.\n", dbName) // Log

	collection, err := getCollection(*db, collectionName)

	if err != nil {
		panic(err) // Log
	}

	fmt.Printf("Connected to collection '%s'.\n", collectionName) // Log

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
