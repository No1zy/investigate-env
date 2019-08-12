#!/bin/bash

URL=$1
LANGS=("go" "java" "php" "python" "ruby" "perl")

go run main.go $URL
docker-compose up --build > /dev/null
for lang in ${LANGS[@]}; do
    docker-compose logs $lang
done
yes | docker-compose rm -v
