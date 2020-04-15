package regs

import (
	"fmt"
    "application/pkg/utils"
    "application/pkg/logger"
)

//var addrMaps map[string][]register32

//func addAddrMap(name string, regs []register32) {
//	if len(addrMaps) == 0 {
//		addrMaps = make(map[string][]register32)
//	}
//	addrMaps[name] = regs
//}

type value struct {
	value uint32
	name  string
	desc  string
}

type field struct {
	bitStart uint8
	bitEnd   uint8
	name     string
	desc     string
	values   []value
}

func (f *field) getName() string {
	if f == nil {
		return ""
	}
	return f.name
}

func (f *field) getDesc() string {
	if f == nil {
		return ""
	}
	return f.desc
}

func (f *field) getValueName(value uint32) string {

	if f == nil {
		return ""
	}
	if len(f.values) == 0 {
		return ""
	}
	for i := 0; i < len(f.values); i++ {
		if f.values[i].value == value {
			return f.values[i].name
		}
	}
	return "unknown"
}

func (f *field) Set(value uint32) {
    //TODO
}

func (f *field) Get() uint32 {
    //TODO
    return 0
}

type register32 struct {
	addr   uint32 //full address
	name   string
	desc   string
	fields []field
}

func (r *register32) getName() string {
	if r == nil {
		return ""
	}
	return r.name
}

func (r *register32) getDesc() string {
	if r == nil {
		return ""
	}
	return r.desc
}

func (r *register32) Set(value uint32) {
    if r == nil {
        //TODO
        return
    }
    utils.WriteDevMem32(r.addr, value)
}

func (r *register32) Get() uint32 {
    if r == nil {
        //TODO
        return 0
    }
    return utils.ReadDevMem32(r.addr)
}

func (r *register32) Field(n string) *field {
    //TOOD
    return nil
}

func (r *register32) Fields(bitStart, bitEnd uint8) []*field {
	fields := make([]*field, 0)

	if r == nil {
		return fields
	}

	for i := 0; i < len(r.fields); i++ {
		field := &r.fields[i]
		if field.bitStart >= bitStart {
			if field.bitEnd <= bitEnd {
				fields = append(fields, field)
			}
		}
	}
	return fields
}
////
func (r *register32) Dump() {
    if r == nil {
        //TODO
        return
    }
    value   := r.Get()

    fmt.Printf("Register %s (%s)\n", r.getName(), r.getDesc() )

    fields := r.Fields(0, 32)
    if len(fields) > 0 {
        for _, field := range fields {
            fmt.Printf("Field %s[%d:%d] (%s) ", field.getName(), field.bitStart, field.bitEnd, field.getDesc())

            fieldValue := ((value << 0) << (31 - field.bitEnd) >> (31 - field.bitEnd)) >> field.bitStart
            recognizedValue := field.getValueName(fieldValue)
            if recognizedValue != "" {
                fmt.Printf("val = 0x%X (%s)\n", fieldValue, recognizedValue)
            } else {
                fmt.Printf("val = 0x%X\n", fieldValue)
            }
        }
    } else {
        fmt.Printf("Fields not found\n")
    }
}
////

func ByAddr(r uint32) *register32 {
	var found *register32

	for _, reg := range Registers {
		if reg.addr == r {
			found = &reg
			break
		}
	}
    if found == nil {
        logger.Log.Fatal().
            Uint32("addr", r).
            Msg("No such reg")
    }
	return found
}

func ByNameStr(n string) *register32 {
    var found *register32

    for _, reg := range Registers {
        if reg.name == n {
            found = &reg
            break
        }
    }
    if found == nil {
        logger.Log.Fatal().
            Str("name", n).
            Msg("No such reg")
    }
    return found
}

func ByNameConst(n uint) *register32 {
    if n > uint(len(Registers)) {
        return nil
    }
    return &Registers[n]
}

