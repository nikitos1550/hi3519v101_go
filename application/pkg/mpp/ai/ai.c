#include "ai.h"

#if HI_MPP == 2 || HI_MPP == 3

#define SAMPLES_PER_FRAME   960

#define OPUS
//#define AAC

#if defined(OPUS)
    #include <opus.h>
    static OpusEncoder *opus_enc;
#elif defined(AAC)
    #include "aacenc_lib.h"
    static HANDLE_AACENCODER *fdkaac_enc;
#endif

//void test_opus_init() {
//    int error;
//    //OpusEncoder *enc;
//    //enc = opus_encoder_create(Fs, channels, application, &error);
//    opus_enc = opus_encoder_create(48000, 1, OPUS_APPLICATION_AUDIO, &error);
//    if (error <0 ) {
//      printf("failed to create opus encoder: %s\n", opus_strerror(error));
//      //return EXIT_FAILURE;
//    }
//
//    error = opus_encoder_ctl(opus_enc, OPUS_SET_BITRATE(256000));
//    if (error<0) {
//      printf("failed to set bitrate: %s\n", opus_strerror(error));
//      //return EXIT_FAILURE;
//    }
//}

static pthread_t mpp_ai_thread_pid;

void* mpp_ai_thread(HI_VOID *param){   //now we start it from go space

    #if defined(OPUS)
        int error_opus;
        opus_enc = opus_encoder_create(48000, 1, OPUS_APPLICATION_AUDIO, &error_opus);
        if (error_opus < 0 ) {
            printf("failed to create opus encoder: %s\n", opus_strerror(error_opus));
        }

        error_opus = opus_encoder_ctl(opus_enc, OPUS_SET_BITRATE(256000));
        if (error_opus < 0 ) {
            printf("failed to set bitrate: %s\n", opus_strerror(error_opus));
        }
    #elif defined(AAC)
        *fdkaac_enc = 0;
        if (aacEncOpen(fdkaac_enc, 0, 0) != AACENC_OK) {
            printf("ERROR: aacEncOpen() failed\n");
        }

        if (aacEncEncode(*fdkaac_enc, 0, 0, 0, 0) != AACENC_OK) {
            printf("ERROR: encoder initialization failed\n");
            //goto FAIL;
        }
    #endif

    //GO_LOG_AI(LOGGER_TRACE, "AI thread run");
    printf("AI thread run\n");

    int fd = 0;
    unsigned int error = 0;
        
    AUDIO_FRAME_S stFrame;
    AEC_FRAME_S   stAecFrm;

    //#if defined(OPUS)
    //    FILE *fout;
    //    fout = fopen("/opt/nfs/test.opus", "w");
    //#elif defined(AAC)
    //    FILE *fout;
    //    fout = fopen("/opt/nfs/test.aac", "w");
    //#endif

#if 1
    fd = HI_MPI_AI_GetFd(0, 0);
    if (fd < 0) {
        printf("HI_MPI_AI_GetFd error %d\n", fd);
        return NULL;
    }

    memset(&stAecFrm, 0, sizeof(AEC_FRAME_S));

    while (1) {
        error = HI_MPI_AI_GetFrame(0, 0, &stFrame, &stAecFrm, -1);
        if (error != 0) {
            printf("HI_MPI_AI_GetFrame error %u\n", error);
            return NULL;
        }

        audio_info_from_c info;
        audio_data_from_c data;

        data.data   = stFrame.pVirAddr[0];
        data.length = stFrame.u32Len;

        info.timestamp = stFrame.u64TimeStamp;

        go_callback_raw_tmp(&info, &data);

        //printf("new audio frame ts %llu seq %u length %u vir %p phy %u\n", stFrame.u64TimeStamp, stFrame.u32Seq, stFrame.u32Len, stFrame.pVirAddr[0], stFrame.u32PhyAddr[0]);
        
    #if defined(OPUS)

        #define MAX_PACKET_SIZE  (4000)

        unsigned char cbits[MAX_PACKET_SIZE];
        int nbBytes;
        //printf("sizeof(opus_int16) = %d\n", sizeof(opus_int16));
        //opus_int16 data[SAMPLES_PER_FRAME];
    
        #if HI_MPP == 2
            opus_int16 *raw_data = stFrame.pVirAddr[0];
        #elif HI_MPP == 3
            opus_int16 *raw_data = stFrame.pVirAddr[0];
        #elif HI_MPP == 4
            //TODO not working
            opus_int16 *raw_data = stFrame.u64VirAddr[0];
        #endif

        nbBytes = opus_encode(opus_enc, raw_data, SAMPLES_PER_FRAME, cbits, MAX_PACKET_SIZE);
        if (nbBytes<0) {
            printf("encode failed: %s\n", opus_strerror(nbBytes));
            //return EXIT_FAILURE;
        } else {
            //printf("encoded %d bytes\n", nbBytes);
            //fwrite(cbits, sizeof(unsigned char), nbBytes, fout);
            //fflush(fout);
            
            data.data   = cbits;
            data.length = nbBytes;

            go_callback_opus_tmp(&info, &data);

        }
    #elif defined(AAC)
        //TODO
    #endif

        /*    
        HI_U8 *pUserPageAddrV = stFrame.pVirAddr[0];

        unsigned int i;
        for (i=0;i<(stFrame.u32Len/2);i++) {
            short *tv = &pUserPageAddrV[i*4];
            printf("i=%d Virt: %d\n", i, *tv);
        } 
        */

        //s32Ret = HI_MPI_AENC_SendFrame(pstAiCtl->AencChn, &stFrame, &stAecFrm);

        error = HI_MPI_AO_SendFrame(0, 0, &stFrame, 1000);

        error = HI_MPI_AI_ReleaseFrame(0, 0, &stFrame, &stAecFrm);
        if (error != 0) {
            printf("HI_MPI_AI_GetFrame error %u\n", error);
            return NULL;
        } 
    }
#endif
      
    //GO_LOG_AI(LOGGER_ERROR, "AI thread failed");
    printf("AI thread failed\n");

    return NULL;   
}


