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

    type SubRegister struct {
        startBit    uint8
        endBit      uint8
        size        uint8
        valids      []uint32 //??? valid values array
        presets     map[string]uint32
    }
    func (r * SubRegister) setValue(v uint32) {}
    func (r * SubRegister) setPreset(p string) {}
    func (r * SubRegister) getValue() uint32 {}

    type Register32 struct {
        base            uint64      //base address
        offset          uint64      //offset address
        reset           uint32      //reset value
        apll_postdiv2   SubRegsiter //
        apll_postdiv1   SubRegister //
        apll_frac       SubRegister //
    }


*/
