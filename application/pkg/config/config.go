package config

import (
	"log"
	"os"
	"flag"
)

var flagConfigFile *string

func init() {
    flagConfigFile = flag.String("config", "/opt/config.json", "Config file path")
}

//TODO make config package looks similar as flag package
//other packages can add values that will be parsed and than they will get them in their scope vars

func Init() {
	//log.Println("ENVIRON")
	for _, pair := range os.Environ() {
		log.Println(pair)
	}
	//log.Println("-----")
}
