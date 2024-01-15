package db

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

func Init(uri string) {
	Client = Connect("mongodb://mongo_user:mongo_password@mongo:27017")

	TransitCollection = CreateCollection(Client, "traffic", "transits")
	UserCollection = CreateCollection(Client, "traffic", "users")
	VehicleCollection = CreateCollection(Client, "traffic", "vehicles")
}

func Deinit() {
	err := Client.Disconnect(context.TODO())

	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func Connect(uri string) *mongo.Client {
	log.Println("Connecting to database...")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Error(err)
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Error(err)
		panic(err)
	}

	log.Println("Connected to database")

	return client
}

func Disconnect(client *mongo.Client) {
	log.Println("Disconnecting to database...")

	client.Disconnect(context.TODO())

	log.Println("Disconnected from database")
}

func CreateCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
	mongoCollection := client.Database(database).Collection(collection)

	log.Println("Created collection")

	return mongoCollection
}

func Add(collection *mongo.Collection, document interface{}) primitive.ObjectID {
	id, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	return id.InsertedID.(primitive.ObjectID)
}

func Delete(collection *mongo.Collection, id primitive.ObjectID) *mongo.DeleteResult {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	log.Printf("Deleted %d rows", result.DeletedCount)

	return result
}

func Find(collection *mongo.Collection) *mongo.Cursor {
	result, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Error(err)
		panic(err)
	}

	return result
}
