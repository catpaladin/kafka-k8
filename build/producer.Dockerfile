FROM golang:1.11.2 as build

# install dependencies
RUN go get github.com/Shopify/sarama

# copy app
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0

# build
RUN go build -o producer src/pub/producer.go

# second stage
FROM alpine:3.8 as run

WORKDIR /app

COPY --from=build /app/producer /app/

CMD ["./producer"]