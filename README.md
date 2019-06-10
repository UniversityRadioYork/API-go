# API-go
New implementation of API in Go, completely separated from MyRadio.

## Requirements
Before setup make sure to have these installed:
- Golang
- Postgres (Can be on networked machine)
- Memcached (Can be on networked machine)

## Setup
### Getting the code
```bash
go get github.com/UniversityRadioYork/API-go
```

### SQL setup
Setup your SQL database with the appropriate structure (.sql file to be made)

### Config
Change the config.toml.example file to config.toml and edit where necesary, it should look something like this:
```
host = "localhost"
port = "5432"
user = "username"
password = "password"
dbname = "dbname"
sslmode = "disable"
memcache = "localhost:11211"
```

### Running the API
```bash
go build
```
Then, run the executable built for your platform