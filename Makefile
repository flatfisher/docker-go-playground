build:
	docker-compose build

run:
	go run ./

lint:
	go vet

test:
	go test -v