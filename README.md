# Message broker GRPC

This repo contains solutions of "Message broker" on Golang

## DESCRIPTION
Write a simple message broker, SDK for it in the form of a library and examples of client applications (producer, consumer).
Operation protocol - gRPC


## How 2RUN

1. Clone project

```
git clone git@github.com:rabdavinci/broker.git .
```

2. Run microservice broker

```
$ go run broker/cmd/main.go
```

3. Run producer

```
$ go run producer/main.go

```
## TODO

1. Add consumer
2. Optimize broker. Migrate to RabbitMQ
3. Optimize producer
4. Add tests
5. Dockerize, move configurations to `ENV` or `YML`
6. Configure `CI/CD`
7. Optimize broker. Add Scalability
8. Add documentation
