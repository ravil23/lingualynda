#!/bin/sh

export POSTGRES_PASSWORD=test
docker-compose run -p 5432:5432 db
