package main

import (
	//    "fmt"
	"flag"
	_ "net"
	//_ "expvar"
	"log"
	_ "net/http"
	_ "os"
	_ "time"

    "openhisiipcam.org/buildinfo"
	//"openhisiipcam.org/hisi"
	"openhisiipcam.org/openapi"

	//_ "openhisiipcam.org/scripts"

	//"github.com/yuin/gopher-lua"

    //_ "openhisiipcam.org/onvif"

    "openhisiipcam.org/hisi"
)

func main() {

    var VVV int
    VVV = 44
    log.Println(VVV)

	log.Println(buildinfo.Date)
	log.Println(buildinfo.Branch)
	log.Println(buildinfo.User)
	//TODO log.Println(himpp3.GetChipFamily())
	//TODO log.Println(himpp3.GetChip())
	//TODO log.Println(himpp3.GetCMOS())

	//portPtr := flag.Int("port", 80, "http port")
	//socketpathPtr := flag.String("socket path", "/tmp/app_minimal.sock", "UDS socket file path")

	flag.Parse()

	log.Printf("minimal application\n")
	//log.Printf("http port %d\n", *portPtr)
	//log.Printf("UDS socket file path %s\n", *socketpathPtr)

	///////////////////////////////////////////////////////////////////////////

    //log.Println("hi3516av200 swig check")
    //log.Println("HI_FALSE: ", hi3516av200.HI_FALSE)

    //hi3516av200.HI_MPI_SYS_Exit()
	hisi.MppInit()

	//himpp3.SysInit()
	//himpp3.MipiIspInit()
	//himpp3.ViInit()
	//himpp3.VpssInit()
	//himpp3.VencInit()

	///////////////////////////////////////////////////////////////////////////

	//router := openapi.NewRouter()
	openapi.Init()

	///////////////////////////////////////////////////////////////////////////

	/*
	   L := lua.NewState()
	   //defer L.Close()
	   if err := L.DoString(`print("hello from LUA\n")`); err != nil {
	       panic(err)
	   }
	   L.Close()
	*/
	///////////////////////////////////////////////////////////////////////////
	/*
	   log.Println("Starting USD HTTP server")

	   os.Remove("/tmp/app_minimal.sock")
	   l, err := net.Listen("unix", "/tmp/app_minimal.sock")
	   if err != nil {
	       log.Printf("error: %v\n", err)
	       return
	   }
	   go http.Serve(l, router)

	   log.Println("Starting NET HTTP server")
	   srv := &http.Server{
	       Addr:           ":80",
	       Handler:        router,
	       ReadTimeout:    10 * time.Second,
	       WriteTimeout:   10 * time.Second,
	       MaxHeaderBytes: 1 << 20,
	   }
	   srv.ListenAndServe()
	*/

	//https://blog.sgmansfield.com/2016/06/how-to-block-forever-in-go/
	select {}
}
