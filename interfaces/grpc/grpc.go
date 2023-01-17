package grpc

import (
	"flag"
	"fmt"
	"net"

	"github.com/viafcccy/garden-be/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/viafcccy/garden-be/interfaces/grpc/registersrv"
)

func InitGrpc() {
	IP := flag.String("ip", "0.0.0.0", "ip 地址")
	Port := flag.Int("port", 0, "端口") // ip, port 这里应该是自动生成。后续注册到 k8s 等
	*IP = "127.0.0.1"
	*Port = 5001

	// 启动 grpc
	server := grpc.NewServer()

	// 注册 pb 服务
	registersrv.RegisterSrv(server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 启动 grpc 服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	global.GLog.Infof("garden-be rpc 服务启动... IP: %s，端口：%d", *IP, *Port)
}
