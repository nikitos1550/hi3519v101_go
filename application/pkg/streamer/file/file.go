//+build streamerFile

package file

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"    
	"strings"
	"time"

	"application/pkg/mpp/venc"
	"application/pkg/openapi"
	"application/pkg/streamer/rtsp"
	 
	"github.com/google/uuid"
)

var flagStoragePath     *string
var downloadLink string

type chunk struct {
	FilePath string
	Size int
	StartTime time.Time 
	EndTime time.Time 
	Md5 string
}

type activeRecord struct {
    //Payload    chan []byte
    Payload    chan venc.ChannelEncoder
	Started    bool
	EncoderId  int
	CurrentFile *os.File
	CurrentFilePath string
	Size int
	StartTime time.Time 
	Chunks []chunk
	Extention string
}

type storedRecord struct {
	EncoderId  int
	RecordId  string
	Size int
	StartTime time.Time 
	EndTime time.Time 
	Chunks []chunk
	Extention string
}

type chunkInfo struct {
	DownloadLink string
	Size int
	StartTime time.Time 
	EndTime time.Time 
	Duration string
	Md5 string
}

type recordInfo struct {
	RecordId  string
	Status    string
	EncoderId  int
	Size int
	Duration string
	StartTime time.Time 
	EndTime time.Time 
	Chunks []chunkInfo
}

type storedRecords struct {
	Records map[string] storedRecord
}

type responseRecord struct {
	RecordId string
	Message string
}

var (
	ActiveRecords map[string] activeRecord
    StoredRecords storedRecords
)

func init() {
    flagStoragePath = flag.String("streamer-file-storage", "/opt/nfs", "files storage path")
	ActiveRecords = make(map[string] activeRecord)
	StoredRecords.Records = make(map[string] storedRecord)
	downloadLink = "/files/record/download"

    openapi.AddApiRoute("apiDescription", "/files/record", "GET", apiDescription)

    openapi.AddApiRoute("startNewRecord", "/files/record/start", "GET", startNewRecord)
    openapi.AddApiRoute("stopRecord", "/files/record/stop", "GET", stopRecord)

	openapi.AddApiRoute("downloadRecord", downloadLink, "GET", downloadRecord)

    openapi.AddApiRoute("listRecord", "/files/record/info", "GET", listRecord)
    openapi.AddApiRoute("listAllRecords", "/files/record/listall", "GET", listAllRecords)
    openapi.AddApiRoute("listActiveRecords", "/files/record/listactive", "GET", listActiveRecords)
    openapi.AddApiRoute("listFinishedRecords", "/files/record/listfinished", "GET", listFinishedRecords)
    openapi.AddApiRoute("removeRecord", "/files/record/remove", "GET", removeRecord)

	readStoredRecords()
}

func Init() {
}

func getStoredPath() string {
	return path.Join(*flagStoragePath, "records.json")
}

func writeStoredRecords() {
	json, _ := json.MarshalIndent(StoredRecords, "", "  ")
    err := ioutil.WriteFile(getStoredPath(), json, 0644)	
    if err != nil {
		log.Println("Failed to write records to file " + getStoredPath())
	}
}

func readStoredRecords() {
    data, err := ioutil.ReadFile(getStoredPath())
    if err != nil {
		log.Println("Failed to read records from file " + getStoredPath())
		return
    }
    
	err = json.Unmarshal(data, &StoredRecords)
    if err != nil {
        log.Println("Failed to parse records from file " + getStoredPath())
    }
}

func keyFrame(encoder string, data []byte) bool {
	sps := rtsp.ExtractSps(encoder, data);
	return len(sps) > 0
}

func saveRecord(uuid string, record activeRecord, md5Hash string) {
	venc.RemoveSubscription(record.EncoderId, record.Payload)
	record.CurrentFile.Close()

	finalFilePath := record.CurrentFilePath
	if (strings.HasSuffix(finalFilePath, "." + record.Extention + ".tmp")){
		finalFilePath = finalFilePath[0:len(finalFilePath)-len("." + record.Extention + ".tmp")]
		finalFilePath += strconv.Itoa(len(record.Chunks) - 1) + "." + record.Extention
		err := os.Rename(record.CurrentFilePath, finalFilePath)
		if err != nil {
			log.Println("Failed to move file " + record.CurrentFilePath + " to " + finalFilePath)
		}
	}

	record.Chunks[len(record.Chunks) - 1].FilePath = finalFilePath
	record.Chunks[len(record.Chunks) - 1].Md5 = md5Hash

	StoredRecords.Records[uuid] = storedRecord{
		EncoderId: record.EncoderId,
		RecordId: uuid,
		Size: record.Size,
		StartTime: record.StartTime,
		EndTime: record.Chunks[len(record.Chunks) - 1].EndTime,
		Chunks: record.Chunks,
		Extention: record.Extention,
	}

	delete(ActiveRecords, uuid)

	writeStoredRecords()
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Records api:\n\n", "/files/record")
}

