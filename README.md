# Palu Covid


## Dependency
- [Golang Version 1.13](https://golang.org/doc/install)
- [Postgres Version 11 or above](https://www.postgresql.org/download/)


## How to Build
```
make build
```

## How to Run
Add your config to `sample.config.yml` then run
```
make copy-config
```

For first time run, create the database and set up migration
```
createdb "covid"
./out/palu-covid migrate
```

Run the service
```
./out/palu-covid server
```

## Sample Request via Postman
Import these collections to postman
```
https://www.getpostman.com/collections/35fe323d100ea4803c57
```
