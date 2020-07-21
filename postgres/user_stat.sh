#!/bin/sh

psql -v USER_ID="$1" -U lingualynda -d lingualynda <<-EOSQL
    SELECT
      user_id,
      date(timestamp) as date,
      count(1) as answers_count,
      round(100. * sum(correctly_translated::int) / count(1), 1)::text || '%' as correct_answers_rate
    FROM usermemorizedterm
    WHERE user_id = :USER_ID
    GROUP BY user_id, date
    ORDER BY date;
EOSQL
