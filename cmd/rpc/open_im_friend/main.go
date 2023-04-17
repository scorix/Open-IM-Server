package main

import (
	"flag"
	"fmt"

	"github.com/OpenIMSDK/Open-IM-Server/internal/rpc/friend"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImFriendPort
	rpcPort := flag.Int("port", defaultPorts[0], "get RpcFriendPort from cmd,default 12000 as port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.FriendPrometheusPort[0], "friendPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start friend rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := friend.NewFriendServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
