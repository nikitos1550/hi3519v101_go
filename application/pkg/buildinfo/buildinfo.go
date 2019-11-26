// This is special file, vars values are filled during build time

package buildinfo

var (
	GoVersion  string
	GccVersion string

	BuildDateTime string
	BuildTags     string

	BuildUser   string
	BuildCommit string
	BuildBranch string

	BoardProfile string
	BoardVendor  string
	BoardModel   string

	Chip string

	CmosProfile string

	TotalRam string
	LinuxRam string
	MppRam   string
)

type BuildInfo struct {
	GoVersion  string
	GccVersion string

	BuildDateTime string
	BuildTags     string

	BuildUser   string
	BuildCommit string
	BuildBranch string

	BoardProfile string
	BoardVendor  string
	BoardModel   string

	Chip string

	CmosProfile string

	TotalRam string
	LinuxRam string
	MppRam   string
}

func CopyAll(out *BuildInfo) {
	out.GoVersion = GoVersion
	out.GccVersion = GccVersion

	out.BuildDateTime = BuildDateTime
	out.BuildTags = BuildTags

	out.BuildUser = BuildUser
	out.BuildCommit = BuildCommit
	out.BuildBranch = BuildBranch

	out.BoardProfile = BoardProfile
	out.BoardVendor = BoardVendor
	out.BoardModel = BoardModel

	out.Chip = Chip

	out.CmosProfile = CmosProfile

	out.TotalRam = TotalRam
	out.LinuxRam = LinuxRam
	out.MppRam = MppRam
}
