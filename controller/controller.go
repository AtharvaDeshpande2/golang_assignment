package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github/atharvadeshpande/mongoapi/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017" //to be filled
const dbName = "Item"
const colName = "list"
const colName2 = "Register Users"

var collection *mongo.Collection
var xyz *mongo.Collection

//connecting with database

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)
	xyz = client.Database(dbName).Collection(colName2)

	fmt.Println("Collection instance is ready")
}

//helper (sep files)

func insertOneItem(item model.Items) {
	inserted, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one item", inserted.InsertedID)

}

func insertOneUser(item model.RegisterUser) {
	inserted, err := xyz.InsertOne(context.Background(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one item", inserted.InsertedID)
}

//update one record helper

func updateOneItem(itemId string, quant int) {
	id, _ := primitive.ObjectIDFromHex(itemId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"quantity": quant}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)

}

// delete
func deleteOneItem(itemId string) {
	id, err := primitive.ObjectIDFromHex(itemId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item Deleted ", deleteCount)
}

//all delete

func deleteAllItem() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of items deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all items

func getAllItems() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var items []primitive.M

	for cur.Next(context.Background()) {
		var item bson.M
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	defer cur.Close(context.Background())
	return items

}

func getAllUsers() []primitive.M {
	cur, err := xyz.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var users []primitive.M

	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer cur.Close(context.Background())
	return users

}

// controller functions
func GetMyAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allItems := getAllItems()
	json.NewEncoder(w).Encode(allItems)

}
func GetMyAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allUsers := getAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func CreateItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var item model.Items
	_ = json.NewDecoder(r.Body).Decode(&item)
	insertOneItem(item)
	json.NewEncoder(w).Encode(item)
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.RegisterUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertOneUser(user)
	json.NewEncoder(w).Encode(user)
}
func UpdatedItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var item model.Items
	json.NewDecoder(r.Body).Decode(&item)
	params := mux.Vars(r)
	updateOneItem(params["id"], item.Quantity)
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteOneItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneItem(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllItem()
	json.NewEncoder(w).Encode(count)
}
