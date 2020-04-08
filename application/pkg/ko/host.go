//+build 386 amd64
//+build host

//g-o:generate rm -f kobin_host.go
//g-o:generate go run -tags "generate host" ./generate.go --output kobin_host.go --tag host --dir ../../sdk/host/ko/ --pkg ko --source ./host.go

package ko

var (
        ModulesList = [...][2]string{}
        MinimalModulesList = [...]string{}
)

