//+build !gohisiprobe

//go:generate rm -f kobin_hi3516av100.go
//go:generate go run -tags "generate hi3516av100" ./generator.go --output kobin_hi3516av100.go --tag hi3516av100 --dir ../../sdk/hi3516av100/ko/ --pkg ko --source ./hi3516av100.go

//go:generate rm -f kobin_hi3516av200.go
//go:generate go run -tags "generate hi3516av200" ./generator.go --output kobin_hi3516av200.go --tag hi3516av200 --dir ../../sdk/hi3516av200/ko/ --pkg ko --source ./hi3516av200.go

//go:generate rm -f kobin_hi3516cv100.go
//go:generate go run -tags "generate hi3516cv100" ./generator.go --output kobin_hi3516cv100.go --tag hi3516cv100 --dir ../../sdk/hi3516cv100/ko/ --pkg ko --source ./hi3516cv100.go

//go:generate rm -f kobin_hi3516cv200.go
//go:generate go run -tags "generate hi3516cv200" ./generator.go --output kobin_hi3516cv200.go --tag hi3516cv200 --dir ../../sdk/hi3516cv200/ko/ --pkg ko --source ./hi3516cv200.go

//go:generate rm -f kobin_hi3516cv300.go
//go:generate go run -tags "generate hi3516cv300" ./generator.go --output kobin_hi3516cv300.go --tag hi3516cv300 --dir ../../sdk/hi3516cv300/ko/ --pkg ko --source ./hi3516cv300.go

//go:generate rm -f kobin_hi3516cv500.go
//go:generate go run -tags "generate hi3516cv500" ./generator.go --output kobin_hi3516cv500.go --tag hi3516cv500 --dir ../../sdk/hi3516cv500/ko/ --pkg ko --source ./hi3516cv500.go

//go:generate rm -f kobin_hi3516ev200.go
//go:generate go run -tags "generate hi3516ev200" ./generator.go --output kobin_hi3516ev200.go --tag hi3516ev200 --dir ../../sdk/hi3516ev200/ko/ --pkg ko --source ./hi3516ev200.go

//go:generate rm -f kobin_hi3519av100.go
//go:generate go run -tags "generate hi3519av100" ./generator.go --output kobin_hi3519av100.go --tag hi3519av100 --dir ../../sdk/hi3519av100/ko/ --pkg ko --source ./hi3519av100.go

//go:generate rm -f kobin_hi3559av100.go
//go:generate go run -tags "generate hi3559av100" ./generator.go --output kobin_hi3559av100.go --tag hi3559av100 --dir ../../sdk/hi3559av100/ko/ --pkg ko --source ./hi3559av100.go

package ko

