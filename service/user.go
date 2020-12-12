package service

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserCreate Create
func UserCreate(input model.NewUser) (*model.User, error) {
	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")

	user := model.User{
		Name:      input.Name,
		CreatedAt: &timeNow,
	}

	result, err := client.Database("training").Collection("user").InsertOne(ctx, user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return &user, nil
}

//UserCreateBatch Create Batch
func UserCreateBatch(input []*model.NewUser) ([]*model.User, error) {
	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	var interfaceBatch []interface{}
	var userBatch []*model.User
	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")

	for _, val := range input {
		user := model.User{
			Name:      val.Name,
			CreatedAt: &timeNow,
		}
		interfaceBatch = append(interfaceBatch, &user)
		userBatch = append(userBatch, &user)
	}

	result, err := client.Database("training").Collection("user").InsertMany(ctx, interfaceBatch)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for index, val := range result.InsertedIDs {
		userBatch[index].ID = val.(primitive.ObjectID).Hex()
	}

	return userBatch, nil
}

//UserUpdate Update
func UserUpdate(input model.UpdateUser) (*model.User, error) {
	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")

	objectID, _ := primitive.ObjectIDFromHex(input.ID)

	filter := bson.M{
		"_id": objectID,
	}

	user := model.User{
		ID:        input.ID,
		Name:      input.Name,
		UpdatedAt: &timeNow,
	}

	_, err := client.Database("training").Collection("user").UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: input.Name},
			{Key: "updated_at", Value: &timeNow},
		}},
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err != nil {
		log.Println(err)
		panic(err)
	}

	return &user, nil
}

//UserDelete Delete
func UserDelete(id string) (string, error) {

	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	_, err = client.Database("training").Collection("user").DeleteOne(ctx, filter, options.Delete())

	if err != nil {
		fmt.Println(err)
		return "Fail", err
	}

	return "Success", nil
}

// //UserDeleteAll Delete All
// func UserDeleteAll() (string, error) {

// 	client := config.ConnectMongo()
// 	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
// 	defer cancel()

// 	filter := bson.M{}

// 	res, err := client.Database("training").Collection("user").DeleteOne(ctx, filter, options.DeleteOptions)

// 	if err != nil {
// 		fmt.Println(err)
// 		return "Fail", err
// 	}

// 	fmt.Println(res.DeletedCount)

// 	return "Success", nil
// }

//UserGetByID Get By ID
func UserGetByID(id string) (*model.User, error) {
	var user model.User

	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	err = client.Database("training").Collection("user").FindOne(ctx, filter, options.FindOne()).Decode(&user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserGetAll Get All
func UserGetAll() ([]*model.User, error) {

	var users []*model.User

	client := config.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()

	cursor, err := client.Database("training").Collection("user").Find(ctx, bson.D{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		//Create a value into which the single document can be decoded
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, &user)

	}

	return users, nil
}
