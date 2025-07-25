// Copyright 2016-2018, Pulumi Corporation.
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
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/hashicorp/terraform-provider-vault/schema"
	"github.com/hashicorp/terraform-provider-vault/vault"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-vault/provider/v7/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "vault"

	// modules:
	mainMod           = "index"
	adMod             = "AD"
	aliCloudMod       = "AliCloud"
	appRoleMod        = "AppRole"
	awsMod            = "Aws"
	azureMod          = "Azure"
	consulMod         = "Consul"
	databaseMod       = "Database"
	gcpMod            = "Gcp"
	genericMod        = "Generic"
	githubMod         = "GitHub"
	identityMod       = "Identity"
	jwtMod            = "Jwt"
	kmipMod           = "Kmip"
	kubernetesMod     = "Kubernetes"
	kvMod             = "kv"
	ldapMod           = "Ldap"
	managedMod        = "Managed"
	mongoDBAtlasMod   = "MongoDBAtlas"
	oktaMod           = "Okta"
	pkiSecretMod      = "PkiSecret"
	rabbitMqMod       = "RabbitMQ"
	samlMod           = "Saml"
	secretsMod        = "Secrets"
	sshMod            = "Ssh"
	terraformCloudMod = "TerraformCloud"
	tokenMod          = "TokenAuth"
	transformMod      = "Transform"
	transitMod        = "Transit"
	templateMod       = "Template"
)

var moduleMap = map[string]string{
	"ad":              adMod,
	"alicloud":        aliCloudMod,
	"approle":         appRoleMod,
	"aws":             awsMod,
	"azure":           azureMod,
	"consul":          consulMod,
	"config":          "Config",
	"database":        databaseMod,
	"gcp":             gcpMod,
	"generic":         genericMod,
	"github":          githubMod,
	"identity":        identityMod,
	"jwt":             jwtMod,
	"kmip":            kmipMod,
	"kubernetes":      kubernetesMod,
	"kv":              kvMod,
	"ldap":            ldapMod,
	"managed":         managedMod,
	"mongodbatlas":    mongoDBAtlasMod,
	"okta":            oktaMod,
	"pki_secret":      pkiSecretMod,
	"rabbitmq":        rabbitMqMod,
	"saml":            samlMod,
	"secrets":         secretsMod,
	"ssh":             sshMod,
	"terraform_cloud": terraformCloudMod,
	"token":           tokenMod,
	"transform":       transformMod,
	"transit":         transitMod,
	"template":        templateMod,
}

var namespaceMap = map[string]string{
	mainPkg: "Vault",
}

// Override legacy name in JS and Python that were used instead of lowercase.
var specialNamesMap = map[string]string{
	"AppRole":   "appRole",
	"PkiSecret": "pkiSecret",
	"RabbitMQ":  "rabbitMq",
}

func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	if value, exist := specialNamesMap[moduleTitle]; exist {
		moduleName = value
	}
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(mainPkg + ":" + token + ":" + mem)
}

func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

func ref[T any](v T) *T { return &v }

