package main

import (
	"application/pkg/buildinfo"
	"application/pkg/ko"
	"application/pkg/mpp/utils"
	"application/pkg/utils/chip"
	"application/pkg/utils/temperature"
	"encoding/json"
	"flag"
	"fmt"
	//"log"
	"net/http"
	"strings"
	"strconv"
	//"time"

	//"os"
	//"github.com/rs/zerolog"
    	//"github.com/rs/zerolog/log"
	"application/pkg/logger"
)

type answerSchema struct {
	App string `json:"appName"`

	//ChipDetectedReg string `json:"chipDetectedReg"`
	ChipDetectedMpp string          `json:"chipDetectedMpp"`

	Mpp string `json:"mppVersion"`

	//SysIdReg uint32 `json:"chipIdReg"`
	//SysIdMpp        uint32          `json:"chipIdMpp"`

	TempVal float32 `json:"temperature,omitempty"`
	TempHW  bool  `json:"temperatureHW"`

	Info buildinfo.Info `json:"buildInfo"`
}

var (
	memTotal uint
	memLinux uint
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("ytytyt")

	var schema answerSchema

	schema.App = "GoHisiProbe"

	//schema.ChipDetectedReg = chip.Detect(chip.RegId())
	schema.ChipDetectedMpp  = chip.Detect(utils.MppId())

	schema.Mpp = utils.Version()

	//schema.SysIdReg = chip.RegId()
	//schema.SysIdMpp = utils.MppId()

	var err error
	schema.TempVal, err = temperature.Get()

	if err != nil {
		schema.TempHW = false
	} else {
		schema.TempHW = true
	}

	buildinfo.CopyAll(&schema.Info)
	schema.Info.CmosProfile = ""

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
	fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    	//log.Print("hello world")

	//log := zerolog.New(os.Stdout)
	//logWriter := zerolog.MultiLevelWriter(os.Stdout, zerolog.ConsoleWriter{Out: os.Stderr})
	//log := zerolog.New(logWriter).With().Timestamp().Logger()
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//log.Print("GoHisiProbe")

        memTotal := flag.String("mem-total", "32M", "Total RAM size") //&ko.MemTotal
        memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest should be used for MPP") //ko.MemLinux
        memMpp   := flag.String("mem-mpp", "12M", "RAM size passed to MPP") //ko.MemMpp
	
	httpPort := flag.Uint("http-port", 80, "Web server port")

        flag.Parse()

	logger.Init()

        //TODO make correct memory size siffix handle
        *memTotal = strings.Trim(*memTotal, "M")
        ko.MemTotal, _ = strconv.ParseUint(*memTotal, 10, 64)
        
        *memLinux = strings.Trim(*memLinux, "M")
        ko.MemLinux, _  = strconv.ParseUint(*memLinux, 10, 64)

        *memMpp = strings.Trim(*memMpp, "M")
        ko.MemMpp, _     = strconv.ParseUint(*memMpp, 10, 64)

        //log.Println("mem-total", ko.MemTotal)
        //log.Println("mem-linux", ko.MemLinux)
        //log.Println("mem-mpp", ko.MemMpp)

	logger.Info().
		Uint64("mem-total", ko.MemTotal).
		Uint64("mem-linux", ko.MemLinux).
                Uint64("mem-mm", ko.MemMpp).
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

	//log.Print("Loading modules...")
	ko.UnloadAll()
	ko.LoadMinimal()
	//log.Print("Loading modules done")

	http.HandleFunc("/", apiHandler)

	if *httpPort == 0 {
		*httpPort = 80
	}
	if *httpPort > 65536 {
		*httpPort = 80
	}
	port := strconv.Itoa(int(*httpPort))
	logger.Log.Info().
        	Uint("port", *httpPort).
                Msg("Starting http server")
	
	//logger.Log.Panic().Msg("Test panic")

	//TODO check errors
	http.ListenAndServe(":"+port, nil)

        logger.Log.Error().
                Msg("TODO Something wrong with http server")

	//time.Sleep(20*time.Millisecond)
}
