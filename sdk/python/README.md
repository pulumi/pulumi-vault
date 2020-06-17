[![Build Status](https://travis-ci.com/pulumi/pulumi-vault.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-vault)

# Hashicorp Vault Resource Provider

The Vault resource provider for Pulumi lets you manage Vault resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/vault

or `yarn`:

    $ yarn add @pulumi/vault

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_vault

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-vault/sdk/v2/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Vault

## Configuration

The following configuration points are available:

- `vault:address` - (Required) Origin URL of the Vault server. This is a URL with a scheme, a hostname and a port but with no path.
  May be set via the `VAULT_ADDR` environment variable.
- `vault:token` - (Required) Vault token that will be used by the provider to authenticate. May be set via the `VAULT_TOKEN`
  environment variable. If none is otherwise supplied, the provider will attempt to read it from ~/.vault-token (where the vault
  command stores its current token). The provider will issue itself a new token that is a child of the one given, with a short TTL
  to limit the exposure of any requested secrets. Note that the given token must have the update capability on the `auth/token/create`
  path in Vault in order to create child tokens.
- `vault:tokenName` - (Optional) Token name to use for creating the Vault child token. May be set via the `VAULT_TOKEN_NAME`
  environment variable. 
- `vault:ca_cert_file` - (Optional) Path to a file on local disk that will be used to validate the certificate presented by
  the Vault server. May be set via the `VAULT_CACERT` environment variable.
- `vault:ca_cert_dir` - (Optional) Path to a directory on local disk that contains one or more certificate files that will
  be used to validate the certificate presented by the Vault server. May be set via the `VAULT_CAPATH` environment variable.
- `vault:client_auth` - (Optional) A configuration block, described below, that provides credentials used by the provider
  to authenticate with the Vault server. At present there is little reason to set this, because the provider does not 
  support the TLS certificate authentication mechanism.
    - `vault:cert_file` - (Required) Path to a file on local disk that contains the PEM-encoded certificate to present to the server.
    - `vault:key_file` - (Required) Path to a file on local disk that contains the PEM-encoded private key for which the 
    authentication certificate was issued.
- `vault:skip_tls_verify` - (Optional) Set this to true to disable verification of the Vault server's TLS certificate. This
  is strongly discouraged except in prototype or development environments, since it exposes the possibility that the provider
  can be tricked into writing secrets to a server controlled by an intruder. May be set via the `VAULT_SKIP_VERIFY` environment variable.
- `vault:max_lease_ttl_seconds` - (Optional) Used as the duration for the intermediate Vault token the provider issues itself,
  which in turn limits the duration of secret leases issued by Vault. Defaults to `20` minutes and may be set via the 
  `TERRAFORM_VAULT_MAX_TTL` environment variable. See the section above on Using Vault credentials in the provider configuration
  for the implications of this setting.
- `vault:max_retries` - (Optional) Used as the maximum number of retries when a 5xx error code is encountered. Defaults to `2`
  retries and may be set via the VAULT_MAX_RETRIES environment variable.
- `vault:namespace` - (Optional) Set the namespace to use. May be set via the `VAULT_NAMESPACE` environment variable. Available
  only for Vault Enterprise.

## Reference

For further information, please visit [the Vault provider docs](https://www.pulumi.com/docs/intro/cloud-providers/vault) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/vault).
