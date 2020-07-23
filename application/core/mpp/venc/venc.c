#include "venc.h"

#if 1
#include <string.h>

#define FrameSavingMode 0
#define MJPEGBUFK       3
#define H26XK           1.5

int mpp_venc_mjpeg_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_encoder_in *in) {
    switch (in->bitrate_control) {
        case VENC_RC_MODE_MJPEGCBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;

            #if HI_MPP == 1
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32StatTime          = in->stat_time;           //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32ViFrmRate         = in->in_fps;              //mpp1: (0; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.fr32TargetFrmRate    = in->out_fps;             //mpp1: (0; u32ViFrmRate]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel    = in->fluctuate_level;     //mpp1: [0; 5] (0 is recommended)
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32BitRate           = in->bitrate;             //mpp1: [2; 40960]
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32StatTime          = in->stat_time;           //mpp2: [1; 60]                 //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32SrcFrmRate        = in->in_fps;              //mpp2: [1; 240]                //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.fr32DstFrmRate       = in->out_fps;             //mpp2: (0; u32SrcFrmRate]      //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel    = in->fluctuate_level;     //mpp2: [1; 5] (1 recommended)  //mpp3: [1; 5]
                stVencChnAttr->stRcAttr.stAttrMjpegeCbr.u32BitRate           = in->bitrate;             //mpp2: [2; 102400]             //mpp3: [2; 102400]
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stMjpegCbr.u32StatTime               = in->stat_time;           //mpp4: [1;60]
                stVencChnAttr->stRcAttr.stMjpegCbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1;240]
                stVencChnAttr->stRcAttr.stMjpegCbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stMjpegCbr.u32BitRate                = in->bitrate;             //mpp4: Hi3559A V100ES/Hi3559A V100: [2, 409600] Hi3519A V100/Hi3556A V100/Hi3516C V500/Hi3516D V300: [2, 204800]
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300/Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
            break;
        case VENC_RC_MODE_MJPEGVBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_MJPEGVBR;

            #if HI_MPP == 1
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32StatTime          = in->stat_time;           //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32ViFrmRate         = in->in_fps;              //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.fr32TargetFrmRate    = in->out_fps;             //mpp1: (0; u32ViFrmRate]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MinQfactor        = in->min_q_factor;        //mpp1: [1; u32MaxQfactor)
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MaxQfactor        = in->max_q_factor;        //mpp1: [1; 99]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MaxBitRate        = in->bitrate;             //mpp1: [2; 40960]
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32StatTime          = in->stat_time;           //mpp2: [1; 60]                 //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32SrcFrmRate        = in->in_fps;              //mpp2: [1; 240]                //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.fr32DstFrmRate       = in->out_fps;             //mpp2: (0; u32SrcFrmRate]      //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MinQfactor        = in->min_q_factor;        //mpp2: [1; u32MaxQfactor)      //mpp3: [1; u32MaxQfactor]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MaxQfactor        = in->max_q_factor;        //mpp2: [1; 99]                 //mpp3: [1; 99]
                stVencChnAttr->stRcAttr.stAttrMjpegeVbr.u32MaxBitRate        = in->bitrate;             //mpp2: [2; 102400]             //mpp3: [2; 102400]
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stMjpegVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60] 
                stVencChnAttr->stRcAttr.stMjpegVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stMjpegVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stMjpegVbr.u32MaxBitRate             = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300/Hi3556V200/Hi3559V200：[2, 102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
            break;
        case VENC_RC_MODE_MJPEGFIXQP:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_MJPEGFIXQP;

            #if HI_MPP == 1
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.u32Qfactor         = in->q_factor;            //mpp1: [1; 99]
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.u32ViFrmRate       = in->in_fps;              //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.fr32TargetFrmRate  = in->out_fps;             //mpp1: (0, u32ViFrmRate]
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.u32Qfactor         = in->q_factor;            //mpp2: [1; 99]                 //mpp3: [1; 99]
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.u32SrcFrmRate      = in->in_fps;              //mpp2: [1; 240]                //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrMjpegeFixQp.fr32DstFrmRate     = in->out_fps;             //mpp2: (0; u32ViFrmRate]       //mpp3: [1/16; u32ViFrmRate]
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stMjpegFixQp.u32Qfactor             = in->q_factor;             //mpp4: [1; 99]
                stVencChnAttr->stRcAttr.stMjpegFixQp.u32SrcFrameRate        = in->in_fps;               //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stMjpegFixQp.fr32DstFrameRate       = in->out_fps;              //mpp4: [1/64; u32SrcFrmRate]
            #endif
            break;
        default:
            ;;;//TODO error
    }

    return ERR_NONE;
}

