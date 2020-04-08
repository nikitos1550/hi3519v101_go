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
	"log"
	"net/http"
	"strings"
	"strconv"
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
	log.Println("GoHisiProbe")

        memTotal := flag.String("mem-total", "32M", "Total RAM size") //&ko.MemTotal
        memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest should be used for MPP") //ko.MemLinux
        memMpp   := flag.String("mem-mpp", "12M", "RAM size passed to MPP") //ko.MemMpp

        flag.Parse()

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

	log.Println("CMD parsed params:")
	log.Println("Total board RAM ", ko.MemTotal, "MB")
	log.Println("Linux RAM ", ko.MemLinux, "MB")
	log.Println("MPP RAM", ko.MemMpp, "MB")

	log.Println("Build time info:")
	log.Println("Go: ", buildinfo.GoVersion)
	log.Println("Gcc: ", buildinfo.GccVersion)
	log.Println("Date: ", buildinfo.BuildDateTime)
	log.Println("Tags: ", buildinfo.BuildTags)
	log.Println("Branch: ", buildinfo.BuildBranch)
	log.Println("Cmos: ", buildinfo.CmosProfile)
	log.Println("")

	log.Println("Loading modules...")
	ko.UnloadAll()
	ko.LoadMinimal()
	log.Println("Loading modules done")

	log.Println("Starting http server :80")
	http.HandleFunc("/", apiHandler)
	
	//TODO check errors
	http.ListenAndServe(":80", nil)
}
