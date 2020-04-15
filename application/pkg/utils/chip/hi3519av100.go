//+build arm
//+build hi3519av100

package chip

var (
    chips = [...]string {
        "hi3559av100",
    }
)

func RegId() uint32 {
    return 0
}

