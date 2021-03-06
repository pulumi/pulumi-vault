// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
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
	case "vault:identity/entity:Entity":
		r = &Entity{}
	case "vault:identity/entityAlias:EntityAlias":
		r = &EntityAlias{}
	case "vault:identity/entityPolicies:EntityPolicies":
		r = &EntityPolicies{}
	case "vault:identity/group:Group":
		r = &Group{}
	case "vault:identity/groupAlias:GroupAlias":
		r = &GroupAlias{}
	case "vault:identity/groupMemberEntityIds:GroupMemberEntityIds":
		r = &GroupMemberEntityIds{}
	case "vault:identity/groupPolicies:GroupPolicies":
		r = &GroupPolicies{}
	case "vault:identity/oidc:Oidc":
		r = &Oidc{}
	case "vault:identity/oidcKey:OidcKey":
		r = &OidcKey{}
	case "vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID":
		r = &OidcKeyAllowedClientID{}
	case "vault:identity/oidcRole:OidcRole":
		r = &OidcRole{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := vault.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"vault",
		"identity/entity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/entityAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/entityPolicies",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/groupAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/groupMemberEntityIds",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/groupPolicies",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/oidc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/oidcKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/oidcKeyAllowedClientID",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"identity/oidcRole",
		&module{version},
	)
}
