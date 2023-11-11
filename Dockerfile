FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/app/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

ENV SERVER_ADDR="localhost:8080" 
ENV PSQL_HOST="localhost"
ENV PSQL_PORT="5432"
ENV PSQL_USER="postgres"
ENV PSQL_PASSWORD="postgres"
ENV PSQL_DBNAME="pdd-datastore"
ENV PSQL_SSLMODE="disable"
ENV PSQL_CONNECT_TIMEOUT=10
ENV PSQL_CONNECT_WAIT_TIME=3
ENV PSQL_CONNECT_ATTEMPTS=3
ENV PSQL_CONNECT_BLOCKS=false
ENV PSQL_CLOSE_TIMEOUT=10

EXPOSE 8080
CMD [ "/app/main"]