#ifndef ERROR_H_
#define ERROR_H_

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct error_in_struct {
    unsigned int mpp; 
    int general;
} error_in;  

typedef unsigned int error_mpp;
typedef int error_general;

#endif // ERROR_H_
