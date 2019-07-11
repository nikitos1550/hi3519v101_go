#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/ioctl.h>

#include "hi.h"
#include "hi_reg_user.h"

static int fd = -1;

/*
 * read a u32 value by physics address
 * 
 */
int hi_reg_read(unsigned int address, U32 *value)
{
	struct hi_reg_handle handle;
	int ret;

	if (-1 == fd) {
		fd = open("/dev/hi_reg", O_RDWR);
		if (fd < 0) {
			printf("open \"/dev/hi_reg\" failed!\n");
			return -EFAULT;
		}
	}
	handle.phys_addr = address;
	handle.data = NULL;
	handle.size = sizeof(unsigned int);

	ret = ioctl(fd, HI_REG_READ, &handle);
	if (!ret)
		*value = (U32)handle.data;
	return ret;
}

/*
 * write a value by physics address, 
 * which size should be 1, 2 or 4 bytes.
 *
 */
int hi_reg_write(unsigned int address, unsigned int value, int size)
{
	struct hi_reg_handle handle;

    if (-1 == fd) {
		fd = open("/dev/hi_reg", O_RDWR);
		if (fd < 0) {
			printf("open \"/dev/hi_reg\" failed!\n");
			return -EFAULT;
		}
	}
	handle.phys_addr = address;
	handle.data = (void *)value;
	handle.size = size;

	return ioctl(fd, HI_REG_WRITE, &handle);
}

