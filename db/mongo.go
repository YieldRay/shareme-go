package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	coll *mongo.Collection
}

func (this mongodb) Get(namespace string) (content string, err error) {
	var result map[string]string

	err = this.coll.FindOne(context.TODO(), bson.D{{Key: "namespace", Value: namespace}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		return
	}
	return result["data"], nil

}

func (this mongodb) Set(namespace, content string) bool {
	var err error

	this.coll.UpdateOne(
		context.TODO(),
		bson.D{{Key: "namespace", Value: namespace}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "namespace", Value: namespace}, {Key: "data", Value: content}}}},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return false
	}
	return true
}

func MongoDB(srv, dbName, colName string) (db IDB) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(srv))

	if err != nil {
		panic(err)
	}
	
	coll := client.Database(dbName).Collection(colName)
	return mongodb{coll}

}
