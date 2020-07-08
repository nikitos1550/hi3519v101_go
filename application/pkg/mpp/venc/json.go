package venc

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "reflect"
    "strings"
)

func (this BitrateControlParameters) MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer

    buffer.WriteString("{")

    v := reflect.ValueOf(this)

    fields, err := intsToJson(v)
    if err != nil {
        return nil, errors.New("intsToJson error TODO")
    }

    buffer.Write(fields)

    buffer.WriteString("}")

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func (this GopParameters) MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer

    buffer.WriteString("{")

    v := reflect.ValueOf(this)

    fields, err := intsToJson(v)
    if err != nil {
        return nil, errors.New("intsToJson error TODO")
    }

    buffer.Write(fields)

    buffer.WriteString("}")

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func intsToJson(v reflect.Value) ([]byte, error) {
    var buffer bytes.Buffer

    typeOfS := v.Type()

    var first bool = true
    for i := 0; i < v.NumField(); i++ {
        val := v.Field(i).Interface().(int)

        var name string
        var omitempty bool = false

        tag, ok := typeOfS.Field(i).Tag.Lookup("json")
        if ok {
            items := strings.Split(tag, ",")
            name = items[0]
            if len(items) == 2 && items[1] == "omitempty" {
                omitempty = true
            }
        } else {
            name = typeOfS.Field(i).Name
        }

        if !first && (val != invalidValue || !omitempty) {
                buffer.WriteString(",")
        }

        if val != invalidValue {
            buffer.WriteString(fmt.Sprintf("\"%s\":%d", name, val))
            first = false
        } else {
            if !omitempty {
                buffer.WriteString(fmt.Sprintf("\"%s\":\"n/a\"", name))
                first = false
            }
        }
    }

    return buffer.Bytes(), nil
}

func (this Codec) MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer
    switch this {
        case MJPEG:
            buffer.WriteString("\"mjpeg\"")
        case H264:
            buffer.WriteString("\"h264\"")
        case H265:
            buffer.WriteString("\"h265\"")
        default:
            buffer.WriteString("\"n/a\"")
            //return nil, errors.New("Unknown codec")
    }

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func (this *Codec) UnmarshalJSON(b []byte) error {
    var value string
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }

    //fmt.Println("b:", value)

    switch value {
        case "mjpeg":
            *this = MJPEG
        case "h264":
            *this = H264
        case "h265":
            *this = H265
        default:
            *this = Codec(invalidValue)
    }
    return nil
}

func (this BitrateControl) MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer
    switch this {
        case Cbr:
            buffer.WriteString("\"cbr\"")
        case Vbr:
            buffer.WriteString("\"vbr\"")
        case FixQp:
            buffer.WriteString("\"fixqp\"")
        case CVbr:
            buffer.WriteString("\"cvbr\"")
        case AVbr:
            buffer.WriteString("\"avbr\"")
        case QVbr:
            buffer.WriteString("\"qvbr\"")
        default:
            buffer.WriteString("\"n/a\"")
            //return nil, errors.New("Unknown codec")
    }

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func (this *BitrateControl) UnmarshalJSON(b []byte) error {
    var value string
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    switch value {
        case "cbr":
            *this = Cbr
        case "vbr":
            *this = Vbr
        case "fixqp":
            *this = FixQp
        case "cvbr":
            *this = CVbr
        case "avbr":
            *this = AVbr
        case "qvbr":
            *this = QVbr
        default:
            *this = BitrateControl(invalidValue)
    }
    return nil
}


func (this Profile)MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer
    switch this {
        case Baseline:
            buffer.WriteString("\"baseline\"")
        case Main:
            buffer.WriteString("\"main\"")
        case High:
            buffer.WriteString("\"high\"")
        default:
            buffer.WriteString("\"n/a\"")
            //return nil, errors.New("Unknown profile")
    }

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func (this *Profile) UnmarshalJSON(b []byte) error {
    var value string
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    switch value {
        case "baseline":
            *this = Baseline
        case "main":
            *this = Main
        case "high":
            *this = High
        default:
            *this = Profile(invalidValue)
    }
    return nil
}

func (this GopStrategyType)MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer
    switch this {
        case NormalP:
            buffer.WriteString("\"normalp\"")
        case DualP:
            buffer.WriteString("\"dualp\"")
        case SmartP:
            buffer.WriteString("\"smartp\"")
        case AdvSmartP:
            buffer.WriteString("\"advsmartp\"")
        case BipredB:
            buffer.WriteString("\"bipredb\"")
        case IntraR:
            buffer.WriteString("\"intrar\"")
        default:
            buffer.WriteString("\"n/a\"")
            //return nil, errors.New("Unknown profile")
    }

    //fmt.Println(buffer.String())

    return buffer.Bytes(), nil
}

func (this *GopStrategyType) UnmarshalJSON(b []byte) error {
    var value string
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    switch value {
        case "normalp":
            *this = NormalP
        case "dualp":
            *this = DualP
        case "smartp":
            *this = SmartP
        case "advsmartp":
            *this = AdvSmartP
        case "bipredb":
            *this = BipredB
        case "intrar":
            *this = IntraR
        default:
            *this = GopStrategyType(invalidValue)
    }
    return nil
}

