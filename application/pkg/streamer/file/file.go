//+build streamerFile

package file

import (
    "bytes"
	"encoding/json"
    "flag"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"path"
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
	F *os.File
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

    openapi.AddApiRoute("startNewRecord", "/files/record/start", "GET", startNewRecord)
    openapi.AddApiRoute("stopRecord", "/files/record/stop", "GET", stopNewRecord)

	readStoredRecords()
}

func Init() {
}

func getStoredPath() string {
	return path.Join(*flagStoragePath, "records.json")
}

func writeStoredRecords() {
	json, _ := json.Marshal(StoredRecords)	
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
	record.F.Close()

	StoredRecords.Records[uuid] = storedRecord{
		EncoderId: record.EncoderId,
		RecordId: uuid,
		Size: record.Size,
		StartTime: record.StartTime,
		EndTime: time.Now(),
		Files: []string{},
	}

	delete(ActiveRecords, uuid)

	writeStoredRecords()
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
		F: f,
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

			ActiveRecords[uuid].F.Write(data)

			record := ActiveRecords[uuid]
			record.Size += len(data)
			ActiveRecords[uuid] = record
		}
    }()

	openapi.ResponseSuccessWithDetails(w, responseRecord{RecordId: uuid, Message: "Record was started"})
}

func stopNewRecord(w http.ResponseWriter, r *http.Request)  {
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