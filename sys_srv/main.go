package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sys_srv/global"
	"sys_srv/handler"
	"sys_srv/initialize"
	"sys_srv/proto"
)

func main() {
	//init logger
	initialize.InitZapLogger()
	//init Nacos config
	initialize.InitNacosConfig()
	//init sys config
	initialize.InitServerConfigFromNacos()
	//init multiple languages support
	initialize.InitI18n()
	//init gorm for mysql database
	initialize.InitGorm()

	//init sys server
	conn, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port))
	if err != nil {
		zap.S().Fatalf("fail to listen ip:%s port:%d", global.ServerConfig.Host)
	}
	server := grpc.NewServer()

	//init consul
	initialize.InitConsul(server)

	proto.RegisterSystemServer(server, &handler.SystemServer{})
	err = server.Serve(conn)
	if err != nil {
		zap.S().Fatalf("fail to start server: %s", err.Error())
	}
}
