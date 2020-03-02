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


#endif //_LOOP_H
