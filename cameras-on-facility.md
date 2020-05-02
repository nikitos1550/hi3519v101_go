
# Facility cameras

## 12.04.2020

|FSlot|Vendor|Model|SoC|CMOS|Comment|
|-----|------|-----|---|----|-------|
|1|JVT|S274H19V-L29|hi3519v101|imx274|-|
|2|XM|IVG-85HF20PYA-S|hi3516ev200|imx307|-|
|3|XM|53H20-S|hi3516cv100|imx122|-|
|4|XM|IVG-HP203Y-SE|hi3516cv300|imx291|-|
|5|XM|IVG-HP201Y-SE|hi3516cv300|imx323|-|
|6|JVT|S323H16XF|hi3516cv300|imx323|-|
|7|RUISION|RS-H622QM-B0|hi3516cv300|imx323|
|8|XM|IVG-85HF30PS-S|hi3516ev200|sc4239|check cmos|             
|9|XM|IPG-83H50P-B|hi3516av100|imx178|-|
|10|XM|IPG-83HE20PY-S|hi3516ev100|imx323|-|
|11|XM|IVG-83H80NV-BE|hi3516av200|OS08A10|-|
|12|TOPSEE|TH38J29|hi3516ev200|imx307|cmos plastic case|
|13|SSQVISION|ON335H16D|hi3516dv300|imx335|-|
|14|JVT|S226H19V-L29|hi3519v101|imx226|-|
|15|XM|53H20-AE|hi3516cv100|imx222|-|
|16|XM|83H40PL-B|hi3516av100|OV4689|-|
|17|SSQVISION|unknown|hi3519v101|imx326|-|
|18|SSQVISION|ON290H16D|hi3516dv100|imx290|rev.2|
|..|...      |...      |...        |...   |19 usb uart adapter broken!|
|20|unknown|unknown|hi3518ev201|unknown|soih65|
|21|TOPSEE|TH38D16|hi3516ev200|sc3235|check cmos!!!|
|22|YI|YHS-113-IR|hi3518ev201|unknown|-|
|23|XM|IPG-53H13PL-S|hi3518cv100|AR0130|k4b1g164...|
|24|TOPSEE|TH38C21|hi3518ev200|ar0130|-| //uboot can`t be stoped
|25|unknown|unknown|hi3518ev200|unknown|-|
|..|...|...|...|...|...|
|27|unknown|unknown|hi3516dv100|it6801fn|hdmi encoder|
|..|...|...|...|...|...|
|30|unknown|unknown|xm530|unknown|-|
|31|HISILICON|DEMB|hi3516dv300|imx327|official dev board|
|32|HISILICON|DEMBVERC|hi3559av100|imx334|official dev board|

## Checklist

- hi3516av100 family
	- [X] hi3516av100
	- [X] hi3516dv100                              
- hi3516av200 family
	- [X] hi3519v101
	- [X] hi3516av200
- hi3516cv100 family
	- [X] hi3516cv100
	- [X] hi3518cv100
	- [ ] hi3518ev100
- hi3516cv200 family
	- [ ] hi3516cv200 **N/A**
	- [X] hi3518ev200 
	- [X] hi3518ev201         
- hi3516cv300 family
	- [X] hi3516cv300
	- [X] hi3516ev100
- hi3516cv500 family
	- [ ] hi3516cv500 **N/A**
	- [X] hi3516dv300
	- [ ] hi3516av300 **Not working UART/uboot** SSQVISION_unknown_hi3516av300_imx334
- hi3516ev200 family
	- [X] hi3516ev300
	- [X] hi3516ev200
	- [ ] hi3516dv200 **N/A**
	- [ ] hi3518ev300 **N/A**
- hi3519av100 family
	- [ ] hi3519av100 **Pcb seems to be broken**
- hi3559av100 family
	- [X] hi3559av100
