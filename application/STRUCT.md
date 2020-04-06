## Idea

* family related code separated on files level
* host is special family, that allow run code on any system
* chip difference inside family should be solved in runtime 
* file names:
  * *.go file will be compiled in all cases
  * *_hi35****.go is just mark, it should contains //+build arm,hi35****
  * *_host.go should contains //+build 386 amd64

## ...

golang 
	OS always Linux
	ARCH
		host 
		arm
			5
			7
			8
		arm64

file level compile time code separation
	file suffix
		_arm (all 32bit arms)
		_arm64 (64bit arm)
		_386/_amd64 (host)
	tags
		//+build inside file

CGO C code is placed inside go files
	it is possible to use defines in C code

Also it is possible to have global variables/constants family/chip name 
to let us switch behaviour in runtime


Compile time
	File level
		filename suffix			- not reasonable
		build tags
	*Code level
		#define ONLY FOR C code
Runtime level
	Code level
		global variable/constant 

buildtags
	os
	arch
	hisi family
	hisi chip



## NOTES

* https://golang.org/pkg/go/build/
* https://github.com/golang/go/wiki/GoArm
* https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
