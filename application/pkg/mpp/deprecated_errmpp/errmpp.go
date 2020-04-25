package errmpp

import (
    "application/pkg/logger"
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
    return resolveFunc(e.f) + " " + name + " (" + desc + ")"
}

type codeInfo struct {
    name string
    desc string
}

func resolveCode(code uint) (string, string) {

    if val, ok := codes[code]; ok {
        return val.name, val.desc
    }

    logger.Log.Warn().
        Uint("code", code).
        Msg("ERRMPP missed desc")

    return "unknown", "unknown"
}

