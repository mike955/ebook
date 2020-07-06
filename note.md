protoc --gofast_out=plugins=grpc:user user.proto
protoc -I. --grpc-gateway_out=logtostderr=true:. user.proto


protoc --gofast_out=plugins=grpc:ebook ebook.proto
protoc -I. --grpc-gateway_out=logtostderr=true:ebook ebook.proto

todo 
  - 统一处理响应
  - 统一注册服务
  - 优雅重启服务

grpc-gateway support not formdata