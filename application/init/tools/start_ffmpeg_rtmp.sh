#!/bin/sh

ffmpeg -i /tmp/pipe -c:v copy -f flv rtmp://192.168.10.2/hls/pipe
