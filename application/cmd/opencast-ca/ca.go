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
    "net/url"

    "os"
    "mime/multipart"
    //"io"

    //"time"

    //dac "github.com/xinsnake/go-http-digest-auth-client"

    //"github.com/delphinus/go-digest-request"
    //"golang.org/x/net/context"

    "github.com/bobziuchkovski/digest"

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
    client0              *http.Client

    captureAdminHost    string
    captureAdminPath    string

    schedulerHost       string
    schedulerPath       string

    ingestHost          string
    ingestPath          string

    //ctx                 context.Context
    //rd                  *digestRequest.DigestRequest
    //t                   dac.DigestTransport

    //c                   *http.Client
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

    c.client0            = &http.Client{}

    //c.ctx = context.Background()
    //c.rd = digestRequest.New(ctx, c.user, c.pass)
    //c.rd = digestRequest.New(c.ctx, "opencast_system_account", "CHANGE_ME123")

    //c.t = dac.NewTransport(c.user, c.pass)
    //c.t = dac.NewTransport("opencast_system_account", "CHANGE_ME123")

    //t := digest.NewTransport("opencast_system_account", "CHANGE_ME123")
    t := digest.NewTransport(c.user, c.pass)
    c.client, _ = t.Client()
}

func (c *opencastClient) getService(service string) (string, string, error) {
    url := "http://"+c.host+":"+c.port+"/services/available.json?serviceType="+service
    request, _ := http.NewRequest("GET", url, nil)
    //request.SetBasicAuth(c.user, c.pass)

    request.Header.Set("X-Requested-Auth", "Digest")

    responce, err := c.client.Do(request)
    //responce, err := c.rd.Do(request)
    //responce, err := c.t.RoundTrip(request)

    if err != nil {
        fmt.Println("err:", err)
        return "", "", err
    }

    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)

    fmt.Println(string(body))

    serviceAnswer, _ := getResponce1(body)

    fmt.Println(serviceAnswer)

    return serviceAnswer.Services.Service.Host, serviceAnswer.Services.Service.Path, nil
}

func (c *opencastClient) register(state string) error {
    var dataStr = []byte(`state=`+state+`&address=http://test`)

    url := c.captureAdminHost+c.captureAdminPath+"/agents/"+c.name
    fmt.Println("url:", url)

    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataStr))
    //request.SetBasicAuth(c.user, c.pass)
    request.Header.Set("X-Requested-Auth", "Digest")

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
    //request.SetBasicAuth(c.user, c.pass)
    request.Header.Set("X-Requested-Auth", "Digest")

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

func (c *opencastClient) createMediaPackage() (string, error) {
    ///ingest/createMediaPackage
    url := c.ingestHost+c.ingestPath+"/createMediaPackage"
    fmt.Println("url:", url)

    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("X-Requested-Auth", "Digest")

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return "", err
    }
    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)
    fmt.Println(string(body))

    return string(body), nil
}

func (c *opencastClient) addDCCatalog(xml string, dublinCore string) (string, error) {
    //# Add DC catalog
    //curl -f --digest -u ${USER}:${PASSWORD} -H "X-Requested-Auth: Digest" \
    //"${HOST}/ingest/addDCCatalog" -F "mediaPackage=<${TMP_MP}" \
    //-F "dublinCore=<${TMP_DC}" -o "${TMP_MP}"

//    dublinCoreTmp := "<?xml version=\"1.0\" encoding=\"ISO-8859-1\" standalone=\"no\"?>\n"+
    dublinCoreTmp := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n"+
        "<dublincore xmlns=\"http://www.opencastproject.org/xsd/1.0/dublincore/\" "+
        "xmlns:dcterms=\"http://purl.org/dc/terms/\" "+
        "xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\">\n"+
        "<dcterms:creator>demo123</dcterms:creator>\n"+
        "<dcterms:contributor>demo123</dcterms:contributor>\n"+
        "<dcterms:created xsi:type=\"dcterms:W3CDTF\">2020-06-16T19:30Z</dcterms:created>\n"+
        "<dcterms:temporal xsi:type=\"dcterms:Period\">start=2020-06-16T19:30Z; end=2020-06-16T19:31Z; scheme=W3C-DTF;</dcterms:temporal>\n"+
        "<dcterms:description>demo123</dcterms:description>\n"+
        "<dcterms:subject>demo123</dcterms:subject>\n"+
        "<dcterms:language>demo123</dcterms:language>\n"+
        "<dcterms:spatial>pyca123</dcterms:spatial>\n"+
        "<dcterms:title>Demo event 123</dcterms:title>\n"+
        "</dublincore>"

    uri := c.ingestHost+c.ingestPath+"/addDCCatalog"
    fmt.Println("url:", uri)

    params := url.Values{}
	params.Add("mediaPackage", xml)
	params.Add("dublinCore", dublinCoreTmp)
    params.Add("flavor", "dublincore/episode")   

	var dataStr = []byte(params.Encode())

    fmt.Println(string(dataStr))

    //var dataStr = []byte(`mediaPackage=`+xml+`&dublinCore=`+dublinCoreTmp+`&flavor=dublincore/episode`)

    //request, _ := http.NewRequest("POST", url, ioutil.NopCloser(body))
    request, _ := http.NewRequest("POST", uri, ioutil.NopCloser(bytes.NewBuffer(dataStr)))
    //request, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataStr))
    //request.Header.Set("X-Requested-Auth", "Digest")
    //request.Header.Set("Content-type", "application/x-www-form-urlencoded")
    request.SetBasicAuth("admin", "opencast123")

    //responce, err := c.client.Do(request)
    responce, err := c.client0.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return "", err
    }
    defer responce.Body.Close()

    fmt.Println("responce code:", responce.Status)

    body, _ := ioutil.ReadAll(responce.Body)
    fmt.Println(string(body))

    return string(body), nil

}


