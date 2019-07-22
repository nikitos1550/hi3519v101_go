#!/bin/bash

#java -jar openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate \
#  -i http://petstore.swagger.io/v2/swagger.json \
#  -g go-server \
#  -o samples/server/petstore/go-server

#java -jar ./openapi-generator-cli.jar generate \
#  -i ./api.yaml \
#  -g go-gin-server \
#  -o go-gin-server

# https://github.com/deepmap/oapi-codegen
 ~/go/bin/oapi-codegen pet.yaml  > pet.gen.go
