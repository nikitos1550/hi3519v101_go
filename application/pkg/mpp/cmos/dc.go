package cmos

//typedef struct hiVI_SYNC_CFG_S { 
//  VI_VSYNC_E             enVsync;   
//  VI_VSYNC_NEG_E         enVsyncNeg;
//  VI_HSYNC_E             enHsync;   
//  VI_HSYNC_NEG_E         enHsyncNeg;  
//  VI_VSYNC_VALID_E       enVsyncValid;   
//  VI_VSYNC_VALID_NEG_E   enVsyncValidNeg;
//  VI_TIMING_BLANK_S      stTimingBlank;
//} VI_SYNC_CFG_S;

//typedef struct hiVI_TIMING_BLANK_S {
//    HI_U32 u32HsyncHfb ;    /* Horizontal front blanking width */
//    HI_U32 u32HsyncAct ;    /* Horizontal effetive width */
//    HI_U32 u32HsyncHbb ;    /* Horizontal back blanking width */
//    HI_U32 u32VsyncVfb ;    /* Vertical front blanking height of one frame or odd-field frame picture */
//    HI_U32 u32VsyncVact ;   /* Vertical effetive width of one frame or odd-field frame picture */
//    HI_U32 u32VsyncVbb ;    /* Vertical back blanking height of one frame or odd-field frame picture */
//    HI_U32 u32VsyncVbfb ;   /* Even-field vertical front blanking height when input mode is interlace (invalid when progressive input mode) */
//    HI_U32 u32VsyncVbact ;  /* Even-field vertical effetive width when input mode is interlace (invalid when progressive input mode) */
//    HI_U32 u32VsyncVbbb ;   /* Even-field vertical back blanking height when input mode is interlace (invalid when progressive input mode) */
//}VI_TIMING_BLANK_S;


//typedef enum hiVI_VSYNC_E {
//    VI_VSYNC_FIELD = 0,           /* Field/toggle mode:a signal reversal means a new frame or a field */
//    VI_VSYNC_PULSE,               /* Pusle/effective mode:a pusle or an effective signal means a new frame or a field */
//} VI_VSYNC_E;
type dcVSync uint
const (
	DCVSyncField	dcVSync = 1
	DCVSyncPulse	dcVSync = 2
)

//typedef enum hiVI_VSYNC_NEG_E {
//    VI_VSYNC_NEG_HIGH = 0,        /*if VIU_VSYNC_E = VIU_VSYNC_FIELD,then the vertical synchronization signal of even field is high-level,
//                                    if VIU_VSYNC_E = VIU_VSYNC_PULSE,then the vertical synchronization pulse is positive pulse.*/
//    VI_VSYNC_NEG_LOW              /*if VIU_VSYNC_E = VIU_VSYNC_FIELD,then the vertical synchronization signal of even field is low-level,
//                                    if VIU_VSYNC_E = VIU_VSYNC_PULSE,then the vertical synchronization pulse is negative pulse.*/
//} VI_VSYNC_NEG_E;
type dcVSyncNeg uint
const (
	DCVSyncNegHigh		dcVSyncNeg = 1
	DCVSyncNegLow		dcVSyncNeg = 2
)

//typedef enum hiVI_HSYNC_E {
//    VI_HSYNC_VALID_SINGNAL = 0,   /* the horizontal synchronization is valid signal mode */
//    VI_HSYNC_PULSE,               /* the horizontal synchronization is pulse mode, a new pulse means the beginning of a new line */
//} VI_HSYNC_E;
type dcHSync uint
const (
	DCHSyncSignal	dcHSync = 1
	DCHSyncPulse	dcHSync = 2
)

//typedef enum hiVI_HSYNC_NEG_E {
//    VI_HSYNC_NEG_HIGH = 0,        /*if VI_HSYNC_E = VI_HSYNC_VALID_SINGNAL,then the valid horizontal synchronization signal is high-level;
//                                    if VI_HSYNC_E = VI_HSYNC_PULSE,then the horizontal synchronization pulse is positive pulse */
//    VI_HSYNC_NEG_LOW              /*if VI_HSYNC_E = VI_HSYNC_VALID_SINGNAL,then the valid horizontal synchronization signal is low-level;
//                                    if VI_HSYNC_E = VI_HSYNC_PULSE,then the horizontal synchronization pulse is negative pulse */
//} VI_HSYNC_NEG_E;
type dcHSyncNeg uint 
const (
	DCHSyncNegHigh	dcHSyncNeg = 1
	DCHSyncNegLow	dcHSyncNeg = 2
)

//typedef enum hiVI_VSYNC_VALID_E {
//    VI_VSYNC_NORM_PULSE = 0,      /* the vertical synchronization is pusle mode, a pusle means a new frame or field  */
//    VI_VSYNC_VALID_SINGAL,        /* the vertical synchronization is effective mode, a effective signal means a new frame or field */
//} VI_VSYNC_VALID_E;
type dcVSyncValid uint
const (
	DCVSyncValidPulse	dcVSyncValid = 1
	DCVSyncValidSignal	dcVSyncValid = 2
)

//typedef enum hiVI_VSYNC_VALID_NEG_E {
//    VI_VSYNC_VALID_NEG_HIGH = 0,  /*if VI_VSYNC_VALID_E = VI_VSYNC_NORM_PULSE,a positive pulse means vertical synchronization pulse;
//                                    if VI_VSYNC_VALID_E = VI_VSYNC_VALID_SINGAL,the valid vertical synchronization signal is high-level */
//    VI_VSYNC_VALID_NEG_LOW        /*if VI_VSYNC_VALID_E = VI_VSYNC_NORM_PULSE,a negative pulse means vertical synchronization pulse;
//                                    if VI_VSYNC_VALID_E = VI_VSYNC_VALID_SINGAL,the valid vertical synchronization signal is low-level */
//} VI_VSYNC_VALID_NEG_E;
type dcVSyncValidNeg uint
const (
	DCVSyncValidNegHigh	dcVSyncValidNeg = 1
	DCVSyncValidNegLow	dcVSyncValidNeg = 2
)

type dcSyncAttr struct {
	VSync			dcVSync
	VSyncNeg		dcVSyncNeg
	HSync			dcHSync
	HSyncNeg		dcHSyncNeg
	VSyncValid		dcVSyncValid
	VSyncValidNeg	dcVSyncValidNeg

    TimingHfb      uint
    TimingAct      uint
    TimingHbb      uint
    TimingVfb      uint
    TimingVact     uint
    TimingVbb      uint
    TimingVbfb     uint
    TimingVbact    uint
    TimingVbbb     uint

	/*
    timingHsyncHfb      uint
    timingHsyncAct      uint
    timingHsyncHbb      uint
    timingVsyncVfb      uint
    timingVsyncVact     uint
    timingVsyncVbb      uint
    timingVsyncVbfb     uint
    timingVsyncVbact    uint
    timingVsyncVbbb     uint
	*/
}