//go:embed cmd/pulumi-resource-vault/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                shimv2.NewProvider(schema.NewProvider(vault.Provider()).SchemaProvider()),
		Name:             "vault",
		DisplayName:      "HashiCorp Vault",
		Description:      "A Pulumi package for creating and managing HashiCorp Vault cloud resources.",
		Keywords:         []string{"pulumi", "vault"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		GitHubOrg:        "hashicorp",
		Repository:       "https://github.com/pulumi/pulumi-vault",
		Version:          version.Version,
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		UpstreamRepoPath: "./upstream",
		DocRules:         &tfbridge.DocRuleInfo{EditRules: docEditRules},

		Config: map[string]*tfbridge.SchemaInfo{
			"skip_tls_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"VAULT_SKIP_VERIFY",
					},
				},
			},
			"max_lease_ttl_seconds": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"TERRAFORM_VAULT_MAX_TTL",
					},
					Value: 1200,
				},
			},
			"max_retries": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"VAULT_MAX_RETRIES",
					},
					Value: 2,
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Main
			"vault_audit":                  {Tok: makeResource(mainMod, "Audit")},
			"vault_audit_request_header":   {Tok: makeResource(mainMod, "AuditRequestHeader")},
			"vault_auth_backend":           {Tok: makeResource(mainMod, "AuthBackend")},
			"vault_cert_auth_backend_role": {Tok: makeResource(mainMod, "CertAuthBackendRole")},
			"vault_egp_policy":             {Tok: makeResource(mainMod, "EgpPolicy")},
			"vault_mfa_duo":                {Tok: makeResource(mainMod, "MfaDuo")},
			"vault_mfa_okta":               {Tok: makeResource(mainMod, "MfaOkta")},
			"vault_mfa_pingid":             {Tok: makeResource(mainMod, "MfaPingid")},
			"vault_mfa_totp":               {Tok: makeResource(mainMod, "MfaTotp")},
			"vault_mount":                  {Tok: makeResource(mainMod, "Mount")},
			"vault_namespace": {
				Tok: makeResource(mainMod, "Namespace"),
				Fields: map[string]*tfbridge.SchemaInfo{"namespace": {
					// error CS0542: 'Namespace': member names cannot be the same as their enclosing type
					CSharpName: "TargetNamespace",
				}},
			},
			"vault_policy": {
				Tok: makeResource(mainMod, "Policy"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"policy": {
						CSharpName: "PolicyContents",
					},
				},
			},
			"vault_rgp_policy":                 {Tok: makeResource(mainMod, "RgpPolicy")},
			"vault_token":                      {Tok: makeResource(mainMod, "Token")},
			"vault_quota_rate_limit":           {Tok: makeResource(mainMod, "QuotaRateLimit")},
			"vault_nomad_secret_backend":       {Tok: makeResource(mainMod, "NomadSecretBackend")},
			"vault_nomad_secret_role":          {Tok: makeResource(mainMod, "NomadSecretRole")},
			"vault_quota_lease_count":          {Tok: makeResource(mainMod, "QuotaLeaseCount")},
			"vault_raft_snapshot_agent_config": {Tok: makeResource(mainMod, "RaftSnapshotAgentConfig")},
			"vault_raft_autopilot":             {Tok: makeResource(mainMod, "RaftAutopilot")},

			// AD
			"vault_ad_secret_backend": {Tok: makeResource(adMod, "SecretBackend")},
			"vault_ad_secret_role":    {Tok: makeResource(adMod, "SecretRole")},
			"vault_ad_secret_library": {
				Tok: makeResource(adMod, "SecretLibrary"),
				Docs: &tfbridge.DocInfo{
					Source: "ad_secret_backend_library.html.md",
				},
			},

			// AppRole
			"vault_approle_auth_backend_role":  {Tok: makeResource(appRoleMod, "AuthBackendRole")},
			"vault_approle_auth_backend_login": {Tok: makeResource(appRoleMod, "AuthBackendLogin")},
			"vault_approle_auth_backend_role_secret_id": {
				Tok: makeResource(appRoleMod, "AuthBackendRoleSecretId"),
				Aliases: []tfbridge.AliasInfo{
					{Type: ref(makeResource(appRoleMod, "AuthBackendRoleSecretID").String())},
				},
			},

			// Azure
			"vault_azure_secret_backend":      {Tok: makeResource(azureMod, "Backend")},
			"vault_azure_secret_backend_role": {Tok: makeResource(azureMod, "BackendRole")},

			// Database
			"vault_database_secret_backend_connection": {
				Tok: makeResource(databaseMod, "SecretBackendConnection"),
				Docs: &tfbridge.DocInfo{
					Source: "database_secret_backend_connection.md",
				},
			},
			"vault_database_secret_backend_role": {
				Tok: makeResource(databaseMod, "SecretBackendRole"),
				Docs: &tfbridge.DocInfo{
					Source: "database_secret_backend_role.md",
				},
			},
			"vault_database_secret_backend_static_role": {
				Tok: makeResource(databaseMod, "SecretBackendStaticRole"),
				Docs: &tfbridge.DocInfo{
					Source: "database_secret_backend_static_role.md",
				},
			},
			"vault_database_secrets_mount": {Tok: makeResource(databaseMod, "SecretsMount")},

			// GCP
			"vault_gcp_auth_backend":                {Tok: makeResource(gcpMod, "AuthBackend")},
			"vault_gcp_auth_backend_role":           {Tok: makeResource(gcpMod, "AuthBackendRole")},
			"vault_gcp_secret_backend":              {Tok: makeResource(gcpMod, "SecretBackend")},
			"vault_gcp_secret_impersonated_account": {Tok: makeResource(gcpMod, "SecretImpersonatedAccount")},
			"vault_gcp_secret_roleset":              {Tok: makeResource(gcpMod, "SecretRoleset")},
			"vault_gcp_secret_static_account":       {Tok: makeResource(gcpMod, "SecretStaticAccount")},

			// Generic
			"vault_generic_endpoint": {Tok: makeResource(genericMod, "Endpoint")},
			"vault_generic_secret":   {Tok: makeResource(genericMod, "Secret")},

			// Github
			"vault_github_auth_backend": {Tok: makeResource(githubMod, "AuthBackend")},
			"vault_github_team": {
				Tok: makeResource(githubMod, "Team"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"team": {
						CSharpName: "TeamCity",
					},
				},
			},
			"vault_github_user": {
				Tok: makeResource(githubMod, "User"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"user": {
						CSharpName: "UserName",
					},
				},
			},

			// SAML
			"vault_saml_auth_backend":      {Tok: makeResource(samlMod, "AuthBackend")},
			"vault_saml_auth_backend_role": {Tok: makeResource(samlMod, "AuthBackendRole")},

			// Identity
			"vault_identity_entity":       {Tok: makeResource(identityMod, "Entity")},
			"vault_identity_entity_alias": {Tok: makeResource(identityMod, "EntityAlias")},
			"vault_identity_group":        {Tok: makeResource(identityMod, "Group")},
			"vault_identity_group_alias": {
				Tok: makeResource(identityMod, "GroupAlias"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Fixes https://github.com/pulumi/pulumi-vault/issues/11
					"name": {
						Name: "name",
					},
				},
			},
			"vault_identity_group_member_group_ids": {Tok: makeResource(identityMod, "GroupMemberGroupIds")},
			"vault_identity_group_policies": {
				Tok: makeResource(identityMod, "GroupPolicies"),
				Docs: &tfbridge.DocInfo{
					Source: "database_secret_backend_role.md",
				},
			},
			"vault_identity_mfa_duo":                    {Tok: makeResource(identityMod, "MfaDuo")},
			"vault_identity_mfa_login_enforcement":      {Tok: makeResource(identityMod, "MfaLoginEnforcement")},
			"vault_identity_mfa_okta":                   {Tok: makeResource(identityMod, "MfaOkta")},
			"vault_identity_mfa_pingid":                 {Tok: makeResource(identityMod, "MfaPingid")},
			"vault_identity_mfa_totp":                   {Tok: makeResource(identityMod, "MfaTotp")},
			"vault_identity_oidc":                       {Tok: makeResource(identityMod, "Oidc")},
			"vault_identity_oidc_key":                   {Tok: makeResource(identityMod, "OidcKey")},
			"vault_identity_oidc_key_allowed_client_id": {Tok: makeResource(identityMod, "OidcKeyAllowedClientID")},
			"vault_identity_oidc_role":                  {Tok: makeResource(identityMod, "OidcRole")},
			"vault_identity_entity_policies": {
				Tok: makeResource(identityMod, "EntityPolicies"),
				Docs: &tfbridge.DocInfo{
					Source: "identity_entity_policies.html.md",
				},
			},
			"vault_identity_group_member_entity_ids": {Tok: makeResource(identityMod, "GroupMemberEntityIds")},
			"vault_identity_oidc_assignment":         {Tok: makeResource(identityMod, "OidcAssignment")},
			"vault_identity_oidc_client":             {Tok: makeResource(identityMod, "OidcClient")},
			"vault_identity_oidc_provider":           {Tok: makeResource(identityMod, "OidcProvider")},
			"vault_identity_oidc_scope":              {Tok: makeResource(identityMod, "OidcScope")},

			// JWT
			"vault_jwt_auth_backend":      {Tok: makeResource(jwtMod, "AuthBackend")},
			"vault_jwt_auth_backend_role": {Tok: makeResource(jwtMod, "AuthBackendRole")},

			// KMIP
			"vault_kmip_secret_backend": {Tok: makeResource(kmipMod, "SecretBackend")},
			"vault_kmip_secret_role":    {Tok: makeResource(kmipMod, "SecretRole")},
			"vault_kmip_secret_scope":   {Tok: makeResource(kmipMod, "SecretScope")},

			// Kubernetes
			"vault_kubernetes_auth_backend_config": {
				Tok: makeResource(kubernetesMod, "AuthBackendConfig"),
				Docs: &tfbridge.DocInfo{
					Source: "kubernetes_auth_backend_config.md",
				},
			},
			"vault_kubernetes_auth_backend_role": {
				Tok: makeResource(kubernetesMod, "AuthBackendRole"),
				Docs: &tfbridge.DocInfo{
					Source: "kubernetes_auth_backend_role.html.md",
				},
			},
			"vault_kubernetes_secret_backend": {
				Tok: makeResource(kubernetesMod, "SecretBackend"),
			},
			"vault_kubernetes_secret_backend_role": {
				Tok: makeResource(kubernetesMod, "SecretBackendRole"),
			},

			// KV
			"vault_kv_secret":            {Tok: makeResource(kvMod, "Secret")},
			"vault_kv_secret_backend_v2": {Tok: makeResource(kvMod, "SecretBackendV2")},
			"vault_kv_secret_v2":         {Tok: makeResource(kvMod, "SecretV2")},

			// LDAP
			"vault_ldap_auth_backend":                {Tok: makeResource(ldapMod, "AuthBackend")},
			"vault_ldap_auth_backend_group":          {Tok: makeResource(ldapMod, "AuthBackendGroup")},
			"vault_ldap_auth_backend_user":           {Tok: makeResource(ldapMod, "AuthBackendUser")},
			"vault_ldap_secret_backend":              {Tok: makeResource(ldapMod, "SecretBackend")},
			"vault_ldap_secret_backend_dynamic_role": {Tok: makeResource(ldapMod, "SecretBackendDynamicRole")},
			"vault_ldap_secret_backend_library_set":  {Tok: makeResource(ldapMod, "SecretBackendLibrarySet")},
			"vault_ldap_secret_backend_static_role":  {Tok: makeResource(ldapMod, "SecretBackendStaticRole")},

			// Managed keys
			"vault_managed_keys": {Tok: makeResource(managedMod, "Keys")},

			// MongoDBAtlas
			"vault_mongodbatlas_secret_backend": {Tok: makeResource(mongoDBAtlasMod, "SecretBackend")},
			"vault_mongodbatlas_secret_role":    {Tok: makeResource(mongoDBAtlasMod, "SecretRole")},

			// Okta
			"vault_okta_auth_backend":       {Tok: makeResource(oktaMod, "AuthBackend")},
			"vault_okta_auth_backend_group": {Tok: makeResource(oktaMod, "AuthBackendGroup")},
			"vault_okta_auth_backend_user":  {Tok: makeResource(oktaMod, "AuthBackendUser")},

			// PKI
			"vault_pki_secret_backend_cert":           {Tok: makeResource(pkiSecretMod, "SecretBackendCert")},
			"vault_pki_secret_backend_config_ca":      {Tok: makeResource(pkiSecretMod, "SecretBackendConfigCa")},
			"vault_pki_secret_backend_config_issuers": {Tok: makeResource(pkiSecretMod, "SecretBackendConfigIssuers")},
			"vault_pki_secret_backend_config_urls":    {Tok: makeResource(pkiSecretMod, "SecretBackendConfigUrls")},
			"vault_pki_secret_backend_intermediate_cert_request": {
				Tok: makeResource(pkiSecretMod, "SecretBackendIntermediateCertRequest"),
			},
			"vault_pki_secret_backend_intermediate_set_signed": {
				Tok: makeResource(pkiSecretMod, "SecretBackendIntermediateSetSigned"),
			},
			"vault_pki_secret_backend_issuer": {Tok: makeResource(pkiSecretMod, "SecretBackendIssuer")},
			"vault_pki_secret_backend_key":    {Tok: makeResource(pkiSecretMod, "SecretBackendKey")},
			"vault_pki_secret_backend_role": {
				Tok: makeResource(pkiSecretMod, "SecretBackendRole"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ou": {
						// Renaming this field avoids an pluralization error.
						Name: "organizationUnit",
					},
				},
			},
			"vault_pki_secret_backend_root_cert": {Tok: makeResource(pkiSecretMod, "SecretBackendRootCert")},
			"vault_pki_secret_backend_root_sign_intermediate": {
				Tok: makeResource(pkiSecretMod, "SecretBackendRootSignIntermediate"),
			},
			"vault_pki_secret_backend_sign":       {Tok: makeResource(pkiSecretMod, "SecretBackendSign")},
			"vault_pki_secret_backend_crl_config": {Tok: makeResource(pkiSecretMod, "SecretBackendCrlConfig")},

			// Transform
			"vault_transform_alphabet": {
				Tok: makeResource(transformMod, "Alphabet"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"alphabet": {
						CSharpName: "AlphabetSet",
					},
				},
			},

			// Transit
			"vault_transit_secret_cache_config": {
				Docs: &tfbridge.DocInfo{
					Source: "transit_secret_backend_cache_config.html.md",
				},
			},
			"vault_config_ui_custom_message": {Docs: missingDocs},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Main
			"vault_policy_document": {
				Tok: makeDataSource(mainMod, "getPolicyDocument"),
				Docs: &tfbridge.DocInfo{
					Source: "policy_document.md",
				},
			},
			"vault_auth_backend": {
				Tok: makeDataSource(mainMod, "getAuthBackend"),
				Docs: &tfbridge.DocInfo{
					Source: "auth_backend.html.md",
				},
			},
			"vault_nomad_access_token":   {Tok: makeDataSource(mainMod, "getNomadAccessToken")},
			"vault_auth_backends":        {Tok: makeDataSource(mainMod, "getAuthBackends")},
			"vault_raft_autopilot_state": {Tok: makeDataSource(mainMod, "getRaftAutopilotState")},

			// AppRole
			"vault_approle_auth_backend_role_id": {
				Docs: &tfbridge.DocInfo{Source: "approle_auth_backend_role_id.md"},
			},

			// Kubernetes
			"vault_kubernetes_auth_backend_config": {
				Docs: &tfbridge.DocInfo{
					Source: "kubernetes_auth_backend_config.md",
				},
			},
			"vault_kubernetes_auth_backend_role": {
				Docs: &tfbridge.DocInfo{
					Source: "kubernetes_auth_backend_role.md",
				},
			},
			"vault_kubernetes_service_account_token": {
				Docs: &tfbridge.DocInfo{Source: "kubernetes_credentials.html.md"},
			},

			// KV
			"vault_kv_secret_subkeys_v2": {
				Docs: &tfbridge.DocInfo{Source: "kv_subkeys_v2.html.md"},
			},

			"vault_aws_static_access_credentials": {Docs: missingDocs},
			"vault_ldap_dynamic_credentials":      {Docs: missingDocs},
			"vault_ldap_static_credentials":       {Docs: missingDocs},
			"vault_namespace": {
				Tok:  makeDataSource(mainMod, "getNamespace"),
				Docs: missingDocs,
			},
			"vault_namespaces": {
				Tok:  makeDataSource(mainMod, "getNamespaces"),
				Docs: missingDocs,
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				RespectSchemaVersion: true,
			}
			i.PyProject.Enabled = true
			return i
		})(),

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
	}

	prov.MustComputeTokens(tks.MappedModules("vault_", "index", moduleMap,
		func(module, name string) (string, error) {
			return string(makeResource(module, name)), nil
		}))
	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	edits := []tfbridge.DocsEdit{
		// These sections would trigger other edit rules so they must run first for discovery.
		cleanUpSecretsWarnings,
	}
	edits = append(edits,
		defaults...,
	)
	return append(edits,
		oktaAuthBackedUserImport,
		skipConfiguringAndPopulatingSection,
		skipBestPracticesSection,
		skipTutorialsSection,
		skipNamespacesSection,
	)
}

