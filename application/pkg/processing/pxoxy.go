//+build processing

package processing

/*
void proxyCallback() {
	printf("proxyCallback\n");
}

void* getCallback(){
	return proxyCallback;
}
*/
import "C"

import (
	"log"
	"unsafe"
)

func init() {
	log.Println("processing init")
	var c unsafe.Pointer
	c = C.getCallback()
	register("proxy", c)
}
