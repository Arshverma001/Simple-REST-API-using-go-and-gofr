package main

import (
	"fmt"

	"github.com/Arshverma001/controller"
	"gofr.dev/pkg/gofr"
)

func main() {
	fmt.Println("Routes")
	app := gofr.New()
	app.GET("/api/movies", controller.GettheMovies)
	app.GET("/api/movie/{id}", controller.GetMovieId)
	app.POST("/api/movie", controller.InsertSingleMovie)
	app.PUT("/api/movie/{id}", controller.UpdateSingleMovie)
	app.DELETE("/api/movie/{id}", controller.DeleteSingleMovie)
	app.Start()

}
