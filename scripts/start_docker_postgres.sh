#!bin/bash
docker run -d \
    --name alcohol-app \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=alcohol-app \
    -v /home/tom/Code/alcohol-app/backend/data/postgres:/var/lib/postgresql/data \
    -v /home/tom/Code/alcohol-app/backend/configs/init.sql:/docker-entrypoint-initdb.d/init.sql \
    -p 5432:5432 \
    postgres
