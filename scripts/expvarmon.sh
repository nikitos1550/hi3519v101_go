#!/bin/sh

#go get github.com/divan/expvarmon

#./expvarmon -ports="http://213.141.129.12:8080" -endpoint="/cam1/api/debug/vars"
~/go/bin/expvarmon -ports="http://192.168.10.101:80" -endpoint="/api/debug/vars" -i=1s
