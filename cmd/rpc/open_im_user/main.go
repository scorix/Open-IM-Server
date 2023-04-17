package main

import (
	"flag"
	"fmt"

	"github.com/OpenIMSDK/Open-IM-Server/internal/rpc/user"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImUserPort
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.UserPrometheusPort[0], "userPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start user rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := user.NewUserServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
