package controllers

import (
	"net/http"
	"strings"
	"log"

        "github.com/astaxie/beego"
        "github.com/projectcalico/libcalico-go/lib/api"
        "github.com/projectcalico/libcalico-go/lib/client"
        libnet "github.com/projectcalico/libcalico-go/lib/net"
)

// Operations about ippool
type IppoolController struct {
        beego.Controller
}

var ippoolClient client.IPPoolInterface

//func init() {
//        calioClient, err := client.NewFromEnv()
//        if err != nil {
//                log.Println("Error calico client initialization.")
//        }
//	
//        ippoolClient = calioClient.IPPools()
//}

// @Title List ippools.
// @Description List ippools.
// @Success 200 {object} IppoolList 
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [get]
func (c *IppoolController) List() {
        log.Println("Invoke Calico-apiserver Ippool List Api. Request Header: ", c.Ctx.Request.Header)

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
	log.Println("Invoke Calico-apiserver Ippool Get Api. Request Header: ", c.Ctx.Request.Header)
	
	ippoolStringName := c.GetString (":ippool")
        ippoolStringName = strings.Replace(ippoolStringName, "-", "/", 1)

        _, ippoolCidrName, _ := libnet.ParseCIDR(ippoolStringName)
        result, err := ippoolClient.Get(api.IPPoolMetadata{CIDR: *ippoolCidrName})
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
	log.Println("Invoke Calico-apiserver Ippool Create Api. Request Header: ", c.Ctx.Request.Header)
	
	var req api.IPPool
	err := BodyToObject(c.Ctx.Request, &req) 
	if err != nil {
		log.Printf("Failed to unmarshall reques: %v", err)
		writeErrResponse(c, http.StatusBadRequest, err)
		return
	}
	
	result, err := ippoolClient.Create(&req)
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
	log.Println("Invoke Calico-apiserver Ippool Update Api. Request Header: ", c.Ctx.Request.Header)	

	ippoolStringName := c.GetString (":ippool")
        ippoolStringName = strings.Replace(ippoolStringName, "-", "/", 1)	

	_, ippoolCidrName, _ := libnet.ParseCIDR(ippoolStringName)
        result, err := ippoolClient.Get(api.IPPoolMetadata{CIDR: *ippoolCidrName})
        if err != nil {
		writeErrResponse(c, http.StatusBadRequest, err)
	}

	var req api.IPPool
        err = BodyToObject(c.Ctx.Request, &req)
        if err != nil {
                log.Printf("Failed to unmarshall reques: %v", err)
                writeErrResponse(c, http.StatusBadRequest, err)
                return
        }

        result, err = ippoolClient.Update(&req)
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
// @router / [put]
func (c *IppoolController) Apply() {
	log.Println("Invoke Calico-apiserver Ippool Apply Api. Request Header: ", c.Ctx.Request.Header)

        var req api.IPPool
        err := BodyToObject(c.Ctx.Request, &req)
        if err != nil {
                log.Printf("Failed to unmarshall reques: %v", err)
                writeErrResponse(c, http.StatusBadRequest, err)
                return
        }

        result, err := ippoolClient.Apply(&req)
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
	log.Println("Invoke Calico-apiserver Ippool Get Api. Request Header: ", c.Ctx.Request.Header)
	ippoolStringName := c.GetString (":ippool")
        ippoolStringName = strings.Replace(ippoolStringName, "-", "/", 1)

        _, ippoolCidrName, _ := libnet.ParseCIDR(ippoolStringName)
        err := ippoolClient.Delete(api.IPPoolMetadata{CIDR: *ippoolCidrName})
        if err == nil {
                writeResponse(c, http.StatusOK, *ippoolCidrName)
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
	c.ServeJSON()
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeErrResponse(c *IppoolController, code int, err error) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = err.Error()
	c.ServeJSON()
}
