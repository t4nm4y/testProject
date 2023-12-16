
# goFr CRUD Application
A REST API built using [GoFr](https://gofr.dev/). It demonstrates all the CRUD functions.
## Prerequisites

This project requires [GoLang](https://go.dev/dl/), [GoFr](https://gofr.dev/) and [MySQL](https://dev.mysql.com/downloads/mysql/).

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/t4nm4y/testProject.git
    cd your-repository
    ```

2. Setting up the database:

-  Go to the "db" directory and run these commands:

    ```go
    docker build -t custom-mysql-image .
    docker run -d --name mysql-container -p 3306:3306 custom-mysql-image
    ```
    This will setup a SQL database with temporary data for testing.

3. Run the main.go file:
    ```go
    go run main.go
    ```
    
4. At last, Cleanup:
    
    To stop and remove the Docker container, run:

    ```go
    docker stop mysql-container
    docker rm mysql-container
    ```


