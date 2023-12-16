package routes

import (
	"fmt"
	"strings"
	"errors"
	"gofr.dev/pkg/gofr"
	"github.com/t4nm4y/testProject/model"
)


func AllMovies(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM movies")

	if err != nil {
		return nil, err
	}
    var movies []model.Movie

	// Iterate over the rows and scan data into Movie objects
	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Language, &movie.Genre)
		if err != nil {
			return nil, err
		}
		// Append the Movie object to the slice
		movies = append(movies, movie)
	}
	return movies, nil
}

func DeleteMovie(ctx *gofr.Context) (interface{}, error) {
	var name=ctx.Param("name")
	fmt.Println("data:", name)
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM movies WHERE name=?", name);
	if err != nil {
		return nil, err
	}

	//delete only returns null(status: 204) if no error
	return nil, nil
}

func AddMovie(ctx *gofr.Context) (interface{}, error) {
    var movie model.Movie
    if err := ctx.Bind(&movie); err != nil {
        return nil, err
    }

	if movie.Name == "" {
        return nil, errors.New("Name is required!")
    }

    _, err := ctx.DB().ExecContext(ctx, "INSERT INTO movies (name, language, genre) VALUES (?, ?, ?)", movie.Name, movie.Language, movie.Genre)
    if err != nil {
        return nil, err
    }
	message := fmt.Sprintf("%s added successfully", movie.Name)
    return message, nil
}

//you can only update the language or genre of the movie
func UpdateMovie(ctx *gofr.Context) (interface{}, error) {
    var movie model.Movie
    if err := ctx.Bind(&movie); err != nil {
        return nil, err
    }

	if movie.Name == "" {
        return nil, errors.New("Name is required!")
    }

    // Retrieve the existing movie details from the database
    existingMovie, err := getMovieByName(ctx, movie.Name)
    if err != nil {
        return nil, err
    }

    // Update the fields only if they are provided in the parameters
    if movie.Language != "" {
        existingMovie.Language = movie.Language
    }

    if movie.Genre != "" {
        existingMovie.Genre = movie.Genre
    }

    // Perform the update in the database
    _, err = ctx.DB().ExecContext(ctx, "UPDATE movies SET language=?, genre=? WHERE name=?", existingMovie.Language, existingMovie.Genre, existingMovie.Name)
    if err != nil {
        return nil, err
    }
	message := fmt.Sprintf("%s updated successfully", existingMovie.Name)
    return message, nil
}

// Helper function to retrieve a movie by name from the database
func getMovieByName(ctx *gofr.Context, name string) (*model.Movie, error) {
    var movie model.Movie
    err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM movies WHERE name=?", name).Scan(&movie.ID, &movie.Name, &movie.Language, &movie.Genre)
    if err != nil {
        return nil, err
    }
    return &movie, nil
}


//for fetching all the movies of a particular genre
func MoviesByGenre(ctx *gofr.Context) (interface{}, error) {
    genre := strings.ToLower(ctx.Param("genre")) // Convert to lowercase for case insensitive search

    rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM movies WHERE LOWER(genre) = ?", genre)
    if err != nil {
        return nil, err
    }

    var movies []model.Movie
    for rows.Next() {
        var movie model.Movie
        if err := rows.Scan(&movie.ID, &movie.Name, &movie.Language, &movie.Genre); err != nil {
            return nil, err
        }
        movies = append(movies, movie)
    }

    return movies, nil
}


//for fetching all the movies of a particular language
func MoviesByLanguage(ctx *gofr.Context) (interface{}, error) {
    language := strings.ToLower(ctx.Param("language")) // Convert to lowercase for case insensitive search


    rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM movies WHERE language=?", language)
    if err != nil {
        return nil, err
    }

    var movies []model.Movie
    for rows.Next() {
        var movie model.Movie
        if err := rows.Scan(&movie.ID, &movie.Name, &movie.Language, &movie.Genre); err != nil {
            return nil, err
        }
        movies = append(movies, movie)
    }

    return movies, nil
}



