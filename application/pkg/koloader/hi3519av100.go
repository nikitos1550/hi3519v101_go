// +build hi3519av100

//go:generate rm -f hi3519av100_kobin.go
//go:generate go run -tags "generate hi3519acv100" ./generate.go --output hi3519av100_kobin.go --tag hi3519av100 --dir ../../sdk/hi3519av100/ko --pkg koloader --source ./hi3519av100.go

package koloader

//TODO
