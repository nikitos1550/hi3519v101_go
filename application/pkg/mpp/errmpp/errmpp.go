package errmpp

import (
    "strconv"
)

type errorMpp struct {
    f uint
    c uint
    //name string
    //desc string
}

func New(f uint, c uint) errorMpp {
    var e errorMpp

    e.f = f
    e.c = c

    return e
}

func (e errorMpp) Error() string {
    name, desc := resolveCode(e.c)
    return resolveFunc(e.f) + ": 0x" + strconv.FormatInt(int64(e.c), 16) + " " + name + " (" + desc + ")"
}
