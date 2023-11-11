startdb:
	docker-compose up -d
run:
	go run cmd/app/main.go

test:
	go test -cover ./... 
	
.PHONY: startdb run test