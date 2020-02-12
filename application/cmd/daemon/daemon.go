package main

import (
	"flag"
	"log"
	"os"
    "fmt"
    //"io/ioutil"

	"application/pkg/buildinfo"
	//"application/pkg/config"
	"application/pkg/openapi"
	"application/pkg/scripts"
	"application/pkg/mpp"

	"application/pkg/streamer"

    _"application/pkg/debug"
	_"application/pkg/utils/temperature"
	_"application/pkg/utils/chip"
)

func main() {
    //log.SetOutput(os.Stdout)
    //log.SetOutput(ioutil.Discard)
    flag.Usage = usage
	//flagVersion := flag.Bool("version", false, "Prints application version information")

	//log.Println("application daemon")
    flag.Parse()

    /*
    if *flagVersion {
		printVersion()
		os.Exit(0)
	}
    */

	//config.Init() //deprecated, use cmd and scripts

	openapi.Init() 	//openapi init should go first, becasue of -openapi-routes flag
					//same time, it will start serve requests immediately, but 
					//some requests need mpp and other initilization

	scripts.Init()	//

	mpp.Init()      //Init mpp and all subsystems
	streamer.Init() //Init streamers

    //Start serving after everything inited and setuped
	scripts.Start()
	openapi.Start()

	log.Println("daemon init done")
	select {} //pause this routine forever
}

func usage() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        printVersion()
        flag.PrintDefaults()
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