int mpp_venc_h264_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_encoder_in *in) {
    switch (in->bitrate_control) {
        case VENC_RC_MODE_H264CBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;

            #if HI_MPP == 1
                stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264CBRv2;   //!!!
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32Gop                 = in->gop;                 //mpp1: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32StatTime            = in->stat_time;           //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32ViFrmRate           = in->in_fps;              //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.fr32TargetFrmRate      = in->out_fps;             //mpp1: (0; u32ViFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32BitRate             = in->bitrate;             //mpp1: [2; 40960]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32FluctuateLevel      = in->fluctuate_level;     //mpp1: [1; 5]
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32Gop                 = in->gop;                 //mpp2: [1; 65536]              //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32StatTime            = in->stat_time;           //mpp2: [1; 60]                 //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32SrcFrmRate          = in->in_fps;              //mpp2: [1; 240]                //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.fr32DstFrmRate         = in->out_fps;             //mpp2: (0; u32SrcFrmRate]      //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32BitRate             = in->bitrate;             //mpp2: [2, 102400]             //mpp3: 
                                                                                                                                        //Hi3519 V100/Hi3519 V101/Hi3531D V100/Hi3521DV100/Hi3536C V100: [2, 102400]
                                                                                                                                        //Hi3516C V300/Hi3516E V100: [2, 30720]
                stVencChnAttr->stRcAttr.stAttrH264Cbr.u32FluctuateLevel      = in->fluctuate_level;     //mpp2: [0; 5]                  //mpp3: [1; 5]
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH264Cbr.u32Gop                     = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH264Cbr.u32StatTime                = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH264Cbr.u32SrcFrameRate            = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH264Cbr.fr32DstFrameRate           = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH264Cbr.u32BitRate                 = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300/：[2,51200]
                                                                                                        //Hi3556V200/ Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
            break;
        case VENC_RC_MODE_H264VBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264VBR;

            #if HI_MPP == 1
                stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264VBRv2;   //!!!
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32Gop                 = in->gop;                 //mpp1: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32StatTime            = in->stat_time;           //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32ViFrmRate           = in->in_fps;              //mpp1: [1, 60]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.fr32TargetFrmRate      = in->out_fps;             //mpp1: (0; u32ViFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MinQp               = in->min_qp;              //mpp1: [0; 51]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MaxQp               = in->max_qp;              //mpp1: (u32MinQp; 51]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MaxBitRate          = in->bitrate;             //mpp1: [2; 40960]
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32Gop                 = in->gop;                 //mpp2: [1; 65536]              //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32StatTime            = in->stat_time;           //mpp2: [1; 60]                 //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32SrcFrmRate          = in->in_fps;              //mpp2: [1; 240]                //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.fr32DstFrmRate         = in->out_fps;             //mpp2: (0; u32SrcFrmRate]      //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MinQp               = in->min_qp;              //mpp2: [0; 51]                 //mpp3: [0; 51]
                #if HI_MPP == 3
                    stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MinIQp          = in->min_i_qp;                                            //mpp3: [u32MinQp, u32MaxQp]
                #endif
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MaxQp               = in->max_qp;              //mpp2: (u32MinQp; 51]          //mpp3: [u32MinQp; 51]
                stVencChnAttr->stRcAttr.stAttrH264Vbr.u32MaxBitRate          = in->bitrate;             //mpp2: [2; 102400]             //mpp3: 
																																		//Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
																																		//Hi3516C V300/Hi3516E V100: [2, 30720]<Paste>
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH264Vbr.u32Gop                     = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH264Vbr.u32StatTime                = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH264Vbr.u32SrcFrameRate            = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH264Vbr.fr32DstFrameRate           = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH264Vbr.u32MaxBitRate              = in->bitrate;             //mpp4: 
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
            break;
        case VENC_RC_MODE_H264FIXQP:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264FIXQP;

    		#if HI_MPP == 1
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32Gop               = in->gop;                 //mpp1: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32ViFrmRate         = in->in_fps;              //mpp1: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.fr32TargetFrmRate    = in->out_fps;             //mpp1: (0; u32ViFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32IQp               = in->i_qp;                //mpp1: [0; 51]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32PQp               = in->p_qp;                //mpp1: [0; 51]
          	#elif HI_MPP == 2 || \
				  HI_MPP == 3
        		stVencChnAttr->stRcAttr.stAttrH264FixQp.u32Gop               = in->gop;                 //mpp2: [1; 65536]				//mpp3: [1; 65536]
            	stVencChnAttr->stRcAttr.stAttrH264FixQp.u32SrcFrmRate        = in->in_fps;              //mpp2: [1; 240]				//mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.fr32DstFrmRate       = in->out_fps;             //mpp2: (0; u32SrcFrmRate]		//mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32IQp               = in->i_qp;                //mpp2: [0; 51]					//mpp3: [0; 51]
                stVencChnAttr->stRcAttr.stAttrH264FixQp.u32PQp               = in->p_qp;                //mpp2: [0; 51]					//mpp3: [0; 51]
				#if HI_MPP == 3
                	stVencChnAttr->stRcAttr.stAttrH264FixQp.u32BQp           = in->b_qp;												//mpp3: [0; 51]
				#endif
			#elif HI_MPP == 4
    			stVencChnAttr->stRcAttr.stH264FixQp.u32Gop                   = in->gop;                 //mpp4: [1; 65536]
    			stVencChnAttr->stRcAttr.stH264FixQp.u32SrcFrameRate          = in->in_fps;              //mpp4: [1; 240]
    			stVencChnAttr->stRcAttr.stH264FixQp.fr32DstFrameRate         = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
    			stVencChnAttr->stRcAttr.stH264FixQp.u32IQp                   = in->i_qp;                //mpp4: [0; 51]
    			stVencChnAttr->stRcAttr.stH264FixQp.u32PQp                   = in->p_qp;                //mpp4: [0; 51]
    			stVencChnAttr->stRcAttr.stH264FixQp.u32BQp                   = in->b_qp;                //mpp4: [0; 51]
    		#endif
            break;
		case VENC_RC_MODE_H264AVBR:
			stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264AVBR;

			#if HI_MPP == 1
				;;;//TODO error
			#elif HI_MPP == 2 || \
				  HI_MPP == 3
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32Gop                = in->gop;                 //mpp2: [1; 65536]         		//mpp3: [1; 65536] 
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32StatTime           = in->stat_time;           //mpp2: [1; 60]					//mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32SrcFrmRate         = in->in_fps;              //mpp2: [1; 240]               	//mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.fr32DstFrmRate        = in->out_fps;             //mpp2: (0; u32SrcFrmRate]      //mpp3: (0; u32SrcFrmRate]  
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32MaxBitRate         = in->bitrate;             //mpp2: [2; 102400]				//mpp3:
																																		//Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
																																		//Hi3516C V300/Hi3516E V100: [2, 30720]
    		#elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH264AVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH264AVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH264AVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH264AVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH264AVbr.u32MaxBitRate             = in->bitrate;             //mpp4: 
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
			break;
        case VENC_RC_MODE_H264CVBR:
			stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264CVBR;

			#if HI_MPP == 1 || \
				HI_MPP == 2 || \
				HI_MPP == 3 
				;;;//TODO error
			#elif HI_MPP == 4
            	stVencChnAttr->stRcAttr.stH264CVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH264CVbr.u32StatTime               = in->stat_time;           //mpp4: [1, 60]
                stVencChnAttr->stRcAttr.stH264CVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH264CVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH264CVbr.u32LongTermStatTime       = 1;                       //mpp4: [1, 1440]
                stVencChnAttr->stRcAttr.stH264CVbr.u32ShortTermStatTime      = 1;                       //mpp4: [1; 120]
                stVencChnAttr->stRcAttr.stH264CVbr.u32MaxBitRate             = in->bitrate;             //mpp4: 
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
                stVencChnAttr->stRcAttr.stH264CVbr.u32LongTermMaxBitrate     = in->bitrate;             //mpp4: [2; u32MaxBitRate]  //1024 + 512*u32FrameRate/30; //TODO
                stVencChnAttr->stRcAttr.stH264CVbr.u32LongTermMinBitrate     = 256;                     //mpp4: [0; u32LongTermMaxBitrate]
         	#endif
			break;
		case VENC_RC_MODE_H264QVBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H264QVBR;

            #if HI_MPP == 1 || \
                HI_MPP == 2
				;;; //TODO error
			#elif HI_MPP == 3 
            	stVencChnAttr->stRcAttr.stAttrH264QVbr.u32Gop                = in->gop;                 //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264QVbr.u32StatTime           = in->stat_time;           //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264QVbr.u32SrcFrmRate         = in->in_fps;              //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264QVbr.fr32DstFrmRate        = in->out_fps;             //mpp3: (0; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264QVbr.u32TargetBitRate      = in->bitrate;             //mpp3: 
                                                                                                        //Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
                                                                                                        //Hi3516C V300/Hi3516E V100: [2, 30720]
          	#elif HI_MPP == 4
            	stVencChnAttr->stRcAttr.stH264QVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH264QVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH264QVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH264QVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH264QVbr.u32TargetBitRate          = in->bitrate;             //mpp4: 
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
          	#endif
			break;
        default:
            ;;;//TODO error
    }

	#if HI_MPP == 4
    	stVencChnAttr->stVencAttr.stAttrH264e.bRcnRefShareBuf        = FrameSavingMode;
    #endif

    return ERR_NONE;
}

