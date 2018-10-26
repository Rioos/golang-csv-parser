# Neoway CSV Parser

A fast and simple parser written in Go for Neoway CSV files.

## Features

- [x] Dockerized service
- [x] Saves on a Postgres Database
- [x] Faster than ~~lightspeed~~ a bullet

## What do you need

- [x] Docker and Docker-Compose

## How to run

```
git clone git@github.com:Rioos/golang-csv-parser.git
cd golang-csv-parser
docker-compose up -d
```

## How to use

- Service will listen on port 80 and wait for a CSV/TXT file on the route '/csv' (must be form-data and must have the file as "file" input name)
- Database Server will listen on port 5432 and has set postgres/postgres as username/password
- Database name is neoway_csv 
