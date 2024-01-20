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

    TransitCollection *mongo.Collection
    UserCollection    *mongo.Collection
    VehicleCollection *mongo.Collection
)

func Init(uri string) error {
	client, err := Connect("mongodb://mongo_user:mongo_password@mongo:27017")

	if err != nil {
		return err
	}

	Client = client

    TransitCollection = CreateCollection(Client, "traffic", "transits")
    UserCollection = CreateCollection(Client, "traffic", "users")
    VehicleCollection = CreateCollection(Client, "traffic", "vehicles")

	return nil
}

func Deinit() error {
    return Client.Disconnect(context.TODO())
}

func Connect(uri string) (*mongo.Client, error) {
    log.Info("Connecting to database...")

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

    if err != nil {
		return nil, err
    }

    if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
    }

    log.Info("Connected to database")

    return client, nil
}

func Disconnect(client *mongo.Client) error {
    log.Info("Disconnecting to database...")

    err := client.Disconnect(context.TODO())

    log.Info("Disconnected from database")

	return err
}

func CreateCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
    mongoCollection := client.Database(database).Collection(collection)

    log.Info("Created collection")

    return mongoCollection
}

func Add(collection *mongo.Collection, document interface{}) (primitive.ObjectID, error) {
    id, err := collection.InsertOne(context.TODO(), document)

    if err != nil {
		// TODO: Fix.
		// return nil, err
    }

    return id.InsertedID.(primitive.ObjectID), nil
}

func Delete(collection *mongo.Collection, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

    if err != nil {
		return nil, err
    }

    return result, nil
}

func Find(collection *mongo.Collection) (*mongo.Cursor, error) {
    result, err := collection.Find(context.TODO(), bson.M{})

    if err != nil {
		return nil, err
    }

    return result, nil
}
