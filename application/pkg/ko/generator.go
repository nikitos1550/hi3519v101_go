//The following directive is necessary to make the package coherent:

//+build ignore

//This program generates kobin_hi35****.go. It can be invoked by running
//make generate

package main

import (
    "log"
    "github.com/shuLhan/go-bindata"
    "path/filepath"
    "regexp"
    "flag"
    "application/pkg/ko" //MAYBE replace by go/parser?
    "os"
    "strconv"
)

var (
    output  string
    tag     string
    source  string
    pkg     string
    dir     string
    mode    string
)

func main() {
    log.Println("Ko binary embedded generator (go-bindata based)");

    flag.StringVar(&output, "output",   "", "help")
    flag.StringVar(&tag,    "tag",      "", "help")
    flag.StringVar(&source, "source",   "", "help")
    flag.StringVar(&pkg,    "pkg",      "", "help")
    flag.StringVar(&dir,    "dir",      "", "help")
    flag.StringVar(&mode,   "mode",     "", "help")

    flag.Parse()

    //dir     = filepath.Clean(dir)
    output  = filepath.Clean(output)
    source  = filepath.Clean(source)

    //TODO check input options

    cfg := bindata.NewConfig()

    cfg.Debug       = false
    cfg.Dev         = false
    cfg.MD5Checksum = false
    cfg.NoCompress  = true
    cfg.NoMemCopy   = true
    cfg.NoMetadata  = true
    cfg.Split       = false

    cfg.Output      = output
    cfg.Package     = pkg
    cfg.Tags        = tag

    cfg.Prefix, _   = regexp.Compile(dir)

    //var list *[][2]string
    //list := ko.ModulesList
    //list := ko.MinimalModulesList
    	//count := len(ko.ModulesList)

	var list []string
	if mode != "minimal" {
        	count := len(ko.ModulesList)
        	list = make([]string, count)
        	for i := range(list) {
                	list[i] = ko.ModulesList[i][0]
        	}
	} else {
	
		count := len(ko.MinimalModulesList)
		list = make([]string, count)
		for i := range(list) {
			list[i] = ko.MinimalModulesList[i]
		}
	}

    //if mode == "minimal" {
    //	list = ko.MinimalModulesList
    //}

    log.Println("Expecting " + strconv.Itoa(len(list)) + " files...")

    cfg.Input = make([]bindata.InputConfig, len(list))

    for i := range(list) {
        _, err := os.Stat(dir+"/"+list[i])

        if os.IsNotExist(err) {
            log.Fatal("File "+dir+"/"+list[i]+" doesn`t exist!")
        } else {
            log.Println("Adding file "+dir+"/"+list[i])
        }

        var inputConfigTmp bindata.InputConfig
        inputConfigTmp.Path = dir+"/"+list[i]
        cfg.Input[i] = inputConfigTmp
    }

    err := bindata.Translate(cfg)
    if err != nil {
        log.Fatal("go-bindata error: ", err)
    }

    log.Println("Output file "+output)
}

