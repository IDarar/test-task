run:
	go run ./cmd/main.go
compose:
	sudo docker-compose up -d
verbosetest:
	go test -count=1 -v ./...
covertest:
	go test -count=1 -cover ./...
proto:
	protoc -I . service.proto --go_out=plugins=grpc:.
enterpostgres:
	sudo docker exec -it postgres13 psql -U root
servicemock:
	mockgen -source=internal/service/users.go -destination internal/service/mocks/mock.go
dockerbuild:
	sudo docker build -t test-task .
tag:
	sudo docker tag test-task:latest aince/test-task
push:
	sudo docker push aince/test-task