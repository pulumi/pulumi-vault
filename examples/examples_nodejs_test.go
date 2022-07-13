// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccPolicy(t *testing.T) {
	test := getJSBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "policy"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"vault:skipChildToken": "true",
		},
		Dependencies: []string{
			"@pulumi/vault",
		},
	})

	return baseJS
}
