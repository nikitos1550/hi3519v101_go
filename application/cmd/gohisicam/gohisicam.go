package main

import (
	"flag"
	"fmt"
	//"io/ioutil"
	//"log"
	"os"
	//"regexp"
	//"strings"
	//"strconv"

	"application/pkg/buildinfo"
	//"application/pkg/config"
	"application/pkg/mpp"
	"application/pkg/openapi"
	"application/pkg/scripts"

	//"application/pkg/mpp/cmos"

	"application/pkg/processing"

	"application/pkg/streamer"

	_ "application/pkg/debug"
	//"application/pkg/ko"
	_ "application/pkg/utils/chip"
	_ "application/pkg/utils/temperature"

	"application/pkg/logger"

    "application/pkg/utils/memparse"
)

func main() {
	flag.Usage = usage
	//flagVersion := flag.Bool("version", false, "Prints application version information")

	memTotal := flag.String("mem-total", "32M", "Total RAM size") //&ko.MemTotal
	memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest will be used for MPP") //ko.MemLinux
	memMpp	 := flag.String("mem-mpp", "12M", "RAM size passed to MPP") //ko.MemMpp

	chip	 := flag.String("chip", buildinfo.Family, "Chip app will be running on")
	
	flag.Parse()

	logger.Init()

    var devInfo mpp.DeviceInfo

    devInfo.MemTotalSize = memparse.Str(*memTotal)
    devInfo.MemLinuxSize = memparse.Str(*memLinux)
    devInfo.MemMppSize = memparse.Str(*memMpp)

    devInfo.Chip = *chip

    logger.Info().
        Uint64("mem-total", devInfo.MemTotalSize).
        Uint64("mem-linux", devInfo.MemLinuxSize).
        Uint64("mem-mpp", devInfo.MemMppSize).
        Str("chip", *chip).
        Msg("cmdline mem params")

    logger.Log.Info().
        Str("go", buildinfo.GoVersion).
        Str("gcc", buildinfo.GccVersion).
        Str("date", buildinfo.BuildDateTime).
        Str("tags", buildinfo.BuildTags).
        Str("branch", buildinfo.BuildBranch).
        Str("sdk", buildinfo.SDK).
        Str("cmos", buildinfo.CmosProfile).
        Msg("build info")

	//ko.CmosName = buildinfo.CmosProfile


	/*
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
	*/

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

	mpp.Init(devInfo)      //Init mpp and all subsystems
	streamer.Init() //Init streamers
	processing.Init()

	//Start serving after everything inited and setuped
	scripts.Start()
	openapi.Start()

	//log.Println("daemon init done")
	logger.Log.Info().Msg("GoHisiCam started")
	select {} //pause this routine forever
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	//printVersion()
	//openapi.PrintInfo()
	//cmos.PrintInfo()
	flag.PrintDefaults()
}

