package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/Arshverma001/model"
	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/errors"
)

const connectionString = "mongodb+srv://arshverma303:arsh30301@cluster0.g1tllll.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected succesfully")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// Get all movies
func GetAllMovies() ([]primitive.M, error) {
	var movies []primitive.M
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies, nil
}

// Add one movie
func InsertOneMovie(movie model.Netflix) error {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		return err
	}

	fmt.Println("Inserted 1 movie successfully", inserted.InsertedID)
	return nil
}

//Update one movie

// func UpdateOneMovie(movieId string) (int64, error) {
// 	id, err := primitive.ObjectIDFromHex(movieId)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"watched": true}}

// 	result, err := collection.UpdateOne(context.Background(), filter, update)
// 	// ans:=result.ModifiedCount

// 	if err != nil {
// 		return -1, err
// 	}

// 	fmt.Println("modified count:", result.ModifiedCount)
// 	return result.ModifiedCount, nil
// }

func UpdateOneMovie(movieId string, updateItems model.Netflix) (primitive.M, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	name := mapper.ConvertStructToBSONMap(updateItems, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	update := bson.M{"$set": name}
	filter := bson.M{"_id": id}

	result, err := collection.UpdateMany(context.Background(), filter, update)
	// ans:=result.ModifiedCount

	if err != nil {
		return nil, err
	}

	fmt.Println("modified count:", result.ModifiedCount)

	movie, err := GetMovieById(movieId)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

// Delete One movies
func DeleteOneMovie(movieId string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	// deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	// fmt.Println("Number of Movies deleted:", deleteCount)
	return deleteCount, nil

}

func GetMovieById(movieId string) (primitive.M, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.D{{Key: "_id", Value: id}}
	cursor, err := collection.Find(context.Background(), filter)

	fmt.Println("Id to be found is", movieId)

	if err != nil {
		fmt.Println("Mongo Db error", err)
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Println("Error:", err)
		}
	}()

	var movie primitive.M

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Error", err)
			return nil, err
		}
	}
	fmt.Println("Id that is be found is", movieId)
	fmt.Println("Id that is be found is", movie)

	if movie == nil {
		return nil, &errors.Response{
			Reason: "No movie with the given id",
		}
	}

	return movie, nil

}
