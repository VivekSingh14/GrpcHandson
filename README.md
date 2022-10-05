# GrpcHandson


# CMD

protoc -I greet/proto --go_out=. --go_opt=module=github.com/VivekSingh14/GrpcHandson --go-grpc_out=. --go-grpc_opt=module=github.com/VivekSingh14/GrpcHandson  greet/proto/dummy.proto