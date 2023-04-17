package main

import (
	"flag"
	"fmt"

	rpcConversation "github.com/OpenIMSDK/Open-IM-Server/internal/rpc/conversation"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImConversationPort
	rpcPort := flag.Int("port", defaultPorts[0], "RpcConversation default listen port 11300")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.ConversationPrometheusPort[0], "conversationPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start conversation rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := rpcConversation.NewRpcConversationServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()

}
