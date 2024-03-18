package grpc

import (
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
)

type T struct {
	Consul struct {
		Ip          string `json:"Ip"`
		Port        int    `json:"Port"`
		SerivceName string `json:"SerivceName"`
		Wait        string `json:"Wait"`
		With        string `json:"With"`
		Service     string `json:"Service"`
	} `json:"Consul"`
}

func Client() (*grpc.ClientConn, error) {
	//config, err := nacos.GetConfig()
	//if err != nil {
	//	logs.Info(err, "errrrrrr")
	//	return nil, err
	//}
	//var t T
	//json.Unmarshal([]byte(config), &t)
	return grpc.Dial("consul://192.168.118.9:8500/"+"Grpc"+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy":"round_robin"}`))

}
