#!bin/bash
docker run -d \
    --name alcohol-app-backend \
    -p 8080:8080 \
    --env-file ../.env.docker \
    --network alcohol-app-network \
    alcohol-app-backend:latest
