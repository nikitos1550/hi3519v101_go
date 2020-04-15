package ko

import (
    "strconv"
)

type Parameter string
type Parameters map[string]*Parameter

func (p Parameters) Add(name string) *Parameter {
    if p[name] == nil {
        var tmp Parameter
        p[name] = &tmp
    }
    return p[name]
}

//func (p *Parameter) Prefix(prefix string) *Parameter {
//    *p = prefix + *p
//    return p  
//}

func (p *Parameter) Str(value string) *Parameter {
    *p = *p + Parameter(value)
    return p
}

func (p *Parameter) Uint64(value uint64) *Parameter {
    *p = *p + Parameter(strconv.FormatUint(value, 10))
    return p
}

func (p *Parameter) Uint64Hex(value uint64) *Parameter {
    *p = *p + Parameter(strconv.FormatUint(value, 16))
    return p
}

func (p *Parameter) Bool(value bool) *Parameter {
    if value == true {
        *p = *p + "1"
    } else {
        *p = *p + "0"
    }
    return p
}


//func (p *Parameter) Suffix(suffix string) *Parameter {
//    *p = *p + suffix
//    return p
//}


