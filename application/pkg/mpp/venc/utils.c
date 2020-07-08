#include "venc.h"

void invalidate_mpp_venc_create_encoder_in (mpp_venc_create_encoder_in *in) {
    in->id              = INVALID_VALUE;

    in->codec           = INVALID_VALUE;
    in->profile         = INVALID_VALUE;

    in->width           = INVALID_VALUE;
    in->height          = INVALID_VALUE;

    in->in_fps          = INVALID_VALUE;
    in->out_fps         = INVALID_VALUE;

    in->bitrate_control = INVALID_VALUE; 

    in->gop             = INVALID_VALUE;

    in->gop_mode        = INVALID_VALUE;

    in->i_pq_delta      = INVALID_VALUE;
    in->s_p_interval    = INVALID_VALUE;
    in->s_pq_delta      = INVALID_VALUE;
    in->bg_interval     = INVALID_VALUE;
    in->bg_qp_delta     = INVALID_VALUE;
    in->vi_qp_delta     = INVALID_VALUE;
    in->b_frm_num       = INVALID_VALUE;
    in->b_qp_delta      = INVALID_VALUE;

    in->bitrate         = INVALID_VALUE;

    in->stat_time       = INVALID_VALUE;
    in->fluctuate_level = INVALID_VALUE;

    in->q_factor        = INVALID_VALUE;
    in->min_q_factor    = INVALID_VALUE;
    in->max_q_factor    = INVALID_VALUE;

    in->i_qp            = INVALID_VALUE;
    in->p_qp            = INVALID_VALUE;
    in->b_qp            = INVALID_VALUE;

    in->min_qp          = INVALID_VALUE;
    in->max_qp          = INVALID_VALUE;
    in->min_i_qp        = INVALID_VALUE;
}   
