#include "quirc.h"

struct quirc *qr;
uint8_t *image;

int quirc_quirc_init(int w, int h) {

    qr = quirc_new();
    if (!qr) {
        printf("Failed to allocate memory for qr");
        return HI_NULL;
    }

    printf("w=%d, h=%d\n", w, h);

    if (quirc_resize(qr, w, h) < 0) {
        printf("Failed to allocate video memory for qr");
        return HI_NULL;
    }

    return ERR_NONE;
}

int quirc_quirc_deinit() {

    quirc_destroy(qr);

    return ERR_NONE;
}

int quirc_process(error_in *err, void *frame) {

    VIDEO_FRAME_INFO_S *myFrame = frame;

    int w, h;
    w = myFrame->stVFrame.u32Width;
    h = myFrame->stVFrame.u32Height;

    HI_U8 *pUserPageAddr;
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        pUserPageAddr = (HI_U8 *)HI_MPI_SYS_Mmap(myFrame->stVFrame.u32PhyAddr[0], w*h);
    #elif HI_MPP == 4
        pUserPageAddr = (HI_U8 *)HI_MPI_SYS_Mmap(myFrame->stVFrame.u64PhyAddr[0], w*h);
    #endif

    quirc_begin(qr, &w, &h);
    image = quirc_begin(qr, &w, &h);

    memcpy(image, pUserPageAddr, w*h);
    HI_MPI_SYS_Munmap(pUserPageAddr, w*h);

    quirc_end(qr);

    int num_codes;
    //int i;

    num_codes = quirc_count(qr);
    //printf("Found %d qr codes\n", num_codes);

    if (num_codes > 0) {

        //for (i = 0; i < num_codes; i++) {
	    struct quirc_code code;
	    struct quirc_data data;
	    quirc_decode_error_t errqr;

	    quirc_extract(qr, 0, &code); //quirc_extract(qr, i, &code);

	    errqr = quirc_decode(&code, &data);
	    if (errqr) {
		    //printf("DECODE FAILED: %s\n", quirc_strerror(errqr));
        } else {
		    //printf("Data: %s\n", data.payload);
            //printf("point x:%d y:%d\n", code.corners[0].x, code.corners[0].y);
            //printf("point x:%d y:%d\n", code.corners[1].x, code.corners[1].y);
            //printf("point x:%d y:%d\n", code.corners[2].x, code.corners[2].y);
            //printf("point x:%d y:%d\n", code.corners[3].x, code.corners[3].y);
    
            ////////////////////////////////////////////////////////////////////////////
            //TODO rework all this test code
            #if HI_MPP == 3            
            VGS_HANDLE Handle;
            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VGS_BeginJob, &Handle);

            VGS_TASK_ATTR_S stTask;
            memcpy(&stTask.stImgIn, frame, sizeof(VIDEO_FRAME_INFO_S));
            memcpy(&stTask.stImgOut, frame, sizeof(VIDEO_FRAME_INFO_S));

            VGS_DRAW_LINE_S stVgsDrawLine[4];
            
            for(int j=0;j<4;j++) {
                stVgsDrawLine[j].u32Thick = 2; //[2;8]
                stVgsDrawLine[j].u32Color = 0xFF0000;
    
                
                int next = j+1;
                if (j==3) {
                    next = 0;
                }
    
                stVgsDrawLine[j].stStartPoint.s32X = (code.corners[j].x/2)*2;
                stVgsDrawLine[j].stStartPoint.s32Y = (code.corners[j].y/2)*2;
        
                stVgsDrawLine[j].stEndPoint.s32X   = (code.corners[next].x/2)*2;
                stVgsDrawLine[j].stEndPoint.s32Y   = (code.corners[next].y/2)*2;
                
            }
            
            //DO_OR_RETURN_ERR_MPP(err, HI_MPI_VGS_AddDrawLineTask, Handle, &stTask, &stVgsDrawLine);
            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VGS_AddDrawLineTaskArray, Handle, &stTask, stVgsDrawLine, 4); 

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VGS_EndJob, Handle);
            #endif
        }
    }

    ///////////HI_MPI_SYS_Munmap(pUserPageAddr, w*h);

    return ERR_NONE;    
}
