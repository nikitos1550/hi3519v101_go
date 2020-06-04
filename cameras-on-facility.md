
# Facility cameras

## 12.04.2020

|N |Vendor      |Model              |SoC        |CMOS       |Comment                |catch-uboot|uboot-net  |linux  |jpeg|
|--|------------|-------------------|-----------|-----------|-----------------------|-----------|-----------|-------|----|
|1 |JVT         |S274H19V-L29       |hi3519v101 |imx274     |                       |+          |+          |+      |    |
|2 |XM          |IVG-85HF20PYA-S    |hi3516ev200|imx307     |                       |+          |+          |+      |    |
|3 |XM          |53H20-S            |hi3516cv100|imx122     |                       |+          |+          |+      |    |
|4 |XM          |IVG-HP203Y-SE      |hi3516cv300|imx291     |                       |+          |+          |+      |    |
|5 |XM          |IVG-HP201Y-SE      |hi3516cv300|imx323     |                       |+          |+          |+      |    |
|6 |JVT         |S323H16XF          |hi3516cv300|imx323     |                       |+          |+          |+      |    |
|7 |RUISION     |RS-H622QM-B0       |hi3516cv300|imx323     |                       |+          |+          |+      |    |
|8 |XM          |IVG-85HF30PS-S     |hi3516ev200|sc4239     |check cmos             |+          |+          |+      |    |
|9 |XM          |IPG-83H50P-B       |hi3516av100|imx178     |                       |+          |+          |+      |    |
|10|XM          |IPG-83HE20PY-S     |hi3516ev100|imx323     |                       |+          |+          |+      |    |
|11|XM          |IVG-83H80NV-BE     |hi3516av200|OS08A10    |                       |+          |+          |+      |    |
|12|TOPSEE      |TH38J29            |hi3516ev200|imx307     |cmos plastic case      |+          |+          |+      |    |
|13|SSQVISION   |ON335H16D          |hi3516dv300|imx335     |                       |+          |+          |+      |    |
|14|JVT         |S226H19V-L29       |hi3519v101 |imx226     |                       |+          |+          |+      |    |
|15|XM          |53H20-AE           |hi3516cv100|imx222     |                       |+          |+          |+      |    |
|16|XM          |83H40PL-B          |hi3516av100|OV4689     |                       |+          |+          |+      |    |
|17|SSQVISION   |unknown            |hi3519v101 |imx326     |uboot tftp issue       |+          |+          |-      |    |
|18|SSQVISION   |ON290H16D          |hi3516dv100|imx290     |rev.2                  |+          |+          |+      |    |
|..|...         |...                |...        |...        |uart adapter broken    |           |           |       |    |
|20|unknown     |dacility20         |hi3518ev200|h65        |                       |+          |+          |+      |    |
|21|TOPSEE      |TH38D16            |hi3516ev200|sc3235     |check cmos             |+          |+          |+      |    |
|22|YI          |YHS-113-IR         |hi3518ev201|ov9732     |wifi only              |+          |-          |       |    |
|23|XM          |IPG-53H13PL-S      |hi3518cv100|AR0130     |k4b1g164               |+          |+          |+      |    |
|..|...         |...                |...        |...        |uart adapter broken    |           |           |       |    |
|25|RUISION     |RS-H802J-B0        |hi3518ev200|f22        |                       |+          |+          |+      |    |
|26|XM          |ivg-85hg50pya-s    |hi3516ev300|imx335     |                       |+          |+          |+      |    |
|27|unknown     |facility27         |hi3516dv100|it6801fn   |hdmi encoder, mr. ipcam|+          |+          |+      |    |
|28|unknown     |facility28         |hi3516cv500|imx327     |mr. ipcam              |+          |+          |+      |    |
|29|ruision     |rs-h805-a0         |hi3518ev200|ar0237     |                       |           |           |       |    | 
|30|unknown     |unknown            |xm530      |unknown    |no responce            |+          |+          |+      |    |
|31|HISILICON   |DEMB               |hi3516dv300|imx327     |                       |+          |+          |+-     |    |
|32|HISILICON   |DEMBVERC           |hi3559av100|imx334     |                       |+          |+          |+      |    |

## App success table

|family     |chip       |camera                                 |status |
|-----------|-----------|---------------------------------------|-------|
|hi3516cv100|hi3516cv100|xm_53H20-S_hi3516cv100_imx122          |+		|
|           |           |XM_53H20-AE_hi3516cv100_imx222         |+		|
|           |hi3518cv100|XM_IPG-53H13PL-S_hi3518cv100_AR0130    |+		|
|           |hi3518ev100|N/A                                    |		|
|hi3516cv200|hi3516cv200|N/A                                    |		|
|           |hi3518ev200|ruision_rs-h805-a0_hi3518ev200_ar0237  |-		|
|           |           |ruision_rs-h802j-b0_hi3518ev200_f22    |+		|
|           |           |unknown_facility20_hi3518ev200_h65     |?		|
|           |hi3518ev201|YI_YHS-113-IR_hi3518ev201_ov9732       |?		|
|hi3516cv300|hi3516cv300|XM_IVG-HP203Y-SE_hi3516cv300_imx291    |+		|
|           |           |XM_IVG-HP201Y-SE_hi3516cv300_imx323    |+		|
|           |           |JVT_S323H16XF_hi3516cv300_imx323       |+		|
|           |           |RUISION_RS-H622QM-B0_hi3516cv300_imx323|+		|
|           |hi3516ev100|XM_IPG-83HE20PY-S_hi3516ev100_imx323   |+		|
|hi3516av100|hi3516av100|XM_IPG-83H50P-B_hi3516av100_imx178     |+		|
|           |           |XM_83H40PL-B_hi3516av100_OV4689        |+		|
|           |hi3516dv100|SSQVISION_ON290H16D_hi3516dv100_imx290 |+		|
|           |           |unknown_facility27_hi3516dv100_it6801fn|?		|
|hi3516av200|hi3516av200|XM_IVG-83H80NV-BE_hi3516av200_OS08A10  |+		|
|           |hi3519v101 |JVT_S274H19V-L29_hi3519v101_imx274     |+		|
|           |           |JVT_S226H19V-L29_hi3519v101_imx226     |+		|
|           |           |SSQVISION_unknown_hi3519v101_imx326    |?      |
|hi3516ev200|hi3516ev200|XM_IVG-85HF20PYA-S_hi3516ev200_imx307  |+      |
|           |           |XM_IVG-85HF30PS-S_hi3516ev200_sc4239   |?      |
|           |           |TOPSEE_TH38J29_hi3516ev200_imx307      |+      |
|           |           |TOPSEE_TH38D16_hi3516ev200_sc3235      |?      |
|           |hi3516ev300|XM_ivg-85hg50pya-s_hi3516ev300_imx335  |+      |
|           |hi3518ev300|N/A                                    |		|
|           |hi3516dv200|N/A                                    |		|
|hi3516cv500|hi3516cv500|unknown_facility28_hi3516cv500_imx327  |+		|
|           |hi3516dv300|SSQVISION_ON335H16D_hi3516dv300_imx335 |+		|
|           |           |HISILICON_DEMB_hi3516dv300_imx327      |+		|
|hi3519av100|hi3519av100|N/A                                    |		|
|hi3559av100|hi3559av100|HISILICON_DEMBVERC_hi3559av100_imx334  |?		|


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
