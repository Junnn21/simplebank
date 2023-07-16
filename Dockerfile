# Build stage
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/LICENSE ./migrate
COPY --from=builder /app/README.md ./migrate
COPY --from=builder /app/migrate ./migrate
COPY start.sh .
COPY app.env .
COPY db/migration ./migration
RUN ["chmod", "+x", "./start.sh"]

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
