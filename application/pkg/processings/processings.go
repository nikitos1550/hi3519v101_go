//+build processing

package processings

import (
    "errors"
    "sync"

    "application/pkg/processings/processing"

    //"application/pkg/processing/bind"
    //"application/pkg/processing/proxy"    //TODO rename to forward
    //"application/pkg/processing/qr"
    //"application/pkg/processing/schedule"
    //"application/pkg/processing/yuv"
    "application/pkg/processings/sample"

    "application/pkg/mpp/connection"
    //"application/pkg/mpp/vpss"

    "application/pkg/logger"
    //"application/pkg/common"    //TODO remove
)

//type ActiveProcessing struct {
//    Proc common.Processing
//    InputChannel int
//    InputProcessing int
//    Encoders map[int] common.Encoder
//    Processings map[int] bool
//}

type Source struct {

}

type ProcessingInstance struct {
    Id          int
    Processing  processing.Processing

    Created     bool

    mutex       sync.RWMutex

    source      Source
    clients     map[connection.Client] bool//Connection
}

const (
	instancesAmount = 1024
)

var (
    instances   map[int]*ProcessingInstance
    mutex       sync.RWMutex
)

func init() {
    instances = make(map[int]*ProcessingInstance)
}

func Init() {
    sample.Init()

    //bind.Init()
    //proxy.Init()
    //qr.Init()
    //schedule.Init()
    //yuv.Init()

    logger.Log.Debug().
        Msg("Processing inited")
}

func GetInstance(id int) (*ProcessingInstance, error) {
    mutex.Lock()
    defer mutex.Unlock()

    _, exist := instances[id]
    if !exist {
        return nil, errors.New("No such instance")
    }

    return instances[id], nil
}

func CreateInstance(name string) (int, error) {
    var err error

    maker, err := processing.GetMaker(name)
    if err != nil {
        return 0, errors.New("No such processing type")
    }

	mutex.Lock()
    defer mutex.Unlock()

	var id int = -1
	for i:=0; i < instancesAmount; i++ {
		_, exist := instances[i]
		if !exist {
			id = i
			break
		}
	}

	if id == -1 {
		return 0, errors.New("No more processings can be created")
	}

    p := maker.Create()

    err = p.Init()
    if err != nil {
        return 0, err
    }

    var instance = ProcessingInstance{}
    instance.Id = id
    instance.Processing = p

    instances[id] = &instance

	return id, nil
}

func DestroyInstance(id int) error {
    instance, exist := instances[id]
    if !exist {
        return errors.New("No such processing instance")
    }

	mutex.Lock()
	defer mutex.Unlock()

    err := instance.Processing.DeInit()
    if err != nil {
        return err
    }

    delete(instances, id)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

//func (p *ProcessingInstance) Init() error {
//
//    return nil
//}
//
//func (p *ProcessingInstance) DeInit() error {
//    return nil
//}

func (i *ProcessingInstance) AddClient() error {
    i.mutex.Lock()
    defer i.mutex.Unlock()

    return nil
}

func (i *ProcessingInstance) RemoveClient() error {
    i.mutex.Lock()
    defer i.mutex.Unlock()

    return nil
}

func (i *ProcessingInstance) BindChannel() error {
    i.mutex.Lock()
    defer i.mutex.Unlock()

    var err error

    //c, err := vpss.GetChannel(0)
    if err != nil {
        return err
    }

    //err = c.AddClient(connection.Client(i.Processing))
    //err = c.AddClient(i.Processing.(connection.Client))
    if err != nil {
        return err
    }


    return nil
}

func (i *ProcessingInstance) UnbindChannel() error {
    i.mutex.Lock()
    defer i.mutex.Unlock()

    return nil
}

////////////////////////////////////////////////////////////////////////////////

//func SubscribeEncoderToProcessing(processingId int, encoder common.Encoder)  (int, string)  {
    /*
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoder.GetId()]
	if (exists) {
		return -1, "Already subscribed"
	}
	
	processing.Encoders[encoder.GetId()] = encoder
	ActiveProcessings[processingId] = processing
    */
//	return 0, ""
//}

//func UnsubscribeEncoderToProcessing(processingId int, encoder common.Encoder)  (int, string)  {
    /*
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoder.GetId()]
	if (!exists) {
		return -1, "Encoder not subscribed"
	}

	delete(processing.Encoders, encoder.GetId())	
    */
//	return 0, ""
//}

/*
func SubscribeProcessingToProcessing(processingId int, subscribeProcessingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Main processing not created"
	}

	subscribeProcessing, subscribeProcessingExists := ActiveProcessings[subscribeProcessingId]
	if (!subscribeProcessingExists) {
		return -1, "Subscribe processing not created"
	}

	_, exists := activeProcessing.Processings[subscribeProcessingId]
	if (exists) {
		return -1, "Already subscribed"
	}

	activeProcessing.Processings[subscribeProcessingId] = true
	ActiveProcessings[processingId] = activeProcessing

	subscribeProcessing.InputProcessing = processingId
	ActiveProcessings[subscribeProcessingId] = subscribeProcessing

	return 0, ""
}
*/
/*
func UnsubscribeProcessingToProcessing(processingId int, subscribeProcessingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Main processing not created"
	}

	_, subscribeProcessingExists := ActiveProcessings[subscribeProcessingId]
	if (!subscribeProcessingExists) {
		return -1, "Subscribe processing not created"
	}

	_, exists := activeProcessing.Processings[subscribeProcessingId]
	if (!exists) {
		return -1, "Processing not subscribed"
	}

	delete(activeProcessing.Processings, subscribeProcessingId)	

	return 0, ""
}
*/
