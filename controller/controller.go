package controller

import (
	"encoding/json"

	"github.com/Arshverma001/model"
	"github.com/Arshverma001/mongo"
	"gofr.dev/pkg/gofr"
)

// Show all the Movies
func GettheMovies(ctx *gofr.Context) (interface{}, error) {
	movies, err := mongo.GetAllMovies()
	if err != nil {
		return nil, err
	}
	var allmovies []model.Netflix
	movieByte, _ := json.Marshal(movies)
	json.Unmarshal(movieByte, &allmovies)
	return allmovies, nil
}

//Insert one Movie

func InsertSingleMovie(ctx *gofr.Context) (interface{}, error) {
	var movie model.Netflix
	ctx.Bind(&movie)

	//insertion
	err := mongo.InsertOneMovie(movie)
	if err != nil {
		return nil, err
	}

	var allmovies model.Netflix
	movieByte, _ := json.Marshal(movie)
	json.Unmarshal(movieByte, &allmovies)
	return allmovies, nil
}

//Update data

func UpdateSingleMovie(ctx *gofr.Context) (interface{}, error) {
	var movies model.Netflix
	ctx.Bind(&movies)

	id := ctx.PathParam("id")

	// if movies.Movie == "" {
	// 	return nil, &errors.Response{
	// 		Reason: "Check the name of the fields",
	// 	}
	// }

	upMovie, err := mongo.UpdateOneMovie(id, movies)
	if err != nil {
		return nil, err
	}

	var allmovies model.Netflix
	movieByte, _ := json.Marshal(upMovie)
	json.Unmarshal(movieByte, &allmovies)
	return allmovies, nil
}

//Delete One Movie

func DeleteSingleMovie(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	numberOfDeleted, err := mongo.DeleteOneMovie(id)

	if err != nil {
		return nil, err
	}
	// fmt.Println("Number of Movies deleted:", numberOfDeleted)

	return numberOfDeleted, err
}

func GetMovieId(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	movie, err := mongo.GetMovieById(id)

	if err != nil {
		return nil, err
	}
	var themovie model.Netflix
	movieByte, _ := json.Marshal(movie)
	json.Unmarshal(movieByte, &themovie)
	return themovie, nil
}
