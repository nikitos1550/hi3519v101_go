package main

//Reference
//https://docs.opencast.org/r/8.x/developer/#modules/capture-agent/capture-agent/
//https://github.com/opencast/pyCA/blob/master/pyca/utils.py

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
    _"net/url"

    //"time"

    _"strings"
    _"github.com/arran4/golang-ical"
)

//84.201.135.192
//http://84.201.135.192:8080/services/available.json?serviceType=org.opencastproject.capture.admin

/*
{
   "services":{
      "service":{
         "type":"org.opencastproject.capture.admin",
         "host":"http://84.201.135.192:8080",
         "path":"/capture-admin",
         "active":true,
         "online":true,
         "maintenance":false,
         "jobproducer":false,
         "onlinefrom":"2020-06-06T13:01:36.694Z",
         "service_state":"NORMAL",
         "state_changed":"2020-06-06T13:01:36.694Z",
         "error_state_trigger":0,
         "warning_state_trigger":0
      }
   }
}
*/

type Responce1JSON struct {
    Services ServicesJSON `json:"services"`
}

type ServicesJSON struct {
    Service ServiceJSON `json:"service"`
}

type ServiceJSON struct {
    ServiceType             string  `json:"type"`
    Host                    string  `json:"host"`
    Path                    string  `json:"path"`
    Active                  bool    `json:"active"`
    Online                  bool    `json:"online"`
    Maintenance             bool    `json:"maintenance"`
    Jobproducer             bool    `json:"jobproducer"`
    Onlinefrom              string  `json:"onlinefrom"`
    Service_state           string  `json:"service_state"`
    StateChanged            string  `json:"state_changed"`
    ErrorStateTrigger       uint    `json:"error_state_trigger"`
    WarningStateTrigger     uint    `json:"warning_state_trigger"`
}

func getResponce1(body []byte) (*Responce1JSON, error) {
    var s = new(Responce1JSON)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

type opencastClient struct {
    host                string
    port                string

    user                string
    pass                string

    name                string

    client              *http.Client

    captureAdminHost    string
    captureAdminPath    string

    schedulerHost       string
    schedulerPath       string

    ingestHost          string
    ingestPath          string

}

func (c *opencastClient) create(host string, port string, user string, pass string, name string) {
    c.host              = host
    c.port              = port

    //tmp
    //c.captureAdminHost  = host
    //c.captureAdminPath  = port
    //c.schedulerHost     = host
    //c.schedulerPath     = port
    //c.ingestHost        = host
    //c.ingestPath        = port

    c.user              = user
    c.pass              = pass

    c.name              = name

    c.client            = &http.Client{}
}

func (c *opencastClient) getService(service string) (string, string, error) {
    url := "http://"+c.host+":"+c.port+"/services/available.json?serviceType="+service
    request, _ := http.NewRequest("GET", url, nil)
    request.SetBasicAuth(c.user, c.pass)

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return "", "", err
    }

    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)

    fmt.Println(string(body))

    serviceAnswer, _ := getResponce1(body)

    fmt.Println(serviceAnswer)

    return serviceAnswer.Services.Service.Host,serviceAnswer.Services.Service.Path,nil
}

func (c *opencastClient) register(state string) error {
    var dataStr = []byte(`state=`+state+`&address=http://test`)

    url := c.captureAdminHost+c.captureAdminPath+"/agents/"+c.name
    fmt.Println("url:", url)

    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataStr))
    request.SetBasicAuth(c.user, c.pass)

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return err
    }

    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)
    
    fmt.Println(string(body))

    return nil
}

func (c *opencastClient) getSchedule() error {

    //${SCHEDULER-ENDPOINT}/calendars?agentid=&cutoff=
    //Use the HTTP ETag and If-Not-Modified header to have Opencast only sent schedules when they have actually changed.

    url := c.schedulerHost+c.schedulerPath+"/calendars?agentid="+c.name
    fmt.Println("url:", url)

    request, _ := http.NewRequest("GET", url, nil)
    request.SetBasicAuth(c.user, c.pass)

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return err
    }
    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)
    fmt.Println(string(body))

    //cal, _ := ics.ParseCalendar(strings.NewReader(string(body5)))
    //fmt.Println(cal)
    
    return nil
}

func (c *opencastClient) testSendMedia() error {

    return nil
}

func main() {
    fmt.Println("CA")

    var c opencastClient
    c.create("84.201.135.192", "8080", "admin", "opencast123", "MY-TEST-CA")

    c.captureAdminHost, c.captureAdminPath, _ = c.getService("org.opencastproject.capture.admin")
    c.schedulerHost,    c.schedulerPath,    _ = c.getService("org.opencastproject.scheduler")
    c.ingestHost,       c.ingestPath,       _ = c.getService("org.opencastproject.ingest")

    c.register("idle")
    c.getSchedule()


//////////////---------------
//    params := url.Values{}
//	params.Add("configuration", "{'capture.device.names':'MOCK_SCREEN,MOCK_PRESENTER,MOCK_MICROPHONE'}")
//
//    //fmt.Println(url.QueryEscape(`configuration={'capture.device.names':'MOCK_SCREEN,MOCK_PRESENTER,MOCK_MICROPHONE2'}`))
//    fmt.Println(params.Encode())
//    var confStr = []byte(params.Encode())
//    //var confStr = []byte(`configuration=%7B'capture.device.names'%3A'MOCK_SCREEN,MOCK_PRESENTER,MOCK_MICROPHONE2'%7D`)
//    //var confStr = []byte(`%7B'capture.device.names'%3A'MOCK_SCREEN,MOCK_PRESENTER,MOCK_MICROPHONE222'%7D`)
//
//    //capture-admin/agents/$AGENT_NAME/configuration
//    req3url := responce.Services.Service.Host+responce.Services.Service.Path+"/agents/"+name+"/configuration"
//    fmt.Println("req3url:", req3url)
//
//    req3, _ := http.NewRequest("POST", req3url, bytes.NewBuffer(confStr))
//    req3.SetBasicAuth(user, pass)
//    //req3.Header.Set("Content-type", "application/json")
//
//    resp3, err3 := client.Do(req3)
//    if err3 != nil {
//        fmt.Println("err3:", err3)
//        return
//    }
//    defer resp3.Body.Close()
//    body3, _ := ioutil.ReadAll(resp3.Body)
//    fmt.Println(string(body3))
/////////////-------------------------



}
