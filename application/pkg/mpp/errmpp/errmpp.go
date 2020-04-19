package errmpp

import (
    "strconv"
)

type errorMpp struct {
    method string
    code uint
    //name string
    //desc string
}

func New(m string, c uint) errorMpp {
    var e errorMpp

    e.method = m
    e.code = c

    return e
}

func (e errorMpp) Error() string {
    name, desc := resolve(e.code)
    return e.method + ": 0x" + strconv.FormatInt(int64(e.code), 16) + " " + name + " (" + desc + ")"
}
