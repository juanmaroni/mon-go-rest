package mongodb

import (
	"context"
	"errors"
	"fmt"
	"mon-go-rest/config"
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

func NewConnection(ctx context.Context, dbName, collectionName string) MongoConnection {
	logger := config.Logger
	mongoDbUri := config.MongoDb.Uri

	// Set client options
	opts := options.Client()
	opts.ApplyURI(mongoDbUri)
	opts.SetServerSelectionTimeout(1 * time.Second)
    opts.SetConnectTimeout(1 * time.Second)
	
	client, err := mongo.Connect(opts)

	if err != nil {
		logger.Error(fmt.Sprintf("Error: couldn't connect to host '%s'.", mongoDbUri))
	}
    
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
        logger.Error(fmt.Sprintf("Error: unreachable host '%s'.", mongoDbUri))
    }

	logger.Info("Connected to MongoDB.")

	db, err := getDatabase(ctx, *client, dbName)

	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info(fmt.Sprintf("Connected to database '%s'.", dbName))

	collection, err := getCollection(ctx, *db, collectionName)

	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info(fmt.Sprintf("Connected to collection '%s'.", collectionName))

	return MongoConnection {
		Client: client,
		Db: db,
		Collection: collection,
		Connected: true,
	}
}

func (conn *MongoConnection) CloseConnection(ctx context.Context) {
	logger := config.Logger

	if err := conn.Client.Disconnect(ctx); err != nil {
		logger.Error(err.Error())
	}

	conn.Connected = false
	logger.Info("Disconnected from MongoDB.")
}

func GetRecordById[T interface{}](ctx context.Context, coll mongo.Collection, idField string, value int) (T, error) {
	var record T

	if err := coll.FindOne(ctx, bson.D{{Key: idField, Value: value}}).Decode(&record); err != nil {
		return record, err
	}

	return record, nil
}

func GetAllRecords[T interface{}](ctx context.Context, coll mongo.Collection) ([]T, error) {
	var records []T
	cursor, err := coll.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func getDatabase(ctx context.Context, client mongo.Client, dbName string) (*mongo.Database, error) {
	dbList, err := client.ListDatabaseNames(ctx, bson.M{"name": dbName})

	if err != nil || len(dbList) != 1 || dbName != dbList[0] {
		return nil, errors.New(fmt.Sprintf("Error: couldn't find database '%s'.", dbName))
	}

	return client.Database(dbName), nil
}

func getCollection(ctx context.Context, db mongo.Database, collectionName string) (*mongo.Collection, error) {
	collectionList, err := db.ListCollectionNames(ctx, bson.M{"name": collectionName})

	if err != nil || len(collectionList) != 1 || collectionName != collectionList[0] {
		return nil, errors.New(fmt.Sprintf("Error: couldn't find collection '%s'.", collectionName))
	}

	return db.Collection(collectionName), nil
}
