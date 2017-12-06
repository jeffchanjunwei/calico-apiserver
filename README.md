# Introduction
calico-apiserver is a self-developed module which invoke the interfaces of calico and supply Restful apis. Now it can be used to edit the calico ippool and workloadendpoint resources. 

# Functional structure
<img src="https://github.com/jeffchanjunwei/calico-apiserver/raw/master/paas-net.png" width = "500" height = "500" alt="paas-net structure" align=center />

# Version of dependent packages
+ libcalico-go:1.7.1
+ beego:1.9.0
+ go:v1.9.1

# Installation instructions
- docker container  
  docker build . -t jeffchanjunwei/calico-apiserver
- local compile
  bee run -gendoc=true -downdoc=false

