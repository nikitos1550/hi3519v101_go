THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := $(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/go

test:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build ./app_sample.go

clean:
	rm ./app_sample
