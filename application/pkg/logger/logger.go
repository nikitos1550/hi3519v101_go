package logger

import (
	"flag"

	"os"
	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/diode"
	//"fmt"
	//"time"

	"io"
	"io/ioutil"
)

var (
	Log zerolog.Logger

	defaultLevel int //zerolog.Level

	outputConsole bool
	outputFile string
	//outputApi bool
)

func init() {
	flag.IntVar(&defaultLevel, "logger-level", -1, "Logger level [-1 ... 5]") 

	flag.BoolVar(&outputConsole, "logger-output-console", true, "Write logger output to stdout")
	flag.StringVar(&outputFile, "logger-output-file", "", "Write logger output to file, assumed file name")
	//flag.BoolVar(&outputApi, "logger-output-api", true, "Expose http API for logger")

}

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if defaultLevel > 5 {
		defaultLevel = 5
	}
        if defaultLevel < -1 {
                defaultLevel = -1
        }

	zerolog.SetGlobalLevel(zerolog.Level(defaultLevel))
	//switch defaultLevel {
	//	case 0:
	//		zerolog.SetGlobalLevel(zerolog.DebugLevel)	
	//}

        //log.Print("hello world")
	/*
	if outputFile != "" {
		logFile, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			//do something
		}
	}
	*/

	initApiLog()

	var consoleWriter io.Writer
	var fileWriter io.Writer
	//var apiWriter io.Writer

	if outputConsole == true {
		//this doesn`t work on hi3516cv100, hi3519v101, host ok, other didn`t check
		/*
		wr := diode.NewWriter(os.Stdout, 1000, 0*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
		
		consoleWriter = zerolog.ConsoleWriter{Out: wr}
		*/
		consoleWriter = zerolog.ConsoleWriter{Out: os.Stdout}
	} else {
		consoleWriter = ioutil.Discard
	}
	fileWriter = ioutil.Discard
	//if outputApi == true {
	//	apiWriter = ioutil.Discard //&apilog
	//} else {
	//	apiWriter = ioutil.Discard
	//}

        //log := zerolog.New(os.Stdout)
	logWriter := zerolog.MultiLevelWriter(consoleWriter, fileWriter, apiWriter)
        Log = zerolog.New(logWriter).With().Timestamp().Logger()
        //log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	Log.Debug().
		Int("logger-level", defaultLevel).
		Bool("logger-output-console", outputConsole).
		Str("logger-output-file", outputFile).
		Msg("Logger started")
}

