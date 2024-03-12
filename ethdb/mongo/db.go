package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	client *mongo.Client
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

	return &Db{client}

}

func (db *Db) InsertOne() error {

}

func (db *Db) FindOne() error {
	//coll := client.Database("newtest").Collection("myCollection")
	//title := "wrdup"
	//var result bson.M
	//err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	//if err == mongo.ErrNoDocuments {
	//fmt.Printf("No document was found with the title %s\n", title)
	//return
	//}
	//if err != nil {
	//panic(err)
	//}
	//jsonData, err := json.MarshalIndent(result, "", "    ")
	//if err != nil {
	//panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)

}
