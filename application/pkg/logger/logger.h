#pragma once

#define LOGGER_PANIC 	5
#define LOGGER_FATAL 	4
#define LOGGER_ERROR 	3
#define LOGGER_WANR 	2
#define LOGGER_INFO 	1
#define LOGGER_DEBUG 	0
#define LOGGER_TRACE 	-1


#ifdef LOGGER_C

#define GO_LOG_AI(level, msg) go_logger_ai(level, msg)
#define GO_LOG_VI(level, msg) go_logger_vi(level, msg)
#define GO_LOG_ISP(level, msg) go_logger_isp(level, msg)
#define GO_LOG_MIPI(level, msg) go_logger_mipi(level, msg)
#define GO_LOG_SYS(level, msg) go_logger_sys(level, msg)
#define GO_LOG_VENC(level, msg) go_logger_venc(level, msg)
#define GO_LOG_VPSS(level, msg) go_logger_vpss(level, msg)
//#define GO_LOG_LOOP(level, msg) go_logger_loop(level, msg)

#else

#define GO_LOG_AI(level, msg) ;;
#define GO_LOG_VI(level, msg) ;;
#define GO_LOG_ISP(level, msg) ;;
#define GO_LOG_MIPI(level, msg) ;;
#define GO_LOG_SYS(level, msg) ;;
#define GO_LOG_VENC(level, msg) ;;
#define GO_LOG_VPSS(level, msg) ;;
#define GO_LOG_LOOP(level, msg) ;;

#endif

void go_logger_ai(int level, char * msg);
void go_logger_vi(int level, char * msg);
void go_logger_isp(int level, char * msg);
void go_logger_mipi(int level, char * msg);
void go_logger_sys(int level, char * msg);
void go_logger_venc(int level, char * msg);
void go_logger_vpss(int level, char * msg);
//void go_logger_loop(int level, char * msg);

