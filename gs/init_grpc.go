package gs

import (
	"reflect"
	"runtime"

	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/web"
	g "google.golang.org/grpc"
)

func init() {
	Object(new(initGrpcContainer)).Init(func(init *initGrpcContainer) {
		init.webServer.SetupGrpc(func(svr *g.Server) {
			server := reflect.ValueOf(svr)
			srvMap := make(map[string]reflect.Value)

			init.grpcServers.ForEach(func(serviceName string, rpcServer *grpc.Server) {
				service := reflect.ValueOf(rpcServer.Service)
				srvMap[serviceName] = service
				fn := reflect.ValueOf(rpcServer.Register)
				fn.Call([]reflect.Value{server, service})
			})

			for service, info := range svr.GetServiceInfo() {
				srv := srvMap[service]
				for _, method := range info.Methods {
					m, _ := srv.Type().MethodByName(method.Name)
					fnPtr := m.Func.Pointer()
					fnInfo := runtime.FuncForPC(fnPtr)
					file, line := fnInfo.FileLine(fnPtr)
					log.Infof("/%s/%s %s:%d ", service, method.Name, file, line)
				}
			}

		})
	})
}

type initGrpcContainer struct {
	webServer   web.Server   `autowire:""`
	grpcServers *GrpcServers `autowire:""`
}
