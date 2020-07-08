package connection

import (
    "bytes"
)

func (this ConnectionType) MarshalJSON() ([]byte, error) {
    var buffer bytes.Buffer
    switch this {
        case None:
            buffer.WriteString("\"none\"")
        case BindEncoder:
            buffer.WriteString("\"bindencoder\"")
        case BindDisplay:
            buffer.WriteString("\"binddisplay\"")
        case Push:
            buffer.WriteString("\"push\"")
        default:
            buffer.WriteString("\"n/a\"")
    }

    return buffer.Bytes(), nil
}

