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

### Running the API
```bash
go build
```
Then, run the executable built for your platform