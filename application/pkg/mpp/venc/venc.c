#if 1
#include "venc.h"

#include <string.h>

int mpp_venc_mjpeg_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_in *in) {
    switch (in->bitrate_control) {
        case VENC_RC_MODE_MJPEGCBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;

            #if HI_MPP == 1
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime          = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32ViFrmRate         = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.fr32TargetFrmRate    = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel    = in->fluctuate_level;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate           = in->bitrate;
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime          = in->stat_time;  
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32SrcFrmRate        = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.fr32DstFrmRate       = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel    = in->fluctuate_level;
                stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate           = in->bitrate;
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
                stVencChnAttr.stRcAttr.stMjpegeCbr.u32StatTime              = in->stat_time;
                stVencChnAttr.stRcAttr.stMjpegeCbr.u32SrcFrameRate          = in->in_fps;
                stVencChnAttr.stRcAttr.stMjpegeCbr.fr32DstFrameRate         = in->out_fps;
                stVencChnAttr.stRcAttr.stMjpegeCbr.u32BitRate               = in->bitrate;
            #endif
            break;
        case VENC_RC_MODE_MJPEGVBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGVBR;

            #if HI_MPP == 1
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGVBR;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32StatTime          = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32ViFrmRate         = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.fr32TargetFrmRate    = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MinQfactor        = in->min_q_factor;//50;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxQfactor        = in->max_q_factor;//95;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxBitRate        = in->bitrate;//TODO
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGVBR;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32StatTime          = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32SrcFrmRate        = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.fr32DstFrmRate       = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MinQfactor        = in->min_q_factor;//50;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxQfactor        = in->max_q_factor;//95;
                stVencChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxBitRate        = in->bitrate;//TODO 
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGVBR;
                stVencChnAttr.stRcAttr.stMjpegVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stMjpegVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stMjpegVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stMjpegVbr.u32MaxBitRate             = in->bitrate;//TODO
            #endif
            break;
        case VENC_RC_MODE_MJPEGFIXQP:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGFIXQP;

            #if HI_MPP == 1
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGFIXQP;
                stVencChnAttr.stRcAttr.stAttrMjpegeFixQp.u32Qfactor         = in->q_factor;
                stVencChnAttr.stRcAttr.stAttrMjpegeFixQp.u32ViFrmRate       = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrMjpegeFixQp.fr32TargetFrmRate  = in->out_fps;
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGFIXQP;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.u32Qfactor             = in->q_factor;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.u32SrcFrmRate          = in->in_fps;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.fr32DstFrmRate         = in->out_fps;    
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGFIXQP;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.u32Qfactor             = in->q_factor;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.u32SrcFrameRate        = in->in_fps;
                stVencChnAttr.stRcAttr.stMjpegeFixQp.fr32DstFrameRate       = in->out_fps;
            #endif
            break;
        default:
            ;;;//TODO error
    }

    return ERR_NONE;
}

