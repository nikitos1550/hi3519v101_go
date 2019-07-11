#ifndef __HI_REG_USER_H__
#define __HI_REG_USER_H__

#include <linux/ioctl.h>
#include "hi_type.h"

struct hi_reg_handle {
	unsigned int phys_addr;
	unsigned int size;
	void *data;
	int flags;
};


#define HI_REG_BASE  'R'

#define HI_REG_READ  \
	_IOW(HI_REG_BASE, 1, struct hi_reg_handle)
#define HI_REG_WRITE  \
	_IOW(HI_REG_BASE, 2, struct hi_reg_handle)


int hi_reg_read(unsigned int address, U32 *value);
int hi_reg_write(unsigned int address, unsigned int value, int size);

#endif   /* __HI_REG_USER_H__ */
