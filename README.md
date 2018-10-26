# Neoway CSV Parser

A fast and simple parser written in Go for Neoway CSV files.

## Features

- [x] Dockerized service
- [x] Saves on a Postgres Database
- [x] Faster than ~~lightspeed~~ a bullet

## What do you need

- [x] Docker and Docker-Compose [(install here)](https://docs.docker.com/install/)

## How to run

1 - Download the project
```
git clone git@github.com:Rioos/golang-csv-parser.git
cd golang-csv-parser
```

2 - (Optional) You can setup your prefered .env configuration modifying .env

3 - Make sure nothing is running on port 5432 or 80

4 - Run docker-compose
```
docker-compose up -d
```



## How to use

- Service will listen on port 80 and wait for a CSV/TXT file on the route '/csv' (must be form-data and must have the file as "file" input name)
- Database Server will listen on port 5432 and has set postgres/postgres as username/password by default
- Database name is neoway_csv by default

- You can send a request through Postman just like that

![Postman example](https://i.imgur.com/gmJJjHv.png)
