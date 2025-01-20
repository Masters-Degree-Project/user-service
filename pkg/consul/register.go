package consul

import (
	"fmt"
	"log"
	"strconv"
	"user/pkg/config"

	"github.com/hashicorp/consul/api"
)

func RegisterService() error {
	consulHost := config.Config("CONSUL_HOST")
	consulPort := config.Config("CONSUL_PORT")

	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("http://%s:%s", consulHost, consulPort)
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return fmt.Errorf("failed to create consul client: %v", err)
	}

	serviceName := config.Config("SERVICE_NAME")
	serviceId := config.Config("SERVICE_ID")
	serviceIp := config.Config("SERVICE_IP")
	servicePort, err := strconv.Atoi(config.Config("SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Port number could not be converted to integer: %v", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Address: serviceIp,
		Port:    servicePort,
		Tags: []string{
			"traefik.enable=true",
			fmt.Sprintf("traefik.http.routers.%s-router.rule=Headers(`X-Service`, `%s`)", serviceName, serviceName),
			fmt.Sprintf("traefik.http.routers.%s-router.service=%s", serviceName, serviceName),
			fmt.Sprintf("traefik.http.routers.%s-router.entryPoints=web", serviceName),
			fmt.Sprintf("traefik.http.services.%s.loadBalancer.server.port=%d", serviceName, servicePort),
		},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", serviceIp, servicePort),
			Interval: "10s",
		},
	}

	// Consul'e servisi kaydet
	if err := client.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("service registration failed: %v", err)
	}

	log.Printf("✅ Consul service registered successfully on port!")
	return nil
}
