//+build arm
//+build hi3516av200
//+build openapi

package isp

import (
    "log"
    "fmt"
    "net/http"
    "application/pkg/openapi"
    "encoding/json"
)

/*
HI_S32 HI_MPI_ISP_GetFMWState(ISP_DEV IspDev, ISP_FMW_STATE_E *penState);
HI_S32 HI_MPI_ISP_GetModuleControl(ISP_DEV IspDev, ISP_MODULE_CTRL_U *punModCtrl);
HI_S32 HI_MPI_ISP_QueryInnerStateInfo(ISP_DEV IspDev, ISP_INNER_STATE_INFO_S *pstInnerStateInfo)
*/

func init() {
    openapi.AddApiRoute("serveIspStat", "/mpp/isp/stat", "GET", serveIspStat)

    //openapi.AddApiRoute("serveIspModules", "/mpp/isp/modules", "GET", serveIspModules)
    //openapi.AddApiRoute("setIspModules", "/mpp/isp/modules", "POST", setIspModules)

    //openapi.AddApiRoute("serveIspState", "/mpp/isp/state", "GET", serveIspState)

}

type serveIspStatSchema struct {
    FMWState    bool    `json:"FMWState"`
}

func serveIspStat(w http.ResponseWriter, r *http.Request) {
    log.Println("mpp.isp.serveIspStat")

    var schema serveIspStatSchema

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

