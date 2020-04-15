#!/bin/sh

curl -i -H "Accept: application/json" -H "Content-Type: application/json" --unix-socket /tmp/application.sock -X GET "http://localhost/api/rtsp/stop?streamName=stream"

