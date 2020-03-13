#!/bin/sh

npm run compile:cli
chmod +x ./cli/index.js

rm -rf ./bundles
rm -rf ./typings
cp -r ./_typings ./typings

rm -rf ./cli/node_modules

npm run bundle

cd cli; npm install; cd ..
rm -rf ./cli/node_modules/redoc/bundles
rm -rf ./cli/node_modules/redoc/typings
cp -r ./bundles ./cli/node_modules/redoc/
cp -r ./typings ./cli/node_modules/redoc/

