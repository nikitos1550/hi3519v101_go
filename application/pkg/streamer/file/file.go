//+build streamerFile

package file

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"application/pkg/openapi"
	"application/pkg/mpp/venc"
	 
	"github.com/google/uuid"
)

var flagStoragePath     *string

type activeRecord struct {
    Payload    chan []byte
	Started    bool
	EncoderId  int
	CurrentFile *os.File
	CurrentFilePath string
	Size int
	StartTime time.Time 
}

type storedRecord struct {
	EncoderId  int
	RecordId  string
	Size int
	StartTime time.Time 
	EndTime time.Time 
	Files []string
}

type recordInfo struct {
	RecordId  string
	Status    string
	EncoderId  int
	Size int
	Duration string
	StartTime time.Time 
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

    openapi.AddApiRoute("apiDescription", "/files/record", "GET", apiDescription)

    openapi.AddApiRoute("startNewRecord", "/files/record/start", "GET", startNewRecord)
    openapi.AddApiRoute("stopRecord", "/files/record/stop", "GET", stopRecord)

	openapi.AddApiRoute("downloadRecord", "/files/record/download", "GET", downloadRecord)

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

func keyFrame(data []byte) bool {
	keyData := []byte{0x00, 0x00, 0x00, 0x01, 0x67}
	return bytes.HasPrefix(data, keyData)
}

func saveRecord(uuid string, record activeRecord) {
	venc.RemoveSubscription(record.EncoderId)
	record.CurrentFile.Close()

	finalFilePath := record.CurrentFilePath
	if (strings.HasSuffix(finalFilePath, ".tmp")){
		finalFilePath = finalFilePath[0:len(finalFilePath)-len(".tmp")]
		err := os.Rename(record.CurrentFilePath, finalFilePath)
		if err != nil {
			log.Println("Failed to move file " + record.CurrentFilePath + " to " + finalFilePath)
		}
	}

	StoredRecords.Records[uuid] = storedRecord{
		EncoderId: record.EncoderId,
		RecordId: uuid,
		Size: record.Size,
		StartTime: record.StartTime,
		EndTime: time.Now(),
		Files: []string{finalFilePath},
	}

	delete(ActiveRecords, uuid)

	writeStoredRecords()
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Records api:\n\n", "/files/record")
}

func startNewRecord(w http.ResponseWriter, r *http.Request)  {
	uuid := uuid.New().String()
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	recordFolder := path.Join(*flagStoragePath, uuid)
	folderErr := os.MkdirAll(recordFolder, os.ModePerm)
	if folderErr != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: "", Message: "Failed to create folder " + recordFolder})
		return
    }
	
	file := path.Join(recordFolder, "out.h264.tmp")
	f,err := os.Create(file)
	if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: "", Message: "Failed to create file " + file})
		return
    }
	
	ActiveRecords[uuid] = activeRecord{
		Payload: make(chan []byte, 100),
		Started: true,
		EncoderId: encoderId,
		CurrentFile: f,
		CurrentFilePath: file,
		Size: 0,
		StartTime: time.Now(),
	}

	venc.SubsribeEncoder(encoderId, ActiveRecords[uuid].Payload)
	
    go func() {
		for {
			if (!ActiveRecords[uuid].Started){
				saveRecord(uuid, ActiveRecords[uuid])
				break
			}

			data := <- ActiveRecords[uuid].Payload
			if (ActiveRecords[uuid].Size == 0){
				if (!keyFrame(data)){
					continue
				}

				record := ActiveRecords[uuid]
				record.StartTime = time.Now()
				ActiveRecords[uuid] = record
			}

			ActiveRecords[uuid].CurrentFile.Write(data)

			record := ActiveRecords[uuid]
			record.Size += len(data)
			ActiveRecords[uuid] = record
		}
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

	record, exists := StoredRecords.Records[recordId]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found"})
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=" + recordId + ".h264")
	http.ServeFile(w, r, record.Files[0])
}

func addActiveRecord(records *[]recordInfo, record activeRecord, uuid string)  {
	info := recordInfo{
		RecordId: uuid,
		Status: "recording",
		EncoderId: record.EncoderId,
		Size: record.Size,
		Duration: fmt.Sprintf("%v", time.Now().Sub(record.StartTime)),
		StartTime: record.StartTime,
	}
	
	*records = append(*records, info)
}

func addActiveRecords(records *[]recordInfo)  {
	for uuid, record := range ActiveRecords {
		addActiveRecord(records, record, uuid)
	}
}

func addFinishedRecord(records *[]recordInfo, record storedRecord)  {
	info := recordInfo{
		RecordId: record.RecordId,
		Status: "finished",
		EncoderId: record.EncoderId,
		Size: record.Size,
		Duration: fmt.Sprintf("%v", record.EndTime.Sub(record.StartTime)),
		StartTime: record.StartTime,
	}
	
	*records = append(*records, info)
}

func addFinishedRecords(records *[]recordInfo)  {
	for _, record := range StoredRecords.Records {
		addFinishedRecord(records, record)
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
		addActiveRecord(&records, active, recordId)
		openapi.ResponseSuccessWithDetails(w, records)
		return
	}

	stored, storedExists := StoredRecords.Records[recordId]
	if (storedExists) {
		addFinishedRecord(&records, stored)
		openapi.ResponseSuccessWithDetails(w, records)
		return
	}

	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{RecordId: recordId, Message: "Record not found"})
}

func listAllRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addActiveRecords(&records)
	addFinishedRecords(&records)
	openapi.ResponseSuccessWithDetails(w, records)
}

func listActiveRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addActiveRecords(&records)
	openapi.ResponseSuccessWithDetails(w, records)
}

func listFinishedRecords(w http.ResponseWriter, r *http.Request)  {
	var records []recordInfo
	addFinishedRecords(&records)
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
	for _, file := range record.Files {
		err = os.RemoveAll(file)
        if err != nil {
            log.Println("Failed to remove file " + file)
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
