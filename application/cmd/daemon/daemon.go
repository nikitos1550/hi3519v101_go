package main

import (
	"application/pkg/buildinfo"
	_ "application/pkg/koloader"
	_ "application/pkg/mpp"
	"application/pkg/openapi"
	_ "application/pkg/utils/chip"
	"application/pkg/utils/temperature"
	"flag"
	"log"
	"os"
)

func main() {

	flagVersion := flag.Bool("version", false, "Prints application version information")

	log.Println("application daemon")
    flag.Parse()
    
    if *flagVersion {
		printVersion()
		os.Exit(0)
	}

	//koloader.LoadMinimal()
	//mpp.Init()
	openapi.Init()
	temperature.Init()

	log.Println("loaded")

	select {}
}

func printVersion() {
	log.Println("Go: ", buildinfo.GoVersion)
	log.Println("Gcc: ", buildinfo.GccVersion)
	log.Println("Date: ", buildinfo.BuildDateTime)
	log.Println("Tags: ", buildinfo.BuildTags)
	log.Println("User: ", buildinfo.BuildUser)
	log.Println("Commit: ", buildinfo.BuildCommit)
	log.Println("Branch: ", buildinfo.BuildBranch)
	log.Println("Vendor: ", buildinfo.BoardVendor)
	log.Println("Model: ", buildinfo.BoardModel)
	log.Println("Chip: ", buildinfo.Chip)
	log.Println("Cmos: ", buildinfo.CmosProfile)
	log.Println("Total ram: ", buildinfo.TotalRam)
	log.Println("Linux ram: ", buildinfo.LinuxRam)
	log.Println("Mpp ram: ", buildinfo.MppRam)
}
