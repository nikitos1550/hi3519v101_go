package record

import (
    "io"
)

type RecordMP4 Record

func NewMP4(r *Record) *RecordMP4 {
    return (*RecordMP4)(r)
}

func (r *RecordMP4) Read() (int, error) {

    return 0, nil
}

func (r *RecordMP4) WriteTo(w io.Writer) (int, error) {

    return 0, nil
}
