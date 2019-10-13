// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestAccPolicy(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "policy"),
		})

	integration.ProgramTest(t, &test)
}

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
	return integration.ProgramTestOptions{
		Tracing: "https://tracing.pulumi-engineering.com/collector/api/v1/spans",
	}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	token := getToken(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"vault:address": "http://127.0.0.1:8200",
			"vault:token": token,
		},
		Dependencies: []string{
			"@pulumi/vault",
		},
	})

	return baseJS
}
