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

	ChipDetectedReg string `json:"chipDetectedReg"`
	ChipDetectedMpp string          `json:"chipDetectedMpp"`

	Mpp string `json:"mppVersion"`

	SysIdReg uint32 `json:"chipIdReg"`
	SysIdMpp        uint32          `json:"chipIdMpp"`

	TempVal float32 `json:"temperature"`
	TempHW  string  `json:"temperatureHW"`

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

	schema.ChipDetectedReg = chip.Detect(chip.RegId())
	schema.ChipDetectedMpp  = chip.Detect(utils.MppId())

	schema.Mpp = utils.Version()

	schema.SysIdReg = chip.RegId()
	schema.SysIdMpp = utils.MppId()

	var err error
	schema.TempVal, err = temperature.Get()

	if err != nil {
		schema.TempHW = "not availible"
	} else {
		schema.TempHW = "availible"
	}

	buildinfo.CopyAll(&schema.Info)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
	fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
	log.Println("GoHisiProbe")

        memTotal := flag.String("mem-total", "32M", "Total RAM size") //&ko.MemTotal
        memLinux := flag.String("mem-linux", "20M", "RAM size passed to Linux kernel, rest will be used for MPP") //ko.MemLinux
        memMpp   := flag.String("mem-mpp", "12M", "RAM size passed to MPP") //ko.MemMpp

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


	log.Println("CMD parsed params:")
	log.Println("Total board RAM ", ko.MemTotal, "MB")
	log.Println("Linux RAM ", ko.MemLinux, "MB")
	log.Println("")

	log.Println("Build time info:")
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
