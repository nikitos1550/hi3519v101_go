//The following directive is necessary to make the package coherent:

//+build ignore

//This program generates embed_*.go. It can be invoked by running
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

    "io"
    //"os/exec"
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

    cfg.Output      = "embed_"+output+".go"
    cfg.Package     = pkg
    cfg.Tags        = tag

    cfg.Prefix, _   = regexp.Compile(dir)


	var list []string

    //if mode != "minimal" {
        count := len(ko.ModulesList)
        list = make([]string, count)
        for i := range(list) {
            list[i] = ko.ModulesList[i][0]
        }
	//} else {
    //    count := len(ko.MinimalModulesList)
	//	list = make([]string, count)
	//	for i := range(list) {
	//		list[i] = ko.MinimalModulesList[i]
	//	}
	//}

    log.Println("Expecting " + strconv.Itoa(len(list)) + " files...")

    cfg.Input = make([]bindata.InputConfig, len(list))


    //TODO create append binary and go files

    appendGoFile, err := os.Create("./append_"+output+".go")
    if err != nil {
        panic(err)
    }

    appendBinFile, err := os.Create("./append_"+output+".bin")
    if err != nil {
        panic(err)
    }

    appendGoFile.WriteString("//+build !ignore,!generate\n")
    appendGoFile.WriteString("//+build "+output+",koAppend\n\n")
    appendGoFile.WriteString("package ko\n")
    appendGoFile.WriteString("var (\n")
    appendGoFile.WriteString("ModulesInfo = map[string] [2]uint{\n")

    //cmdTouch := exec.Command("touch", "./append_"+output+".bin")
    //cmdTouch.Run()

    var offsetAccum int64 = 0

    for i := range(list) {
        fileInfo, err := os.Stat(dir+"/"+list[i])

        if os.IsNotExist(err) {
            log.Fatal("File "+dir+""+list[i]+" doesn`t exist!")
        } else {
            log.Println("Adding file "+dir+""+list[i])
        }

        var inputConfigTmp bindata.InputConfig
        inputConfigTmp.Path = dir+""+list[i]
        cfg.Input[i] = inputConfigTmp

        //log.Println("cat "+dir+""+list[i]+" >> ./append_"+output+".bin")
        //cmdAppend := exec.Command("cat "+dir+""+list[i]+" >> ./append_"+output+".bin")
        //err = cmdAppend.Run()
        //if err != nil {
        //    panic(err)
        //}

        source, err2 := os.Open(dir+""+list[i])
        if err2 != nil {
                panic(err2)
        }

        buf := make([]byte, 1024*64)
        for {
                n, err := source.Read(buf)
                if err != nil && err != io.EOF {
                        panic(err)
                }
                if n == 0 {
                        break
                }

                if _, err := appendBinFile.Write(buf[:n]); err != nil {
                        panic(err)
                }
        }

        source.Close()

        appendGoFile.WriteString("\""+list[i]+"\": [2]uint{"+strconv.FormatInt(offsetAccum, 10)+", "+strconv.FormatInt(fileInfo.Size(), 10)+"},\n")
        offsetAccum = offsetAccum + fileInfo.Size()
    }

    appendGoFile.WriteString("}\n)\n")
    appendGoFile.Sync()

    err = bindata.Translate(cfg)
    if err != nil {
        log.Fatal("go-bindata error: ", err)
    }


    log.Println("Output file "+output)
}

