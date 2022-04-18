protoc -I=/root/go/go-grpc-assignment/protos/ --go_out=/root/go/go-grpc-assignment/protos/api/proto/v1 --go-grpc_out=/root/go/go-grpc-assignment/protos/api/proto/v1 /root/go/go-grpc-assignment/protos/user.proto


grpcurl --plaintext localhost:9092 list

grpcurl --plaintext localhost:9092 list UserDetail

grpcurl --plaintext localhost:9092 describe UserDetail.GetUsers

grpcurl --plaintext --msg-template localhost:9092 describe .UserDetailRequest

grpcurl --plaintext -d '{"id": [1,2]}' localhost:9092 UserDetail/GetUsers