var oktaAuthBackedUserImport = tfbridge.DocsEdit{
	Path: "okta_auth_backend_user.html.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		content = append(content, []byte("\n\n## Import\n\n"+
			"Okta authentication backend users can be imported using its `path/user` ID format, e.g.\n"+
			"```\n"+
			"$ terraform import vault_okta_auth_backend_user.example okta/foo\n"+
			"```\n")...)
		return content, nil
	},
}

var missingDocs = &tfbridge.DocInfo{AllowMissing: true}

// Pulumi does encrypt secrets, so redact TF specific warnings.
var cleanUpSecretsWarnings = tfbridge.DocsEdit{
	Path: "index.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		files := []string{
			"using-credentials",
			"overview",
		}

		replacesDir := "provider/installation-replaces/"
		for _, file := range files {
			input, err := os.ReadFile(replacesDir + file + "-input.md")
			if err != nil {
				return nil, err
			}
			replace, err := os.ReadFile(replacesDir + file + "-desired.md")
			if err != nil {
				return nil, err
			}
			if bytes.Contains(content, input) {
				content = bytes.ReplaceAll(
					content,
					input,
					replace)
			} else {
				// Hard error to ensure we keep this content up to date
				return nil, fmt.Errorf("could not find text in upstream index.html.markdown, "+
					"verify file content at %s", replacesDir+file+"-input.md")
			}

		}
		return content, nil
	},
}

// Removes a "Best Practices" section that includes TF-specific recommendations
var skipBestPracticesSection = tfbridge.DocsEdit{
	Path: "index.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Best Practices"
		})
	},
}

// Removes a "Configuring and Populating Vault" section that talks about secrets only
var skipConfiguringAndPopulatingSection = tfbridge.DocsEdit{
	Path: "index.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Configuring and Populating Vault"
		})
	},
}

// This part of the documentation is very TF-specific.
// We should consider creating our own tutorial if the need arises.
// See https://github.com/pulumi/pulumi-vault/issues/618.
var skipNamespacesSection = tfbridge.DocsEdit{
	Path: "index.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Namespace support"
		})
	},
}

// Removes a TF-specific "Tutorials" section
var skipTutorialsSection = tfbridge.DocsEdit{
	Path: "index.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Tutorials"
		})
	},
}