int mpp_venc_h265_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_encoder_in *in) {                                                  
    switch (in->bitrate_control) {
		case VENC_RC_MODE_H265CBR:																		
			stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;

            #if HI_MPP == 2 || \
                HI_MPP == 3
                stVencChnAttr->stRcAttr.stAttrH265Cbr.u32Gop                 = in->gop;               	//mpp2: [1; 65536]                  //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH265Cbr.u32StatTime            = in->stat_time;           //mpp2: [1; 60]                     //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH265Cbr.u32SrcFrmRate          = in->in_fps;              //mpp2: [1; 240]                    //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH265Cbr.fr32DstFrmRate         = in->out_fps;           	//mpp2: (0; u32SrcFrmRate]          //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH265Cbr.u32BitRate             = in->bitrate;             //mpp2: [2; 102400]                 //mpp3:
                                                                                                                                            //Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
                                                                                                                                            //Hi3516C V300/Hi3516E V100: [2, 30720]
                stVencChnAttr->stRcAttr.stAttrH265Cbr.u32FluctuateLevel      = in->fluctuate_level;   	//mpp2: [0, 5] (0 is recommended)???//mpp3: [1, 5] (1 is recommended)
			#elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH265Cbr.u32Gop                     = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH265Cbr.u32StatTime                = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH265Cbr.u32SrcFrameRate            = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265Cbr.fr32DstFrameRate           = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH265Cbr.u32BitRate                 = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
          	#endif
			break;
		case VENC_RC_MODE_H265VBR:
			stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265VBR;

			#if HI_MPP == 1
				;;;//TODO error
			#elif HI_MPP == 2 || \
				  HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrH265Vbr.u32Gop                 = in->gop;                 //mpp2: [1; 65536]                  //mpp3: [1; 65536] 
                stVencChnAttr->stRcAttr.stAttrH265Vbr.u32StatTime            = in->stat_time;           //mpp2: [1; 60]                     //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH265Vbr.u32SrcFrmRate          = in->in_fps;              //mpp2: [1; 240]                    //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH265Vbr.fr32DstFrmRate         = in->out_fps;             //mpp2: (0; u32SrcFrmRate]          //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH265Vbr.u32MinQp               = in->min_qp;              //mpp2: [0; 51]                     //mpp3: [0; 51]
				#if HI_MPP == 3
                    stVencChnAttr->stRcAttr.stAttrH265Vbr.u32MinIQp          = in->min_i_qp;                                                //mpp3: [u32MinQp; u32MaxQp]
				#endif
                stVencChnAttr->stRcAttr.stAttrH265Vbr.u32MaxQp               = in->max_qp;              //mpp2: (u32MinQp; 51]              //mpp3: [u32MinQp; 51]
        		stVencChnAttr->stRcAttr.stAttrH265Vbr.u32MaxBitRate          = in->bitrate;             //mpp2: [2; 102400]                 //mpp3:
                                                                                                                                            //Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
                                                                                                                                            //Hi3516C V300/Hi3516E V100: [2, 30720]
    		#elif HI_MPP == 4 
            	stVencChnAttr->stRcAttr.stH265Vbr.u32Gop                     = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH265Vbr.u32StatTime                = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH265Vbr.u32SrcFrameRate            = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265Vbr.fr32DstFrameRate           = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH265Vbr.u32MaxBitRate              = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
          	#endif
		case VENC_RC_MODE_H265FIXQP:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265FIXQP;

            #if HI_MPP == 1
            	;;;//TODO error
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrH265FixQp.u32Gop               = in->gop;                 //mpp2: [1; 65536]                  //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH265FixQp.u32SrcFrmRate        = in->in_fps;              //mpp2: [1; 240]                    //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH265FixQp.fr32DstFrmRate       = in->out_fps;             //mpp2: (0; u32SrcFrmRate]          //mpp3: [1/16; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH265FixQp.u32IQp               = in->i_qp;                //mpp2: [0; 51]                     //mpp3: [0; 51]
                stVencChnAttr->stRcAttr.stAttrH265FixQp.u32PQp               = in->p_qp;                //mpp2: [0; 51]                     //mpp3: [0; 51]
				#if HI_MPP == 3
					stVencChnAttr->stRcAttr.stAttrH265FixQp.u32BQp           = in->b_qp;                                                    //mpp3: [0; 51]
				#endif
			#elif HI_MPP == 4
            	stVencChnAttr->stRcAttr.stH265FixQp.u32Gop                   = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH265FixQp.u32SrcFrameRate          = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265FixQp.fr32DstFrameRate         = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH265FixQp.u32IQp                   = in->i_qp;                //mpp4: [0; 51]
                stVencChnAttr->stRcAttr.stH265FixQp.u32PQp                   = in->p_qp;                //mpp4: [0; 51]
                stVencChnAttr->stRcAttr.stH265FixQp.u32BQp                   = in->b_qp;                //mpp4: [0; 51]
           	#endif
		case VENC_RC_MODE_H265AVBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265AVBR;

            #if HI_MPP == 1
            	;;;//TODO error
            #elif HI_MPP == 2 || \
            	  HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32Gop                = in->gop;                 //mpp2: [1; 65536]                  //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32StatTime           = in->stat_time;           //mpp2: [1; 60]                     //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32SrcFrmRate         = in->in_fps;              //mpp2: [1; 240]                    //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.fr32DstFrmRate        = in->out_fps;             //mpp2: (0; u32SrcFrmRate]          //mpp3: (0; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH264AVbr.u32MaxBitRate         = in->bitrate;             //mpp2: [2; 102400]                 //mpp3: 
                                                                                                                                            //Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
                                                                                                                                            //Hi3516C V300/Hi3516E V100: [2, 30720]
    		#elif HI_MPP == 4
    			stVencChnAttr->stRcAttr.stH265AVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
        		stVencChnAttr->stRcAttr.stH265AVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60]
            	stVencChnAttr->stRcAttr.stH265AVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265AVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH265AVbr.u32MaxBitRate             = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
         	#endif
        case VENC_RC_MODE_H265CVBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265CVBR;

            #if HI_MPP == 1 || \
                HI_MPP == 2 || \
                HI_MPP == 3 
                ;;;//TODO error
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH265CVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH265CVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH265CVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265CVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stH265CVbr.u32LongTermStatTime       = 1;                       //mpp4: [1; 1440]
                stVencChnAttr->stRcAttr.stH265CVbr.u32ShortTermStatTime      = 1;                       //mpp4: [1; 120]
                stVencChnAttr->stRcAttr.stH265CVbr.u32MaxBitRate             = in->bitrate;             //mpp4: 
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]。
                stVencChnAttr->stRcAttr.stH265CVbr.u32LongTermMaxBitrate     = in->bitrate;             //mpp4: [2; u32MaxBitRate]          //1024 + 512*u32FrameRate/30; //TODO
                stVencChnAttr->stRcAttr.stH265CVbr.u32LongTermMinBitrate     = 256;                     //mpp4: [0; u32LongTermMaxBitrate]
            #endif
            break;
        case VENC_RC_MODE_H265QVBR:
            stVencChnAttr->stRcAttr.enRcMode = VENC_RC_MODE_H265QVBR;

            #if HI_MPP == 1 || \
                HI_MPP == 2
                ;;;//TODO error
            #elif HI_MPP == 3 
                stVencChnAttr->stRcAttr.stAttrH265QVbr.u32Gop                = in->gop;                 //mpp3: [1; 65536]
                stVencChnAttr->stRcAttr.stAttrH265QVbr.u32StatTime           = in->stat_time;           //mpp3: [1; 60]
                stVencChnAttr->stRcAttr.stAttrH265QVbr.u32SrcFrmRate         = in->in_fps;              //mpp3: [1; 240]
                stVencChnAttr->stRcAttr.stAttrH265QVbr.fr32DstFrmRate        = in->out_fps;             //mpp3: (0; u32SrcFrmRate]
                stVencChnAttr->stRcAttr.stAttrH265QVbr.u32TargetBitRate      = in->bitrate;             //mpp3: 
                                                                                                        //Hi3519 V100/Hi3519 V101/Hi3531DV100/Hi3521D V100/Hi3536C V100: [2, 102400]
                                                                                                        //Hi3516C V300/Hi3516E V100: [2, 30720]
            #elif HI_MPP == 4
                stVencChnAttr->stRcAttr.stH265QVbr.u32Gop                    = in->gop;                 //mpp4: [1; 65536]
                stVencChnAttr->stRcAttr.stH265QVbr.u32StatTime               = in->stat_time;           //mpp4: [1; 60]
                stVencChnAttr->stRcAttr.stH265QVbr.u32SrcFrameRate           = in->in_fps;              //mpp4: [1; 240]
                stVencChnAttr->stRcAttr.stH265QVbr.fr32DstFrameRate          = in->out_fps;             //mpp4: [1/64; u32SrcFrmRate] 
                stVencChnAttr->stRcAttr.stH265QVbr.u32TargetBitRate          = in->bitrate;             //mpp4:
                                                                                                        //Hi3559AV100ES/Hi3559AV100：[2, 614400]
                                                                                                        //Hi3519AV100/Hi3556AV100：[2, 204800]
                                                                                                        //Hi3516CV500/Hi3516DV300/Hi3516AV300：[2,51200]
                                                                                                        //Hi3556V200/Hi3559V200：[2,102400]
                                                                                                        //Hi3516EV200/Hi3516EV300/Hi3518EV300/Hi3516DV200：[2,61440]
            #endif
            break;
		default:
			;;;//TODO error   
    }   
    
    #if HI_MPP == 4
        stVencChnAttr->stVencAttr.stAttrH265e.bRcnRefShareBuf        = FrameSavingMode;
   	#endif

    return ERR_NONE;
}


