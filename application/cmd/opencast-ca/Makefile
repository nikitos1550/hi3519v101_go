all:
	@echo "TODO"

run:
	go run ./ca.go

build:
	go build ./ca.go

testvideo.mp4:
	ffmpeg -f lavfi -i testsrc=size=640x360 -t 30 -pix_fmt yuv420p -c:v libx264 -tune zerolatency -profile:v main -g 25 testvideo.mp4
