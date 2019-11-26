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

//TODO make insternal structs by goups (build group, toolchain group, board group, etc)
type Info struct {
	GoVersion  		string	`json:"goversion"`
	GccVersion 		string	`json:"gccversion"`
	BuildDateTime 	string	`json:"builddatetime"`
	BuildTags     	string	`json:"buildtags"`
	BuildUser   	string	`json:"builduser"`
	BuildCommit 	string	`json:"buildcommit"`
	BuildBranch 	string	`json:"buildbranch"`
	BoardProfile 	string	`json:"boardprofile"`
	BoardVendor  	string	`json:"boardvendor"`
	BoardModel   	string	`json:"boardmodel"`
	Chip 			string	`json:"chip"`
	CmosProfile 	string	`json:"cmos"`
	TotalRam 		string	`json:"totalram"`
	LinuxRam 		string	`json:"linuxram"`
	MppRam   		string	`json:"mppram"`
}

func CopyAll(out *Info) {
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
