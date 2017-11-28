package controllers

import (
	"net/http"
	"net"
	"strings"

        "github.com/astaxie/beego"
        "github.com/astaxie/beego/logs"
        "github.com/projectcalico/libcalico-go/lib/api"
        "github.com/projectcalico/libcalico-go/lib/client"
)

// Operations about ippool
type IppoolController struct {
        beego.Controller
}

var ippoolClient interface{}

func init() {
        calioClient, err := client.NewFromEnv()
        if err != nil {
                logs.Error("Error calico client initialization.")
        }
	
        ippoolClient = calioClient.IPPools()
}

// @Title List ippools.
// @Description List ippools.
// @Success 200 {object} IppoolList 
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [get]
func (c *IppoolController) List() {
        logs.Info("Invoke Calico-apiserver Ippool List Api. Request Header: ", c.Ctx.Request.Header)

        result, err := ippoolClient.List(api.IPPoolMetadata{})
	if err == nil {
		writeResponse(c, http.StatusOK, result)
	} else {
		writeErrResponse(c, http.StatusBadRequest, err)
	}
}

// @Title Get specific ippool.
// @Description Get specific ippool.
// @Param  ippool path string true "ippool"
// @Success 200 {object} Ippool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:ippool [get]
func (c *IppoolController) Get() {
	logs.Info("Invoke Calico-apiserver Ippool Get Api. Request Header: ", c.Ctx.Request.Header)
	ippoolStringName := c.GetString (":ippool")
 	ippoolCidrName := net.IPNet{ net.IP{[]byte(strings.Split(ippoolStringName, "/")[0])}, net.IPMask {[]byte((strings.Split(ippoolStringName, "/"))[1])}}

	result, err := ippoolClient.Get(api.IPPoolMetadata{ippoolCidrName})
	if err == nil {
                writeResponse(c, http.StatusOK, result)
        } else {
                writeErrResponse(c, http.StatusBadRequest, err)
        }
}

// @Title create ippool.
// @Description create  ippool.
// @Success 200 {object} ippool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [put]
func (c *IppoolController) Create() {
	logs.Info("Invoke Calico-apiserver Ippool Create Api. Request Header: ", c.Ctx.Request.Header)
	
	var req api.IPPool
	err := BodyToObject(c.Ctx.Request, &req) 
	if err != nil {
		glog.Errorf("Failed to unmarshall reques: %v", err)
		writeErrResponse(c, http.StatusBadRequest, err)
		return
	}
	
	result, err := ippoolClient.Create(req)
	if err == nil {
                writeResponse(c, http.StatusOK, result)
        } else {
                writeErrResponse(c, http.StatusBadRequest, err)
        }	
}

// @Title update specific ippool.
// @Description update specific ippool.
// @Param  ippool path string true "ippool"
// @Success 200 {object} ippool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:ippool [put]
func (c *IppoolController) Update() {
	logs.Info("Invoke Calico-apiserver Ippool Update Api. Request Header: ", c.Ctx.Request.Header)	

	ippoolStringName := c.GetString (":ippool")
        //ippoolCidrName := net.IPNet{
        //        IP: []byte(strings.Split(ippoolStringName, "/")[0])
        //        Mask: []byte(strings.Split(ippoolStringName, "/")[1])
        //}
        ippoolCidrName := net.IPNet{ net.IP{[]byte(strings.Split(ippoolStringName, "/")[0])}, net.IPMask {[]byte((strings.Split(ippoolStringName, "/"))[1])}}

        result, err := ippoolClient.Get(api.IPPoolMetadata{ippoolCidrName})

        if result == nil {
		writeErrResponse(c, http.StatusBadRequest, errors.New("Error when update ippool. Ippool does not exists"))
	}

	var req api.IPPool
        err = BodyToObject(c.Ctx.Request, &req)
        if err != nil {
                glog.Errorf("Failed to unmarshall reques: %v", err)
                writeErrResponse(c, http.StatusBadRequest, err)
                return
        }

        result, err = ippoolClient.Update(req)
        if err == nil {
                writeResponse(c, http.StatusOK, result)
        } else {
                writeErrResponse(c, http.StatusBadRequest, err)
        }

} 

// @Title apply ippool.
// @Description  apply updates an IP pool if it exists, or create a new pool if it does not exists.
// @Success 200 {object} ippool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:ippool [put]
func (c *IppoolController) Apply() {
	logs.Info("Invoke Calico-apiserver Ippool Apply Api. Request Header: ", c.Ctx.Request.Header)

        var req api.IPPool
        err := BodyToObject(c.Ctx.Request, &req)
        if err != nil {
                glog.Errorf("Failed to unmarshall reques: %v", err)
                writeErrResponse(c, http.StatusBadRequest, err)
                return
        }

        result, err := ippoolClient.Apply(req)
        if err == nil {
                writeResponse(c, http.StatusOK, result)
        } else {
                writeErrResponse(c, http.StatusBadRequest, err)
        }
	
}

// @Title delete specific ippool.
// @Description delete specific ippool.
// @Param  ippool path string true "ippool"
// @Success 200 {object} ippool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:ippool [delete]
func (c *IppoolController) Delete() {
	logs.Info("Invoke Calico-apiserver Ippool Get Api. Request Header: ", c.Ctx.Request.Header)
        ippoolStringName := c.GetString (":ippool")
        //ippoolCidrName := net.IPNet{
        //        IP: []byte(strings.Split(ippoolStringName, "/")[0])
        //        Mask: []byte(strings.Split(ippoolStringName, "/")[1])
        //}
	ippoolCidrName := net.IPNet{ net.IP{[]byte(strings.Split(ippoolStringName, "/")[0])}, net.IPMask {[]byte((strings.Split(ippoolStringName, "/"))[1])}}

        result, err := ippoolClient.Delete(api.IPPoolMetadata{ippoolCidrName})
        if err == nil {
                writeResponse(c, http.StatusOK, result)
        } else {
                writeErrResponse(c, http.StatusBadRequest, err)
        }

}


// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeResponse(c *IppoolController, code int, object interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = object
	c.ServerJSON()
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeErrResponse(c *IppoolController, code int, err error) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = err.Error()
	c.ServerJSON()
}
