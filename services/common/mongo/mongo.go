package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	log "github.com/sirupsen/logrus"
)

var (
    Client *mongo.Client

    ItemCollection *mongo.Collection
)

func Connect() error {
    log.Info("Connecting to database...")

    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongo_user:mongo_password@mongo:27017"))

    if err != nil {
		return  err
    }

    if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return  err
    }

    log.Info("Connected to database")

	Client = client

    ItemCollection = CreateCollection(Client, "go_demo", "items")

    return nil
}

func Disconnect() error {
    log.Info("Disconnecting to database...")

    err := Client.Disconnect(context.Background())

    log.Info("Disconnected from database")

	return err
}

func CreateCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
    mongoCollection := client.Database(database).Collection(collection)

    log.Info("Created collection")

    return mongoCollection
}

func Add(collection *mongo.Collection, document interface{}) (primitive.ObjectID, error) {
    id, err := collection.InsertOne(context.Background(), document)

    if err != nil {
		// TODO: Fix.
		// return nil, err
    }

    return id.InsertedID.(primitive.ObjectID), nil
}

func Delete(collection *mongo.Collection, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})

    if err != nil {
		return nil, err
    }

    return result, nil
}

func Find(collection *mongo.Collection) (*mongo.Cursor, error) {
    result, err := collection.Find(context.Background(), bson.M{})

    if err != nil {
		return nil, err
    }

    return result, nil
}