int mpp_venc_create_encoder(error_in *err, mpp_venc_create_encoder_in *in) {
    VENC_CHN_ATTR_S stVencChnAttr;

    #if HI_MPP == 1 || \
        HI_MPP == 2 || \
        HI_MPP == 3 

    switch (in->codec) {
        case PT_MJPEG:
            stVencChnAttr.stVeAttr.enType   = PT_MJPEG;

            #if HI_MPP == 1
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32BufSize               = in->width * in->height * MJPEGBUFK;
                stVencChnAttr.stVeAttr.stAttrMjpeg.bByFrame                 = HI_TRUE;
                stVencChnAttr.stVeAttr.stAttrMjpeg.bMainStream              = HI_TRUE;  //TODO
                stVencChnAttr.stVeAttr.stAttrMjpeg.bVIField                 = HI_FALSE; //TODO
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32Priority              = 0;        //TODO
            #elif HI_MPP == 2
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32BufSize               = in->width * in->height * MJPEGBUFK;
                stVencChnAttr.stVeAttr.stAttrMjpeg.bByFrame                 = HI_TRUE;
            #elif HI_MPP == 3
                stVencChnAttr.stVeAttr.stAttrMjpege.u32MaxPicWidth          = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32MaxPicHeight         = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32PicWidth             = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32PicHeight            = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32BufSize              = in->width * in->height * MJPEGBUFK;
                stVencChnAttr.stVeAttr.stAttrMjpege.bByFrame                = HI_TRUE;
            #endif
            break;
        case PT_H264:
            stVencChnAttr.stVeAttr.enType   = PT_H264;

            #if HI_MPP == 1
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32BufSize               = in->width * in->height * H26XK;
                stVencChnAttr.stVeAttr.stAttrH264e.u32Profile               = in->profile;
                stVencChnAttr.stVeAttr.stAttrH264e.bByFrame                 = HI_TRUE;
                stVencChnAttr.stVeAttr.stAttrH264e.bField                   = HI_FALSE;
                stVencChnAttr.stVeAttr.stAttrH264e.bMainStream              = HI_TRUE;  //TODO
                stVencChnAttr.stVeAttr.stAttrH264e.u32Priority              = 0;        //TODO
                stVencChnAttr.stVeAttr.stAttrH264e.bVIField                 = HI_FALSE; //TODO
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32BufSize               = in->width * in->height * H26XK;
                stVencChnAttr.stVeAttr.stAttrH264e.u32Profile               = in->profile;  
                stVencChnAttr.stVeAttr.stAttrH264e.bByFrame                 = HI_TRUE;   
                #if HI_MPP == 2
                    stVencChnAttr.stVeAttr.stAttrH264e.u32BFrameNum         = 0;
                    stVencChnAttr.stVeAttr.stAttrH264e.u32RefNum            = 1;
                #endif 
            #endif

            break;
        case PT_H265:
            stVencChnAttr.stVeAttr.enType   = PT_H265;

            #if HI_MPP == 1
                ;;;//TODO error
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr.stVeAttr.stAttrH265e.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrH265e.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrH265e.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrH265e.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrH265e.u32BufSize               = in->width * in->height * H26XK;
                stVencChnAttr.stVeAttr.stAttrH265e.u32Profile               = in->profile;
                stVencChnAttr.stVeAttr.stAttrH265e.bByFrame                 = HI_TRUE;
                #if HI_MPP == 2
                    stVencChnAttr.stVeAttr.stAttrH265e.u32BFrameNum         = 0;
                    stVencChnAttr.stVeAttr.stAttrH265e.u32RefNum            = 1;
                #endif
            #endif

            break;
        default:
            ;;;//TODO error
    }
    #endif

    #if HI_MPP == 4
        stVencChnAttr.stVencAttr.u32MaxPicWidth                     = in->width;
        stVencChnAttr.stVencAttr.u32MaxPicHeight                    = in->height;
        stVencChnAttr.stVencAttr.u32PicWidth                        = in->width;
        stVencChnAttr.stVencAttr.u32PicHeight                       = in->height;
        switch (in->codec) {
            case PT_MJPEG:
                stVencChnAttr.stVencAttr.enType                     = PT_MJPEG;
                stVencChnAttr.stVencAttr.u32BufSize                 = in->width * in->height * MJPEGBUFK;
                break;
            case PT_H264:
                stVencChnAttr.stVencAttr.enType                     = PT_H264;
                stVencChnAttr.stVencAttr.u32BufSize                 = in->width * in->height * H26XK;
                break;
            case PT_H265:
                stVencChnAttr.stVencAttr.enType                     = PT_H265;
                stVencChnAttr.stVencAttr.u32BufSize                 = in->width * in->height * H26XK;
                break;
        }
        stVencChnAttr.stVencAttr.u32Profile                         = in->profile;
        stVencChnAttr.stVencAttr.bByFrame                           = HI_TRUE;
    #endif

    switch (in->codec) {
        case PT_MJPEG:
            mpp_venc_mjpeg_params(&stVencChnAttr, in);
            break;
        case PT_H264:
            mpp_venc_h264_params(&stVencChnAttr, in);
            break;
        case PT_H265:
            mpp_venc_h265_params(&stVencChnAttr, in);
            break;
        default:
            ;;;//TODO error
    }

    //////////////////

    #if HI_MPP == 3 \
        || HI_MPP == 4
    
    switch (in->gop_mode) {
        case VENC_GOPMODE_NORMALP:
			stVencChnAttr.stGopAttr.enGopMode  					= VENC_GOPMODE_NORMALP;      
            stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta      = in->i_pq_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            break;
        case VENC_GOPMODE_DUALP:
			stVencChnAttr.stGopAttr.enGopMode  					= VENC_GOPMODE_DUALP;
            stVencChnAttr.stGopAttr.stDualP.u32SPInterval		= in->s_p_interval;         //mpp3: [0, 1)U(1, u32Gop – 1] (u32Gop indicates the interval of the I-frames)  //mpp4: 
            stVencChnAttr.stGopAttr.stDualP.s32SPQpDelta		= in->s_pq_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            stVencChnAttr.stGopAttr.stDualP.s32IPQpDelta		= in->i_pq_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            break;
        case VENC_GOPMODE_SMARTP:
			stVencChnAttr.stGopAttr.enGopMode 					= VENC_GOPMODE_SMARTP;
			stVencChnAttr.stGopAttr.stSmartP.u32BgInterval		= in->bg_interval;          //mpp3: u32BgInterval must be greater than or equal to                          //mpp4: [u32Gop, 65536]
                                                                                            //u32Gop and must be an integral multiple of u32Gop.
            stVencChnAttr.stGopAttr.stSmartP.s32BgQpDelta		= in->bg_qp_delta;          //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            stVencChnAttr.stGopAttr.stSmartP.s32ViQpDelta		= in->vi_qp_delta;          //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            break;
        case VENC_GOPMODE_ADVSMARTP:
            #if HI_MPP == 3
                ;;; //TODO error
            #elif HI_MPP == 4
			stVencChnAttr.stGopAttr.enGopMode 					= VENC_GOPMODE_ADVSMARTP;
            stVencChnAttr.stGopAttr.stAdvSmartP.u32BgInterval	= in->bg_interval;          //mpp4: [u32Gop; 65536]                                                         
            stVencChnAttr.stGopAttr.stAdvSmartP.s32BgQpDelta	= in->bg_qp_delta;          //mpp4: [-10; 30]                                                               
            stVencChnAttr.stGopAttr.stAdvSmartP.s32ViQpDelta	= in->vi_qp_delta;          //mpp4: [-10; 30]                                                               
            #endif
            break;
        case VENC_GOPMODE_BIPREDB:
			//TODO hi3516cv300 doesn`t support it
			stVencChnAttr.stGopAttr.enGopMode 					= VENC_GOPMODE_BIPREDB;
            stVencChnAttr.stGopAttr.stBipredB.u32BFrmNum		= in->b_frm_num;            //mpp3: [1; 3]	                                                                //mpp4: [1; 3]
            stVencChnAttr.stGopAttr.stBipredB.s32BQpDelta		= in->b_qp_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            stVencChnAttr.stGopAttr.stBipredB.s32IPQpDelta		= in->i_pq_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            break;
        case VENC_GOPMODE_INTRAREFRESH:
            stVencChnAttr.stGopAttr.enGopMode                   = VENC_GOPMODE_NORMALP;
            stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta      = -1;//in->i_pq_delta;           //mpp3: [-10; 30]                                                               //mpp4: [-10; 30]
            /*
			#if HI_MPP == 3
            //typedef struct hiVENC_PARAM_INTRA_REFRESH_S {
            //    HI_BOOL bRefreshEnable;
            //    HI_BOOL bISliceEnable;
            //    HI_U32 u32RefreshLineNum;
            //    HI_U32 u32ReqIQp;
            //}VENC_PARAM_INTRA_REFRESH_S;
            //H264: u32RefreshLineNum*MaxRefreshFrameInGop ≥ (u32PicHeight + 15)/16
            //H265: u32RefreshLineNum*MaxRefreshFrameInGop ≥ (u32PicHeight + 63)/64
			//If advanced frame skipping reference is not used: MaxRefreshFrameInGop = Gop;
            //If advanced frame skipping reference is used: MaxRefreshFrameInGop = (Gop + (u32Base*(u32Enhance+1) - 1))/(u32Base*(u32Enhance+1))
            //HI_S32 HI_MPI_VENC_SetIntraRefresh(VENC_CHN VeChn,VENC_PARAM_INTRA_REFRESH_S *pstIntraRefresh)
            
            VENC_PARAM_INTRA_REFRESH_S stIntraRefresh;
            
            stIntraRefresh.bRefreshEnable = HI_TRUE;
            stIntraRefresh.bISliceEnable = HI_TRUE;
            stIntraRefresh.u32RefreshLineNum = (in->height + 15) / 16;
            stIntraRefresh.u32ReqIQp = HI_FALSE;

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SetIntraRefresh, in->id, &stIntraRefresh);
			#endif
            */
        default:
            ;;;//TODO error
    }

	//stVencChnAttr.stGopAttr.enGopMode = VENC_GOPMODE_LOWDELAYB
   
	#endif 
	///////////////////////////////////////////////////////

    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateGroup, in->id);
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateChn, in->id, &stVencChnAttr);

    if (in->gop_mode == VENC_GOPMODE_INTRAREFRESH) {
        #if HI_MPP == 3
            //typedef struct hiVENC_PARAM_INTRA_REFRESH_S {
            //    HI_BOOL bRefreshEnable;
            //    HI_BOOL bISliceEnable;
            //    HI_U32 u32RefreshLineNum;
            //    HI_U32 u32ReqIQp;
            //}VENC_PARAM_INTRA_REFRESH_S;
            //H264: u32RefreshLineNum*MaxRefreshFrameInGop ≥ (u32PicHeight + 15)/16
            //H265: u32RefreshLineNum*MaxRefreshFrameInGop ≥ (u32PicHeight + 63)/64
            //If advanced frame skipping reference is not used: MaxRefreshFrameInGop = Gop;
            //If advanced frame skipping reference is used: MaxRefreshFrameInGop = (Gop + (u32Base*(u32Enhance+1) - 1))/(u32Base*(u32Enhance+1))
            //HI_S32 HI_MPI_VENC_SetIntraRefresh(VENC_CHN VeChn,VENC_PARAM_INTRA_REFRESH_S *pstIntraRefresh)
            
            VENC_PARAM_INTRA_REFRESH_S stIntraRefresh;
            
            stIntraRefresh.bRefreshEnable = HI_TRUE;
            stIntraRefresh.bISliceEnable = HI_TRUE;
            stIntraRefresh.u32RefreshLineNum = (in->height) / in->gop;
            stIntraRefresh.u32ReqIQp = HI_FALSE;

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SetIntraRefresh, in->id, &stIntraRefresh);
            printf("HI_MPI_VENC_SetIntraRefresh ok\n");
        #elif HI_MPP == 4
            VENC_INTRA_REFRESH_S stIntraRefresh;

            stIntraRefresh.bRefreshEnable = HI_TRUE;
            //stIntraRefresh.enInraRefreshMode = INTRA_REFRESH_COLUMN;
            stIntraRefresh.enIntraRefreshMode = INTRA_REFRESH_ROW;
            stIntraRefresh.u32RefreshNum = (in->height) / in->gop;
            stIntraRefresh.u32ReqIQp = HI_FALSE;

			DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SetIntraRefresh, in->id, &stIntraRefresh);
            printf("HI_MPI_VENC_SetIntraRefresh ok\n");
		#endif
    }

    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_RegisterChn, in->id, in->id);
        //s32Ret = HI_MPI_VENC_RegisterChn(VencGrp, VencChn);
    #endif

    //#if HI_MPP == 1 \
    //    || HI_MPP == 2 \
    //    || HI_MPP == 3
    //    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvPic, in->id);
    //#elif HI_MPP == 4
    //    VENC_RECV_PIC_PARAM_S  stRecvParam;
    //    stRecvParam.s32RecvPicNum = -1;
    //
    //    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvFrame, in->id, &stRecvParam);
    //#endif

    return ERR_NONE;
}

