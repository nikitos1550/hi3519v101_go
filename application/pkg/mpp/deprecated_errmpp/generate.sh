#!/bin/sh

#//+build arm
#//+build debug
#        
#package errmpp
#        
#/*
##include "errmpp.h"
#*/
#import "C"
#    
#import (
#    "application/pkg/logger"
#)

#func resolveFunc(f uint) string {
#    switch f {

#    case C.ERR_F_HI_MPI_VENC_StopRecvPic:
#        return "HI_MPI_VENC_StopRecvPic"
#    default:
#        logger.Log.Warn().
#            Uint("func", f).
#            Msg("ERRMPP missed desc")
#        return "unknown"
#    }
#}

