//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"sys_srv/global"
)

func InitConsul(server *grpc.Server) {
	//register health service
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//init service
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host,
		global.ServerConfig.ConsulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID, err := uuid.NewUUID()
	if err != nil {
		zap.S().Fatal("fail to generate service UUID, err:", err.Error())
	}
	registration.ID = serviceID.String()
	registration.Port = int(global.ServerConfig.Port)
	registration.Tags = global.ServerConfig.Tags
	registration.Address = global.ServerConfig.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Fatal("fail to register service to consul, err:", err.Error())
	}
}
