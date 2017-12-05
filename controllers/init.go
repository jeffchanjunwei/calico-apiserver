package controllers

import (
	"github.com/projectcalico/libcalico-go/lib/client"
)

func init() {
	calicoClient, err := client.NewFromEnv()
	if err != nil {
		log.Println("Error calico client initialization.")
	}

	ippoolClient = calicoClient.IPPools()
	workloadendpointClient = calioClient.WorkloadEndpoints()
}
