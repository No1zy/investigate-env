#!/bin/bash

URL=$1

go run main.go $URL
docker-compose up --build > /dev/null
docker-compose logs go
docker-compose logs java
docker-compose logs php
yes | docker-compose rm -v
