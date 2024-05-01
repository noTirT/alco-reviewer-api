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

- `PG_HOST`: Host URL of the postgres database
- `PG_PORT`: Port of the postgres database on the host
- `PG_USER`: Username to log into the postgres database
- `PG_PASSWORD`: Password to log into the postgres database
- `PG_DATABASE`: Database used in the project

## Run locally

### Using ngrok

- Run `ngrok http http://localhost:8080` to enable tunneling for your backend
- The API can be called through the URL provided in the output of the command

### Using docker

- Run `scripts/build_image.sh` to build the docker image for the backend
- Run `scripts/start_docker_postgres.sh` to start a postgres instance in docker
- Run `scripts/start_docker_backend.sh` to start a docker container running the golang backend

### Using docker-compose

- Run `docker compose up <-d>` to start the application

## TODO

- Better error detection and handling (not always expired when verification fails)