#define SAMPLE_AUDIO_PTNUMPERFRM   320

int mpp_ai_test(error_in *err) {

    AIO_ATTR_S stAioAttr;


    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_8000;
    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_16000;
    stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_48000;
    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_96000;

    stAioAttr.enBitwidth     = AUDIO_BIT_WIDTH_16;
    //stAioAttr.enBitwidth     = AUDIO_BIT_WIDTH_8; //not working
    stAioAttr.enWorkmode     = AIO_MODE_I2S_MASTER;
    stAioAttr.enSoundmode    = AUDIO_SOUND_MODE_MONO;
    stAioAttr.u32EXFlag      = 0;
    stAioAttr.u32FrmNum      = MAX_AUDIO_FRAME_NUM;//30;

    //Number of sampling points for one frame
    //Value range: In G711, G726, and ADPCM_DVI4, the value is 80,
    //160, 240, 320, or 480; in ADPCM_IMA, the value is 81, 161, 241,
    //321, or 481.
    //The value range of the length of an AI frame is [80, 2048]. The value
    //range of the length of an AO frame is [80, 4096].

    //The value of u32PtNumPerFrm (number of sampling points in each frame) and that of the
    //enSamplerate (sampling rate) determine the frequency of interrupt generation of hardware.
    //When the frequency is too high, the system performance and other services will be affected.
    //You are advised to set values for the two parameters based on the following formula:
    //(u32PtNumPerFrm x 1000)/enSamplerate ≥ 10. For example, when the sampling rate is 16
    //kHz, you are advised to set the number of sampling points to a value that is greater than or
    //equal to 160.

    stAioAttr.u32PtNumPerFrm = SAMPLES_PER_FRAME; //800;//320;//SAMPLE_AUDIO_PTNUMPERFRM;
    stAioAttr.u32ChnCnt      = 1;
    stAioAttr.u32ClkSel      = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_SetPubAttr, 0, &stAioAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_Enable, 0);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_EnableChn, 0, 0);

    AUDIO_TRACK_MODE_E enTrackMode;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_GetTrackMode, 0, &enTrackMode);
    printf("AUDIO_TRACK_MODE_E %d\n", enTrackMode);

    #if 0
    if (HI_TRUE == bResampleEn) {
        s32Ret = HI_MPI_AI_EnableReSmp(AiDevId, i, enOutSampleRate);^M
        if (s32Ret) {
            printf("%s: HI_MPI_AI_EnableReSmp(%d,%d) failed with %#x\n", __FUNCTION__, AiDevId, i, s32Ret);
            return s32Ret;
        }
    }
    #endif

    #if 0
    if (NULL != pstAiVqeAttr) {
        HI_BOOL bAiVqe = HI_TRUE;
        switch (u32AiVqeType) {
            case 0:
                s32Ret = HI_SUCCESS;
                bAiVqe = HI_FALSE;
                break;
            case 1:
                s32Ret = HI_MPI_AI_SetTalkVqeAttr(AiDevId, i, SAMPLE_AUDIO_AO_DEV, i, (AI_TALKVQE_CONFIG_S *)pstAiVqeAttr);
                break;
            case 2:
                s32Ret = HI_MPI_AI_SetHiFiVqeAttr(AiDevId, i, (AI_HIFIVQE_CONFIG_S *)pstAiVqeAttr);
                break;
            case 3:
                s32Ret = HI_MPI_AI_SetRecordVqeAttr(AiDevId, i, (AI_RECORDVQE_CONFIG_S *)pstAiVqeAttr);
                break;
            default:
                s32Ret = HI_FAILURE;
                break;
        }
        
        if (s32Ret) {
            printf("%s: SetAiVqe%d(%d,%d) failed with %#x\n", __FUNCTION__, u32AiVqeType, AiDevId, i, s32Ret);
            return s32Ret;
        }
        
        if (bAiVqe) {
            s32Ret = HI_MPI_AI_EnableVqe(AiDevId, i);
            if (s32Ret) {
                printf("%s: HI_MPI_AI_EnableVqe(%d,%d) failed with %#x\n", __FUNCTION__, AiDevId, i, s32Ret);
                return s32Ret;
            }
        }
    }
    #endif

    /*
    MPP_CHN_S stSrcChn, stDestChn;

    stSrcChn.enModId = HI_ID_AI;
    stSrcChn.s32ChnId = 0;
    stSrcChn.s32DevId = 0;
    stDestChn.enModId = HI_ID_AO;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);
    */


    #if 1
    AI_CHN_PARAM_S stAiChnPara;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_GetChnParam, 0, 0, &stAiChnPara);

    stAiChnPara.u32UsrFrmDepth = 30;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_AI_SetChnParam, 0, 0, &stAiChnPara);
    #endif 

    DO_OR_RETURN_ERR_GENERAL(err, pthread_create, &mpp_ai_thread_pid, 0, (void* (*)(void*))mpp_ai_thread, NULL);
    DO_OR_RETURN_ERR_GENERAL(err, pthread_setname_np, mpp_ai_thread_pid, "AI");

    return ERR_NONE;
}

