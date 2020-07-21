![CI](https://github.com/ravil23/lingualynda/workflows/CI/badge.svg?branch=master)

# Lingua Lynda
AI powered dialog system for learning languages

## Requirements
- `Go v1.13`
- `Docker Compose` (optional)

## Running
Specify `BOT_TOKEN` environment variable and run next command:
```
docker-compose up --build -d
```

## Testing on localhost
Run development environment on localhost:
```
docker-compose up --build postgres pgadmin
```
It is useful for debugging and writing integration tests.

## Show reports
Get list of available reports:
```
docker exec postgres_container ls /reports
```
Show list of users:
```
docker exec postgres_container /reports/users_list.sh
```
Show statistic of user:
```
docker exec postgres_container /reports/user_stat.sh <USER_ID>
```
