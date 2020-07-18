#ifdef SVP

#include "../mpp/include/mpp.h"
#include <stdio.h>

int svp_rt_init() {
    int error_code = 0;

    error_code = HI_SVPRT_RUNTIME_Init("", NULL);
    printf("HI_SVPRT_RUNTIME_Init returns %d\n", error_code);
    //////
    FILE *fp = HI_NULL;

    //segnet_inst.wk  segnet_inst_for_nnie.wk
    fp = fopen("/opt/nfs/runtime_alexnet_no_group_inst.wk", "rb");
    //fp = fopen("/opt/svp/segnet_inst.wk", "rb");
    //fp = fopen("/opt/svp/segnet_inst_for_nnie.wk", "rb");
    printf("fopen returns %d\n", fp);

    error_code = fseek(fp, 0L, SEEK_END);
    printf("fseek returns %d\n", error_code);

    HI_S32 s32RuntimeWkLen = ftell(fp);
    printf("s32RuntimeWkLen is %d\n", s32RuntimeWkLen);

    error_code = fseek(fp, 0L, SEEK_SET);
    printf("fseek returns %d\n", error_code);

    HI_RUNTIME_WK_INFO_S astWkInfo[1];
    memset(&astWkInfo[0], 0, sizeof(astWkInfo));
    strncpy(astWkInfo[0].acModelName, "alexnet", 20);
    
    //HI_RUNTIME_MEM_S stMemInfo;
    //stMemInfo.u32Size = s32RuntimeWkLen;
    astWkInfo[0].stWKMemory.u32Size = s32RuntimeWkLen;

    //s32Ret = HI_MPI_SYS_MmzAlloc_Cached(&pstMemInfo->u64PhyAddr, (HI_VOID **)&(pstMemInfo->u64VirAddr), NULL, HI_NULL, pstMemInfo->u32Size);
    error_code = HI_MPI_SYS_MmzAlloc(&astWkInfo[0].stWKMemory.u64PhyAddr, (HI_VOID **)&(astWkInfo[0].stWKMemory.u64VirAddr), NULL, HI_NULL, astWkInfo[0].stWKMemory.u32Size);
    printf("HI_MPI_SYS_MmzAlloc returns %d\n", error_code);

    error_code = fread((HI_VOID *)((uintptr_t)astWkInfo[0].stWKMemory.u64VirAddr), s32RuntimeWkLen, 1, fp);
    printf("fread returns %d\n", error_code);

    HI_RUNTIME_GROUP_INFO_S stGroupInfo;
    memset(&stGroupInfo, 0, sizeof(HI_RUNTIME_GROUP_INFO_S));

    stGroupInfo.stWKsInfo.u32WKNum = 1;
    stGroupInfo.stWKsInfo.pstAttrs = &(astWkInfo[0]);

    printf("astWkInfo[0].acModelName %s\n", astWkInfo[0].acModelName);
    printf("astWkInfo[0].stWKMemory.u32Size %d\n", astWkInfo[0].stWKMemory.u32Size);


    //////

    
    //HI_RUNTIME_GROUP_HANDLE *phGroupHandle;
    HI_RUNTIME_GROUP_HANDLE hGroupHandle = HI_NULL;

    
    char *config = "name: \"single_alexnet\"      \
                    priority: 1                 \
                    max_tmpbuf_size_mb: 1024    \
                    input {                     \
                        name: \"data\"            \
                    }                           \
                    model {                     \
                        name: \"alexnet\"         \
                        bottom: {name: \"data\"}  \
                    top: {name: \"prob\"}         \
                    }";
    
    /*
    char *config = "name: \"single_ssd\"                                \
                    priority: 1                                         \
                    max_tmpbuf_size_mb: 3072                            \   
                    input {                                             \
                        name: \"data\"                                  \
                    }                                                   \
                    model {                                             \
                        name: \"ssd\"                                   \
                        bottom: {name: \"data\"}                        \
                        top: {name: \"conv4_3_norm_mbox_loc_perm\"}     \
                        top: {name: \"conv4_3_norm_mbox_conf_perm\"}    \
                        top: {name: \"fc7_mbox_loc_perm\"}              \
                        top: {name: \"fc7_mbox_conf_perm\"}             \
                        top: {name: \"conv6_2_mbox_loc_perm\"}          \
                        top: {name: \"conv6_2_mbox_conf_perm\"}         \
                        top: {name: \"conv7_2_mbox_loc_perm\"}          \
                        top: {name: \"conv7_2_mbox_conf_perm\"}         \
                        top: {name: \"conv8_2_mbox_loc_perm\"}          \
                        top: {name: \"conv8_2_mbox_conf_perm\"}         \
                        top: {name: \"conv9_2_mbox_loc_perm\"}          \
                        top: {name: \"conv9_2_mbox_conf_perm\"}         \
                    }";
    */

    printf("%s\n", config);

    error_code = HI_SVPRT_RUNTIME_LoadModelGroup(config, &stGroupInfo, &hGroupHandle);
    //error_code = HI_SVPRT_RUNTIME_LoadModelGroupSync(config, &stGroupInfo, &hGroupHandle); //non exist
    printf("HI_SVPRT_RUNTIME_LoadModelGroup returns %d\n", error_code);

    error_code = HI_SVPRT_RUNTIME_UnloadModelGroup(hGroupHandle);
    printf("HI_SVPRT_RUNTIME_UnloadModelGroup returns %d\n", error_code);

    error_code = HI_SVPRT_RUNTIME_DeInit();
    printf("HI_SVPRT_RUNTIME_DeInit returns %d\n", error_code);

    return 0;
}

#endif
