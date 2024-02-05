package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	log "github.com/sirupsen/logrus"
)

var (
    client         *mongo.Client
    ItemCollection *mongo.Collection
)

func Connect() error {
    log.Info("Connecting to MongoDB...")

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	database := os.Getenv("MONGO_DATABASE")

    client_, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)))

    if err != nil {
		return  err
    }

	client = client_

    if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return  err
    }

    log.Info("Connected to MongoDB")

    ItemCollection = CreateCollection(database, "items")

    return nil
}

func Disconnect() error {
    log.Info("Disconnecting from MongoDB...")

    err := client.Disconnect(context.Background())

    log.Info("Disconnected from MongoDB")

	return err
}

func CreateCollection(database string, collection string) *mongo.Collection {
    mongoCollection := client.Database(database).Collection(collection)

    log.Info("Created collection")

    return mongoCollection
}

func Add(collection *mongo.Collection, document interface{}) (primitive.ObjectID, error) {
    id, err := collection.InsertOne(context.Background(), document)

    if err != nil {
		return primitive.ObjectID{}, err
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