int mpp_venc_update_encoder(error_in *err, mpp_venc_create_encoder_in *in) {
    VENC_CHN_ATTR_S stAttr;

    //HI_S32 HI_MPI_VENC_GetChnAttr(VENC_CHN VeChn, VENC_CHN_ATTR_S*pstAttr);
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_GetChnAttr, in->id, &stAttr);

    //TODO CHANGE

    //HI_S32 HI_MPI_VENC_SetChnAttr(VENC_CHN VeChn, const VENC_CHN_ATTR_S* pstAttr);
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SetChnAttr, in->id, &stAttr);
    
    return ERR_NONE;
}

int mpp_venc_start_encoder(error_in *err, mpp_venc_start_encoder_in *in) { 
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvPic, in->id);
    #elif HI_MPP == 4
        VENC_RECV_PIC_PARAM_S  stRecvParam;
        stRecvParam.s32RecvPicNum = -1;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvFrame, in->id, &stRecvParam);
    #endif

    return ERR_NONE; 
}

int mpp_venc_stop_encoder(error_in *err, mpp_venc_stop_encoder_in *in) { 
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StopRecvPic, in->id);
    #elif HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StopRecvFrame, in->id);
    #endif
    
    return ERR_NONE; 
}

//int mpp_venc_update_encoder(error_in *err, mpp_venc_create_encoder_in *in) { return ERR_NONE; }

