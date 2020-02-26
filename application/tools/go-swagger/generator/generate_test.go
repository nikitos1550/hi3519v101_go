package generator

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateAndTest(t *testing.T) {
	if runtime.GOOS == "windows" {
		// don't run race tests on Appveyor CI
		t.SkipNow()
	}
	defer func() {
		log.SetOutput(os.Stdout)
	}()

	cases := map[string]struct {
		name      string
		spec      string
		target    string
		template  string
		wantError bool
		prepare   func(opts *GenOpts)
		verify    func(testing.TB, string)
	}{
		"issue 1943": {
			spec:   "../fixtures/bugs/1943/fixture-1943.yaml",
			target: "../fixtures/bugs/1943",
			prepare: func(opts *GenOpts) {
				opts.ExcludeSpec = false
			},
			verify: func(t testing.TB, target string) {
				packages := filepath.Join(target, "...")
				testPrg := filepath.Join(target, "datarace_test.go")

				if p, err := exec.Command("go", "get", packages).CombinedOutput(); err != nil {
					if !assert.NoError(t, err, "go get %s: %s\n%s", packages, err, p) {
						return
					}
				}

				t.Log("running data race test on generated server")
				if p, err := exec.Command("go", "test", "-v", "-race", testPrg).CombinedOutput(); err != nil {
					if !assert.NoError(t, err, "go test -race %s: %s\n%s", packages, err, p) {
						return
					}
				}

			},
		},
		"packages_mangling": {
			spec: "../fixtures/bugs/2111/fixture-2111.yaml",
			prepare: func(opts *GenOpts) {
				opts.IncludeMain = true
			},
			verify: func(t testing.TB, target string) {
				require.True(t, fileExists(target, defaultServerTarget))
				assert.True(t, fileExists(filepath.Join(target, "cmd", "unsafe-tag-names-server"), "main.go"))

				target = filepath.Join(target, defaultServerTarget)

				buf, err := ioutil.ReadFile(filepath.Join(target, "configure_unsafe_tag_names.go"))
				require.NoError(t, err)

				target = filepath.Join(target, defaultOperationsTarget)
				require.True(t, fileExists(target, ""))

				assert.True(t, fileExists(target, "abc_linux"))
				assert.True(t, fileExists(target, "abc_test"))
				assert.True(t, fileExists(target, "api"))
				assert.True(t, fileExists(target, "custom"))
				assert.True(t, fileExists(target, "hash_tag_donuts"))
				assert.True(t, fileExists(target, "nr123abc"))
				assert.True(t, fileExists(target, "nr_at_donuts"))
				assert.True(t, fileExists(target, "plus_donuts"))
				assert.True(t, fileExists(target, "strfmt"))
				assert.True(t, fileExists(target, "forced"))
				assert.True(t, fileExists(target, "gtl"))
				assert.True(t, fileExists(target, "nr12nasty"))
				assert.True(t, fileExists(target, "override"))
				assert.True(t, fileExists(target, "get_notag.go"))
				assert.True(t, fileExists(target, "operationsops"))

				buf2, err := ioutil.ReadFile(filepath.Join(target, "unsafe_tag_names_api.go"))
				require.NoError(t, err)

				// assert imports, with deconfliction
				code := string(buf)
				baseImport := `github.com/go-swagger/go-swagger/fixtures/bugs/2111/packages_mangling/restapi/operations`
				assertImports(t, baseImport, code)

				assertInCode(t, `api.APIGetConflictHandler = apiops.GetConflictHandlerFunc(`, code)
				assertInCode(t, `api.StrfmtGetAnotherConflictHandler = strfmtops.GetAnotherConflictHandlerFunc(`, code)
				assertInCode(t, `api.GetNotagHandler = operations.GetNotagHandlerFunc(`, code)

				api := string(buf2)
				assertImports(t, baseImport, api)

				assertInCode(t, `APIGetConflictHandler: apiops.GetConflictHandlerFunc(func(params apiops.GetConflictParams) middleware.Responder {`, api)
				assertInCode(t, `StrfmtGetAnotherConflictHandler: strfmtops.GetAnotherConflictHandlerFunc(func(params strfmtops.GetAnotherConflictParams) middleware.Responder {`, api)
				assertInCode(t, `GetNotagHandler: GetNotagHandlerFunc(func(params GetNotagParams) middleware.Responder {`, api)

				assertInCode(t, `OverrideDeleteTestOverrideHandler override.DeleteTestOverrideHandler`, api)
				assertInCode(t, `StrfmtGetAnotherConflictHandler strfmtops.GetAnotherConflictHandler`, api)
				assertInCode(t, `APIGetConflictHandler apiops.GetConflictHandler`, api)
				assertInCode(t, `CustomGetCustomHandler custom.GetCustomHandler`, api)
				assertInCode(t, `AbcLinuxGetMultipleHandler abc_linux.GetMultipleHandler`, api)
				assertInCode(t, `GetNotagHandler GetNotagHandler`, api)
				assertInCode(t, `AbcLinuxGetOtherReservedHandler abc_linux.GetOtherReservedHandler`, api)
				assertInCode(t, `PlusDonutsGetOtherUnsafeHandler plus_donuts.GetOtherUnsafeHandler`, api)
				assertInCode(t, `AbcTestGetReservedHandler abc_test.GetReservedHandler`, api)
				assertInCode(t, `GtlGetTestOverrideHandler gtl.GetTestOverrideHandler`, api)
				assertInCode(t, `HashTagDonutsGetUnsafeHandler hash_tag_donuts.GetUnsafeHandler`, api)
				assertInCode(t, `NrAtDonutsGetYetAnotherUnsafeHandler nr_at_donuts.GetYetAnotherUnsafeHandler`, api)
				assertInCode(t, `ForcedPostTestOverrideHandler forced.PostTestOverrideHandler`, api)
				assertInCode(t, `Nr12nastyPutTestOverrideHandler nr12nasty.PutTestOverrideHandler`, api)
				assertInCode(t, `Nr123abcTestIDHandler nr123abc.TestIDHandler`, api)
			},
		},
		"packages_flattening": {
			spec: "../fixtures/bugs/2111/fixture-2111.yaml",
			prepare: func(opts *GenOpts) {
				opts.SkipTagPackages = true
			},
			verify: func(t testing.TB, target string) {
				require.True(t, fileExists(target, defaultServerTarget))

				target = filepath.Join(target, defaultServerTarget)
				buf, err := ioutil.ReadFile(filepath.Join(target, "configure_unsafe_tag_names.go"))
				require.NoError(t, err)

				target = filepath.Join(target, defaultOperationsTarget)
				require.True(t, fileExists(target, ""))

				assert.False(t, fileExists(target, "abc_linux"))
				assert.False(t, fileExists(target, "abc_test"))
				assert.False(t, fileExists(target, "api"))
				assert.False(t, fileExists(target, "custom"))
				assert.False(t, fileExists(target, "hash_tag_donuts"))
				assert.False(t, fileExists(target, "nr123abc"))
				assert.False(t, fileExists(target, "nr_at_donuts"))
				assert.False(t, fileExists(target, "plus_donuts"))
				assert.False(t, fileExists(target, "strfmt"))
				assert.False(t, fileExists(target, "forced"))
				assert.False(t, fileExists(target, "gtl"))
				assert.False(t, fileExists(target, "nr12nasty"))
				assert.False(t, fileExists(target, "override"))
				assert.False(t, fileExists(target, "operationsops"))

				assert.True(t, fileExists(target, "get_notag.go"))

				buf2, err := ioutil.ReadFile(filepath.Join(target, "unsafe_tag_names_api.go"))
				require.NoError(t, err)

				code := string(buf)
				baseImport := `github.com/go-swagger/go-swagger/fixtures/bugs/2111/packages_flattening/restapi/operations`
				assertRegexpInCode(t, baseImport, code)

				assertInCode(t, `api.GetConflictHandler = operations.GetConflictHandlerFunc(`, code)
				assertInCode(t, `api.GetAnotherConflictHandler = operations.GetAnotherConflictHandlerFunc(`, code)
				assertInCode(t, `api.GetNotagHandler = operations.GetNotagHandlerFunc(`, code)

				api := string(buf2)
				assertInCode(t, `GetConflictHandler: GetConflictHandlerFunc(func(params GetConflictParams) middleware.Responder {`, api)
				assertInCode(t, `GetAnotherConflictHandler: GetAnotherConflictHandlerFunc(func(params GetAnotherConflictParams) middleware.Responder {`, api)
				assertInCode(t, `NotagHandler: GetNotagHandlerFunc(func(params GetNotagParams) middleware.Responder {`, api)

				assertInCode(t, `DeleteTestOverrideHandler DeleteTestOverrideHandler`, api)
				assertInCode(t, `GetAnotherConflictHandler GetAnotherConflictHandler`, api)
				assertInCode(t, `GetConflictHandler GetConflictHandler`, api)
				assertInCode(t, `GetCustomHandler GetCustomHandler`, api)
				assertInCode(t, `GetMultipleHandler GetMultipleHandler`, api)
				assertInCode(t, `GetNotagHandler GetNotagHandler`, api)
				assertInCode(t, `GetOtherReservedHandler GetOtherReservedHandler`, api)
				assertInCode(t, `GetOtherUnsafeHandler GetOtherUnsafeHandler`, api)
				assertInCode(t, `GetReservedHandler GetReservedHandler`, api)
				assertInCode(t, `GetTestOverrideHandler GetTestOverrideHandler`, api)
				assertInCode(t, `GetUnsafeHandler GetUnsafeHandler`, api)
				assertInCode(t, `GetYetAnotherUnsafeHandler GetYetAnotherUnsafeHandler`, api)
				assertInCode(t, `PostTestOverrideHandler PostTestOverrideHandler`, api)
				assertInCode(t, `PutTestOverrideHandler PutTestOverrideHandler`, api)
				assertInCode(t, `TestIDHandler TestIDHandler`, api)
			},
		},
		"main_package": {
			spec: "../fixtures/bugs/2111/fixture-2111.yaml",
			prepare: func(opts *GenOpts) {
				opts.IncludeMain = true
				opts.MainPackage = "custom-api"
				opts.SkipTagPackages = true
			},
			verify: func(t testing.TB, target string) {
				assert.True(t, fileExists(filepath.Join(target, "cmd", "custom-api"), "main.go"))
			},
		},
	}

	for name, cas := range cases {
		thisCas := cas

		t.Run(name, func(t *testing.T) {
			var captureLog bytes.Buffer
			log.SetOutput(&captureLog)

			defer func() {
				if t.Failed() {
					t.Logf("ERROR: generation failed:\n%s", captureLog.String())
				}
			}()
			spec := filepath.FromSlash(thisCas.spec)
			opts := testGenOpts()
			if thisCas.target == "" {
				opts.Target = filepath.Join(filepath.Dir(spec), opts.LanguageOpts.ManglePackageName(name, "server_test"))
			} else {
				opts.Target = thisCas.target
			}
			_ = os.Mkdir(opts.Target, 0755)

			opts.Spec = spec

			if thisCas.prepare != nil {
				thisCas.prepare(opts)
			}
			defer func() {
				if thisCas.target == "" {
					_ = os.RemoveAll(filepath.Join(opts.Target))
				} else {
					_ = os.RemoveAll(filepath.Join(opts.Target, defaultServerTarget))
					_ = os.RemoveAll(filepath.Join(opts.Target, "cmd"))
					_ = os.RemoveAll(filepath.Join(opts.Target, defaultModelsTarget))
				}
			}()

			t.Logf("generating test server at: %s", opts.Target)
			err := GenerateServer("", nil, nil, opts)
			if thisCas.wantError {
				require.Errorf(t, err, "expected an error for server build fixture: %s", opts.Spec)
			} else {
				require.NoError(t, err, "unexpected error for server build fixture: %s", opts.Spec)
			}

			if thisCas.verify != nil {
				thisCas.verify(t, opts.Target)
			}
		})
	}
}
