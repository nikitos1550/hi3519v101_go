package main

import (
	"flag"
	"log"
	"os"

	"application/pkg/buildinfo"
	"application/pkg/config"
	"application/pkg/openapi"
	"application/pkg/scripts"
	"application/pkg/mpp"
	"application/pkg/streamer"
	
	//TODO avoid implicit
	//packages with implicit init ( func init() )
	_"application/pkg/utils/temperature"
	_"application/pkg/utils/chip"
)

func main() {

	flagVersion := flag.Bool("version", false, "Prints application version information")

	log.Println("application daemon")
    flag.Parse()
    
    if *flagVersion {
		printVersion()
		os.Exit(0)
	}

	config.Init()

	openapi.Init() 	//openapi init should go first, becasue of -openapi-routes flag
					//same time, it will start serve requests immediately, but 
					//some requests need mpp and other initilization
	scripts.Init()	//

	mpp.Init()
	streamer.Init()

	openapi.Start()

	log.Println("daemon init done")

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
