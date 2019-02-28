build:
	docker build -t go.playground:0.1 .

run:
	docker run -p 8080:8080 --rm -it go.playground:0.1

stop:
	docker stop go.playground:0.1

lint:
	go vet

test:
	go test