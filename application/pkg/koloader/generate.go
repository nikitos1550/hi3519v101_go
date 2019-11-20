// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates hi35*******_kodata.go. It can be invoked by running
// make generate

package main

import (
    "log"
    "github.com/shuLhan/go-bindata"
    "path/filepath"
    "regexp"
    "flag"
    "application/pkg/koloader" //MAYBE replace by go/parser?
    "os"
    "strconv"
)

var (
    output  string
    tag     string
    source  string
    pkg     string
    dir     string
)

func main() {
    log.Println("Ko binary embedded generator (go-bindata based)");

    flag.StringVar(&output, "output",   "", "help")
    flag.StringVar(&tag,    "tag",      "", "help")
    flag.StringVar(&source, "source",   "", "help")
    flag.StringVar(&pkg,    "pkg",      "", "help")
    flag.StringVar(&dir,    "dir",      "", "help")

    flag.Parse()

    dir     = filepath.Clean(dir)
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

    log.Println("Expecting " + strconv.Itoa(len(koloader.Modules)) + " files...")

    cfg.Input = make([]bindata.InputConfig, len(koloader.Modules))

    for i := range(koloader.Modules) {
        _, err := os.Stat(dir+"/"+koloader.Modules[i][0])

        if os.IsNotExist(err) {
            log.Fatal("File "+dir+"/"+koloader.Modules[i][0]+" doesn`t exist!")
        } else {
            log.Println("Adding file "+dir+"/"+koloader.Modules[i][0])
        }

        var inputConfigTmp bindata.InputConfig
        inputConfigTmp.Path = dir+"/"+koloader.Modules[i][0]
        cfg.Input[i] = inputConfigTmp
    }

    err := bindata.Translate(cfg)
    if err != nil {
        log.Fatal("go-bindata error: ", err)
    }

    log.Println("Output file "+output)
}

