package generator

import (
	"os"
	"path/filepath"
	goruntime "runtime"
	"testing"
)

var checkprefixandfetchrelativepathtests = []struct {
	childpath  string
	parentpath string
	ok         bool
	path       string
}{
	// Positive
	{"/", "/", true, "."},
	{"/User/Gopher", "/", true, "User/Gopher"},
	{"/User/Gopher/Go", "/User/Gopher/Go", true, "."},
	{"/User/../User/Gopher", "/", true, "User/Gopher"},
	// Negative cases
	{"/", "/var", false, ""},
	{"/User/Gopher", "/User/SomethingElse", false, ""},
	{"/var", "/etc", false, ""},
	{"/mnt/dev3", "/mnt/dev3/dir", false, ""},
}

var tempdir = os.TempDir()

var checkbaseimporttest = []struct {
	path         []string
	gopath       string
	targetpath   string
	symlinksrc   string
	symlinkdest  string // symlink is the last dir in targetpath
	expectedpath string
}{
	// No sym link. Positive Test Case
	{[]string{tempdir + "/root/go/src/github.com/go-swagger"}, tempdir + "/root/go/", tempdir + "/root/go/src/github.com/go-swagger", "", "", "github.com/go-swagger"},
	// Symlink points inside GOPATH
	{[]string{tempdir + "/root/go/src/github.com/go-swagger"}, tempdir + "/root/go/", tempdir + "/root/symlink", tempdir + "/root/symlink", tempdir + "/root/go/src/", "."},
	// Symlink points inside GOPATH
	{[]string{tempdir + "/root/go/src/github.com/go-swagger"}, tempdir + "/root/go/", tempdir + "/root/symlink", tempdir + "/root/symlink", tempdir + "/root/go/src/github.com", "github.com"},
	// Symlink point outside GOPATH : Targets Case 1: in baseImport implementation.
	{[]string{tempdir + "/root/go/src/github.com/go-swagger", tempdir + "/root/gopher/go/"}, tempdir + "/root/go/", tempdir + "/root/go/src/github.com/gopher", tempdir + "/root/go/src/github.com/gopher", tempdir + "/root/gopher/go", "github.com/gopher"},
}

func TestCheckPrefixFetchRelPath(t *testing.T) {

	for _, item := range checkprefixandfetchrelativepathtests {
		actualok, actualpath := checkPrefixAndFetchRelativePath(item.childpath, item.parentpath)

		if goruntime.GOOS == "windows" {
			item.path = filepath.FromSlash(item.path)
		}

		switch {
		case actualok != item.ok:
			t.Errorf("checkPrefixAndFetchRelativePath(%s, %s): expected %v, actual %v", item.childpath, item.parentpath, item.ok, actualok)
		case actualpath != item.path:
			t.Errorf("checkPrefixAndFetchRelativePath(%s, %s): expected %s, actual %s", item.childpath, item.parentpath, item.path, actualpath)
		default:
			continue
		}
	}

}

func TestBaseImport(t *testing.T) {

	// 1. Create a root folder /tmp/root
	// 2. Simulate scenario
	//	2.a No Symlink
	//	2.b Symlink from outside of GOPATH to inside
	//  2.c Symlink from inside of GOPATH to outside.
	// 3. Check results.

	oldgopath := os.Getenv("GOPATH")
	golang := GoLangOpts()
	defer func() {
		_ = os.Setenv("GOPATH", oldgopath)
		_ = os.RemoveAll(filepath.Join(tempdir, "root"))
	}()

	for _, item := range checkbaseimporttest {

		// Create Paths
		for _, paths := range item.path {
			_ = os.MkdirAll(paths, 0777)
		}

		// Change GOPATH
		_ = os.Setenv("GOPATH", item.gopath)

		if item.symlinksrc != "" {
			// Create Symlink
			if err := os.Symlink(item.symlinkdest, item.symlinksrc); err == nil {

				// Test (baseImport always with /)
				actualpath := golang.baseImport(item.targetpath)

				if actualpath != item.expectedpath {
					t.Errorf("baseImport(%s): expected %s, actual %s", item.targetpath, item.expectedpath, actualpath)
				}

				_ = os.RemoveAll(filepath.Join(tempdir, "root"))

			} else {
				t.Logf("WARNING:TestBaseImport with symlink could not be carried on. Symlink creation failed for %s -> %s: %v", item.symlinksrc, item.symlinkdest, err)
				t.Logf("WARNING:TestBaseImport with symlink on Windows requires extended privileges (admin or a user with SeCreateSymbolicLinkPrivilege)")
			}
		}
	}

}
