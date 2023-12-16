package routes

import (
	"gofr.dev/pkg/gofr"
)

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
}

func all(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM movies;")

	if err != nil {
		return nil, err
	}
    var movies []Movie

	// Iterate over the rows and scan data into Movie objects
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Language, &movie.Genre)
		if err != nil {
			return nil, err
		}
		// Append the Movie object to the slice
		movies = append(movies, movie)
	}
	return movies, nil
}