int mpp_venc_h264_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_in *in) {
    switch (in->bitrate_control) {
        case VENC_RC_MODE_H264CBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;

            #if HI_MPP == 1
                stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBRv2;   //!!!
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32Gop                 = in->gop;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32StatTime            = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32ViFrmRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.fr32TargetFrmRate      = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32BitRate             = in->bitrate;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32FluctuateLevel      = in->fluctuate_level;
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32Gop                 = in->gop;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32StatTime            = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32SrcFrmRate          = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.fr32DstFrmRate         = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32BitRate             = in->bitrate;
                stVencChnAttr.stRcAttr.stAttrH264Cbr.u32FluctuateLevel      = in->fluctuate_level;
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
                stVencChnAttr.stRcAttr.stH264Cbr.u32Gop                     = in->gop;
                stVencChnAttr.stRcAttr.stH264Cbr.u32StatTime                = in->stat_time;
                stVencChnAttr.stRcAttr.stH264Cbr.u32SrcFrameRate            = in->in_fps;
                stVencChnAttr.stRcAttr.stH264Cbr.fr32DstFrameRate           = in->out_fps;
                stVencChnAttr.stRcAttr.stH264Cbr.u32BitRate                 = in->bitrate;
            #endif
            break;
        case VENC_RC_MODE_H264VBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264VBR;

            #if HI_MPP == 1
                stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264VBRv2;   //!!!
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32Gop                 = in->gop;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32StatTime            = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32ViFrmRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.fr32TargetFrmRate      = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MinQp               = in->min_qp;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MaxQp               = in->max_qp;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MaxBitRate          = in->bitrate;//TODO
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264VBR;  
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32Gop                 = in->gop;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32StatTime            = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32SrcFrmRate          = in->in_fps; 
                stVencChnAttr.stRcAttr.stAttrH264Vbr.fr32DstFrmRate         = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MinQp               = in->min_qp;
                #if HI_MPP == 3
                    stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MinIQp              = in->min_i_qp;
                #endif
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MaxQp               = in->max_qp;
                stVencChnAttr.stRcAttr.stAttrH264Vbr.u32MaxBitRate          = in->bitrate;//TODO
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264VBR;  
                stVencChnAttr.stRcAttr.stH264Vbr.u32Gop                     = in->gop;
                stVencChnAttr.stRcAttr.stH264Vbr.u32StatTime                = in->stat_time;
                stVencChnAttr.stRcAttr.stH264Vbr.u32SrcFrameRate            = in->in_fps; 
                stVencChnAttr.stRcAttr.stH264Vbr.fr32DstFrameRate           = in->out_fps;
                stVencChnAttr.stRcAttr.stH264Vbr.u32MaxBitRate              = in->bitrate;//TODO
            #endif
            break;
        case VENC_RC_MODE_H264FIXQP:

    		#if HI_MPP == 1
            	stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264FIXQP;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32Gop               = in->gop;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32ViFrmRate         = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.fr32TargetFrmRate    = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32IQp               = in->i_qp;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32PQp               = in->p_qp;
          	#elif HI_MPP == 2 || \
				  HI_MPP == 3
    			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264FIXQP;
        		stVencChnAttr.stRcAttr.stAttrH264FixQp.u32Gop               = in->gop;
            	stVencChnAttr.stRcAttr.stAttrH264FixQp.u32SrcFrmRate        = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.fr32DstFrmRate       = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32IQp               = in->i_qp;
                stVencChnAttr.stRcAttr.stAttrH264FixQp.u32PQp               = in->p_qp;
				#if HI_MPP == 3
                	stVencChnAttr.stRcAttr.stAttrH264FixQp.u32BQp               = in->b_qp;
				#endif
			#elif HI_MPP == 4
    			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264FIXQP;
    			stVencChnAttr.stRcAttr.stH264FixQp.u32Gop                   = in->gop;
    			stVencChnAttr.stRcAttr.stH264FixQp.u32SrcFrameRate          = in->in_fps;
    			stVencChnAttr.stRcAttr.stH264FixQp.fr32DstFrameRate         = in->out_fps;
    			stVencChnAttr.stRcAttr.stH264FixQp.u32IQp                   = in->i_qp;
    			stVencChnAttr.stRcAttr.stH264FixQp.u32PQp                   = in->p_qp;
    			stVencChnAttr.stRcAttr.stH264FixQp.u32BQp                   = in->b_qp;
    		#endif
            break;
		case VENC_RC_MODE_H264AVBR:
			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264AVBR;

			#if HI_MPP == 1
				;;;//TODO error
			#elif HI_MPP == 2 || \\
				  HI_MPP == 3
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264AVBR;                            
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32Gop                = in->gop;                            
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32StatTime           = in->stat_time;                                
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32SrcFrmRate         = in->in_fps;                                    
                stVencChnAttr.stRcAttr.stAttrH264AVbr.fr32DstFrmRate        = in->out_fps;                                        
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32MaxBitRate         = in->bitrate;//TODO    
    		#elif HI_MPP == 4
            	//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264AVBR;
                stVencChnAttr.stRcAttr.stH264AVbr.u32Gop                    = in->gop;
                stVencChnAttr.stRcAttr.stH264AVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stH264AVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH264AVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stH264AVbr.u32MaxBitRate             = in->bitrate;//TODO
            #endif
			break;
		case VENC_RC_MODE_H264CVBR
			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CVBR;

			#if HI_MPP == 1 || \
				HI_MPP == 2 || \
				HI_MPP == 3 
				;;;//TODO error
			#elif HI_MPP == 4
            	//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CVBR;
            	stVencChnAttr.stRcAttr.stH264CVbr.u32Gop                    = in->gop;
                stVencChnAttr.stRcAttr.stH264CVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stH264CVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH264CVbr.fr32DstFrameRate          = in_out_fps;
                stVencChnAttr.stRcAttr.stH264CVbr.u32LongTermStatTime       = 1;
                stVencChnAttr.stRcAttr.stH264CVbr.u32ShortTermStatTime      = xxx;
                stVencChnAttr.stRcAttr.stH264CVbr.u32MaxBitRate             = in->bitrate;//TODO
                stVencChnAttr.stRcAttr.stH264CVbr.u32LongTermMaxBitrate     = 1024 + 512*u32FrameRate/30;
                stVencChnAttr.stRcAttr.stH264CVbr.u32LongTermMinBitrate     = 256;
         	#endif
			break;
		case VENC_RC_MODE_H264QVBR:
            #if HI_MPP == 1 || \
                HI_MPP == 2 || \
                HI_MPP == 3 
            	;;;//TODO error
          	#elif HI_MPP == 4
	  			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264QVBR;
            	stVencChnAttr.stRcAttr.stH264QVbr.u32Gop                    = in->gop;
                stVencChnAttr.stRcAttr.stH264QVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stH264QVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH264QVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stH264QVbr.u32TargetBitRate          = in->bitrate;
          	#endif
			break;
        default:
            ;;;//TODO error
    }

	#if HI_MPP == 4
    	stVencChnAttr.stVencAttr.stAttrH264e.bRcnRefShareBuf        = bRcnRefShareBuf; //only mpp4 or even hi3516cv500 //TODO!!!
    #endif

    return ERR_NONE;
}