int mpp_ai_config_inner(error_in *err) {
    int general_error_code = 0;
    int fd;

    fd = open("/dev/acodec", O_RDWR);
    if (fd < 0) {
        RETURN_ERR_GENERAL(err, "open /dev/acodec", fd);
    }

    general_error_code = ioctl(fd, ACODEC_SOFT_RESET_CTRL);
    if (general_error_code != HI_SUCCESS) {
        close(fd);  
        RETURN_ERR_GENERAL(err, "ACODEC_SOFT_RESET_CTRL", general_error_code);
    }

    #if 0
    switch (enSample) {
        case AUDIO_SAMPLE_RATE_8000:
            i2s_fs_sel = ACODEC_FS_8000;
            break;                       
        case AUDIO_SAMPLE_RATE_16000:
            i2s_fs_sel = ACODEC_FS_16000;
            break;                         
        case AUDIO_SAMPLE_RATE_32000:
            i2s_fs_sel = ACODEC_FS_32000;
            break;                         
        case AUDIO_SAMPLE_RATE_11025:
            i2s_fs_sel = ACODEC_FS_11025;
            break;                         
        case AUDIO_SAMPLE_RATE_22050:    
            i2s_fs_sel = ACODEC_FS_22050;
            break;                         
        case AUDIO_SAMPLE_RATE_44100:
            i2s_fs_sel = ACODEC_FS_44100;
            break;                         
        case AUDIO_SAMPLE_RATE_12000:  
            i2s_fs_sel = ACODEC_FS_12000;
            break;                                         
        case AUDIO_SAMPLE_RATE_24000:  
            i2s_fs_sel = ACODEC_FS_24000;
            break;
        case AUDIO_SAMPLE_RATE_48000:  
            i2s_fs_sel = ACODEC_FS_48000;
            break;
        case AUDIO_SAMPLE_RATE_64000:  
            i2s_fs_sel = ACODEC_FS_64000;
            break;
        case AUDIO_SAMPLE_RATE_96000:  
            i2s_fs_sel = ACODEC_FS_96000;
            break;
        default:
            printf("%s: not support enSample:%d\n", __FUNCTION__, enSample);
            ret = HI_FAILURE;
            break;
    }
    #endif

    //ACODEC_FS_E i2s_fs_sel = ACODEC_FS_8000;
    //ACODEC_FS_E i2s_fs_sel = ACODEC_FS_16000;
    ACODEC_FS_E i2s_fs_sel = ACODEC_FS_48000;
    //ACODEC_FS_E i2s_fs_sel = ACODEC_FS_96000;

    general_error_code = ioctl(fd, ACODEC_SET_I2S1_FS, &i2s_fs_sel);
    if (general_error_code != HI_SUCCESS) {
        close(fd);  
        RETURN_ERR_GENERAL(err, "ACODEC_SET_I2S1_FS", general_error_code);
    }

    #if HI_MPP == 2
        #if defined(HI3516AV100)
            ACODEC_MIXER_E input_mode = ACODEC_MIXER_LINEIN;
        #elif defined(HI3516CV200)
            ACODEC_MIXER_E input_mode = ACODEC_MIXER_IN;
        #endif
    #elif HI_MPP == 3
        ACODEC_MIXER_E input_mode = ACODEC_MIXER_IN0;
        //ACODEC_MIXER_E input_mode = ACODEC_MIXER_IN1;
    #endif

    general_error_code = ioctl(fd, ACODEC_SET_MIXER_MIC, &input_mode);
    if (general_error_code != HI_SUCCESS) {
        close(fd);  
        RETURN_ERR_GENERAL(err, "ACODEC_SET_MIXER_MIC", general_error_code);
    }

    /*
    int output_vol = 6;

    general_error_code = ioctl(fd, ACODEC_SET_OUTPUT_VOL, &output_vol);
    if (general_error_code != HI_SUCCESS) {
        close(fd);  
        RETURN_ERR_GENERAL(err, "ACODEC_SET_OUTPUT_VOL", general_error_code);
    }

    general_error_code = ioctl(fd, ACODEC_GET_OUTPUT_VOL, &output_vol);
    if (general_error_code != HI_SUCCESS) {
        close(fd);  
        RETURN_ERR_GENERAL(err, "ACODEC_GET_OUTPUT_VOL", general_error_code);
    }
    printf("Output volume = %d\n", output_vol);
    */

    #if 0
    if (0) /* should be 1 when micin */ {
        /******************************************************************************************
        The input volume range is [-87, +86]. Both the analog gain and digital gain are adjusted.
        A larger value indicates higher volume.
        For example, the value 86 indicates the maximum volume of 86 dB,
        and the value -87 indicates the minimum volume (muted status).
        The volume adjustment takes effect simultaneously in the audio-left and audio-right channels.
        The recommended volume range is [+10, +56].
        Within this range, the noises are lowest because only the analog gain is adjusted,
        and the voice quality can be guaranteed.
        *******************************************************************************************/
        int iAcodecInputVol = 30;
        
        general_error_code = ioctl(fd, ACODEC_SET_INPUT_VOL, &iAcodecInputVol);  
        if (general_error_code != HI_SUCCESS) {
            close(fd);  
            RETURN_ERR_GENERAL(err, "ACODEC_SET_INPUT_VOL", general_error_code); 
        }
    }
    #endif

    close(fd);

    return ERR_NONE;
}


