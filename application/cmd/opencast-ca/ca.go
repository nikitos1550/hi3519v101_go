package main

//Reference
//https://docs.opencast.org/r/8.x/developer/#modules/capture-agent/capture-agent/
//https://github.com/opencast/pyCA/blob/master/pyca/utils.py

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os/exec"
    "strconv"
    "strings"
    "time"

    "mime/multipart"
    "os"
    //"io"

    //dac "github.com/xinsnake/go-http-digest-auth-client"

    //"github.com/delphinus/go-digest-request"
    //"golang.org/x/net/context"

    "github.com/bobziuchkovski/digest"

    _ "strings"
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
    activeRecords map[string]bool
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
    c.activeRecords = make(map[string]bool)
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

    //fmt.Println(string(body))

    serviceAnswer, _ := getResponce1(body)

    //fmt.Println(serviceAnswer)

    return serviceAnswer.Services.Service.Host, serviceAnswer.Services.Service.Path, nil
}

func (c *opencastClient) register(state string) error {
    var dataStr = []byte(`state=`+state+`&address=http://test`)

    url := c.captureAdminHost+c.captureAdminPath+"/agents/"+c.name
    //fmt.Println("url:", url)

    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(dataStr))
    //request.SetBasicAuth(c.user, c.pass)
    request.Header.Set("X-Requested-Auth", "Digest")

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return err
    }

    defer responce.Body.Close()

    //body, _ := ioutil.ReadAll(responce.Body)
    //fmt.Println(string(body))

    return nil
}

func (c *opencastClient) getSchedule() (error,string) {

    //${SCHEDULER-ENDPOINT}/calendars?agentid=&cutoff=
    //Use the HTTP ETag and If-Not-Modified header to have Opencast only sent schedules when they have actually changed.

    url := c.schedulerHost+c.schedulerPath+"/calendars?agentid="+c.name
    //fmt.Println("url:", url)

    request, _ := http.NewRequest("GET", url, nil)
    //request.SetBasicAuth(c.user, c.pass)
    request.Header.Set("X-Requested-Auth", "Digest")

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return err,""
    }
    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)
    //fmt.Println(string(body))

    //cal, _ := ics.ParseCalendar(strings.NewReader(string(body5)))
    //fmt.Println(cal)
    
    return nil, string(body)
}

func (c *opencastClient) createMediaPackage() (string, error) {
    ///ingest/createMediaPackage
    url := c.ingestHost+c.ingestPath+"/createMediaPackage"
    //fmt.Println("url:", url)

    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("X-Requested-Auth", "Digest")

    responce, err := c.client.Do(request)
    if err != nil {
        fmt.Println("err:", err)
        return "", err
    }
    defer responce.Body.Close()

    body, _ := ioutil.ReadAll(responce.Body)
    //fmt.Println(string(body))

    return string(body), nil
}

