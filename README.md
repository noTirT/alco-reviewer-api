# Alco Reviewer API

## Response status meaning

| HTTP Code | Response Status Code | Definition                                                                                                    |
| --------- | -------------------- | ------------------------------------------------------------------------------------------------------------- |
| 400       | 1                    | The body of the request does not match the requirements                                                       |
|           | 2                    | Missing or malformed Authorization header                                                                     |
|           | 3                    | User ID provided in the access token does not exist in the database                                           |
|           | 4                    | Custom control key provided in the access token does not match the key saved with the User ID in the database |
| 401       | 1                    | Provided access token is expired                                                                              |
|           | 2                    | Provided token is invalid                                                                                     |
| 500       | 1                    | An error occured while parsing the JSON body of a http.Request into a golang struct                           |

## Environment Variables

| Variable           | Type   | Definition                           | Default     |
| ------------------ | ------ | ------------------------------------ | ----------- |
| SERVER_PORT        | String | Port the server runs on              | 8080        |
| PG_HOST            | String | Postgres host                        | localhost   |
| PG_PORT            | Number | Postgres port                        | 5432        |
| PG_USER            | String | Postgres username                    | postgres    |
| PG_PASSWORD        | String | Postgres password                    | password    |
| PG_DATABASE        | String | Postgres database name               | alcohol-app |
| JWT_EXPIRATION     | Number | JWT Token expiration time in minutes | 30          |
| PEM_KEY_EXPIRATION | Number | RSA Key expiration time in days      | 14          |

## Run locally

### Without ngrok

- Run `sh scripts/start_docker_postgres.sh`
- Run `go run main.go`

### Using ngrok

- Run `ngrok http http://localhost:8080` to enable tunneling for your backend
- The API can be called through the URL provided in the output of the command

### Using docker

- Run `scripts/build_image.sh` to build the docker image for the backend
- Run `scripts/start_docker_postgres.sh` to start a postgres instance in docker
- Run `scripts/start_docker_backend.sh` to start a docker container running the golang backend

### Using docker-compose

- Run `docker compose up <-d>` to start the application

### Rerun init.sql file on postgres in docker container

- Enter cli of postgres container with `docker exec -it backend-db-1 /bin/bash`
- `psql -U postgres -d alcohol-app -a -f "/docker-entrypoint-initdb.d/init.sql"`
