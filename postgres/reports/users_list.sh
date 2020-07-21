#!/bin/sh

psql -U lingualynda <<-EOSQL
    SELECT *
    FROM public.user
    ORDER BY created_at DESC;
EOSQL
