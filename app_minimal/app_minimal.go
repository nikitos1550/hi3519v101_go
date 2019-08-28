package main

import (
//    "fmt"
    "flag"
    "net"
    "net/http"
    "time"
    "log"
    "os"
    "./internal/openapi"
    "./internal/himpp3"
    "./internal/info"
)

func main() {

	log.Println(info.DATE)
    log.Println(info.BRANCH)
    log.Println(info.USER)
    log.Println(himpp3.GetChipFamily())
    log.Println(himpp3.GetChip())
    log.Println(himpp3.GetCMOS())


    portPtr := flag.Int("port", 80, "http port")
    socketpathPtr := flag.String("socket path", "/tmp/app_minimal.sock", "UDS socket file path")

    flag.Parse()

    log.Printf("minimal application\n");
    log.Printf("http port %d\n", *portPtr);
    log.Printf("UDS socket file path %s\n", *socketpathPtr);

    ///////////////////////////////////////////////////////////////////////////

	himpp3.SysInit()
	himpp3.MipiIspInit()
	himpp3.ViInit()
	himpp3.VpssInit()
	himpp3.VencInit()

    ///////////////////////////////////////////////////////////////////////////

    router := openapi.NewRouter()

    ///////////////////////////////////////////////////////////////////////////
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

}
