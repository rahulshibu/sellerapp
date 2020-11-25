FROM golang:latest AS builder
ADD / /app/sellerapp
WORKDIR /app/sellerapp
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/scraping ./scraping-service/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/saving ./saving-service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache --upgrade bash
COPY --from=builder /app/sellerapp/bin/scraping ./
COPY --from=builder /app/sellerapp/bin/saving ./
COPY start.sh start.sh

EXPOSE 8080 8081
RUN chmod +x ./scraping
RUN chmod +x ./saving
RUN chmod +x start.sh
CMD ./start.sh