int mpp_venc_h265_params(VENC_CHN_ATTR_S *stVencChnAttr, mpp_venc_create_in *in) {
    switch (in->bitrate_control) {
		case VENC_RC_MODE_H265CBR:
			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;

            #if HI_MPP == 2 || \
                HI_MPP == 3
            	//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;                           
                stVencChnAttr.stRcAttr.stAttrH265Cbr.u32Gop                 = in->gop;                            
                stVencChnAttr.stRcAttr.stAttrH265Cbr.u32StatTime            = in->stat_time;                                
                stVencChnAttr.stRcAttr.stAttrH265Cbr.u32SrcFrmRate          = in->in_fps;                                    
                stVencChnAttr.stRcAttr.stAttrH265Cbr.fr32DstFrmRate         = in->out_fps;                                        
                stVencChnAttr.stRcAttr.stAttrH265Cbr.u32BitRate             = in->bitrate;                                            
                stVencChnAttr.stRcAttr.stAttrH265Cbr.u32FluctuateLevel      = in->fluctuate_level;   
			#elif HI_MPP == 4
            	//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;
                stVencChnAttr.stRcAttr.stH265Cbr.u32Gop                     = in->gop;
                stVencChnAttr.stRcAttr.stH265Cbr.u32StatTime                = in->stat_time;
                stVencChnAttr.stRcAttr.stH265Cbr.u32SrcFrameRate            = in->in_fps;
                stVencChnAttr.stRcAttr.stH265Cbr.fr32DstFrameRate           = in->out_fps;
                stVencChnAttr.stRcAttr.stH265Cbr.u32BitRate                 = in->bitrate;
          	#endif
			break;
		case VENC_RC_MODE_H265VBR:
			stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265VBR;

			#if HI_MPP == 1
				;;;//TODO error
			#elif HI_MPP == 2 || \
				  HI_MPP == 3 
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265VBR;
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32Gop                 = in->gop;
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32StatTime            = in->stat_time;
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32SrcFrmRate          = in->in_fps;
                stVencChnAttr.stRcAttr.stAttrH265Vbr.fr32DstFrmRate         = in->out_fps;
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32MinQp               = in->min_qp;
				#if HI_MPP == 3
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32MinIQp              = in->min_i_qp;
				#endif
                stVencChnAttr.stRcAttr.stAttrH265Vbr.u32MaxQp               = in->max_qp;
        		stVencChnAttr.stRcAttr.stAttrH265Vbr.u32MaxBitRate          = in->bitrate;//TODO
    		#elif HI_MPP == 4 
            	//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265VBR;
            	stVencChnAttr.stRcAttr.stH265Vbr.u32Gop                     = in->gop;
                stVencChnAttr.stRcAttr.stH265Vbr.u32StatTime                = in->stat_time;
                stVencChnAttr.stRcAttr.stH265Vbr.u32SrcFrameRate            = in->in_fps;
                stVencChnAttr.stRcAttr.stH265Vbr.fr32DstFrameRate           = in->out_fps;
                stVencChnAttr.stRcAttr.stH265Vbr.u32MaxBitRate              = in->bitrate;//TODO
          	#endif

		case VENC_RC_MODE_H265FIXQP:
            #if HI_MPP == 1
            	;;;//TODO error
            #elif HI_MPP == 2 || \
                  HI_MPP == 3 
                stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265FIXQP;                    
                stVencChnAttr.stRcAttr.stAttrH265FixQp.u32Gop               = in->gop;                        
                stVencChnAttr.stRcAttr.stAttrH265FixQp.u32SrcFrmRate        = in->in_fps;                            
                stVencChnAttr.stRcAttr.stAttrH265FixQp.fr32DstFrmRate       = in->out_fps;                                
                stVencChnAttr.stRcAttr.stAttrH265FixQp.u32IQp               = in->i_qp;                                    
                stVencChnAttr.stRcAttr.stAttrH265FixQp.u32PQp               = in->p_qp;                                        
				#if HI_MPP == 3
					stVencChnAttr.stRcAttr.stAttrH265FixQp.u32BQp               = in->b_qp;   
				#endif
			#elif HI_MPP == 4
            	stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265FIXQP;
            	stVencChnAttr.stRcAttr.stH265FixQp.u32Gop                   = in->gop;
                stVencChnAttr.stRcAttr.stH265FixQp.u32SrcFrameRate          = in->in_fps;
                stVencChnAttr.stRcAttr.stH265FixQp.fr32DstFrameRate         = in->out_fps;
                stVencChnAttr.stRcAttr.stH265FixQp.u32IQp                   = in->i_qp;
                stVencChnAttr.stRcAttr.stH265FixQp.u32PQp                   = in->p_qp;
                stVencChnAttr.stRcAttr.stH265FixQp.u32BQp                   = in->b_qp;
           	#endif
		case VENC_RC_MODE_H265AVBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265AVBR;

            #if HI_MPP == 1
            	;;;//TODO error
            #elif HI_MPP == 2 || \
            	  HI_MPP == 3 
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265AVBR;        
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32Gop                = in->gop;            
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32StatTime           = in->stat_time;                
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32SrcFrmRate         = in->in_fps;                    
                stVencChnAttr.stRcAttr.stAttrH264AVbr.fr32DstFrmRate        = in->out_fps;                        
                stVencChnAttr.stRcAttr.stAttrH264AVbr.u32MaxBitRate         = in->bitrate;//TODO   <Paste>
    		#elif HI_MPP == 4
        		//stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265AVBR;
    			stVencChnAttr.stRcAttr.stH265AVbr.u32Gop                    = in->gop;
        		stVencChnAttr.stRcAttr.stH265AVbr.u32StatTime               = in->stat_time;
            	stVencChnAttr.stRcAttr.stH265AVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH265AVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stH265AVbr.u32MaxBitRate             = in->bitrate;
         	#endif

        case VENC_RC_MODE_H265CVBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CVBR;

            #if HI_MPP == 1 || \
                HI_MPP == 2 || \
                HI_MPP == 3 
                ;;;//TODO error
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CVBR;
                stVencChnAttr.stRcAttr.stH265CVbr.u32Gop                    = in->gop;
                stVencChnAttr.stRcAttr.stH265CVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stH265CVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH265CVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stH265CVbr.u32LongTermStatTime       = 1;
                stVencChnAttr.stRcAttr.stH265CVbr.u32ShortTermStatTime      = xxxx;
                stVencChnAttr.stRcAttr.stH265CVbr.u32MaxBitRate             = in->bitrate;//TODO
                stVencChnAttr.stRcAttr.stH265CVbr.u32LongTermMaxBitrate     = 1024  + 512*u32FrameRate/30;
                stVencChnAttr.stRcAttr.stH265CVbr.u32LongTermMinBitrate     = 256;
            #endif
            break;
        case VENC_RC_MODE_H265QVBR:
            stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265QVBR;

            #if HI_MPP == 1 || \
                HI_MPP == 2 || \
                HI_MPP == 3 
                ;;;//TODO error
            #elif HI_MPP == 4
                //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265QVBR;
                stVencChnAttr.stRcAttr.stH265QVbr.u32Gop                    = in->gop;
                stVencChnAttr.stRcAttr.stH265QVbr.u32StatTime               = in->stat_time;
                stVencChnAttr.stRcAttr.stH265QVbr.u32SrcFrameRate           = in->in_fps;
                stVencChnAttr.stRcAttr.stH265QVbr.fr32DstFrameRate          = in->out_fps;
                stVencChnAttr.stRcAttr.stH265QVbr.u32TargetBitRate          = in->bitrate;
            #endif
            break;
		default:
			;;;//TODO error   
    }   
    
    #if HI_MPP == 4
        stVencChnAttr.stVencAttr.stAttrH265e.bRcnRefShareBuf        = bRcnRefShareBuf;//TODO
   	#endif

    return ERR_NONE;
}


