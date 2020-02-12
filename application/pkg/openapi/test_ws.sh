#!/bin/sh

curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Host: 192.168.10.101:80" \
     --header "Origin: http://192.168.10.101:80" \
     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
     --header "Sec-WebSocket-Version: 13" \
     http://192.168.10.101:80/ws/echo
