//go:build !go && !nodejs && !python && !dotnet
// +build !go,!nodejs,!python,!dotnet

package provider

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/optproviderupgrade"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const providerName = "vault"
const defaultBaselineVersion = "6.1.0"

var programs = []string{
	"test-programs/index_authbackend",
	"test-programs/index_policy",
	"test-programs/index_token",
	"test-programs/identity_group",
	"test-programs/generic_secret",
	"test-programs/github_team",
	"test-programs/github_authbackend",
}

func TestUpgradeCoverage(t *testing.T) {
	providertest.ReportUpgradeCoverage(t)
}

type UpgradeTestOpts struct {
	baselineVersion string
	assertFunc      func(*testing.T, auto.PreviewResult)
	config          map[string]string
}

func WithBaselineVersion(baselineVersion string) func(opts *UpgradeTestOpts) {
	return func(opts *UpgradeTestOpts) {
		opts.baselineVersion = baselineVersion
	}
}

func WithAssertFunc(assertFunc func(*testing.T, auto.PreviewResult)) func(opts *UpgradeTestOpts) {
	return func(opts *UpgradeTestOpts) {
		opts.assertFunc = assertFunc
	}
}

func WithConfig(config map[string]string) func(opts *UpgradeTestOpts) {
	return func(opts *UpgradeTestOpts) {
		opts.config = config
	}
}
func testProviderUpgrade(t *testing.T, dir string, opts ...func(*UpgradeTestOpts)) {
	options := &UpgradeTestOpts{}
	for _, o := range opts {
		o(options)
	}
	testProviderUpgradeWithOpts(t, dir, options.baselineVersion, options.config, options.assertFunc)
}

func testProviderUpgradeWithOpts(
	t *testing.T, dir, baselineVersion string, config map[string]string,
	assertFunction func(*testing.T, auto.PreviewResult),
) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without credentials")
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	if baselineVersion == "" {
		baselineVersion = defaultBaselineVersion
	}
	t.Setenv("VAULT_TOKEN", "root")
	t.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	test := pulumitest.NewPulumiTest(t, dir,
		opttest.DownloadProviderVersion(providerName, baselineVersion),
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
	)
	for k, v := range config {
		test.SetConfig(k, v)
	}
	result := providertest.PreviewProviderUpgrade(test, providerName, baselineVersion, optproviderupgrade.DisableAttach())
	if assertFunction != nil {
		assertFunction(t, result)
	} else {
		assertpreview.HasNoReplacements(t, result)
	}
}

func testProgram(t *testing.T, dir string) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without credentials")
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	test := pulumitest.NewPulumiTest(t, dir,
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
		opttest.SkipInstall(),
	)
	test.Up()
}

func TestPrograms(t *testing.T) {
	for _, p := range programs {
		t.Run(p, func(t *testing.T) {
			testProgram(t, p)
		})
	}
}

func TestProgramsUpgrade(t *testing.T) {
	t.Skipf("skip upgrade tests for now as we have not recorded them.")
	for _, p := range programs {
		t.Run(p, func(t *testing.T) {
			testProviderUpgrade(t, p)
		})
	}
}
