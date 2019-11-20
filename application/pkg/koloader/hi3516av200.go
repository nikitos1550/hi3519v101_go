//+build hi3516av200

//go:generate rm -f hi3516av200_kobin.go
//go:generate go run -tags "hi3516av200" generate.go --output hi3516av200_kobin.go --tag hi3516av200 --dir ../../sdk/hi3516av200/ko --pkg koloader --source ./hi3516av200.go

package koloader

var (

    Modules = [...][2]string {
        [2]string{"hi_osal.ko",             "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi3519v101_base.ko",     ""},
        [2]string{"hi3519v101_sys.ko",      "vi_vpss_online=0 sensor=NULL,NULL mem_total={memTotalSize}"},
        //[2]string{"hi3519v101_tde.ko",     ""},
        [2]string{"hi3519v101_region.ko",      ""},
        [2]string{"hi3519v101_fisheye.ko",     ""},
        [2]string{"hi3519v101_vgs.ko",         ""},
        [2]string{"hi3519v101_isp.ko",         "proc_param=30"},
        [2]string{"hi3519v101_viu.ko",         "detect_err_frame=10"},
        [2]string{"hi3519v101_vpss.ko",        ""},
        [2]string{"hi3519v101_vou.ko",         ""},
        //[2]string{"hifb.ko",              "video='hifb:vram0_size:1620'"},
        [2]string{"hi3519v101_rc.ko",        ""},
        [2]string{"hi3519v101_venc.ko",      ""},
        [2]string{"hi3519v101_chnl.ko",      ""},
        [2]string{"hi3519v101_vedu.ko",      ""},
        [2]string{"hi3519v101_h264e.ko",     ""},
        [2]string{"hi3519v101_h265e.ko",     ""},
        [2]string{"hi3519v101_jpege.ko",     ""},
        //[2]string{"hi3519v101_ive.ko",    "save_power=1"},
        //[2]string{"hi3519v101_photo.ko",  ""},
        //[2]string{"hi_sensor_i2c.ko",     ""},
        [2]string{"hi_pwm.ko",               ""},
        //[2]string{"hi_piris.ko",          ""},
        //[2]string{"hi_sil9136.ko",        "norm=12"},
        //[2]string{"gyro_bosch.ko",        ""},
        //[2]string{"hi3519v101_aio.ko",    ""},
        //[2]string{"hi3519v101_ai.ko",     ""},
        //[2]string{"hi3519v101_ao.ko",     ""},
        //[2]string{"hi3519v101_aenc.ko",   ""},
        //[2]string{"hi3519v101_adec.ko",   ""},
        //[2]string{"hi_acodec.ko",         ""},
        //[2]string{"hi_tlv320aic31.ko",    ""},
        [2]string{"hi_mipi.ko",              ""},
        [2]string{"hi_user.ko",              ""},
        [2]string{"hi_ssp_sony.ko",          ""},
    }


)
