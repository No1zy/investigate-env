#!/bin/bash

URL=$1
LANG=$2
TEMPLATE=$3
LANGS=("go" "java" "php" "python" "ruby" "perl" "node")

go run main.go --lang $LANG --template $TEMPLATE $URL
docker-compose up --build $LANG > /dev/null
for lang in ${LANGS[@]}; do
    docker-compose logs $lang
done
yes | docker-compose rm -v
sudo docker rmi $(sudo docker images | egrep '^<none>' | awk '{print $3}')
