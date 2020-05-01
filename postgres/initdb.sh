#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 -U postgres <<-EOSQL
    CREATE USER lingualynda WITH PASSWORD '$POSTGRES_PASSWORD';
    CREATE DATABASE lingualynda;
    GRANT ALL PRIVILEGES ON DATABASE lingualynda TO lingualynda;
EOSQL
