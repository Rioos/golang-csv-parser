# Neoway CSV Parser

A fast and simple parser written in Go for Neoway CSV files.

# Features

- [x] Dockerized service
- [x] Saves on a Postgres Database
- [x] Faster than ~~lightspeed~~ a bullet

# What do you need

- [x] Docker and Docker-Compose

# How to run

```
docker-compose up -d
```

- Service will listen on port 80 and waiting for a CSV/TXT file on route '/csv'
- Database Server will listen on port 5432 and has set postgres/postgres as username/password
- Database name is neoway_csv 
