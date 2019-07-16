THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := $(THIS_DIR)../buildroot-2019.05.1/output/host/bin/go

test:
	#GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build ./hello.go
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build -ldflags="-s -w" ./mpp_version.go


clean:
	rm ./mpp_version
