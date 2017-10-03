default: clean dependencies build-linux docker

clean:
	rm -rf api
dependencies:
	go get -v ./...
build-macos:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo  .
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
docker:
	docker build -t go-micro-tech-talk/flight-api .