int mpp_ao_test(error_in *err) {

    AIO_ATTR_S stAioAttr;

    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_8000;
    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_16000;
    stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_48000;
    //stAioAttr.enSamplerate   = AUDIO_SAMPLE_RATE_96000; 

    stAioAttr.enBitwidth     = AUDIO_BIT_WIDTH_16;
    stAioAttr.enWorkmode     = AIO_MODE_I2S_MASTER;
    stAioAttr.enSoundmode    = AUDIO_SOUND_MODE_MONO;
    //stAioAttr.enSoundmode    = AUDIO_SOUND_MODE_STEREO;
    stAioAttr.u32EXFlag      = 0;
    stAioAttr.u32FrmNum      = MAX_AUDIO_FRAME_NUM; //30;
    stAioAttr.u32PtNumPerFrm = SAMPLES_PER_FRAME; //800; //320; //SAMPLE_AUDIO_PTNUMPERFRM;
    stAioAttr.u32ChnCnt      = 1;
    //stAioAttr.u32ChnCnt      = 2;
    stAioAttr.u32ClkSel      = 0;


    DO_OR_RETURN_ERR_GENERAL(err, HI_MPI_AO_SetPubAttr, 0, &stAioAttr);
    DO_OR_RETURN_ERR_GENERAL(err, HI_MPI_AO_Enable, 0);
	DO_OR_RETURN_ERR_GENERAL(err, HI_MPI_AO_EnableChn, 0, 0);

    DO_OR_RETURN_ERR_GENERAL(err, HI_MPI_AO_SetVolume, 0, 6);

    HI_BOOL bEnable;
    AUDIO_FADE_S stFade;

    DO_OR_RETURN_ERR_GENERAL(err, HI_MPI_AO_GetMute, 0, &bEnable, &stFade);
    printf("mute %d\n", bEnable);

    return ERR_NONE;
}

#endif
