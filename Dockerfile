FROM golang:latest AS builder
ADD / /app/sellerapp
WORKDIR /app/sellerapp
RUN go mod download
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/sellerapp ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/sellerapp/bin/sellerapp ./
EXPOSE 8080
RUN chmod +x ./sellerapp
ENTRYPOINT ["./sellerapp"]
