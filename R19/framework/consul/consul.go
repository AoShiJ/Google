package consul

import (
	"demo/lx/framework/nacos"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

type consul struct {
	Consul struct {
		Ip          string `json:"Ip"`
		Port        int    `json:"Port"`
		SerivceName string `json:"SerivceName"`
		Wait        string `json:"Wait"`
		With        string `json:"With"`
		Service     string `json:"Service"`
	} `json:"Consul"`
}

func NewConsul(port int) error {

	config, err := nacos.GetConfig()
	if err != nil {
		return err
	}
	var consuls consul
	json.Unmarshal([]byte(config), &consuls)
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", consuls.Consul.Ip, consuls.Consul.Port),
	})
	if err != nil {
		return nil
	}
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    "Grpc",
		Tags:    []string{"GRPC"},
		Port:    port,
		Address: consuls.Consul.Ip,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			GRPC:                           fmt.Sprintf("%v:%v", consuls.Consul.Ip, port),
			DeregisterCriticalServiceAfter: "30s",
		},
	})
	if err != nil {
		return nil
	}
	return nil
}
