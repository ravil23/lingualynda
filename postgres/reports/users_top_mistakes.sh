#!/bin/sh

USER_ID="$1"
LIMIT="${2:-10}"

if [ -z "$USER_ID" ]
  then
    echo "No USER_ID argument"
    exit 1
fi

psql -v USER_ID="$USER_ID" -v LIMIT="$LIMIT" -U lingualynda <<-EOSQL
    SELECT
      user_id,
      term,
      count(1) as mistakes_count
    FROM usermemorizedterm
    WHERE user_id = :USER_ID AND NOT coalesce(correctly_translated, false)
    GROUP BY user_id, term
    ORDER BY mistakes_count DESC
    LIMIT :LIMIT;
EOSQL
