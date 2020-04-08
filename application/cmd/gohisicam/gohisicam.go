package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"

	"application/pkg/buildinfo"
	//"application/pkg/config"
	"application/pkg/mpp"
	"application/pkg/openapi"
	"application/pkg/scripts"

	"application/pkg/mpp/cmos"

	"application/pkg/streamer"

	_ "application/pkg/debug"
	"application/pkg/ko"
	_ "application/pkg/utils/chip"
	_ "application/pkg/utils/temperature"
)

func main() {
	//log.SetOutput(os.Stdout)
	//log.SetOutput(ioutil.Discard)
	flag.Usage = usage
	//flagVersion := flag.Bool("version", false, "Prints application version information")
	memTotal := flag.String("mem-total", "32M", "Total RAM size") //&ko.MemTotal
	memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest will be used for MPP") //ko.MemLinux
	memMpp	 := flag.String("mem-mpp", "12M", "RAM size passed to MPP") //ko.MemMpp

	//log.Println("application daemon")
	flag.Parse()

	//TODO make correct memory size siffix handle
	*memTotal = strings.Trim(*memTotal, "M")
        ko.MemTotal, _ = strconv.ParseUint(*memTotal, 10, 64)
	
	*memLinux = strings.Trim(*memLinux, "M")
        ko.MemLinux, _  = strconv.ParseUint(*memLinux, 10, 64)

	*memMpp = strings.Trim(*memMpp, "M")
        ko.MemMpp, _     = strconv.ParseUint(*memMpp, 10, 64)

	log.Println("mem-total", ko.MemTotal)
	log.Println("mem-linux", ko.MemLinux)
	log.Println("mem-mpp", ko.MemMpp)

	cmdline, err := ioutil.ReadFile("/proc/cmdline")
	if err != nil {
		log.Println("Can`t read /proc/cmdline")
		os.Exit(0)
	}
	cmdlineStr := string(cmdline)
	re := regexp.MustCompile("mem=([0-9]*)M")
	foundLinuxMem := re.FindStringSubmatch(cmdlineStr)
	log.Println("CMDLINE Linux memory (MB) = ", foundLinuxMem[1])

	foundLinuxMemUint, _ := strconv.ParseUint(foundLinuxMem[1], 10, 32)
	if foundLinuxMemUint != uint64(ko.MemLinux) {
		log.Println("Linux mem mistmatch!")
		os.Exit(0)
	}
	/*
		    if *flagVersion {
				printVersion()
				os.Exit(0)
			}
	*/

	//config.Init() //deprecated, use cmd and scripts

	openapi.Init() //openapi init should go first, becasue of -openapi-routes flag
	//same time, it will start serve requests immediately, but
	//some requests need mpp and other initilization

	scripts.Init() //

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
	//openapi.PrintInfo()
	cmos.PrintInfo()
	flag.PrintDefaults()
}

func printVersion() {
	log.Println("Go: ", buildinfo.GoVersion)
	log.Println("Gcc: ", buildinfo.GccVersion)
	log.Println("Date: ", buildinfo.BuildDateTime)
	log.Println("Tags: ", buildinfo.BuildTags)
	//log.Println("User: ", buildinfo.BuildUser)
	//log.Println("Commit: ", buildinfo.BuildCommit)
	log.Println("Branch: ", buildinfo.BuildBranch)
	//log.Println("Vendor: ", buildinfo.BoardVendor)
	//log.Println("Model: ", buildinfo.BoardModel)
	//log.Println("Chip: ", buildinfo.Chip)
	log.Println("Cmos: ", buildinfo.CmosProfile)
	//log.Println("Total ram: ", buildinfo.TotalRam)
	//log.Println("Linux ram: ", buildinfo.LinuxRam)
	//log.Println("Mpp ram: ", buildinfo.MppRam)
}
