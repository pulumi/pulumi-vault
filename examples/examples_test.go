// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getToken(t *testing.T) string {
	token := os.Getenv("VAULT_DEV_ROOT_TOKEN_ID")
	if token == "" {
		t.Skipf("Skipping test due to missing VAULT_DEV_ROOT_TOKEN_ID environment variable")
	}

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
