// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

// Retrieve the token.
//
// If the test is running in CI, a missing token will fail the test. Otherwise a missing
// token will skip the test.
func getToken(t *testing.T) string {
	env := "VAULT_DEV_ROOT_TOKEN_ID"
	token := os.Getenv(env)
	if token != "" {
		return token
	}
	if os.Getenv("CI") != "" {
		t.Fatalf("%q is expected but not defined", env)
	}
	t.Skipf("Skipping test due to missing %q environment variable", env)
	return token
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{}
}
