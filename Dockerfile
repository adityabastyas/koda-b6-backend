#stage 1

#menggunakan image golang versi 1.25.0 berbasis alpine karna ringan, sebagai encironment untk build
FROM golang:1.26.0-alpine AS build

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o backend cmd/main.go

RUN chmod +x backend


#stage 2
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /workspace/backend /app

EXPOSE 8888

ENTRYPOINT ["/app/backend"]