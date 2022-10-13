// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ad

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
	case "vault:ad/secretBackend:SecretBackend":
		r = &SecretBackend{}
	case "vault:ad/secretLibrary:SecretLibrary":
		r = &SecretLibrary{}
	case "vault:ad/secretRole:SecretRole":
		r = &SecretRole{}
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
		"ad/secretBackend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"ad/secretLibrary",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"ad/secretRole",
		&module{version},
	)
}
