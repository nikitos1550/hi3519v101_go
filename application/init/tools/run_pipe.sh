#!/bin/sh

#http://213.141.129.12:8080/cam1/api/pipe/start?encoderId=H264_1280_720_1M&pipeName=pipe
#/tmp/pipe will be created

#curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost/api/pipe/start?encoderId=H264_1280_720_1M&pipeName=pipe"
#curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost/api/pipe/stop?pipeName=pipe"

curl -i -H "Accept: application/json" -H "Content-Type: application/json" --unix-socket /tmp/application.sock -X GET "http://localhost/api/pipe/start?encoderId=H264_1280_720_1M&pipeName=pipe"
