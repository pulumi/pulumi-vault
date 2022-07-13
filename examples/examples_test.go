// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getVaultAddr() string {
	addr := os.Getenv("VAULT_ADDR")
	if addr == "" {
		return "http://127.0.0.1:8200"
	}

	return addr
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest: false,
	}
}
