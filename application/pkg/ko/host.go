//+build 386 amd64
//+build host

//go:generate rm -f kobin_host.go
//go:generate go run -tags "generate host" ./generate.go --output kobin_host.go --tag host --dir ../../sdk/host/ko/ --pkg ko --source ./host.go

package ko

var (
        ModulesList = [...][2]string{}
        minimalModulesList = [...]string{}
)

