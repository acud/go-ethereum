package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ethdb             string = "eth"
	collection_blocks        = "blocks"
)

type Db struct {
	client *mongo.Client
	db     string // the database we're using
}

func (db *Db) Close() error {
	return db.client.Disconnect(context.TODO())
}
func New() *Db {
	uri := "mongodb://localhost:27017/?authSource=admin"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &Db{client, ethdb}
}

func (db *Db) Insert(doc interface{}) ([]byte, error) {
	coll := db.client.Database(db.db).Collection("myCollection")
	res, err := coll.InsertOne()
	v := res.(primitive.ObjectID)
	var vv = [12]byte
	copy(vv, v)
	return vv, nil
}

func (db *Db) Write(coll string, val interface{}) ([]byte, error) {
	res, err := db.client.Database(db.db).Collection(coll).InsertOne(context.TODO(), val)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Document inserted with ID: %s\n", res.InsertedID)
	v := res.(primitive.ObjectID)
	var vv = [12]byte
	copy(vv, v)
	return vv, nil

}

func (db *Db) FindOne() error {
	coll := db.client.Database("newtest").Collection("myCollection")
	title := "wrdup"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

}
