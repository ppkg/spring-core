package option

import "google.golang.org/grpc"

// 服务端可选参数
type GrpcServerOptions []grpc.ServerOption

// 客户端可选参数
type GrpcClientOptions []grpc.DialOption
