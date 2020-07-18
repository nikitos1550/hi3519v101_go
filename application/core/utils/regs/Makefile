all: clean build

systemrdl:
	make -C ./registers-description build

generate: systemrdl
	./rdl_to_go.py ./registers-description/hi3516av200.rdl hi3516av200 --tags hi3516av200 > ./regs_hi3516av200.go
	go fmt ./regs_hi3516av200.go
	#./rdl_to_go.py ./registers-description/hi3516av200.rdl hi3519v101 > ./regs_hi3519v101.go
	#go fmt ./hi3519v101.go

build: generate
	#go build

clean:
	rm -f ./hi35*.go

deps:
	pip3 install systemrdl-compiler
