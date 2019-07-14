THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := $(THIS_DIR)../buildroot-2019.05.1/output/host/bin/go
STRIP := $(THIS_DIR)../buildroot-2019.05.1/output/host/bin/arm-buildroot-linux-uclibcgnueabi-strip

test:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build ./hello.go
	#GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build -ldflags="-s -w" ./hello.go
	#$(STRIP) ./hello

#run:
#	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) run ./hello.go

#$ go run hello-world.go
#
#$ go build hello-world.go
#
#$ ./hello-world

