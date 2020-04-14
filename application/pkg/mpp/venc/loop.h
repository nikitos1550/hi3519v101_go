#ifndef _LOOP_H
#define _LOOP_H

//forward declarations
int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id);
int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id);
void * mpp_data_loop_thread();

//external funcs
void mpp_data_loop_get_data(unsigned int venc_channel);
int mpp_venc_getfd(int venc_channel_id);

typedef struct st_data_from_c {
    char            *data;
    int             length;
} data_from_c;

void go_callback_receive_data(int venc, unsigned int seq, data_from_c * data_pointer, int num);

#define ERR_NONE                0
#define ERR_MPP                 2

int mpp3_venc_sample_mjpeg(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h264(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h265(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_delete_encoder(unsigned int *error_code, int channelId);

#endif //_LOOP_H