func (c *opencastClient) addDCCatalog(xml string, startTimestamp uint64, endTimestamp uint64) (string, error) {
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
        "<dcterms:created xsi:type=\"dcterms:W3CDTF\">" + time.Now().Format(time.RFC3339)+ "</dcterms:created>\n"+
        "<dcterms:temporal xsi:type=\"dcterms:Period\">start=" + time.Unix(0, int64(startTimestamp*1000)).Format(time.RFC3339) + "; end=" + time.Unix(0, int64(endTimestamp*1000)).Format(time.RFC3339) + "; scheme=W3C-DTF;</dcterms:temporal>\n"+
        "<dcterms:description>demo123</dcterms:description>\n"+
        "<dcterms:subject>demo123</dcterms:subject>\n"+
        "<dcterms:language>demo123</dcterms:language>\n"+
        "<dcterms:spatial>pyca123</dcterms:spatial>\n"+
        "<dcterms:title>Demo event 123</dcterms:title>\n"+
        "</dublincore>"

    uri := c.ingestHost+c.ingestPath+"/addDCCatalog"
    //fmt.Println("url:", uri)

    params := url.Values{}
	params.Add("mediaPackage", xml)
	params.Add("dublinCore", dublinCoreTmp)
    params.Add("flavor", "dublincore/episode")   

	var dataStr = []byte(params.Encode())

    //fmt.Println(string(dataStr))

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

    //fmt.Println("responce code:", responce.Status)

    body, _ := ioutil.ReadAll(responce.Body)
    //fmt.Println(string(body))

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
    //fmt.Println("url:", uri)

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
    //fmt.Println("url:", uri)

    paramsC := url.Values{}
    paramsC.Add("mediaPackage", xml)
    paramsC.Add("workflowDefinitionId", "")
    paramsC.Add("workflowInstanceId", "")

    var dataStr = []byte(paramsC.Encode())

    //fmt.Println(string(dataStr))

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

func startRecord(api *CameraApi, startTimestamp uint64, stopTimestamp uint64) (string,int,int,int) {
    channelId := 1
    res := api.CreateChannel(channelId, 1920, 1080, 30)
    if (!res){
        fmt.Println("CreateChannel failed ")
        return "",0,0,0
    }
    fmt.Println("Channel was created ", channelId)

    params := make(map[string]string)
    params["StartTimestamp"] = strconv.FormatUint(startTimestamp, 10)
    params["StopTimestamp"] = strconv.FormatUint(stopTimestamp, 10)
    res, processingId := api.CreateProcessing("schedule", params)
    if (!res){
        fmt.Println("Processing failed ")
        return "",0,0,0
    }
    fmt.Println("Processing was created ", processingId)

    res, encoderId := api.CreateEncoder("H264_1920_1080_1M")
    if (!res){
        fmt.Println("Encoder failed ")
        return "",0,0,0
    }
    fmt.Println("Encoder was created ", encoderId)

    res = api.SubscribeChannel(processingId, channelId)
    if (!res){
        fmt.Println("SubscribeChannel failed ")
        return "",0,0,0
    }
    fmt.Println("Channel was subscribed ")

    res = api.SubscribeProcessing(processingId, encoderId)
    if (!res){
        fmt.Println("SubscribeProcessing failed ")
        return "",0,0,0
    }
    fmt.Println("Processing was subscribed ")

    res, recordId := api.StartRecording(encoderId)
    if (!res){
        fmt.Println("StartRecording failed ")
        return "",0,0,0
    }
    fmt.Println("Record was started ", recordId)
    return recordId,channelId,processingId,encoderId
}

func stopRecord(api *CameraApi, recordId string) {
    var prevSize uint = 0
    for {
        res, records := api.GetAllRecords()
        if (!res){
            fmt.Println("GetAllRecords failed ")
            return
        }
        for _, record := range records.Details {
            if (record.RecordId == recordId){
                if (prevSize != 0 && prevSize == record.Size){
                    res = api.StopRecording(recordId)
                    if res {
                        return
                    }
                }

                prevSize = record.Size
            }
        }

        time.Sleep(time.Second)
    }
}

func waitForFinish(api *CameraApi, recordId string) {
    for {
        res, records := api.GetAllRecords()
        if (!res){
            fmt.Println("GetAllRecords failed ")
            return
        }
        for _, record := range records.Details {
            if (record.RecordId == recordId && record.Status == "finished"){
                return
            }
        }

        time.Sleep(time.Second)
    }
}

func resetCam(api *CameraApi, channelId int, processingId int, encoderId int) {
    api.StopEncoder(encoderId)
    api.StopProcessing(processingId)
    api.StopChannel(channelId)
}

func packVideo(recordId string) bool {
    ffmpegPath := "ffmpeg"
    videoFolder := "/home/cam/cam1/" + recordId + "/"
    videoPath := videoFolder + "out0.h264"
    outPath := "out.mp4"

    cmd := exec.Command(ffmpegPath, "-i", videoPath, "-c:v", "copy", "-f", "mp4", outPath)
    out, err := cmd.CombinedOutput()
    if (err != nil){
        log.Println("ffmpeg failed", err)
        log.Println(string(out))
        return false
    }

    log.Println("Video packed")
    log.Println(string(out))
    return true
}

func convertTime(original string) string{
    return original[:4] + "-" + original[4:6] + "-" + original[6:11] + ":" + original[11:13] + ":" + original[13:]
}

func parseSchedule(schedule string, activeRecords map[string]bool) (string, uint64, uint64){
    uid := ""
    var start uint64 = 0
    var stop uint64 = 0
    lines := strings.Split(schedule, "\r\n")
    for _, line := range lines {
        if (line == "END:VEVENT"){
            if (uid != "" && start != 0 && stop != 0){
                break
            }
        }
       values := strings.Split(line, ":")
        if (values[0] == "UID"){
            _, exists := activeRecords[values[1]]
            if (exists) {
                continue
            }
            uid = values[1]
        }
        if (values[0] == "DTSTART"){
            t, err := time.Parse(time.RFC3339, convertTime(values[1]))
            if (err == nil){
                start = uint64(t.UnixNano() / 1000)
            } else {
                fmt.Println(err)
            }
        }
        if (values[0] == "DTEND"){
            t, err := time.Parse(time.RFC3339, convertTime(values[1]))
            if (err == nil){
                stop = uint64(t.UnixNano() / 1000)
            } else {
                fmt.Println(err)
            }
        }
    }

    fmt.Println(schedule)
    return uid,start,stop
}

func main() {
    fmt.Println("CA")

    var api CameraApi
    api.Create( "test", "hisilicon123","http://213.141.129.12:8080/cam1/")

    var c opencastClient
    //c.create("84.201.135.192", "8080", "admin", "opencast123", "MY-TEST-CA")
    c.create("130.193.39.114", "8080", "opencast_system_account", "CHANGE_ME123", "MY-TEST-CA")

    c.captureAdminHost, c.captureAdminPath, _ = c.getService("org.opencastproject.capture.admin")
    c.schedulerHost,    c.schedulerPath,    _ = c.getService("org.opencastproject.scheduler")
    c.ingestHost,       c.ingestPath,       _ = c.getService("org.opencastproject.ingest")

    c.register("idle")

    for {
        time.Sleep(10 * time.Second)

        err, schedule := c.getSchedule()
        if err != nil{
            log.Println("schedule load failed", err)
            continue
        }

        uid,startTimestamp,endTimestamp := parseSchedule(schedule, c.activeRecords)
        if (uid == "" || startTimestamp == 0 || endTimestamp == 0){
            continue
        }

        _, exists := c.activeRecords[uid]
        if (exists) {
            fmt.Println("Already processed ", uid)
            continue
        }
        c.activeRecords[uid] = true

        recordId,channelId,processingId,encoderId := startRecord(&api, startTimestamp, endTimestamp)
        if (recordId == ""){
         fmt.Println("Record was not started")
         continue
        }
        //move logic to camera
        stopRecord(&api, recordId)
        waitForFinish(&api, recordId)
        resetCam(&api, channelId, processingId, encoderId)
        packVideo(recordId)

        xml, _ := c.createMediaPackage()
        //fmt.Println(xml)

        xml2, _ := c.addDCCatalog(xml,startTimestamp,endTimestamp)
        //fmt.Println(xml2)

        xml3, err := c.addTrack(xml2, "presenter/source", "out.mp4")
        //xml3, err := c.addTrack(xml2, "presenter/source", "./video.webm")
        if err != nil {
            fmt.Println("err:", err)
        }
        //} else {
        //    fmt.Println(xml3)
        //}

        c.ingest(xml3)
        //xml4, _ := c.ingest(xml3)
        //fmt.Println(xml4)
    }
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
