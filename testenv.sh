#!/bin/sh

docker-compose build postgres
docker-compose run -p 5432:5432 postgres
