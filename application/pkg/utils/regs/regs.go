/////go:generate goyacc -l -o parser.go parser.y

package regs

/*
    Thinking about usage interface for regs util

    Sample reg:
    reg {
        desc = "APLL configuration register 0";

        field {
            desc = "Reserved";
        } reserved2[31:31] = 0;
        field {
            desc = "Level-2 output frequency divider of the APLL";
        } apll_postdiv2[30:28];
        field {
            desc = "Reserved";
        } reserved1[27:27];
        field {
            desc = "Level-1 output frequency divider of the APLL";
        } apll_postdiv1[26:24];
        field {
            desc = "Decimal part of the APLL frequency multiplication coefficient";
        } apll_frac[23:0];

    } PERI_CRG_PLL0 @ 0x0000;


    [regs.]REG_NAME.reset()
    [regs.]REG_NAME.set()
    [regs.]REG_NAME.get()
    //[regs.]REG_NAME.apll_frac.reset()
    [regs.]REG_NAME.apll_frac.set()
    [regs.]REG_NAME.apll_frac.get()
    [regs.]REG_NAME.apll_frac.VALUE_PRESET

    type Field struct {
        startBit    uint8
        endBit      uint8
        size        uint8
        valids      []uint32 //??? valid values array
        presets     map[string]uint32
    }
    func (f * Field) setValue(v uint32) {}
    func (f * Field) setPreset(p string) {}
    func (f * Field) getValue() uint32 {}

    type Register32 struct {
        base            uint64      //base address
        offset          uint64      //offset address
        reset           uint32      //reset value
        apll_postdiv2   Field //
        apll_postdiv1   Field //
        apll_frac       Field //
    }

    //////

    type Register32 struct {
        base            uint64                  //base address
        offset          uint64                  //offset address
        reset           uint32                  //reset value
        fields          map[string]Field        //
    }

    var hi3516av200 = map[string]Register32 {
        "testreg" : Register32 {
            base: 0x0,
            offset: 0x0,
            reset: 0x0,
            fields: map[string]Field {
                "apll_postdiv2" : Field {
                    startBit: 0,
                    endbit: 1,
                },
                "apll_postdiv1" : Field {
                    startBit: 2,
                    endbit: 4,
                },
                "apll_frac" : Field {
                    startBit: 2,
                    endbit: 4,
                },
            },
        },
    }


*/


/*
playground tested

package main

import (
    "fmt"
)

type Field struct {
    startBit uint8
    endBit   uint8
    //size        uint8
    //valids      []uint32 //??? valid values array
    //presets     map[string]uint32
}

type Register32 struct {
    base   uint64           //base address
    offset uint64           //offset address
    addr   uint64           //full address
    reset  uint32           //reset value
    fields map[string]Field //
}

var hi3516av200 = map[string]Register32{
    "testreg": Register32{
        base:   0x0,
        offset: 0x0,
        reset:  0x0,
        fields: map[string]Field{
            "apll_postdiv2": Field{
                startBit: 0,
                endBit:   1,
            },
            "apll_postdiv1": Field{
                startBit: 2,
                endBit:   4,
            },
            "apll_frac": Field{
                startBit: 2,
                endBit:   4,
            },
        },
    },
}

//func GetRegValue(name string) uint32 {}
//func SetRegValue(name string, value uint32) {}
//func SetRegValueEx(name string, values []FieldValue) {}

func main() {
    fmt.Println(hi3516av200["testreg"])
    //fmt.Println(romanNumeralDict[900])
}




*/
