package controllers

import (
        "net/http"
        "log"

        "github.com/astaxie/beego"
        "github.com/projectcalico/libcalico-go/lib/api"
        "github.com/projectcalico/libcalico-go/lib/client"
)

type WorkloadendpointController struct {
        beego.Controller
}

var workloadendpointClient client.WorkloadEndpointInterface

func init() {
	workloadendpointClient = calioClient.WorkloadEndpoints()
}

// @Title List workloadendpoints.
// @Description List workloadendpoints.
// @Success 200 {object} WorkloadendpointList
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [get]
func (c *WorkloadendpointController) List() {
	log.Println("Invoke Calico-apiserver Workloadendpoint List Api. Request Header: ", c.Ctx.Request.Header)

	result, err := workloadendpointClient.List(api.WorkloadEndpointMetadata{})
	if err == nil {
		writeWorkloadendpointResponse(c, http.StatusOK, result)
	} else {
		writeWorkloadendpointErrResponse(c, http.StatusBadRequest, err)
	}
}

func writeWorkloadendpointResponse(c *WorkloadendpointController, code int, object interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = object
	c.ServeJSON()
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeWorkloadendpointErrResponse(c *WorkloadendpointController, code int, err error) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = err.Error()
	c.ServeJSON()
}

