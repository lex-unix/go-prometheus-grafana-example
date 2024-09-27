FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o=./api ./main.go


FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /app

COPY --from=builder /app/api ./api

EXPOSE 8080

ENTRYPOINT [ "./api" ]
