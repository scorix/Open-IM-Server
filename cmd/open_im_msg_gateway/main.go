package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/OpenIMSDK/Open-IM-Server/internal/msg_gateway/gate"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/log"
)

func main() {
	log.NewPrivateLog(constant.LogFileName)
	defaultRpcPorts := config.Config.RpcPort.OpenImMessageGatewayPort
	defaultWsPorts := config.Config.LongConnSvr.WebsocketPort
	defaultPromePorts := config.Config.Prometheus.MessageGatewayPrometheusPort
	rpcPort := flag.Int("rpc_port", defaultRpcPorts[0], "rpc listening port")
	wsPort := flag.Int("ws_port", defaultWsPorts[0], "ws listening port")
	prometheusPort := flag.Int("prometheus_port", defaultPromePorts[0], "PushrometheusPort default listen port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("start rpc/msg_gateway server, port: ", *rpcPort, *wsPort, *prometheusPort, ", OpenIM version: ", constant.CurrentVersion)
	gate.Init(*rpcPort, *wsPort)
	gate.Run(*prometheusPort)
	wg.Wait()
}
