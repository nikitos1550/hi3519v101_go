#ifndef _LOGGER_H
#define _LOGGER_H

#define LOGGER_PANIC 	5
#define LOGGER_FATAL 	4
#define LOGGER_ERROR 	3
#define LOGGER_WANR 	2
#define LOGGER_INFO 	1
#define LOGGER_DEBUG 	0
#define LOGGER_TRACE 	-1

static void go_logger(int level, char * msg);

#endif //_LOGGER_H

