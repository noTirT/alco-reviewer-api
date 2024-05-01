#!bin/bash
docker run -d \
    --name backend-db-1 \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=alcohol-app \
    -v /home/tom/Code/alcohol-app/backend/data/postgres:/var/lib/postgresql/data \
    -v /home/tom/Code/alcohol-app/backend/configs/init.sql:/docker-entrypoint-initdb.d/init.sql \
    --network alcohol-app-network \
    -p 5432:5432 \
    postgres
