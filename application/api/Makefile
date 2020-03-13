SPEC	:=	gohisicam.yml

all: build redoctmp
	#@echo "TODO"

build:
	speccy resolve ./src/main.yml -v > $(SPEC)

lint:
	speccy lint $(SPEC) -v


GOPHER200   := $(shell cat ./images/gopher200.png | base64 -w 0)
DOWNLOAD    := $(shell cat ./images/download_button.png | base64 -w 0)
FAVICON     := $(shell cat ./images/favicon.ico | base64 -w 0)

VERSION 	:= 0.0.1

openapipostprocess:
	#echo $(GOPHER200PNG)
	./tools/redoc/cli/index.js bundle $(SPEC) --options ./redoc.conf.json;
	sed "s!GOPHER200!data:image\/png;base64,$(GOPHER200)!g" ./redoc-static.html > ./redoc-static_base64.html
	rm ./redoc-static.html
	sed "s!DOWNLOAD!data:image\/png;base64,$(DOWNLOAD)!g" ./redoc-static_base64.html > ./redoc-static_base64_2.html
	rm ./redoc-static_base64.html
	sed "s!<head>!<head><link rel='shortcut icon' href='data:image\/x-icon;base64,$(FAVICON)'!g" ./redoc-static_base64_2.html > ./redoc-static_base64_3.html
	rm ./redoc-static_base64_2.html
	sed "s!VERSION!$(VERSION)!g" ./redoc-static_base64_3.html > ./redoc-static_base64_4.html
	rm ./redoc-static_base64_3.html
	cp ./redoc-static_base64_4.html ./gohisicam.html
	rm ./redoc-static_base64_4.html

redoctmp: openapipostprocess
	#redoc-cli bundle $(SPEC) 
	cp ./gohisicam.html ~/www/redoc-static.html
