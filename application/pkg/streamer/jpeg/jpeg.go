package jpeg

import (
    "errors"
    "sync"

    "application/pkg/mpp/connection"
    "application/pkg/logger"
)

type jpeg struct {
    sync.RWMutex
    deleted bool

    name    string

    source  connection.SourceEncodedData
}

const (
    maxId   = 1024
)

var (
    jpegs       map[int] *jpeg
    jpegsMutex  sync.RWMutex
    LastId      int
)

func init() {
    jpegs = make(map[int] *jpeg)
}

func Init() {}

func GetById(id int) (*jpeg, error) {
    jpegsMutex.RLock()
    defer jpegsMutex.RUnlock()

    item, exist := jpegs[id]
    if !exist {
        return nil, errors.New("No such instance")
    }

    return item, nil
}

func GetByName(name string) (*jpeg, error) {
    jpegsMutex.RLock()
    defer jpegsMutex.RUnlock()

    for _, item := range(jpegs) {
        if item.Name() == name {
            return item, nil
        }
    }

    return nil, errors.New("No such instance")
}

func Create(name string) (*jpeg, error) {
    jpegsMutex.Lock()
    defer jpegsMutex.Unlock()

    var item jpeg

    id := -1
    for i:=0; i < maxId; i++ {
        _, exist := jpegs[i]
        if !exist {
            id = i
            break
        }
    }

    if id == -1 {
        return nil, errors.New("Max amount reached")
    }

    for _, item := range(jpegs) {
        if item.Name() == name {
            return nil, errors.New("Duplicate name")
        }
    }

    item.name = name

    jpegs[id] = &item

    if id > LastId {
        LastId = id
    }

    logger.Log.Debug().
        Int("id", id).
        Str("name", name).
        Msg("Jpeg created")

    return &item, nil
}

//func DeleteById(id int) error {
//    jpegsMutex.Lock()
//    defer jpegsMutex.Unlock()
//
//    item, exist := jpegs[id]
//    if !exist {
//        return errors.New("No such instance")
//    }
//
//    if item.source != nil {
//        return errors.New("Can`t delete, because sourced")
//    }
//
//    delete(jpegs, id)
//
//    return nil
//}

func Delete(j *jpeg) error {
    jpegsMutex.Lock()
    defer jpegsMutex.Unlock()

    for i:=0; i < maxId; i++ {
        item := jpegs[i]
        if j == item {
            if item.destroy() != nil {
                return errors.New("Can`t delete, because sourced")
            }

            delete(jpegs, i)

            return nil
        }
    }

    return errors.New("No such instance")
}

////////////////////////////////////////////////////////////////////////////////

func (j *jpeg) destroy() error {
    if j == nil {
        return errors.New("Null pointer")
    }

    j.Lock()
    defer j.Unlock()

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return nil
    }

    if j.source != nil {
        return errors.New("Can`t destroy, because sourced")
    }

    j.deleted = true

    return nil
}

func (j *jpeg) Name() string { //(string, error) {
    if j == nil {
        return "" //return nil, errors.New("Null pointer")
    }

    j.RLock()
    defer j.RUnlock()

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return ""
    }

	return j.name //return j.name, nil
}

func (j *jpeg) getSource() (connection.SourceEncodedData, error) {
    if j == nil {
        return nil, errors.New("Null pointer")
    }

    j.RLock()
    defer j.RUnlock()

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return nil, errors.New("Invoked deleted instance")
    }

    if j.source == nil {
        return nil, errors.New("Instance not sourced")
    }

    return j.source, nil
}
