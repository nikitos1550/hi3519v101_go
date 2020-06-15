//+build 386 amd64
//+build host

package utils

func Version() string {
	return "Host"
}

func MppId() uint32 {
	return 0
}

func SyncPTS(pts uint64) {

}
