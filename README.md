# Alco Reviewer API

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

## TODO

- Automatically regenerate key pairs when they are not present
- Also regenerate them after a certain amount of time (create file with next generation date on creation)
- Better error detection and handling (not always expired when verification fails)
