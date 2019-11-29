package config

import (
	"log"
	"os"
)

func Init() {
	//log.Println("ENVIRON")
	for _, pair := range os.Environ() {
		log.Println(pair)
	}
	//log.Println("-----")
}
