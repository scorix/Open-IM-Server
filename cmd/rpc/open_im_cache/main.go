package main

import (
	rpcCache "github.com/OpenIMSDK/Open-IM-Server/internal/rpc/cache"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	promePkg "github.com/OpenIMSDK/Open-IM-Server/pkg/common/prometheus"

	"flag"
	"fmt"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImCachePort
	rpcPort := flag.Int("port", defaultPorts[0], "RpcToken default listen port 10800")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.CachePrometheusPort[0], "cachePrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start cache rpc server, port: ", *rpcPort, ", OpenIM version: ", constant.CurrentVersion)
	rpcServer := rpcCache.NewCacheServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