func writeVideoData(uuid string, chunks int, duration int){
	h := md5.New()
	for {
		if (!ActiveRecords[uuid].Started){
			md5Hash := hex.EncodeToString(h.Sum(nil))
			saveRecord(uuid, ActiveRecords[uuid], md5Hash)
			break
		}

		data := <- ActiveRecords[uuid].Payload
		if (ActiveRecords[uuid].Size == 0){
			if (!keyFrame(ActiveRecords[uuid].Extention, data.Data)){
				continue
			}

			record := ActiveRecords[uuid]
			record.StartTime = time.Now()

			c := chunk{
				FilePath: record.CurrentFilePath,
				Size: 0,
				StartTime: record.StartTime,
				EndTime: time.Now(),
				Md5: "",
			}

			record.Chunks = append(record.Chunks, c)
			ActiveRecords[uuid] = record
		}

		if (chunks != 0){
			record := ActiveRecords[uuid]

			chunkDuration := time.Now().Sub(record.Chunks[len(record.Chunks) - 1].StartTime)
			if (chunkDuration.Seconds() >= float64(duration) && keyFrame(record.Extention, data.Data)){
				record.CurrentFile.Close()
				finalFilePath := record.CurrentFilePath
				if (strings.HasSuffix(finalFilePath, ".tmp")){
					finalFilePath = finalFilePath[0:len(finalFilePath)-len("." + record.Extention + ".tmp")]
					finalFilePath += strconv.Itoa(len(record.Chunks) - 1) + "." + record.Extention
					err := os.Rename(record.CurrentFilePath, finalFilePath)
					if err != nil {
						log.Println("Failed to move file " + record.CurrentFilePath + " to " + finalFilePath)
					}
				}

				f,err := os.Create(record.CurrentFilePath)
				if err != nil {
					log.Println("Failed to create file " + record.CurrentFilePath)
				}

				record.CurrentFile = f
				record.Chunks[len(record.Chunks) - 1].FilePath = finalFilePath
				md5Hash := hex.EncodeToString(h.Sum(nil))
				h.Reset()
				record.Chunks[len(record.Chunks) - 1].Md5 = md5Hash

				c := chunk{
					FilePath: record.CurrentFilePath,
					Size: 0,
					StartTime: time.Now(),
					EndTime: time.Now(),
					Md5: "",
				}
				record.Chunks = append(record.Chunks, c)

				ActiveRecords[uuid] = record
			}
		}

		ActiveRecords[uuid].CurrentFile.Write(data.Data)
		h.Write(data.Data)

		record := ActiveRecords[uuid]
		record.Size += len(data.Data)
		record.Chunks[len(record.Chunks) - 1].Size += len(data.Data)
		record.Chunks[len(record.Chunks) - 1].EndTime = time.Now()
		ActiveRecords[uuid] = record
	}
}

func startNewRecord(w http.ResponseWriter, r *http.Request)  {
	uuid := uuid.New().String()
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}
	encoder, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: "", Message: "Failed to find encoder  " + strconv.Itoa(encoderId)})
		return
	}

	chunks := openapi.GetIntParameterOrDefault(w, r, "chunks", 0)
	chunkDuration := openapi.GetIntParameterOrDefault(w, r, "chunkDuration", 30)

	recordFolder := path.Join(*flagStoragePath, uuid)
	folderErr := os.MkdirAll(recordFolder, os.ModePerm)
	if folderErr != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: "", Message: "Failed to create folder " + recordFolder})
		return
    }
	
	file := path.Join(recordFolder, "out." + encoder.Format + ".tmp")
	f,err := os.Create(file)
	if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: "", Message: "Failed to create file " + file})
		return
    }
	
	ActiveRecords[uuid] = activeRecord{
		//Payload: make(chan []byte, 100),
        Payload: make(chan venc.ChannelEncoder, 100),
		Started: true,
		EncoderId: encoderId,
		CurrentFile: f,
		CurrentFilePath: file,
		Size: 0,
		StartTime: time.Now(),
		Chunks: []chunk{},
		Extention: encoder.Format,
	}

	venc.SubsribeEncoder(encoderId, ActiveRecords[uuid].Payload)
	
    go func() {
		writeVideoData(uuid, chunks, chunkDuration)
    }()

	openapi.ResponseSuccessWithDetails(w, responseRecord{RecordId: uuid, Message: "Record was started"})
}

func stopRecord(w http.ResponseWriter, r *http.Request)  {
	ok, recordId := openapi.GetStringParameter(w, r, "recordId")
	if !ok {
		return
	}

	record, exists := ActiveRecords[recordId]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found or not started"})
		return
	}

	record.Started = false
	ActiveRecords[recordId] = record
	openapi.ResponseSuccessWithDetails(w, responseRecord{RecordId: recordId, Message: "Record was stopped"})
}

