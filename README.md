# User Service

The user service is a gRPC service from which we get user data from database.

## Flags

- `-server-port=9092` To run on port 9092

- `-log-level=debug` To see diffrent level of logs (debug,info,fatal)

## Running on Docker

To run User Service inside docker follow below stps:

1. Build Dockerfile:

   `$ docker build -t grpc-server .`

2. Run Docker image:

   `$ docker run -d --name grpc-server -p 9092:9092 grpc-server:latest`

   #### Or

   To run server on custom port and use different log level there are two flags.

   2.1. `-server-port=8081` To run on port 8081

   2.2. `-log-level=debug` For debug logs

   with flags docker run command

   `$ docker run -d --name grpc-server -p 8081:8081 grpc-server:latest -server-port=8081 -log-level=debug`

## Run on local machine

To run on local machine we need to first follow below steps:

1. Run command to download dependacy:

   `$ go mod download`

2. To build and run:

   `$ GOOS=linux go build -o grpc-server main.go`

   `$ ./grpc-server`

3. To run on system without building executable:

   `$ go run main.go`

## Testing

To test the system install `grpcurl` which is a command line tool which can interact with gRPC API's

https://github.com/fullstorydev/grpcurl

`$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`

### User Data

```
{
    {
      "id": "1",
      "fname": "Tony",
      "city": "Maliby",
      "phone": 1234567890,
      "height": 5.67,
      "Married": true
    },
    {
      "id": "2",
      "fname": "Thor",
      "city": "Asgard",
      "phone": 1234567890,
      "height": 6.67
    },
    {
      "id": "3",
      "fname": "Peter",
      "city": "New York",
      "phone": 1234567890,
      "height": 5.3,
      "Married": true
    }
}
```

### List Services

usign `--plaintext` for not secure stream like tls.

```
$ grpcurl --plaintext localhost:9092 list
UserDetail
grpc.reflection.v1alpha.ServerReflection
```

### List Methods

```
$ grpcurl --plaintext localhost:9092 list UserDetail
UserDetail.GetUser
UserDetail.GetUsersList
```

### Methods detail for GetUser

```
$ grpcurl --plaintext localhost:9092 describe UserDetail.GetUser
UserDetail.GetUser is a method:
rpc GetUser ( .UserDetailRequest ) returns ( .UserDetailResponse );
```

### Methods detail for GetUsersList

```
$ grpcurl --plaintext localhost:9092 describe UserDetail.GetUsersList
UserDetail.GetUsersList is a method:
rpc GetUsersList ( .ListOfUserDetailsRequest ) returns ( .ListOfUserDetailsResponse );
```

### Request to Get User from gRPC service

Success Response:

```
$ grpcurl --plaintext -d '{"id": 1}' localhost:9092 UserDetail/GetUser

{
  "users": {
    "id": "1",
    "fname": "Tony",
    "city": "Maliby",
    "phone": 1234567890,
    "height": 5.67,
    "Married": true
  }
}
```

Error Response:

```
$ grpcurl --plaintext -d '{"id": 4}' localhost:9092 UserDetail/GetUser

ERROR:
  Code: NotFound
  Message: User with id 4 doesn't exist in Database
```

### Request to Get List of User from gRPC service

Success Response:

```
$ grpcurl --plaintext -d '{"ids": [1,2,3]}' localhost:9092 UserDetail/GetUsersList

{
  "users": [
    {
      "id": "1",
      "fname": "Tony",
      "city": "Maliby",
      "phone": 1234567890,
      "height": 5.67,
      "Married": true
    },
    {
      "id": "2",
      "fname": "Thor",
      "city": "Asgard",
      "phone": 1234567890,
      "height": 6.67
    },
    {
      "id": "3",
      "fname": "Peter",
      "city": "New York",
      "phone": 1234567890,
      "height": 5.3,
      "Married": true
    }
  ]
}
```

With invalid user id:

```
$ grpcurl --plaintext -d '{"ids": [1,2,3,4]}' localhost:9092 UserDetail/GetUsersList

{
  "users": [
    {
      "id": "1",
      "fname": "Tony",
      "city": "Maliby",
      "phone": 1234567890,
      "height": 5.67,
      "Married": true
    },
    {
      "id": "2",
      "fname": "Thor",
      "city": "Asgard",
      "phone": 1234567890,
      "height": 6.67
    },
    {
      "id": "3",
      "fname": "Peter",
      "city": "New York",
      "phone": 1234567890,
      "height": 5.3,
      "Married": true
    }
  ],
  "error": {
    "code": 5,
    "message": "Ids doesn't exist [4] in database"
  }
}
```
