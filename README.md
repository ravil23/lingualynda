![CI](https://github.com/ravil23/lingualynda/workflows/CI/badge.svg?branch=master)
![CD](https://github.com/ravil23/lingualynda/workflows/CD/badge.svg?branch=master)

# Lingua Lynda
AI powered dialog system for learning languages

## Requirements
- `Docker Compose`

## Running
Specify `BOT_TOKEN` environment variable and run next command:
```
docker-compose up --build -d
```

## Testing on localhost
Script `./testenv.sh` runs minimal necessary development environment on localhost. It is useful for debugging and writing integration tests.