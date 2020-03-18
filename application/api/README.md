# GoHisiCam Bash client

## Overview

This is a Bash client script for accessing GoHisiCam service.

The script uses cURL underneath for making all REST calls.

## Usage

```shell
# Make sure the script has executable rights
$ chmod u+x 

# Print the list of operations available on the service
$ ./ -h

# Print the service description
$ ./ --about

# Print detailed information about specific operation
$ ./ <operationId> -h

# Make GET request
./ --host http://<hostname>:<port> --accept xml <operationId> <queryParam1>=<value1> <header_key1>:<header_value2>

# Make GET request using arbitrary curl options (must be passed before <operationId>) to an SSL service using username:password
 -k -sS --tlsv1.2 --host https://<hostname> -u <user>:<password> --accept xml <operationId> <queryParam1>=<value1> <header_key1>:<header_value2>

# Make POST request
$ echo '<body_content>' |  --host <hostname> --content-type json <operationId> -

# Make POST request with simple JSON content, e.g.:
# {
#   "key1": "value1",
#   "key2": "value2",
#   "key3": 23
# }
$ echo '<body_content>' |  --host <hostname> --content-type json <operationId> key1==value1 key2=value2 key3:=23 -

# Preview the cURL command without actually executing it
$  --host http://<hostname>:<port> --dry-run <operationid>

```

## Docker image

You can easily create a Docker image containing a preconfigured environment
for using the REST Bash client including working autocompletion and short
welcome message with basic instructions, using the generated Dockerfile:

```shell
docker build -t my-rest-client .
docker run -it my-rest-client
```

By default you will be logged into a Zsh environment which has much more
advanced auto completion, but you can switch to Bash, where basic autocompletion
is also available.

## Shell completion

### Bash

The generated bash-completion script can be either directly loaded to the current Bash session using:

```shell
source .bash-completion
```

Alternatively, the script can be copied to the `/etc/bash-completion.d` (or on OSX with Homebrew to `/usr/local/etc/bash-completion.d`):

```shell
sudo cp .bash-completion /etc/bash-completion.d/
```

#### OS X

On OSX you might need to install bash-completion using Homebrew:

```shell
brew install bash-completion
```

and add the following to the `~/.bashrc`:

```shell
if [ -f $(brew --prefix)/etc/bash_completion ]; then
  . $(brew --prefix)/etc/bash_completion
fi
```

### Zsh

In Zsh, the generated `_` Zsh completion file must be copied to one of the folders under `$FPATH` variable.

## Documentation for API Endpoints

All URIs are relative to **

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChannelApi* | [**vpssIdCreate**](docs/ChannelApi.md#vpssidcreate) | **POST** /api/mpp/channel/{id} | Start channel
*ChannelApi* | [**vpssIdDelete**](docs/ChannelApi.md#vpssiddelete) | **DELETE** /api/mpp/channel/{id} | Stop channel
*ChannelApi* | [**vpssIdShow**](docs/ChannelApi.md#vpssidshow) | **GET** /api/mpp/channel/{id} | Show channel information
*ChannelApi* | [**vpssShow**](docs/ChannelApi.md#vpssshow) | **GET** /api/mpp/channel | Show all channels
*CmosApi* | [**mppCmos**](docs/CmosApi.md#mppcmos) | **GET** /api/mpp/cmos | CMOS information
*DebugApi* | [**debugUmap**](docs/DebugApi.md#debugumap) | **GET** /api/debug/umap | MPP debug avalible information overview
*DebugApi* | [**debugUmapFile**](docs/DebugApi.md#debugumapfile) | **GET** /api/debug/umap/{file} | Show debug information
*JpegApi* | [**serveJpeg1920x1080**](docs/JpegApi.md#servejpeg1920x1080) | **GET** /jpeg/1920_1080.jpg | Get jpeg image 1920x1080 resolution
*JpegApi* | [**serveJpeg3840x2160**](docs/JpegApi.md#servejpeg3840x2160) | **GET** /jpeg/3840_2160.jpg | Get jpeg image 3840x2160 resolution
*MiscApi* | [**openapi**](docs/MiscApi.md#openapi) | **GET** /openapi.json | Get openapi specification
*MiscApi* | [**openapi2**](docs/MiscApi.md#openapi2) | **GET** /openapi.yml | Get openapi specification
*OtherApi* | [**api123**](docs/OtherApi.md#api123) | **GET** / | Get API information
*SystemApi* | [**systemDate**](docs/SystemApi.md#systemdate) | **GET** /api/system/date | Get system date and time
*VencApi* | [**vencCreate**](docs/VencApi.md#venccreate) | **POST** /api/mpp/venc | Create encoder
*VencApi* | [**vencIdChange**](docs/VencApi.md#vencidchange) | **PUT** /api/mpp/venc/{id} | Change encoder runtime settings
*VencApi* | [**vencIdDelete**](docs/VencApi.md#venciddelete) | **DELETE** /api/mpp/venc/{id} | Stop encoder
*VencApi* | [**vencIdShow**](docs/VencApi.md#vencidshow) | **GET** /api/mpp/venc/{id} | Show ecnoder information
*VencApi* | [**vencShow**](docs/VencApi.md#vencshow) | **GET** /api/mpp/venc | Show all encoders


## Documentation For Models

 - [Cat](docs/Cat.md)
 - [CatAllOf](docs/CatAllOf.md)
 - [Dog](docs/Dog.md)
 - [DogAllOf](docs/DogAllOf.md)
 - [EncodersList](docs/EncodersList.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [Lizard](docs/Lizard.md)
 - [LizardAllOf](docs/LizardAllOf.md)
 - [StaticParams](docs/StaticParams.md)
 - [StaticParams2](docs/StaticParams2.md)


## Documentation For Authorization

 All endpoints do not require authorization.

