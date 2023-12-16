
# goFr CRUD Application
A REST API built using [GoFr](https://gofr.dev/). It demonstrates all the CRUD functions.

It is a movie watchlist management application, where you can add, edit, and delete movies that you want to watch. You can fetch movies by name or genre.
## Prerequisites

This project requires [GoLang](https://go.dev/dl/), [GoFr](https://gofr.dev/) and [MySQL](https://dev.mysql.com/downloads/mysql/).

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/t4nm4y/testProject.git
    cd testPoject
    ```

2. Setting up the database:

-  cd in the "db" directory and run these commands:

    ```go
    docker build -t custom-mysql-image .
    docker run -d --name mysql-container -p 3306:3306 custom-mysql-image
    ```
    This will setup a SQL database with temporary data for testing.

3. Run the main.go file:

   cd in the testProject root directory and run this command:
    ```go
    go run main.go
    ```
    It will run on [localhost:8000](http://localhost:8000/)
    
5. At last, Cleanup:
    
    To stop and remove the Docker container, cd in the "db" directory and run:
   
    ```go
    docker stop mysql-container
    docker rm mysql-container
    ```
## API References
A Postman collection is provided in this repo for demonstrating all of these APIs.

Base URL: [localhost:8000](http://localhost:8000/)

You can use Postman to test all these API endpoints.

#### 1. To Fetch all the movies

```
  GET /all
```
#### 2. To Fetch all the movies of a particular genre

```
  GET /movies/genre/
```
In Postman in `Params` set `Key` as "genre" and `Value` as "genre_name" (the genre that you want to search)

#### 3. To Fetch all the movies of a particular language

```
  GET /movies/language/
```
In Postman in `Params` set `Key` as "language" and `Value` as "language_name" (the language you want to search).

#### 4. Add a Movie

```
  POST /add
```

| JSON Key | Type     | Required |
| :-------- | :------- | :------- |
| `name` | `string` | **Required**
| `language` | `string` | - |
| `genre` | `string` | - |

eg: in Postman in the `Body` select `raw` and then put this:
```json
{
    "name": "3 Idiots",
 "language": "Hindi",
    "genre": "Comedy"
}
```

#### 5. Edit/update the details of a movie

```
  PUT /update
```

| JSON Key | Type     | Required |
| :-------- | :------- | :------- |
| `name` | `string` | **Required** |
| `language` | `string` | - |
| `genre` | `string` | - |

eg: in Postman in the `Body` select `raw` and then put this if you want to change the genre:
```json
{
    "name": "3 Idiots",
    "genre": "Drama"
}
```

#### 6. Delete a Movie

```
  DELETE /delete
```
In Postman in `Params` set `Key` as "name" and `Value` as "movie_name" (the name of the movie you want to delete).

### Made with ❤️ by Tanmay Kumar.

