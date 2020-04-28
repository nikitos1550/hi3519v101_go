//+build arm
//+build hi3516cv200
//+build debug

package errmpp

var (
    codes = map[uint] codeInfo {
        //SYS
        0xA0028003: codeInfo{name: "HI_ERR_SYS_ILLEGAL_PARAM", desc: "The parameter configuration is invalid"},
        0xA0028006: codeInfo{name: "HI_ERR_SYS_NULL_PTR", desc: "The pointer is null"},
        0xA0028009: codeInfo{name: "HI_ERR_SYS_NOT_PERM", desc: "The operation is forbidden"},
        0xA0028010: codeInfo{name: "HI_ERR_SYS_NOTREADY", desc: "The system control attributes are not configured"},
        0xA0028012: codeInfo{name: "HI_ERR_SYS_BUSY", desc: "The system is busy"},
        0xA002800C: codeInfo{name: "HI_ERR_SYS_NOMEM", desc: "The memory fails to be allocated due to some causes such as insufficient system memory"},
        //VB
        0xA0018003: codeInfo{name: "HI_ERR_VB_ILLEGAL_PARAM", desc: "The parameter configuration is invalid"},
        0xA0018005: codeInfo{name: "HI_ERR_VB_UNEXIST", desc: "The VB pool does not exist"},
        0xA0018006: codeInfo{name: "HI_ERR_VB_NULL_PTR", desc: "The pointer is null"},
        0xA0018009: codeInfo{name: "HI_ERR_VB_NOT_PERM", desc: "The operation is forbidden"},
        0xA001800C: codeInfo{name: "HI_ERR_VB_NOMEM", desc: "The memory fails to be allocated"},
        0xA001800D: codeInfo{name: "HI_ERR_VB_NOBUF", desc: "The buffer fails to be allocated"},
        0xA0018010: codeInfo{name: "HI_ERR_VB_NOTREADY", desc: "The system control attributes are not configured"},
        0xA0018012: codeInfo{name: "HI_ERR_VB_BUSY", desc: "The system is busy"},
        0xA0018040: codeInfo{name: "HI_ERR_VB_2MPOOLS", desc: "Too many VB pools are created"},
        //VI
        0xA0108001: codeInfo{name: "HI_ERR_VI_INVALID_DEVID", desc: "The VI device ID is invalid"},
        0xA0108002: codeInfo{name: "HI_ERR_VI_INVALID_CHNID", desc: "The VI channel ID is invalid"},
        0xA0108003: codeInfo{name: "HI_ERR_VI_INVALID_PARA", desc: "The VI parameter is invalid"},
        0xA0108006: codeInfo{name: "HI_ERR_VI_INVALID_NULL_PTR", desc: "The pointer of the input parameter is null"},
        0xA0108007: codeInfo{name: "HI_ERR_VI_FAILED_NOTCONFIG", desc: "The attributes of the video device are not set"},
        0xA0108008: codeInfo{name: "HI_ERR_VI_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0108009: codeInfo{name: "HI_ERR_VI_NOT_PERM", desc: "The operation is forbidden"},
        0xA010800C: codeInfo{name: "HI_ERR_VI_NOMEM", desc: "The memory fails to be allocated"},
        0xA010800E: codeInfo{name: "HI_ERR_VI_BUF_EMPTY", desc: "The VI buffer is empty"},
        0xA010800F: codeInfo{name: "HI_ERR_VI_BUF_FULL", desc: "The VI buffer is full"},
        0xA0108010: codeInfo{name: "HI_ERR_VI_SYS_NOTREADY", desc: "The VI system is not initialized"},
        0xA0108012: codeInfo{name: "HI_ERR_VI_BUSY", desc: "The VI system is busy"},
        0xA0108040: codeInfo{name: "HI_ERR_VI_FAILED_NOTENABLE", desc: "The VI device or VI channel is not enabled"},
        0xA0108041: codeInfo{name: "HI_ERR_VI_FAILED_NOTDISABLE", desc: "The VI device or VI channel is not disabled"},
        0xA0108042: codeInfo{name: "HI_ERR_VI_FAILED_CHNOTDISABLE", desc: "The VI channel is not disabled"},
        0xA0108043: codeInfo{name: "HI_ERR_VI_CFG_TIMEOUT", desc: "The video attribute configuration times out"},
        0xA0108044: codeInfo{name: "HI_ERR_VI_NORM_UNMATCH", desc: "Mismatch occurs"},
        0xA0108045: codeInfo{name: "HI_ERR_VI_INVALID_WAYID", desc: "The video channel ID is invalid"},
        0xA0108046: codeInfo{name: "HI_ERR_VI_INVALID_PHYCHNID", desc: "The physical video channel ID is invalid"},
        0xA0108047: codeInfo{name: "HI_ERR_VI_FAILED_NOTBIND", desc: "The video channel is not bound"},
        0xA0108048: codeInfo{name: "HI_ERR_VI_FAILED_BINDED", desc: "The video channel is bound"},
        0xA0108049: codeInfo{name: "HI_ERR_VI_DIS_PROCESS_FAIL", desc: "The DIS fails to run"},
        //VO
        0xA00F8001: codeInfo{name: "HI_ERR_VO_INVALID_DEVID", desc: "The device ID does not fall within the value range"},
        0xA00F8002: codeInfo{name: "HI_ERR_VO_INVALID_CHNID", desc: "The channel ID does not fall within the value range"},
        0xA00F8003: codeInfo{name: "HI_ERR_VO_ILLEGAL_PARAM", desc: "The parameter value does not fall within the value range"},
        0xA00F8006: codeInfo{name: "HI_ERR_VO_NULL_PTR", desc: "The pointer is null"},
        0xA00F8008: codeInfo{name: "HI_ERR_VO_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA00F8009: codeInfo{name: "HI_ERR_VO_NOT_PERMIT", desc: "The operation is forbidden"},
        0xA00F800C: codeInfo{name: "HI_ERR_VO_NO_MEM", desc: "The memory is insufficient"},
        0xA00F8010: codeInfo{name: "HI_ERR_VO_SYS_NOTREADY", desc: "The system is not initialized"},
        0xA00F8012: codeInfo{name: "HI_ERR_VO_BUSY", desc: "The resources are unavailable"},
        0xA00F8040: codeInfo{name: "HI_ERR_VO_DEV_NOT_CONFIG", desc: "The VO device is not configured"},
        0xA00F8041: codeInfo{name: "HI_ERR_VO_DEV_NOT_ENABLE", desc: "The VO device is not enabled"},
        0xA00F8042: codeInfo{name: "HI_ERR_VO_DEV_HAS_ENABLED", desc: "The VO device has been enabled"},
        0xA00F8043: codeInfo{name: "HI_ERR_VO_DEV_HAS_BINDED", desc: "The device has been bound"},
        0xA00F8044: codeInfo{name: "HI_ERR_VO_DEV_NOT_BINDED", desc: "The device is not bound"},
        0xA00F8045: codeInfo{name: "HI_ERR_VO_VIDEO_NOT_ENABLE", desc: "The video layer is not enabled"},
        0xA00F8046: codeInfo{name: "HI_ERR_VO_VIDEO_NOT_DISABLE", desc: "The video layer is not disabled"},
        0xA00F8047: codeInfo{name: "HI_ERR_VO_VIDEO_NOT_CONFIG", desc: "The video layer is not configured"},
        0xA00F8048: codeInfo{name: "HI_ERR_VO_CHN_NOT_DISABLE", desc: "The VO channel is not disabled"},
        0xA00F8049: codeInfo{name: "HI_ERR_VO_CHN_NOT_ENABLE", desc: "No VO channel is enabled"},
        0xA00F804A: codeInfo{name: "HI_ERR_VO_CHN_NOT_CONFIG", desc: "The VO channel is not configured"},
        0xA00F804B: codeInfo{name: "HI_ERR_VO_CHN_NOT_ALLOC", desc: "No VO channel is allocated"},
        0xA00F804C: codeInfo{name: "HI_ERR_VO_INVALID_PATTERN", desc: "The pattern is invalid"},
        0xA00F804D: codeInfo{name: "HI_ERR_VO_INVALID_POSITION", desc: "The cascade position is invalid"},
        0xA00F804E: codeInfo{name: "HI_ERR_VO_WAIT_TIMEOUT", desc: "Waiting times out"},
        0xA00F804F: codeInfo{name: "HI_ERR_VO_INVALID_VFRAME", desc: "The video frame is invalid"},
        0xA00F8050: codeInfo{name: "HI_ERR_VO_INVALID_RECT_PARA", desc: "The rectangle parameter is invalid"},
        0xA00F8051: codeInfo{name: "HI_ERR_VO_SETBEGIN_ALREADY", desc: "The SETBEGIN MPI has been configured"},
        0xA00F8052: codeInfo{name: "HI_ERR_VO_SETBEGIN_NOTYET", desc: "The SETBEGIN MPI is not configured"},
        0xA00F8053: codeInfo{name: "HI_ERR_VO_SETEND_ALREADY", desc: "The SETEND MPI has been configured"},
        0xA00F8054: codeInfo{name: "HI_ERR_VO_SETEND_NOTYET", desc: "The SETEND MPI is not configured"},
        0xA00F8065: codeInfo{name: "HI_ERR_VO_GFX_NOT_DISABLE", desc: "The graphics layer is not disabled"},
        0xA00F8066: codeInfo{name: "HI_ERR_VO_GFX_NOT_BIND", desc: "The graphics layer is not bound"},
        0xA00F8067: codeInfo{name: "HI_ERR_VO_GFX_NOT_UNBIND", desc: "The graphics layer is not unbound"},
        0xA00F8068: codeInfo{name: "HI_ERR_VO_GFX_INVALID_ID", desc: "The graphics layer ID does not fall within the value range"},
        0xA00F806b: codeInfo{name: "HI_ERR_VO_CHN_AREA_OVERLAP", desc: "The VO channel areas overlap"},
        0xA00F806d: codeInfo{name: "HI_ERR_VO_INVALID_LAYERID", desc: "The video layer ID does not fall within the value range"},
        0xA00F806e: codeInfo{name: "HI_ERR_VO_VIDEO_HAS_BINDED", desc: "The video layer has been bound"},
        0xA00F806f: codeInfo{name: "HI_ERR_VO_VIDEO_NOT_BINDED", desc: "The video layer is not bound"},
        //VPSS
        0xA0078001: codeInfo{name: "HI_ERR_VPSS_INVALID_DEVID", desc: "The VPSS group ID is invalid"},
        0xA0078002: codeInfo{name: "HI_ERR_VPSS_INVALID_CHNID", desc: "The VPSS channel ID is invalid"},
        0xA0078003: codeInfo{name: "HI_ERR_VPSS_ILLEGAL_PARAM", desc: "The VPSS parameter is invalid"},
        0xA0078004: codeInfo{name: "HI_ERR_VPSS_EXIST", desc: "A VPSS group is created"},
        0xA0078005: codeInfo{name: "HI_ERR_VPSS_UNEXIST", desc: "No VPSS group is created"},
        0xA0078006: codeInfo{name: "HI_ERR_VPSS_NULL_PTR", desc: "The pointer of the input parameter is null"},
        0xA0078008: codeInfo{name: "HI_ERR_VPSS_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0078009: codeInfo{name: "HI_ERR_VPSS_NOT_PERM", desc: "The operation is forbidden"},
        0xA007800C: codeInfo{name: "HI_ERR_VPSS_NOMEM", desc: "The memory fails to be allocated"},
        0xA007800D: codeInfo{name: "HI_ERR_VPSS_NOBUF", desc: "The buffer pool fails to be allocated"},
        0xA007800E: codeInfo{name: "HI_ERR_VPSS_BUF_EMPTY", desc: "The picture queue is empty"},
        0xA0078010: codeInfo{name: "HI_ERR_VPSS_NOTREADY", desc: "The VPSS is not initialized"},
        0xA0078012: codeInfo{name: "HI_ERR_VPSS_BUSY", desc: "The VPSS is busy"},
        //VENC
        0xA0088002: codeInfo{name: "HI_ERR_VENC_INVALID_CHNID", desc: "The channel ID is invalid"},
        0xA0088003: codeInfo{name: "HI_ERR_VENC_ILLEGAL_PARAM", desc: "The parameter is invalid"},
        0xA0088004: codeInfo{name: "HI_ERR_VENC_EXIST", desc: "The device, channel or resource to be created or applied for exists"},
        0xA0088005: codeInfo{name: "HI_ERR_VENC_UNEXIST", desc: "The device, channel or resource to be used or destroyed does not exist"},
        0xA0088006: codeInfo{name: "HI_ERR_VENC_NULL_PTR", desc: "The parameter pointer is null"},
        0xA0088007: codeInfo{name: "HI_ERR_VENC_NOT_CONFIG", desc: "No parameter is set before use"},
        0xA0088008: codeInfo{name: "HI_ERR_VENC_NOT_SUPPORT", desc: "The parameter or function is not supported"},
        0xA0088009: codeInfo{name: "HI_ERR_VENC_NOT_PERM", desc: "The operation, for example, modifying static parameters, is forbidden"},
        0xA008800C: codeInfo{name: "HI_ERR_VENC_NOMEM", desc: "The memory fails to be allocated due to some causes such as insufficient system memory"},
        0xA008800D: codeInfo{name: "HI_ERR_VENC_NOBUF", desc: "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"},
        0xA008800E: codeInfo{name: "HI_ERR_VENC_BUF_EMPTY", desc: "The buffer is empty"},
        0xA008800F: codeInfo{name: "HI_ERR_VENC_BUF_FULL", desc: "The buffer is full"},
        0xA0088010: codeInfo{name: "HI_ERR_VENC_SYS_NOTREADY", desc: "The system is not initialized or the corresponding module is not loaded"},
        0xA0088012: codeInfo{name: "HI_ERR_VENC_BUSY", desc: "The VENC system is busy"},
        //VDA
        0xA0098001: codeInfo{name: "HI_ERR_VDA_INVALID_DEVID", desc: "The device ID exceeds the valid range"},
        0xA0098002: codeInfo{name: "HI_ERR_VDA_INVALID_CHNID", desc: "The channel ID exceeds the valid range"},
        0xA0098003: codeInfo{name: "HI_ERR_VDA_ILLEGAL_PARAM", desc: "The parameter value exceeds its valid range"},
        0xA0098004: codeInfo{name: "HI_ERR_VDA_EXIST", desc: "The device, channel, or resource to be created or applied for already exists"},
        0xA0098005: codeInfo{name: "HI_ERR_VDA_UNEXIST", desc: "The device, channel, or resource to be used or destroyed does not exist"},
        0xA0098006: codeInfo{name: "HI_ERR_VDA_NULL_PTR", desc: "The pointer is null"},
        0xA0098007: codeInfo{name: "HI_ERR_VDA_NOT_CONFIG", desc: "The system or VDA channel is not configured"},
        0xA0098008: codeInfo{name: "HI_ERR_VDA_NOT_SUPPORT", desc: "The parameter or function is not supported"},
        0xA0098009: codeInfo{name: "HI_ERR_VDA_NOT_PERM", desc: "The operation, for example, attempting to modify the value of a static parameter, is forbidden"},
        0xA009800C: codeInfo{name: "HI_ERR_VDA_NOMEM", desc: "The memory fails to be allocated due to some causes such as insufficient system memory"},
        0xA009800D: codeInfo{name: "HI_ERR_VDA_NOBUF", desc: "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"},
        0xA009800E: codeInfo{name: "HI_ERR_VDA_BUF_EMPTY", desc: "The buffer is empty"},
        0xA009800F: codeInfo{name: "HI_ERR_VDA_BUF_FULL", desc: "The buffer is full"},
        0xA0098010: codeInfo{name: "HI_ERR_VDA_SYS_NOTREADY", desc: "The system is not initialized or the corresponding module is not loaded"},
        0xA0098012: codeInfo{name: "HI_ERR_VDA_BUSY", desc: "The system is busy"},
        //RGN
        0xA0038001: codeInfo{name: "HI_ERR_RGN_INVALID_DEVID", desc: "The device ID exceeds the valid range"},
        0xA0038002: codeInfo{name: "HI_ERR_RGN_INVALID_CHNID", desc: "The channel ID is incorrect or the region handle is invalid"},
        0xA0038003: codeInfo{name: "HI_ERR_RGN_ILLEGAL_PARAM", desc: "The parameter value exceeds its valid range"},
        0xA0038004: codeInfo{name: "HI_ERR_RGN_EXIST", desc: "The device, channel, or resource to be created already exists"},
        0xA0038005: codeInfo{name: "HI_ERR_RGN_UNEXIST", desc: "The device, channel, or resource to be used or destroyed does not exist"},
        0xA0038006: codeInfo{name: "HI_ERR_RGN_NULL_PTR", desc: "The pointer is null"},
        0xA0038007: codeInfo{name: "HI_ERR_RGN_NOT_CONFIG", desc: "The module is not configured"},
        0xA0038008: codeInfo{name: "HI_ERR_RGN_NOT_SUPPORT", desc: "The parameter or function is not supported"},
        0xA0038009: codeInfo{name: "HI_ERR_RGN_NOT_PERM", desc: "The operation, for example, attempting to modify the value of a static parameter, is forbidden"},
        0xA003800C: codeInfo{name: "HI_ERR_RGN_NOMEM", desc: "The memory fails to be allocated due to some causes such as insufficient system memory"},
        0xA003800D: codeInfo{name: "HI_ERR_RGN_NOBUF", desc: "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"},
        0xA003800E: codeInfo{name: "HI_ERR_RGN_BUF_EMPTY", desc: "The buffer is empty"},
        0xA003800F: codeInfo{name: "HI_ERR_RGN_BUF_FULL", desc: "The buffer is full"},
        0xA0038010: codeInfo{name: "HI_ERR_RGN_NOTREADY", desc: "The system is not initialized or the corresponding module is not loaded"},
        0xA0038011: codeInfo{name: "HI_ERR_RGN_BADADDR", desc: "The address is invalid"},
        0xA0038012: codeInfo{name: "HI_ERR_RGN_BUSY", desc: "The system is busy"},
        //AI
        0xA0158001: codeInfo{name: "HI_ERR_AI_INVALID_DEVID", desc: "The AI device ID is invalid"},
        0xA0158002: codeInfo{name: "HI_ERR_AI_INVALID_CHNID", desc: "The AI channel ID is invalid"},
        0xA0158003: codeInfo{name: "HI_ERR_AI_ILLEGAL_PARAM", desc: "The settings of the AI parameters are invalid"},
        0xA0158005: codeInfo{name: "HI_ERR_AI_NOT_ENABLED", desc: "The AI device or AI channel is not enabled"},
        0xA0158006: codeInfo{name: "HI_ERR_AI_NULL_PTR", desc: "The input parameter pointer is null"},
        0xA0158007: codeInfo{name: "HI_ERR_AI_NOT_CONFIG", desc: "The attributes of an AI device are not set"},
        0xA0158008: codeInfo{name: "HI_ERR_AI_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0158009: codeInfo{name: "HI_ERR_AI_NOT_PERM", desc: "The operation is forbidden"},
        0xA015800C: codeInfo{name: "HI_ERR_AI_NOMEM", desc: "The memory fails to be allocated"},
        0xA015800D: codeInfo{name: "HI_ERR_AI_NOBUF", desc: "The AI buffer is insufficient"},
        0xA015800E: codeInfo{name: "HI_ERR_AI_BUF_EMPTY", desc: "The AI buffer is empty"},
        0xA015800F: codeInfo{name: "HI_ERR_AI_BUF_FULL", desc: "The AI buffer is full"},
        0xA0158010: codeInfo{name: "HI_ERR_AI_SYS_NOTREADY", desc: "The AI system is not initialized"},
        0xA0158012: codeInfo{name: "HI_ERR_AI_BUSY", desc: "The AI system is busy"},
        0xA0158041: codeInfo{name: "HI_ERR_AI_VQE_ERR", desc: "A VQE processing error occurs in the AI channel"},
        //AO
        0xA0168001: codeInfo{name: "HI_ERR_AO_INVALID_DEVID", desc: "The AO device ID is invalid"},
        0xA0168002: codeInfo{name: "HI_ERR_AO_INVALID_CHNID", desc: "The AO channel ID is invalid"},
        0xA0168003: codeInfo{name: "HI_ERR_AO_ILLEGAL_PARAM", desc: "The settings of the AO parameters are invalid"},
        0xA0168005: codeInfo{name: "HI_ERR_AO_NOT_ENABLED", desc: "The AO device or AO channel is not enabled"},
        0xA0168006: codeInfo{name: "HI_ERR_AO_NULL_PTR", desc: "The output parameter pointer is null"},
        0xA0168007: codeInfo{name: "HI_ERR_AO_NOT_CONFIG", desc: "The attributes of an AO device are not set"},
        0xA0168008: codeInfo{name: "HI_ERR_AO_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0168009: codeInfo{name: "HI_ERR_AO_NOT_PERM", desc: "The operation is forbidden"},
        0xA016800C: codeInfo{name: "HI_ERR_AO_NOMEM", desc: "The system memory is insufficient"},
        0xA016800D: codeInfo{name: "HI_ERR_AO_NOBUF", desc: "The AO buffer is insufficient"},
        0xA016800E: codeInfo{name: "HI_ERR_AO_BUF_EMPTY", desc: "The AO buffer is empty"},
        0xA016800F: codeInfo{name: "HI_ERR_AO_BUF_FULL", desc: "The AO buffer is full"},
        0xA0168010: codeInfo{name: "HI_ERR_AO_SYS_NOTREADY", desc: "The AO system is not initialized"},
        0xA0168012: codeInfo{name: "HI_ERR_AO_BUSY", desc: "The AO system is busy"},
        0xA0168041: codeInfo{name: "HI_ERR_AO_VQE_ERR", desc: "A VQE processing error occurs in the AO channel"},
        //AENC
        0xA0178001: codeInfo{name: "HI_ERR_AENC_INVALID_DEVID", desc: "The AENC device ID is invalid"},
        0xA0178002: codeInfo{name: "HI_ERR_AENC_INVALID_CHNID", desc: "The AENC channel ID is invalid"},
        0xA0178003: codeInfo{name: "HI_ERR_AENC_ILLEGAL_PARAM", desc: "The settings of the AENC parameters are invalid"},
        0xA0178004: codeInfo{name: "HI_ERR_AENC_EXIST", desc: "An AENC channel is created"},
        0xA0178005: codeInfo{name: "HI_ERR_AENC_UNEXIST", desc: "An AENC channel is not created"},
        0xA0178006: codeInfo{name: "HI_ERR_AENC_NULL_PTR", desc: "The input parameter pointer is null"},
        0xA0178007: codeInfo{name: "HI_ERR_AENC_NOT_CONFIG", desc: "The AENC channel is not configured"},
        0xA0178008: codeInfo{name: "HI_ERR_AENC_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0178009: codeInfo{name: "HI_ERR_AENC_NOT_PERM", desc: "The operation is forbidden"},
        0xA017800C: codeInfo{name: "HI_ERR_AENC_NOMEM", desc: "The system memory is insufficient"},
        0xA017800D: codeInfo{name: "HI_ERR_AENC_NOBUF", desc: "The buffer for AENC channels fails to be allocated"},
        0xA017800E: codeInfo{name: "HI_ERR_AENC_BUF_EMPTY", desc: "The AENC channel buffer is empty"},
        0xA017800F: codeInfo{name: "HI_ERR_AENC_BUF_FULL", desc: "The AENC channel buffer is full"},
        0xA0178010: codeInfo{name: "HI_ERR_AENC_SYS_NOTREADY", desc: "The system is not initialized"},
        0xA0178040: codeInfo{name: "HI_ERR_AENC_ENCODER_ERR", desc: "An AENC data error occurs"},
        0xA0178041: codeInfo{name: "HI_ERR_AENC_VQE_ERR", desc: "A VQE processing error occurs in the AENC channel"},
        //ADEC
        0xA0188001: codeInfo{name: "HI_ERR_ADEC_INVALID_DEVID", desc: "The ADEC device is invalid"},
        0xA0188002: codeInfo{name: "HI_ERR_ADEC_INVALID_CHNID", desc: "The ADEC channel ID is invalid"},
        0xA0188003: codeInfo{name: "HI_ERR_ADEC_ILLEGAL_PARAM", desc: "The settings of the ADEC parameters are invalid"},
        0xA0188004: codeInfo{name: "HI_ERR_ADEC_EXIST", desc: "An ADEC channel is created"},
        0xA0188005: codeInfo{name: "HI_ERR_ADEC_UNEXIST", desc: "An ADEC channel is not created"},
        0xA0188006: codeInfo{name: "HI_ERR_ADEC_NULL_PTR", desc: "The input parameter pointer is null"},
        0xA0188007: codeInfo{name: "HI_ERR_ADEC_NOT_CONFIG", desc: "The attributes of an ADEC channel are not set"},
        0xA0188008: codeInfo{name: "HI_ERR_ADEC_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA0188009: codeInfo{name: "HI_ERR_ADEC_NOT_PERM", desc: "The operation is forbidden"},
        0xA018800C: codeInfo{name: "HI_ERR_ADEC_NOMEM", desc: "The system memory is insufficient"},
        0xA018800D: codeInfo{name: "HI_ERR_ADEC_NOBUF", desc: "The buffer for ADEC channels fails to be allocated"},
        0xA018800E: codeInfo{name: "HI_ERR_ADEC_BUF_EMPTY", desc: "The ADEC channel buffer is empty"},
        0xA018800F: codeInfo{name: "HI_ERR_ADEC_BUF_FULL", desc: "The ADEC channel buffer is full"},
        0xA0188010: codeInfo{name: "HI_ERR_ADEC_SYS_NOTREADY", desc: "The system is not initialized"},
        0xA0188040: codeInfo{name: "HI_ERR_ADEC_DECODER_ERR", desc: "An ADEC data error occurs"},
        0xA0188041: codeInfo{name: "HI_ERR_ADEC_BUF_LACK", desc: "An insufficient buffer occurs in the ADEC channel"},
        //VGS
        0xA02D800E: codeInfo{name: "HI_ERR_VGS_BUF_EMPTY", desc: "The VGS jobs, tasks, or nodes are used up"},
        0xA02D8003: codeInfo{name: "HI_ERR_VGS_ILLEGAL_PARAM", desc: "The VGS parameter value is invalid"},
        0xA02D8006: codeInfo{name: "HI_ERR_VGS_NULL_PTR", desc: "The input parameter pointer is null"},
        0xA02D8008: codeInfo{name: "HI_ERR_VGS_NOT_SUPPORT", desc: "The operation is not supported"},
        0xA02D8009: codeInfo{name: "HI_ERR_VGS_NOT_PERMITTED", desc: "The operation is forbidden"},
        0xA02D800D: codeInfo{name: "HI_ERR_VGS_NOBUF", desc: "The memory fails to be allocated"},
        0xA02D8010: codeInfo{name: "HI_ERR_VGS_SYS_NOTREADY", desc: "The system is not initialized"},
        //ISP
        0xA01C8006: codeInfo{name: "HI_ERR_ISP_NULL_PTR", desc: "The input pointer is null"},
        0xA01C8003: codeInfo{name: "HI_ERR_ISP_ILLEGAL_PARAM", desc: "The input parameter is invalid"},
        0xA01C8008: codeInfo{name: "HI_ERR_ISP_NOT_SUPPORT", desc: "This function is not supported by the ISP"},
        0xA01C8043: codeInfo{name: "HI_ERR_ISP_SNS_UNREGISTER", desc: "The sensor is not registered"},
        0xA01C8041: codeInfo{name: "HI_ERR_ISP_MEM_NOT_INIT", desc: "The external registers are not initialized"},
        0xA01C8040: codeInfo{name: "HI_ERR_ISP_NOT_INIT", desc: "The ISP is not initialized"},
        0xA01C8044: codeInfo{name: "HI_ERR_ISP_INVALID_ADDR", desc: "The address is invalid"},
        0xA01C8042: codeInfo{name: "HI_ERR_ISP_ATTR_NOT_CFG", desc: "The attribute is not configured"},
        0xA01C8045: codeInfo{name: "HI_ERR_ISP_NOMEM", desc: "The memory is insufficient"},
        0xA01C8046: codeInfo{name: "HI_ERR_ISP_NO_INT", desc: "The ISP module has no interrupt"},
        //IVE
        0xA01D8001: codeInfo{name: "HI_ERR_IVE_INVALID_DEVID", desc: "The device ID is invalid"},
        0xA01D8002: codeInfo{name: "HI_ERR_IVE_INVALID_CHNID", desc: "The channel ID or the region handle is invalid"},
        0xA01D8003: codeInfo{name: "HI_ERR_IVE_ILLEGAL_PARAM", desc: "The parameter is invalid"},
        0xA01D8004: codeInfo{name: "HI_ERR_IVE_EXIST", desc: "The device, channel, or resource to be created exists"},
        0xA01D8005: codeInfo{name: "HI_ERR_IVE_UNEXIST", desc: "The device, channel, or resoure to be used or destroyed does not exist"},
        0xA01D8006: codeInfo{name: "HI_ERR_IVE_NULL_PTR", desc: "The pointer is null"},
        0xA01D8007: codeInfo{name: "HI_ERR_IVE_NOT_CONFIG", desc: "The module is not configured"},
        0xA01D8008: codeInfo{name: "HI_ERR_IVE_NOT_SUPPORT", desc: "The parameter or function is not supported"},
        0xA01D8009: codeInfo{name: "HI_ERR_IVE_NOT_PERM", desc: "The operation, for example, attempting to modify the value of a static parameter, is forbidden"},
        0xA01D800C: codeInfo{name: "HI_ERR_IVE_NOMEM", desc: "The memory fails to be allocated for the reasons such as insufficient system memory"},
        0xA01D800D: codeInfo{name: "HI_ERR_IVE_NOBUF", desc: "The buffer fails to be allocated. For example, the requested data buffer is too large"},
        0xA01D800E: codeInfo{name: "HI_ERR_IVE_BUF_EMPTY", desc: "There is no image in the buffer"},
        0xA01D800F: codeInfo{name: "HI_ERR_IVE_BUF_FULL", desc: "The buffer is full of images"},
        0xA01D8010: codeInfo{name: "HI_ERR_IVE_NOTREADY", desc: "The system is not initialized or the corresponding driver is not loaded"},
        0xA01D8011: codeInfo{name: "HI_ERR_IVE_BADADDR", desc: "The address is invalid"},
        0xA01D8012: codeInfo{name: "HI_ERR_IVE_BUSY", desc: "The system is busy"},
        0xA01D8040: codeInfo{name: "HI_ERR_IVE_SYS_TIMEOUT", desc: "The system times out"},
        0xA01D8041: codeInfo{name: "HI_ERR_IVE_QUERY_TIMEOUT", desc: "Querying times out"},
        0xA01D8042: codeInfo{name: "HI_ERR_IVE_OPEN_FILE", desc: "Opening a file fails"},
        0xA01D8043: codeInfo{name: "HI_ERR_IVE_READ_FILE", desc: "Reading a file fails"},
        0xA01D8044: codeInfo{name: "HI_ERR_IVE_WRITE_FILE", desc: "Writing to a file fails"},
        //ODT
        0xA0308002: codeInfo{name: "HI_ERR_ODT_INVALID_CHNID", desc: "The on-die termination (ODT) channel group ID or the region handle is invalid"},
        0xA0308004: codeInfo{name: "HI_ERR_ODT_EXIST", desc: "The device, channel, or resource to be created already exists"},
        0xA0308005: codeInfo{name: "HI_ERR_ODT_UNEXIST", desc: "The device, channel, or resource to be used or destroyed does not exist"},
        0xA0308009: codeInfo{name: "HI_ERR_ODT_NOT_PERM", desc: "The operation, for example, modifying the value of a static parameter, is forbidden"},
        0xA0308010: codeInfo{name: "HI_ERR_ODT_NOTREADY", desc: "The ODT is not initialized"},
        0xA0308012: codeInfo{name: "HI_ERR_ODT_BUSY", desc: "The ODT is busy"},
    }
)