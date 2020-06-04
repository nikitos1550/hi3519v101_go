
//go:generate rm -f embed_hi3516av100.go
//go:generate rm -f append_hi3516av100.go
//go:generate rm -f append_hi3516av100.bin
//go:generate go run -tags "generate hi3516av100" ./generator.go --output hi3516av100 --tag "hi3516av100,koEmbed" --dir ../../sdk/hi3516av100/ko/ --pkg ko --source ./hi3516av100.go
//go:generate go fmt append_hi3516av100.go

//go:generate rm -f embed_hi3516av200.go
//go:generate rm -f append_hi3516av200.go
//go:generate rm -f append_hi3516av200.bin
//go:generate go run -tags "generate hi3516av200" ./generator.go --output hi3516av200 --tag "hi3516av200,koEmbed" --dir ../../sdk/hi3516av200/ko/ --pkg ko --source ./hi3516av200.go
//go:generate go fmt append_hi3516av200.go

//go:generate rm -f embed_hi3516cv100.go
//go:generate rm -f append_hi3516cv100.go 
//go:generate rm -f append_hi3516cv100.bin
//go:generate go run -tags "generate hi3516cv100" ./generator.go --output hi3516cv100 --tag "hi3516cv100,koEmbed" --dir ../../sdk/hi3516cv100/ko/ --pkg ko --source ./hi3516cv100.go
//go:generate go fmt append_hi3516cv100.go

//go:generate rm -f embed_hi3516cv200.go
//go:generate rm -f append_hi3516cv200.go 
//go:generate rm -f append_hi3516cv200.bin
//go:generate go run -tags "generate hi3516cv200" ./generator.go --output hi3516cv200 --tag "hi3516cv200,koEmbed" --dir ../../sdk/hi3516cv200/ko/ --pkg ko --source ./hi3516cv200.go
//go:generate go fmt append_hi3516cv200.go

//go:generate rm -f embed_hi3516cv300.go
//go:generate rm -f append_hi3516cv300.go 
//go:generate rm -f append_hi3516cv300.bin
//go:generate go run -tags "generate hi3516cv300" ./generator.go --output hi3516cv300 --tag "hi3516cv300,koEmbed" --dir ../../sdk/hi3516cv300/ko/ --pkg ko --source ./hi3516cv300.go
//go:generate go fmt append_hi3516cv300.go

//go:generate rm -f embed_hi3516cv500.go
//go:generate rm -f append_hi3516cv500.go 
//go:generate rm -f append_hi3516cv500.bin
//go:generate go run -tags "generate hi3516cv500" ./generator.go --output hi3516cv500 --tag "hi3516cv500,koEmbed" --dir ../../sdk/hi3516cv500/ko/ --pkg ko --source ./hi3516cv500.go
//go:generate go fmt append_hi3516cv500.go

//go:generate rm -f embed_hi3516ev200.go
//go:generate rm -f append_hi3516ev200.go 
//go:generate rm -f append_hi3516ev200.bin
//go:generate go run -tags "generate hi3516ev200" ./generator.go --output hi3516ev200 --tag "hi3516ev200,koEmbed" --dir ../../sdk/hi3516ev200/ko/ --pkg ko --source ./hi3516ev200.go
//go:generate go fmt append_hi3516ev200.go

//go:generate rm -f embed_hi3519av100.go
//go:generate rm -f append_hi3519av100.go 
//go:generate rm -f append_hi3519av100.bin
//go:generate go run -tags "generate hi3519av100" ./generator.go --output hi3519av100 --tag "hi3519av100,koEmbed" --dir ../../sdk/hi3519av100/ko/ --pkg ko --source ./hi3519av100.go
//go:generate go fmt append_hi3519av100.go

//go:generate rm -f embed_hi3559av100.go
//go:generate rm -f append_hi3559av100.go 
//go:generate rm -f append_hi3559av100.bin
//go:generate go run -tags "generate hi3559av100" ./generator.go --output hi3559av100 --tag "hi3559av100,koEmbed" --dir ../../sdk/hi3559av100/ko/ --pkg ko --source ./hi3559av100.go
//go:generate go fmt append_hi3559av100.go

package ko

