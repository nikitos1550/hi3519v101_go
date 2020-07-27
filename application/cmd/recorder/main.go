package main

import (
	"flag"
	"fmt"
	"os"
    "os/signal"
    "syscall"
    //"runtime"

	"application/core/compiletime"
	"application/core/mpp"
	"application/core/utils/chip"
	"application/core/logger"
    "application/core/utils/memparse"
)

func main() {

	flag.Usage = usage

	memTotal := flag.String("mem-total", "32M", "Total RAM size")
	memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest will be used for MPP")
	memMpp	 := flag.String("mem-mpp", "12M", "RAM size passed to MPP")

	chipCmd	 := flag.String("chip", compiletime.Family, "Chip app will be running on")

	flag.Parse()

	logger.Init()

    if compiletime.Family != "hi3516av200" && compiletime.Family != "hi3516cv500" {
        logger.Log.Fatal().
            Msg("This is test recorder only for hi3516av200, hi3516cv500 families")
    }

    //logger.Log.Debug().
    //    Int("NumCPU", runtime.NumCPU()).
    //    Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
    //    Msg("RUNTIME")

    ////////////////////////////////////////////////////////////////////////////

    //TODO move all run env to buildinfo
    var devInfo mpp.DeviceInfo

    devInfo.MemTotalSize = memparse.Str(*memTotal)
    devInfo.MemLinuxSize = memparse.Str(*memLinux)
    devInfo.MemMppSize = memparse.Str(*memMpp)

    devInfo.Chip = *chipCmd

    //devInfo.ViVpssOnline=true

    compiletime.Chip = *chipCmd //TODO temporary

    logger.Info().
        Uint64("mem-total", devInfo.MemTotalSize).
        Uint64("mem-linux", devInfo.MemLinuxSize).
        Uint64("mem-mpp", devInfo.MemMppSize).
        Str("chip", *chipCmd).
        Str("chip_detected", chip.Detect(chip.RegId())).
        Msg("Cmdline params")

    logger.Log.Info().
        Str("go", compiletime.GoVersion).
        Str("gcc", compiletime.GccVersion).
        Str("date", compiletime.BuildDateTime).
        Str("tags", compiletime.BuildTags).
        Str("commit", compiletime.BuildCommit).
        Str("sdk", compiletime.SDK).
        Str("cmos", compiletime.CmosProfile).
        Str("family", compiletime.Family).
        Msg("Compiletime information")


    mpp.Init(devInfo)      //Init mpp and all subsystems

    ////////////////////////////////////////////////////////////////////////////

    initPipeline()
    initArchive()
    initRecorder()

    httpServerStart()

    closeHandler()

	logger.Log.Info().Msg("Recorder ready")

	select {} //pause this routine forever
}

func closeHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
        logger.Log.Info().
            Msg("SIGTERM received")
		os.Exit(0)
	}()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

    fmt.Println("TODO CMOS info")
	//printVersion()
	//openapi.PrintInfo()
	//cmos.PrintInfo()
	flag.PrintDefaults()
}

