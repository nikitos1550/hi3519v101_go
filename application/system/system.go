package system

import (
    "fmt"

    "github.com/vishvananda/netlink"
)

func Test() {
    lo, err := netlink.LinkByName("lo")
    if err != nil {
        fmt.Println(err.Error())
    }

    addr, err := netlink.ParseAddr("169.254.169.254/32")
    if err != nil {
        fmt.Println(err.Error())
    }

    netlink.AddrAdd(lo, addr)
}
