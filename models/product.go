package models

import (
	"context"
	"errors"
	"time"

	"github.com/Ethical-Ralph/simple-crud-go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Price float32 `json:"price" bson:"price,omitempty"`
}


func NewProduct() *Product {
	return &Product{}
}


func handleError(err error) (*Product, error) {
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("document not found")
	} else if err != nil {
		return nil, err
	}    
	return nil, nil
}

func toObjectID(id string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil, errors.New("invalid id")
	}
	return &objectId, nil
}

func (p *Product) Save(d *database.Database) (*mongo.InsertOneResult, error) {
	collection := d.Database.Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	return collection.InsertOne(ctx, p)
}

func GetProduct(id string,d *database.Database) (*Product, error) {
	var result *Product

	objectId, err := toObjectID(id)
	if err != nil{
		return nil, err
	}

	filter := bson.D{{Key:"_id",Value: objectId}}

	collection := d.Database.Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return handleError(err)
	}

	return result, nil
}

func UpdateProduct(id string, data *Product, d *database.Database) (*Product, error) {
	var result *Product

	objectId, err := toObjectID(id)
	if err != nil{
		return nil, err
	}

	filter := bson.D{{Key:"_id",Value: objectId}}

	collection := d.Database.Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	update := bson.M{
        "$set": data,
    }

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&result)

	if err != nil {
		return handleError(err)
	}

	return result, nil
}

func DeleteProduct(id string, d *database.Database) (*Product, error) {
	objectId, err := toObjectID(id)
	if err != nil{
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	collection := d.Database.Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var result *Product
	err = collection.FindOneAndDelete(ctx, filter).Decode(&result)

	if err != nil {
		return handleError(err)
	}
	
	return result, nil
}

func GetProducts(d *database.Database) ([]Product, error) {
	collection := d.Database.Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	defer cur.Close(ctx)


	result := []Product{}

	for cur.Next(ctx) {
		var r Product
		err := cur.Decode(&r)
		if err != nil {
			return nil, err
		}

		result = append(result, r)	
	}

	return result, nil
}