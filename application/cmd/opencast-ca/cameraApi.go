package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
)

type CommonResponse struct {
    InternalError int
    Message string
    Details DatailsJson
}

type DatailsJson struct {
    Message string
}

type PipelineResponse struct {
    InternalError int
    Message string
    Details PipelineDatailsJson
}

type PipelineDatailsJson struct {
    EncoderId int
    Message string
}

type IdResponse struct {
    InternalError int
    Message string
    Details IdDatailsJson
}

type IdDatailsJson struct {
    Id int
    Message string
}

type RecordResponse struct {
    InternalError int
    Message string
    Details RecordDetailsJson
}

type RecordDetailsJson struct {
    RecordId string
    Message string
}

type AllRecordsResponse struct {
    InternalError int
    Message string
    Details []RecordsDetailsJson
}

type RecordsDetailsJson struct {
    RecordId string
    Status string
    EncoderId uint
    Size uint
    Duration string
    StartTime string
    EndTime string
    StartTimestamp uint64
    EndTimestamp uint64
    FrameCount uint
    Chunks []Chunk
}

type Chunk struct {
    DownloadLink string
    Size uint
    StartTime string
    EndTime string
    StartTimestamp uint64
    EndTimestamp uint64
    FrameCount uint
    Duration string
    Md5 string
}

type CameraApi struct {
    user string
    pass string
    camUrl string
    client *http.Client
}

func (c *CameraApi) Create(user string, pass string, camUrl string) {
    c.user = user
    c.pass = pass
    c.camUrl = camUrl
    c.client = &http.Client{}
}

func (c *CameraApi) Request(url string, response interface{}) bool {
    request, _ := http.NewRequest("GET", url, nil)
    request.SetBasicAuth(c.user, c.pass)
    resp, err := c.client.Do(request)
    if err != nil {
        log.Println("Request error ", url, err)
        return false
    }

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, response)
    if(err != nil){
        log.Println("Unmarshal response error ", url, err)
        return false
    }

    return true
}

func (c *CameraApi) CreatePipeline(encoder string) (bool, int) {
    url := c.camUrl + "api/pipeline/create?encoderName=" + encoder

    var pipelineResponse PipelineResponse
    if(!c.Request(url, &pipelineResponse)){
        return false, 0
    }

    if(pipelineResponse.Message != "Success"){
        log.Println("Request failed ", url, pipelineResponse.Message, pipelineResponse.Details.Message)
        return false, 0
    }

    return true, pipelineResponse.Details.EncoderId
}

func (c *CameraApi) CreateChannel(channelId int, width int, height int, fps int) bool {
    url := c.camUrl + "api/mpp/channel/start?channelId=" + strconv.Itoa(channelId)
    url += "&channelId=" + strconv.Itoa(channelId)
    url += "&width=" + strconv.Itoa(width)
    url += "&height=" + strconv.Itoa(height)
    url += "&fps=" + strconv.Itoa(fps)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) CreateProcessing(processingName string, params map[string]string) (bool, int) {
    url := c.camUrl + "api/processing/create?processingName=" + processingName
    for name, value := range params {
        url += "&" + name + "=" + value
    }

    var response IdResponse
    if(!c.Request(url, &response)){
        return false, 0
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false, 0
    }

    return true, response.Details.Id
}

func (c *CameraApi) CreateEncoder(encoder string) (bool, int) {
    url := c.camUrl + "api/encoder/create?encoderName=" + encoder

    var response IdResponse
    if(!c.Request(url, &response)){
        return false, 0
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false, 0
    }

    return true, response.Details.Id
}

func (c *CameraApi) SubscribeChannel(processingId int, channelId int) bool {
    url := c.camUrl + "api/processing/subscribeChannel?processingId=" + strconv.Itoa(processingId)
    url += "&channelId=" + strconv.Itoa(channelId)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) SubscribeProcessing(processingId int, encoderId int) bool {
    url := c.camUrl + "api/encoder/subscribeProcessing?processingId=" + strconv.Itoa(processingId)
    url += "&encoderId=" + strconv.Itoa(encoderId)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) StartRecording(encoderId int) (bool, string) {
    url := c.camUrl + "api/files/record/start?encoderId=" + strconv.Itoa(encoderId)

    var response RecordResponse
    if(!c.Request(url, &response)){
        return false, ""
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false, ""
    }

    return true, response.Details.RecordId
}

func (c *CameraApi) StopRecording(recordId string) bool {
    url := c.camUrl + "api/files/record/stop?recordId=" + recordId

    var response RecordResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) GetAllRecords() (bool, AllRecordsResponse) {
    url := c.camUrl + "api/files/record/listall"

    var response AllRecordsResponse
    if(!c.Request(url, &response)){
        return false, response
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message)
        return false, response
    }

    return true, response
}

func (c *CameraApi) StopChannel(channelId int) bool {
    url := c.camUrl + "api/mpp/channel/stop?channelId=" + strconv.Itoa(channelId)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) StopProcessing(processingId int) bool {
    url := c.camUrl + "api/processing/delete?processingId=" + strconv.Itoa(processingId)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}

func (c *CameraApi) StopEncoder(encoderId int) bool {
    url := c.camUrl + "api/encoder/delete?encoderId=" + strconv.Itoa(encoderId)

    var response CommonResponse
    if(!c.Request(url, &response)){
        return false
    }

    if(response.Message != "Success"){
        log.Println("Request failed ", url, response.Message, response.Details.Message)
        return false
    }

    return true
}
