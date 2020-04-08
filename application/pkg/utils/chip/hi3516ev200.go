//+build arm
//+build hi3516ev200

package chip

var (
    chips = [...]string {
        "hi3516ev300",
        "hi3516ev200",
        "hi3516dv200",
        "hi3518ev300",
    }
)

func RegId() uint32 {
    return 0
}