int mpp_venc_create(error_in *err, mpp_venc_create_in *in) {
    VENC_CHN_ATTR_S stVencChnAttr;

    #if HI_MPP == 4
        stVencChnAttr.stVencAttr.u32MaxPicWidth                     = in->width;
        stVencChnAttr.stVencAttr.u32MaxPicHeight                    = in->height;
        stVencChnAttr.stVencAttr.u32PicWidth                        = in->width;
        stVencChnAttr.stVencAttr.u32PicHeight                       = in->height;
        stVencChnAttr.stVencAttr.u32BufSize                         = "?";
        stVencChnAttr.stVencAttr.u32Profile                         = in->profile;
        stVencChnAttr.stVencAttr.bByFrame                           = HI_TRUE;
    #endif

    switch (in->codec) {
        case PT_MJPEG:
            stVencChnAttr.stVeAttr.enType   = PT_MJPEG;

            #if HI_MPP == 1
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32BufSize               = "?";
                stVencChnAttr.stVeAttr.stAttrMjpeg.bByFrame                 = HI_TRUE;
                stVencChnAttr.stVeAttr.stAttrMjpeg.bMainStream              = HI_TRUE;//TODO
                stVencChnAttr.stVeAttr.stAttrMjpeg.bVIField                 = HI_FALSE;//TODO
                stVencChnAttr.stVeAttr.stAttrMjpeg.u32Priority              = 0;//TODO
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr.stVeAttr.stAttrMjpege.u32MaxPicWidth          = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32MaxPicHeight         = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32PicWidth             = in->width;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32PicHeight            = in->height;
                stVencChnAttr.stVeAttr.stAttrMjpege.u32BufSize              = "?";
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
                stVencChnAttr.stVeAttr.stAttrH264e.u32BufSize               = "?";
                stVencChnAttr.stVeAttr.stAttrH264e.u32Profile               = in->profile;
                stVencChnAttr.stVeAttr.stAttrH264e.bByFrame                 = HI_TRUE;
                stVencChnAttr.stVeAttr.stAttrH264e.bField                   = HI_FALSE;
                stVencChnAttr.stVeAttr.stAttrH264e.bMainStream              = HI_TRUE;//TODO
                stVencChnAttr.stVeAttr.stAttrH264e.u32Priority              = 0;//TODO
                stVencChnAttr.stVeAttr.stAttrH264e.bVIField                 = HI_FALSE;//TODO
            #elif HI_MPP == 2 || \
                  HI_MPP == 3
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicWidth           = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32MaxPicHeight          = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicWidth              = in->width;
                stVencChnAttr.stVeAttr.stAttrH264e.u32PicHeight             = in->height;
                stVencChnAttr.stVeAttr.stAttrH264e.u32BufSize               = "?";    
                stVencChnAttr.stVeAttr.stAttrH264e.u32Profile               = in->profile;  
                stVencChnAttr.stVeAttr.stAttrH264e.bByFrame                 = HI_TRUE;   
                #if HI_MPP == 2
                    stVencChnAttr.stVeAttr.stAttrH264e.u32BFrameNum             = 0;
                    stVencChnAttr.stVeAttr.stAttrH264e.u32RefNum                = 1;
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
                stVencChnAttr.stVeAttr.stAttrH265e.u32BufSize               = "?";
                stVencChnAttr.stVeAttr.stAttrH265e.u32Profile               = in->profile;
                stVencChnAttr.stVeAttr.stAttrH265e.bByFrame                 = HI_TRUE;
                #if HI_MPP == 2
                    stVencChnAttr.stVeAttr.stAttrH265e.u32BFrameNum             = 0;
                    stVencChnAttr.stVeAttr.stAttrH265e.u32RefNum                = 1;
                #endif
            #endif

            break;
        default:
            ;;;//TODO error
    }

    #if HI_MPP == 4
        stVencChnAttr.stVencAttr.u32MaxPicWidth                     = in->width;
        stVencChnAttr.stVencAttr.u32MaxPicHeight                    = in->height;
        stVencChnAttr.stVencAttr.u32PicWidth                        = in->width;
        stVencChnAttr.stVencAttr.u32PicHeight                       = in->height;
        stVencChnAttr.stVencAttr.u32BufSize                         = "?";
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
        default:
            ;;;//TODO error
    }

    //////////////////

    switch (in->gop_mode) {
        case VENC_GOPMODE_NORMALP:
            break;
        case VENC_GOPMODE_DUALP:
            break;
        case VENC_GOPMODE_SMARTP:
            break;
        case VENC_GOPMODE_ADVSMARTP:
            break;
        case VENC_GOPMODE_BIPREDB:
            break;
        default:
            ;;;//TODO error
    }

    #if HI_MPP = 4 
	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
	stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta          = 0;
    #endif
    #if HI_MPP == 3
    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 2;
    #endif

	//-------------------------

    #if HI_MPP = 4
	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_DUALP;
	stVencChnAttr.stGopAttr.stDualP.u32SPInterval;
    stVencChnAttr.stGopAttr.stDualP.s32SPQpDelta;
    stVencChnAttr.stGopAttr.stDualP.s32IPQpDelta;
    #endif
    #if HI_MPP == 3
    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_DUALP;
    stVencChnAttr.stGopAttr.stDualP.s32IPQpDelta  = 4;
    stVencChnAttr.stGopAttr.stDualP.s32SPQpDelta  = 2;
    stVencChnAttr.stGopAttr.stDualP.u32SPInterval = 3;
    #endif

	//-----------------------
    
    #if HI_MPP = 4
	stVencChnAttr.stGopAttr.enGopMode = VENC_GOPMODE_SMARTP;
    stVencChnAttr.stGopAttr.stSmartP.u32BgInterval;
    stVencChnAttr.stGopAttr.stSmartP.s32BgQpDelta;
    stVencChnAttr.stGopAttr.stSmartP.s32ViQpDelta;
    #endif
    #if HI_MPP == 3
    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_SMARTP;
    stVencChnAttr.stGopAttr.stSmartP.s32BgQpDelta = 4;
    stVencChnAttr.stGopAttr.stSmartP.s32ViQpDelta = 2;
    stVencChnAttr.stGopAttr.stSmartP.u32BgInterval = (VIDEO_ENCODING_MODE_PAL == enNorm) ? 75 : 90;

    #endif
	//-----------------------

    #if HI_MPP = 4
	stVencChnAttr.stGopAttr.enGopMode = VENC_GOPMODE_ADVSMARTP;
	stVencChnAttr.stGopAttr.stAdvSmartP.u32BgInterval;
	stVencChnAttr.stGopAttr.stAdvSmartP.s32BgQpDelta;
	stVencChnAttr.stGopAttr.stAdvSmartP.s32ViQpDelta;
    #endif
    

	//-----------------------

    #if HI_MPP = 4
	stVencChnAttr.stGopAttr.enGopMode = VENC_GOPMODE_BIPREDB;
	stVencChnAttr.stGopAttr.stBipredB.u32BFrmNum;
    stVencChnAttr.stGopAttr.stBipredB.s32BQpDelta;
    stVencChnAttr.stGopAttr.stBipredB.s32IPQpDelta;
    #endif

	//-----------------------

	//stVencChnAttr.stGopAttr.enGopMode = VENC_GOPMODE_LOWDELAYB

    
	///////////////////////////////////////////////////////

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateChn, in->venc_id, &stVencChnAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvPic, in->venc_id);

    return ERR_NONE;
}
#endif
