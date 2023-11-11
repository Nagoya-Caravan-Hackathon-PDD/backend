startdb:
	docker-compose up -d
run:
	go run cmd/app/main.go
	
.PHONY: startdb run