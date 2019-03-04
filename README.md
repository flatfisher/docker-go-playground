# docker-go-playground

## Database

```
$ docker-compose build
$ docker-compose up
$ mysql -h 127.0.0.1 -P 3306 -u root -p < sakila-schema.sql
$ mysql -h 127.0.0.1 -P 3306 -u root -p < sakila-data.sql
```

## Test

```
$ go test -v
```