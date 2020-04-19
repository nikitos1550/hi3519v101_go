#ifndef ERROR_H_
#define ERROR_H_

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct error_in_struct {
    unsigned int f;
    unsigned int mpp; 
    int general;
} error_in;  

typedef unsigned int error_mpp;
typedef int error_general;

//#define MPP_F_HI_MPI_SYS_Exit
//HI_MPI_VB_Exit
//#define MPP_F_


#define ERR_F_HI_MPI_SYS_Init                   1
#define ERR_F_HI_MPI_SYS_Exit                   2
#define ERR_F_HI_MPI_SYS_SetConf                3

#define ERR_F_HI_MPI_VB_Init                    4
#define ERR_F_HI_MPI_VB_Exit                    5
#define ERR_F_HI_MPI_VB_SetConf                 6
#define ERR_F_HI_MPI_VB_SetConfig               8

#define ERR_F_HI_MPI_ISP_Run                    9
#define ERR_F_HI_MPI_ISP_Exit                   10
#define ERR_F_HI_MPI_AE_Register                11
#define ERR_F_HI_MPI_AWB_Register               12
#define ERR_F_HI_MPI_AF_Register                13
#define ERR_F_HI_MPI_ISP_MemInit                14
#define ERR_F_HI_MPI_ISP_SetWDRMode             15
#define ERR_F_HI_MPI_ISP_SetPubAttr             16
#define ERR_F_HI_MPI_ISP_Init                   17
#define ERR_F_HI_MPI_ISP_SetImageAttr           18
#define ERR_F_HI_MPI_ISP_SetInputTiming         19



#endif // ERROR_H_
