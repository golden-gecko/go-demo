package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"api/config"
)

var (
	Client *mongo.Client

	TransitCollection *mongo.Collection
	UserCollection    *mongo.Collection
	VehicleCollection *mongo.Collection
)

func Init(uri string) {
	Client = Connect(uri)

	TransitCollection = CreateCollection(Client, config.MONGO_DB_DATABASE, "transits")
	UserCollection = CreateCollection(Client, config.MONGO_DB_DATABASE, "users")
	VehicleCollection = CreateCollection(Client, config.MONGO_DB_DATABASE, "vehicles")
}

func Deinit() {
	err := Client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}

func Connect(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	return client
}

func CreateCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
	mongoCollection := client.Database(database).Collection(collection)

	log.Println("Created collection")

	return mongoCollection
}

func Add(collection *mongo.Collection, document interface{}) primitive.ObjectID {
	id, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		panic(err)
	}

	return id.InsertedID.(primitive.ObjectID)
}

func Find(collection *mongo.Collection) *mongo.Cursor {
	result, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	return result
}

func Delete(collection *mongo.Collection, id primitive.ObjectID) *mongo.DeleteResult {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	log.Printf("Deleted %d rows", result.DeletedCount)

	return result
}

func Disconnect(client *mongo.Client) {
	client.Disconnect(context.TODO())

	log.Println("Disconnected from database")
}
