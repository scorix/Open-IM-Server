package main

import (
	"flag"
	"fmt"

	rpcMessageCMS "github.com/OpenIMSDK/Open-IM-Server/internal/rpc/admin_cms"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImAdminCmsPort
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.AdminCmsPrometheusPort[0], "adminCMSPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start cms rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := rpcMessageCMS.NewAdminCMSServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
