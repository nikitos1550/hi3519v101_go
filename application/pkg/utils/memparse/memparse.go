package memparse

import (
    "strconv"
)

func Str(value string) uint64 {
    var result uint64

    var multiple uint64

    suffix := value[len(value)-1:]
    switch suffix {
        case "K":
            multiple = 1024
        case "M":
            multiple = 1024*1024

        default:
            multiple = 1
    }

    var lastChar int
    if (multiple != 1) {  
        lastChar = 1
    }

    valueUint64, _ := strconv.ParseUint(value[0:len(value)-lastChar], 10, 64) 
    result = valueUint64*multiple

    return result
}
