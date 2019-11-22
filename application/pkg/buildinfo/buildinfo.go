// This is special file, vars values are filled during build time

package buildinfo

var (
    GoVersion       string
    GccVersion      string

    BuildDateTime   string
    BuildTags       string

    BoardVendor     string
    BoardModel      string

    Chip            string

    CmosProfile     string

    TotalRam        uint
    LinuxRam        uint
    MppRam          uint
)
