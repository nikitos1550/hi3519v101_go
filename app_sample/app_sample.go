package main

import (
	"fmt"
	"net/http"
)

// #cgo CFLAGS: -I"../mpp_hi3519_v101/include"
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/libisp.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/libmpi.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/libVoiceEngine.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/lib_hiae.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/lib_hiawb.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/lib_hiaf.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/libupvqe.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/libdnvqe.a
// #cgo LDFLAGS: ${SRCDIR}/../mpp_hi3519_v101/lib/lib_hidefog.a
// #include "hisi.h"
import "C"

// #include "mpi_sys.h"
// #include "mpi_vb.h"
// #include "mpi_isp.h"
// #include "mpi_vpss.h"
// #include "mpi_vi.h"

func main() {
	fmt.Printf("Test hi3519v101 imx174 go application\n")

	C.HI_MPI_SYS_Exit()
	C.HI_MPI_VB_Exit()

	//MPP_VERSION_S stVersion;
	//ret = HI_MPI_SYS_GetVersion(&stVersion);

	/*
		./hi_common.h-1619-#define VERSION_NAME_MAXLEN 64
		./hi_common.h:1650:typedef struct hiMPP_VERSION_S
		./hi_common.h-1681-{
		./hi_common.h-1683-    HI_CHAR aVersion[VERSION_NAME_MAXLEN];
		./hi_common.h:1726:} MPP_VERSION_S;
	*/
	var ver C.MPP_VERSION_S
	C.HI_MPI_SYS_GetVersion(&ver)

	fmt.Printf("MPP version %s\n", C.GoString(&ver.aVersion[0]))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Welcome to my website!")
		var ver C.MPP_VERSION_S
		C.HI_MPI_SYS_GetVersion(&ver)
		fmt.Fprintf(w, "MPP version %s", C.GoString(&ver.aVersion[0]))
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
