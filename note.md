protoc --gofast_out=plugins=grpc:user user.proto
protoc -I. --grpc-gateway_out=logtostderr=true:. user.proto