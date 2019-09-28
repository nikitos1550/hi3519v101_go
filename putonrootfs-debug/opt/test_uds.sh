#!/bin/sh

curl --unix-socket /tmp/app_minimal.sock  http://localhost/
curl --unix-socket /tmp/app_minimal.sock  http://localhost/t