func downloadRecord(w http.ResponseWriter, r *http.Request)  {
	ok, recordId := openapi.GetStringParameter(w, r, "recordId")
	if !ok {
		return
	}

	chunk := openapi.GetIntParameterOrDefault(w, r, "chunk", 0)

	activeRecord, activeExists := ActiveRecords[recordId]
	if (activeExists) {
		if (chunk < len(activeRecord.Chunks) - 1){
			w.Header().Set("Content-Disposition", "attachment; filename=" + recordId + "_chunk_" + strconv.Itoa(chunk) + "." + activeRecord.Extention)
			w.Header().Set("Md5", activeRecord.Chunks[chunk].Md5)
			http.ServeFile(w, r, activeRecord.Chunks[chunk].FilePath)
			return
		}
	}

	storedRecord, storedExists := StoredRecords.Records[recordId]
	if (storedExists) {
		w.Header().Set("Content-Disposition", "attachment; filename=" + recordId + "_chunk_" + strconv.Itoa(chunk) + "." + storedRecord.Extention)
		w.Header().Set("Md5", storedRecord.Chunks[chunk].Md5)
		http.ServeFile(w, r, storedRecord.Chunks[chunk].FilePath)
		return
	}

	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found"})
}

func addChunksInfo(r *http.Request, inputChunks []chunk, chunksCount int, recordId string, outputChunks* []chunkInfo){
	for i := 0; i < chunksCount; i++ {
		link := downloadLink + "?recordId=" + recordId + "&chunk=" + strconv.Itoa(i)
		info := chunkInfo{
			DownloadLink: link,
			Size: inputChunks[i].Size,
			StartTime: inputChunks[i].StartTime,
			EndTime: inputChunks[i].EndTime,
			Duration: fmt.Sprintf("%v", inputChunks[i].EndTime.Sub(inputChunks[i].StartTime)),
			Md5: inputChunks[i].Md5,
		}

		*outputChunks = append(*outputChunks, info)
	}
}

func addActiveRecord(r *http.Request, records *[]recordInfo, record activeRecord, uuid string)  {
	info := recordInfo{
		RecordId: uuid,
		Status: "recording",
		EncoderId: record.EncoderId,
		Size: record.Size,
		Duration: fmt.Sprintf("%v", time.Now().Sub(record.StartTime)),
		StartTime: record.StartTime,
		EndTime: time.Now(),
		Chunks: []chunkInfo{},
	}

	addChunksInfo(r, record.Chunks, len(record.Chunks) - 1, info.RecordId, &info.Chunks)
	
	*records = append(*records, info)
}

func addActiveRecords(r *http.Request, records *[]recordInfo)  {
	for uuid, record := range ActiveRecords {
		addActiveRecord(r, records, record, uuid)
	}
}

func addFinishedRecord(r *http.Request, records *[]recordInfo, record storedRecord)  {
	info := recordInfo{
		RecordId: record.RecordId,
		Status: "finished",
		EncoderId: record.EncoderId,
		Size: record.Size,
		Duration: fmt.Sprintf("%v", record.EndTime.Sub(record.StartTime)),
		StartTime: record.StartTime,
		EndTime: record.EndTime,
		Chunks: []chunkInfo{},
	}
	
	addChunksInfo(r, record.Chunks, len(record.Chunks), info.RecordId, &info.Chunks)

	*records = append(*records, info)
}

func addFinishedRecords(r *http.Request, records *[]recordInfo)  {
	for _, record := range StoredRecords.Records {
		addFinishedRecord(r, records, record)
	}
}

func listRecord(w http.ResponseWriter, r *http.Request)  {
	ok, recordId := openapi.GetStringParameter(w, r, "recordId")
	if !ok {
		return
	}

	var records []recordInfo
	active, activeExists := ActiveRecords[recordId]
	if (activeExists) {
		addActiveRecord(r, &records, active, recordId)
		openapi.ResponseSuccessWithDetails(w, records)
		return
	}

	stored, storedExists := StoredRecords.Records[recordId]
	if (storedExists) {
		addFinishedRecord(r, &records, stored)
		openapi.ResponseSuccessWithDetails(w, records)
		return
	}

	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found"})
}

func listAllRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addActiveRecords(r, &records)
	addFinishedRecords(r, &records)
	openapi.ResponseSuccessWithDetails(w, records)
}

func listActiveRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addActiveRecords(r, &records)
	openapi.ResponseSuccessWithDetails(w, records)
}

func listFinishedRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addFinishedRecords(r, &records)
	openapi.ResponseSuccessWithDetails(w, records)
}

func removeRecord(w http.ResponseWriter, r *http.Request)  {
	ok, recordId := openapi.GetStringParameter(w, r, "recordId")
	if !ok {
		return
	}

	record, recordExists := StoredRecords.Records[recordId]
	if (!recordExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found"})
		return
	}

	var err error
	for _, chunk := range record.Chunks {
		err = os.RemoveAll(chunk.FilePath)
        if err != nil {
            log.Println("Failed to remove file " + chunk.FilePath)
        }
	}

	recordPath := path.Join(*flagStoragePath, recordId)
	err = os.RemoveAll(recordPath)
	if err != nil {
		log.Println("Failed to remove folder " + recordPath)
	}

	delete(StoredRecords.Records, recordId)
	writeStoredRecords()

	openapi.ResponseSuccessWithDetails(w, responseRecord{RecordId: recordId, Message: "Record was removed"})
}
