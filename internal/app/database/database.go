package database

import (
	"context"
	"encoding/json"
	"fmt"
	"grisko-test-app/config"
	"grisko-test-app/internal/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func New(Database string, Collection string) (*DataBase, error) {

	uri := config.DataBaseUri

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	fmt.Println(client.ListDatabases(ctx, bson.D{{}}))

	coll := client.Database(Database).Collection(Collection)
	return &DataBase{
		client: client,
		coll:   coll,
	}, nil
}

func (db *DataBase) InsertOne(product *models.BsonProduct) (*mongo.InsertOneResult, error) {

	doc, err := db.coll.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (db *DataBase) GetAll(filter *models.BsonProduct) (*models.BsonProductList, error) {

	var results models.BsonProductList

	cursor, err := db.coll.Find(context.TODO(), filter)
	if (*filter == models.BsonProduct{}) {
		cursor, err = db.coll.Find(context.TODO(), bson.D{{}})
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (db *DataBase) DeleteProduct(id string) *models.BsonProduct {

	var deleted *models.BsonProduct
	PrimID, _ := primitive.ObjectIDFromHex(id)
	res := db.coll.FindOneAndDelete(context.TODO(), bson.D{{Key: "_id", Value: PrimID}})
	res.Decode(&deleted)

	return deleted
}

func (db *DataBase) UpdateProduct(id string, updates string) *models.BsonProduct {

	var bdoc *bson.D
	var updated *models.BsonProduct

	bson.UnmarshalExtJSON([]byte(updates), true, &bdoc)

	PrimID, _ := primitive.ObjectIDFromHex(id)
	update := bson.D{{Key: "$set", Value: bdoc}}

	res := db.coll.FindOneAndUpdate(
		context.TODO(), bson.D{{Key: "_id", Value: PrimID}}, update,
	)
	res.Decode(&updated)

	return updated
}
