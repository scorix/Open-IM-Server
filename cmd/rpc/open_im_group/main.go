package main

import (
	"flag"
	"fmt"

	"github.com/OpenIMSDK/Open-IM-Server/internal/rpc/group"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImGroupPort
	rpcPort := flag.Int("port", defaultPorts[0], "get RpcGroupPort from cmd,default 16000 as port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.GroupPrometheusPort[0], "groupPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start group rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := group.NewGroupServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
