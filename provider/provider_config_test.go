// Copyright 2016-2026, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-vault/provider/v7/pkg/version"
)

//go:embed cmd/pulumi-resource-vault/schema.json
var testPulumiSchema []byte

func TestProviderConfigAuthLoginAwsObject(t *testing.T) {
	previousVersion := version.Version
	version.Version = "7.11.0"
	t.Cleanup(func() {
		version.Version = previousVersion
	})

	info := Provider()

	server, err := pfbridge.MakeMuxedServer(
		context.Background(), "vault", info, testPulumiSchema,
	)(nil)
	require.NoError(t, err)

	news, err := structpb.NewStruct(map[string]any{
		"address": "https://vault.example.internal",
		"authLoginAws": map[string]any{
			"awsRegion": "us-east-1",
			"role":      "my-role",
		},
		"skipChildToken": true,
		"version":        "7.11.0",
	})
	require.NoError(t, err)

	response, err := server.CheckConfig(context.Background(), &pulumirpc.CheckRequest{
		Urn:  "urn:pulumi:dev::test::pulumi:providers:vault::vault-provider",
		News: news,
	})
	require.NoError(t, err)
	require.Empty(t, response.Failures)
	require.NotNil(t, response.Inputs)
	authLoginAws, ok := response.Inputs.Fields["authLoginAws"]
	require.True(t, ok)
	require.JSONEq(t,
		`{"awsRegion":"us-east-1","role":"my-role"}`,
		authLoginAws.GetStringValue())
}