int mpp_venc_destroy_encoder(error_in *err, mpp_venc_destroy_encoder_in *in) { 
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_DestroyChn, in->id);

    return ERR_NONE; 
}

#else
int mpp_venc_create_encoder(error_in *err, mpp_venc_create_encoder_in *in) { return ERR_NONE; }
int mpp_venc_update_encoder(error_in *err, mpp_venc_create_encoder_in *in) { return ERR_NONE; }
int mpp_venc_destroy_encoder(error_in *err, mpp_venc_destroy_encoder_in *in) { return ERR_NONE; }
#endif

int mpp_send_frame_to_encoder(error_in *err, mpp_send_frame_to_encoder_in *in, void *frame) {
    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, in->id, frame);
    #elif HI_MPP >= 2
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, in->id, frame, -1);
    #endif

    return ERR_NONE;
}

int mpp_venc_request_idr(error_in *err, mpp_venc_request_idr_in *in) {
    #if HI_MPP == 1
        //Not implemented
    #elif   HI_MPP == 2 \
            || HI_MPP == 3 \
            || HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_RequestIDR, in->id, HI_FALSE);
    #endif

    return ERR_NONE;
}

int mpp_venc_reset(error_in *err, mpp_venc_reset_in *in) {
    #if HI_MPP == 1
        //Not implemented
    #elif   HI_MPP == 2 \
            || HI_MPP == 3 \
            || HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_ResetChn, in->id);
    #endif

    return ERR_NONE;
}
