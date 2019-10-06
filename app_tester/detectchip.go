package main

func detectChip(reg uint32) string {
    switch (reg) {
        case 890765568: //0x35180100
            return "hi3518?v100"
        case 890822912: //0x3518E100
            return "hi3518ev100"
        case 890683648: //0x3516C100
            return "hi3516cv100"
        case 890675456: //0x3516A100
            return "hi3516av100"
        case 890687744: //0x3516D100
            return "hi3516dv100"
        case 890831105: //0x35190101
            return "hi3519v101"
        case 890684160: //0x3516C300
            return"hi3516cv300"
        case 890823168: //0x3518E200
            return "hi3518ev200"
        default:
            return "unknown"
    }

}
