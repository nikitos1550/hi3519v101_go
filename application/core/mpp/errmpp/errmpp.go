package errmpp

import (
    "fmt"
    "application/core/logger"
)

type errorMpp struct {
    name string
    code uint
}

func New(name string, code uint) errorMpp {
    var e errorMpp

    e.name = name
    e.code = code

    return e
}

func (e errorMpp) Error() string {
    name, desc := resolveCode(e.code)
    //return resolveFunc(e.f) + " " + name + " (" + desc + ")"
    return e.name + " " + name + " (" + desc + ")"
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
        Str("code", fmt.Sprintf("0x%08X", code)).
        Msg("ERRMPP missed info")

    return "unknown", "unknown"
}

//func resolveFunc(f uint) string {
//    if f == 0 || f >= uint(len(functions)) {
//        logger.Log.Warn().
//            Uint("func", f).
//            Msg("ERRMPP missed info")
//        return "unknown"
//    }
//
//    return functions[f]
//}
