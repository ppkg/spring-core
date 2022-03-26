package option

import "google.golang.org/grpc"

// 服务端可选参数
type GrpcServerOptions struct {
	Options []grpc.ServerOption
}

// 客户端可选参数
type GrpcClientOptions struct {
	Options []grpc.DialOption
}
