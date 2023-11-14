startdb:
	docker-compose up -d
run:
	go run cmd/app/main.go

test:
	go test -cover ./... 

swaginit:
	swag init -g cmd/app/main.go
	
.PHONY: startdb run test