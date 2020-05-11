//+build nobuild

//+build arm
//+build hi3516cv300
//+build imx323,cmos_data_dc,cmos_control_spi,cmos_bus_0

package cmos

var (
    cmosItem = cmos{
        vendor: "Sony",
        model: "imx323",
        dcZeroBitOffset:  4, //typical for hisi
        modes: []cmosMode {
            cmosMode {
                mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                viCrop:     crop{X0: 200, Y0: 20, Width: 1920, Height: 1080,},
                ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                width: 1920,
                height: 1080,
                fps: 30,
                bitness: 12,
                data: DC,
                dcSync: dcSyncAttr {
                    VSync:          DCVSyncPulse,
                    VSyncNeg:       DCVSyncNegHigh,
                    HSync:          DCHSyncSignal,
                    HSyncNeg:       DCHSyncNegHigh,
                    VSyncValid:     DCVSyncValidSignal,
                    VSyncValidNeg:  DCVSyncValidNegHigh,
                    TimingHfb:      0,
                    TimingAct:      1920,
                    TimingHbb:      0,
                    TimingVfb:      0,
                    TimingVact:     1080,
                    TimingVbb:      0,
                    TimingVbfb:     0,
                    TimingVbact:    0,
                    TimingVbbb:     0,
                },
                //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                clock: 37.125,
                wdr: WDRNone,
                description: "normal",
            },
            cmosMode {
                mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                viCrop:     crop{X0: 200, Y0: 20, Width: 1280, Height: 720,},
                ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                width: 1280,
                height: 720,
                fps: 30,
                bitness: 12,
                data: DC,
                dcSync: dcSyncAttr {
                    VSync:          DCVSyncPulse,
                    VSyncNeg:       DCVSyncNegHigh,
                    HSync:          DCHSyncSignal,
                    HSyncNeg:       DCHSyncNegHigh,
                    VSyncValid:     DCVSyncValidSignal,
                    VSyncValidNeg:  DCVSyncValidNegHigh,
                    TimingHfb:      0,
                    TimingAct:      1920,
                    TimingHbb:      0,
                    TimingVfb:      0,
                    TimingVact:     1080,
                    TimingVbb:      0,
                    TimingVbfb:     0,
                    TimingVbact:    0,
                    TimingVbbb:     0,
                },
                //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                clock: 37.125,
                wdr: WDRNone,
                description: "720p 30fps 12bit",
            },
            cmosMode {
                mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                viCrop:     crop{X0: 200, Y0: 20, Width: 1280, Height: 720,},
                ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                width: 1280,
                height: 720,
                fps: 60,
                bitness: 10,
                data: DC,
                dcSync: dcSyncAttr {
                    VSync:          DCVSyncPulse,
                    VSyncNeg:       DCVSyncNegHigh,
                    HSync:          DCHSyncSignal,
                    HSyncNeg:       DCHSyncNegHigh,
                    VSyncValid:     DCVSyncValidSignal,
                    VSyncValidNeg:  DCVSyncValidNegHigh,
                    TimingHfb:      0,
                    TimingAct:      1280,
                    TimingHbb:      0,
                    TimingVfb:      0,
                    TimingVact:     720,
                    TimingVbb:      0,
                    TimingVbfb:     0,
                    TimingVbact:    0,
                    TimingVbbb:     0,
                },
                //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                clock: 37.125,
                wdr: WDRNone,
                description: "720p 60fps 10bit",
            },
        },
        control: cmosControl {
            bus: SPI,
            busNum: 0,
        },
        data: DC,
        bayer: RGGB,
    }
)

