// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "vault:pkiSecret/secretBackendCert:SecretBackendCert":
		r = &SecretBackendCert{}
	case "vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa":
		r = &SecretBackendConfigCa{}
	case "vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers":
		r = &SecretBackendConfigIssuers{}
	case "vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls":
		r = &SecretBackendConfigUrls{}
	case "vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig":
		r = &SecretBackendCrlConfig{}
	case "vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest":
		r = &SecretBackendIntermediateCertRequest{}
	case "vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned":
		r = &SecretBackendIntermediateSetSigned{}
	case "vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer":
		r = &SecretBackendIssuer{}
	case "vault:pkiSecret/secretBackendKey:SecretBackendKey":
		r = &SecretBackendKey{}
	case "vault:pkiSecret/secretBackendRole:SecretBackendRole":
		r = &SecretBackendRole{}
	case "vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert":
		r = &SecretBackendRootCert{}
	case "vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate":
		r = &SecretBackendRootSignIntermediate{}
	case "vault:pkiSecret/secretBackendSign:SecretBackendSign":
		r = &SecretBackendSign{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := vault.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendCert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendConfigCa",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendConfigIssuers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendConfigUrls",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendCrlConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendIntermediateCertRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendIntermediateSetSigned",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendIssuer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendRootCert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendRootSignIntermediate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"pkiSecret/secretBackendSign",
		&module{version},
	)
}
