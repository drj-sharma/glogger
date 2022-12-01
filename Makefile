build-app:
	go build -o mst tests/test.go
run-app:
	./mst
docker-build:
	docker build -t app .
docker-run:
	docker run app
