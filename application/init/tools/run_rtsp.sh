#!/bin/sh

curl -i -H "Accept: application/json" -H "Content-Type: application/json" --unix-socket /tmp/application.sock -X GET "http://localhost/api/rtsp/start?encoderId=H264_1280_720_1M&streamName=stream"

