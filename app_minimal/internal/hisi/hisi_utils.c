#include <stdint.h>
#include <errno.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <string.h>
#include <unistd.h>

int devmem(uint32_t target, uint32_t value, uint32_t * read) {
    //reference https://github.com/pavel-a/devmemX/blob/master/devmem2.c

    unsigned int pagesize = (unsigned)getpagesize(); /* or sysconf(_SC_PAGESIZE)  */
    unsigned int map_size = pagesize;
    int access_size = 4;
    unsigned offset;

    offset = (unsigned int)(target & (pagesize-1));
    if (offset + access_size > pagesize ) {
        // Access straddles page boundary:  add another page:
        map_size += pagesize;
    }

    int fd;
    void *map_base, *virt_addr;

    fd = open("/dev/mem", O_RDWR | O_SYNC);
    if (fd == -1) {
        printf("C DEBUG: Error opening /dev/mem (%d) : %s\n", errno, strerror(errno));
        return 1;
    }

    map_base = mmap(0,
                    map_size,
                    PROT_READ | PROT_WRITE, MAP_SHARED,
                    fd,
                    target & ~((typeof(target))pagesize-1));

    if (map_base == (void *) -1) {
        printf("C DEBUG: Error mapping (%d) : %s\n", errno, strerror(errno));
        return 1;//exit(1);
    }
    //printf("Memory mapped at address %p.\n", map_base);

    virt_addr = map_base + offset;

    //unsigned long read_result;
    if (read == NULL ) {
        *((volatile uint32_t *) virt_addr) = value;
    }

    if (read != NULL) {
        *read = *((volatile uint32_t *) virt_addr);
    }
    //printf("0x%lx value 0x%lx\n", target, read_result);

    if (munmap(map_base, map_size) != 0) {
        printf("C DEBUG: ERROR munmap (%d) %s\n", errno, strerror(errno));
    }

    close(fd);
    return 0;
}