func (c *opencastClient) addTrack(xml string, flavor string, path string) (string, error) {
    //# Add Track
    //curl -f --digest -u ${USER}:${PASSWORD} -H "X-Requested-Auth: Digest" \
    //"${HOST}/ingest/addTrack" -F flavor=presenter/source \
    //-F "mediaPackage=<${TMP_MP}" -F Body=@testvideo.mp4 -o "${TMP_MP}"

    //var err error

    
    file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

    fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	fi, err := file.Stat()
	if err != nil {
		return "", err
	}
	file.Close()
    

    body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

    params := map[string]string{
        "mediaPackage"  : xml,
        "flavor"        : flavor,
    }
    
    for key, val := range params {
        _ = writer.WriteField(key, val)
    }

    //fmt.Println(body.String())
    
    
	part, err := writer.CreateFormFile("BODY", fi.Name())
	if err != nil {
		return "", err
	}
	
    part.Write(fileContents)
  
	err = writer.Close()
	if err != nil {
		return "", err
	}
    fmt.Println(writer.FormDataContentType())

    //fmt.Println(body.String())

    //var dataStr = []byte(`mediaPackage=`+xml+`&flavor=`+flavor)
    //var dataStr = []byte(`mediaPackage=&flavor=`+flavor)
    /*
    paramsC := url.Values{}
    paramsC.Add("mediaPackage", xml)
    paramsC.Add("flavor", flavor)   
    
    var dataStr = []byte(paramsC.Encode())
    */

    uri := c.ingestHost+c.ingestPath+"/addTrack"
    fmt.Println("url:", uri)

    //request, _ := http.NewRequest("POST", uri, body)
    request, _ := http.NewRequest("POST", uri, ioutil.NopCloser(body))
    //request, _ := http.NewRequest("POST", uri, ioutil.NopCloser(bytes.NewBuffer(dataStr)))
    //request.Header.Set("X-Requested-Auth", "Digest")
    //request.Header.Set("Accept","/")
    request.Header.Set("Content-Type", writer.FormDataContentType())
    request.SetBasicAuth("admin", "opencast123")

    //responce, err2 := c.client.Do(request)
    responce, err2 := c.client0.Do(request)
    if err2 != nil {
        fmt.Println("err:", err2)
        return "", err2
    }
    defer responce.Body.Close()

    fmt.Println("responce code:", responce.Status)

    bodya, _ := ioutil.ReadAll(responce.Body)
    fmt.Println(string(bodya))

    return string(bodya), nil
}

func (c *opencastClient) ingest(xml string) (string, error) {
    //curl -f -v -i --digest -u ${USER}:${PASSWORD} \
    //-H "X-Requested-Auth: Digest" \
    //"${HOST}/ingest/ingest" \
    //-F "mediaPackage=<${TMP_MP}" -o "${FINAL}"

    //uri := c.ingestHost+c.ingestPath+"/ingest"
    uri := c.ingestHost+c.ingestPath+"/ingest/schedule-and-upload"
    fmt.Println("url:", uri)

    paramsC := url.Values{}
    paramsC.Add("mediaPackage", xml)
    paramsC.Add("workflowDefinitionId", "")
    paramsC.Add("workflowInstanceId", "")

    var dataStr = []byte(paramsC.Encode())

    fmt.Println(string(dataStr))

    request, _ := http.NewRequest("POST", uri, ioutil.NopCloser(bytes.NewBuffer(dataStr)))
    //request.Header.Set("X-Requested-Auth", "Digest")
    request.Header.Set("Content-type", "application/x-www-form-urlencoded")
    request.SetBasicAuth("admin", "opencast123")

    //responce, err2 := c.client.Do(request)
    responce, err2 := c.client0.Do(request)
    if err2 != nil {
        fmt.Println("err:", err2)
        return "", err2
    }
    defer responce.Body.Close()

    fmt.Println("responce code:", responce.Status)

    bodya, _ := ioutil.ReadAll(responce.Body)
    fmt.Println(string(bodya))

    return string(bodya), nil
}

func main() {
    fmt.Println("CA")

    var c opencastClient
    //c.create("84.201.135.192", "8080", "admin", "opencast123", "MY-TEST-CA")
    c.create("84.201.135.192", "8080", "opencast_system_account", "CHANGE_ME123", "MY-TEST-CA")

    c.captureAdminHost, c.captureAdminPath, _ = c.getService("org.opencastproject.capture.admin")
    c.schedulerHost,    c.schedulerPath,    _ = c.getService("org.opencastproject.scheduler")
    c.ingestHost,       c.ingestPath,       _ = c.getService("org.opencastproject.ingest")

    c.register("idle")
    c.getSchedule()

    xml, _ := c.createMediaPackage()
    fmt.Println(xml)

    xml2, _ := c.addDCCatalog(xml, "")
    fmt.Println(xml2)

    
    xml3, err := c.addTrack(xml2, "presenter/source", "./testvideo.mp4")
    //xml3, err := c.addTrack(xml2, "presenter/source", "./video.webm")
    if err != nil {
        fmt.Println("err:", err)
    } else {
        fmt.Println(xml3)
    }
    
    xml4, _ := c.ingest(xml3)
    fmt.Println(xml4)

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