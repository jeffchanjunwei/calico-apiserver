#!/bin/sh

start_calico_apiserver() {
    echo 'Starting Calico Apiservice ...'
    /go/bin/bee run -gendoc=true -downdoc=false
}

start_calico_apiserver
