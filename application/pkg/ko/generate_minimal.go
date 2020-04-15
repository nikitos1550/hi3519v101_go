//+build gohisiprobe

//go:generate rm -f kobin_hi3516av100.go
//go:generate go run -tags "generate hi3516av100" ./generator.go --output kobin_probe_hi3516av100.go --tag "hi3516av100,gohisiprobe" --dir ../../sdk/hi3516av100/ko/ --pkg ko --source ./hi3516av100.go --mode minimal

//go:generate rm -f kobin_hi3516av200.go
//go:generate go run -tags "generate hi3516av200" ./generator.go --output kobin_probe_hi3516av200.go --tag "hi3516av200,gohisiprobe" --dir ../../sdk/hi3516av200/ko/ --pkg ko --source ./hi3516av200.go --mode minimal

//go:generate rm -f kobin_hi3516cv100.go
//go:generate go run -tags "generate hi3516cv100" ./generator.go --output kobin_probe_hi3516cv100.go --tag "hi3516cv100,gohisiprobe" --dir ../../sdk/hi3516cv100/ko/ --pkg ko --source ./hi3516cv100.go --mode minimal

//go:generate rm -f kobin_hi3516cv200.go
//go:generate go run -tags "generate hi3516cv200" ./generator.go --output kobin_probe_hi3516cv200.go --tag "hi3516cv200,gohisiprobe" --dir ../../sdk/hi3516cv200/ko/ --pkg ko --source ./hi3516cv200.go --mode minimal

//go:generate rm -f kobin_hi3516cv300.go
//go:generate go run -tags "generate hi3516cv300" ./generator.go --output kobin_probe_hi3516cv300.go --tag "hi3516cv300,gohisiprobe" --dir ../../sdk/hi3516cv300/ko/ --pkg ko --source ./hi3516cv300.go --mode minimal

//go:generate rm -f kobin_hi3516cv500.go
//go:generate go run -tags "generate hi3516cv500" ./generator.go --output kobin_probe_hi3516cv500.go --tag "hi3516cv500,gohisiprobe" --dir ../../sdk/hi3516cv500/ko/ --pkg ko --source ./hi3516cv500.go -mode minimal

//go:generate rm -f kobin_hi3516ev200.go
//go:generate go run -tags "generate hi3516ev200" ./generator.go --output kobin_probe_hi3516ev200.go --tag "hi3516ev200,gohisiprobe" --dir ../../sdk/hi3516ev200/ko/ --pkg ko --source ./hi3516ev200.go --mode minimal

//go:generate rm -f kobin_hi3519av100.go
//go:generate go run -tags "generate hi3519av100" ./generator.go --output kobin_probe_hi3519av100.go --tag "hi3519av100,gohisiprobe" --dir ../../sdk/hi3519av100/ko/ --pkg ko --source ./hi3519av100.go --mode minimal

//go:generate rm -f kobin_hi3559av100.go
//go:generate go run -tags "generate hi3559av100" ./generator.go --output kobin_probe_hi3559av100.go --tag "hi3559av100,gohisiprobe" --dir ../../sdk/hi3559av100/ko/ --pkg ko --source ./hi3559av100.go --mode minimal

//go:generate rm -f kobin_host.go
//go:generate go run -tags "generate host" ./generator.go --output kobin_probe_host.go --tag "host,gohisiprobe" --dir ../../sdk/host/ko/ --pkg ko --source ./host.go --mode minimal

package ko
