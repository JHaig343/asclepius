package ipynbparser

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//MakeConnection Connect to MongoDB Atlas
func MakeConnection() (ctx context.Context, cli *mongo.Client) {
	// set up client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + UserName + ":" + DatabasePasswd + "@cluster0-uunik.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	// Set timeout to 10 seconds
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	// TODO: context cancelled as soon as function exits with call below
	// defer cancel()
	//Connect to client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, client
}

//TestPing pings the cluster as a form of health check
func TestPing(ctx context.Context, client *mongo.Client) {
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

//ListDBs lists the databases in mongo cluster
func ListDBs(ctx context.Context, client *mongo.Client) {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

//CloseConnection closes the connection when done
func CloseConnection(ctx context.Context, client *mongo.Client) {
	//This will defer closing connection until main() exits
	defer client.Disconnect(ctx)
}

//InsertNotebook inserts a Notebook object into the MongoDB Atlas collection
func InsertNotebook(ctx context.Context, client *mongo.Client, nb Notebook) {
	//first get the collection we want to insert into
	collection := client.Database(DatabaseName).Collection("ipynbparser")
	res, err := collection.InsertOne(ctx, nb)
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}

//RetrieveNotebook currently retrieves the first notebook object it finds from the database
func RetrieveNotebook(ctx context.Context, client *mongo.Client) Notebook {
	//set context timeout to 30 sec
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database(DatabaseName).Collection("ipynbparser")
	var nb Notebook
	//For now just get first object
	err := collection.FindOne(ctx, bson.D{}).Decode(&nb)
	if err != nil {
		log.Fatal(err)
	}
	return nb

}
