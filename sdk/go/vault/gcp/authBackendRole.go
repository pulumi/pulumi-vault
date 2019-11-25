// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a role in an [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_auth_backend_role.html.markdown.
type AuthBackendRole struct {
	s *pulumi.ResourceState
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOpt) (*AuthBackendRole, error) {
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addGroupAliases"] = nil
		inputs["allowGceInference"] = nil
		inputs["backend"] = nil
		inputs["boundInstanceGroups"] = nil
		inputs["boundLabels"] = nil
		inputs["boundProjects"] = nil
		inputs["boundRegions"] = nil
		inputs["boundServiceAccounts"] = nil
		inputs["boundZones"] = nil
		inputs["maxJwtExp"] = nil
		inputs["maxTtl"] = nil
		inputs["period"] = nil
		inputs["policies"] = nil
		inputs["role"] = nil
		inputs["tokenBoundCidrs"] = nil
		inputs["tokenExplicitMaxTtl"] = nil
		inputs["tokenMaxTtl"] = nil
		inputs["tokenNoDefaultPolicy"] = nil
		inputs["tokenNumUses"] = nil
		inputs["tokenPeriod"] = nil
		inputs["tokenPolicies"] = nil
		inputs["tokenTtl"] = nil
		inputs["tokenType"] = nil
		inputs["ttl"] = nil
		inputs["type"] = nil
	} else {
		inputs["addGroupAliases"] = args.AddGroupAliases
		inputs["allowGceInference"] = args.AllowGceInference
		inputs["backend"] = args.Backend
		inputs["boundInstanceGroups"] = args.BoundInstanceGroups
		inputs["boundLabels"] = args.BoundLabels
		inputs["boundProjects"] = args.BoundProjects
		inputs["boundRegions"] = args.BoundRegions
		inputs["boundServiceAccounts"] = args.BoundServiceAccounts
		inputs["boundZones"] = args.BoundZones
		inputs["maxJwtExp"] = args.MaxJwtExp
		inputs["maxTtl"] = args.MaxTtl
		inputs["period"] = args.Period
		inputs["policies"] = args.Policies
		inputs["role"] = args.Role
		inputs["tokenBoundCidrs"] = args.TokenBoundCidrs
		inputs["tokenExplicitMaxTtl"] = args.TokenExplicitMaxTtl
		inputs["tokenMaxTtl"] = args.TokenMaxTtl
		inputs["tokenNoDefaultPolicy"] = args.TokenNoDefaultPolicy
		inputs["tokenNumUses"] = args.TokenNumUses
		inputs["tokenPeriod"] = args.TokenPeriod
		inputs["tokenPolicies"] = args.TokenPolicies
		inputs["tokenTtl"] = args.TokenTtl
		inputs["tokenType"] = args.TokenType
		inputs["ttl"] = args.Ttl
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("vault:gcp/authBackendRole:AuthBackendRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRole{s: s}, nil
}

// GetAuthBackendRole gets an existing AuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendRoleState, opts ...pulumi.ResourceOpt) (*AuthBackendRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addGroupAliases"] = state.AddGroupAliases
		inputs["allowGceInference"] = state.AllowGceInference
		inputs["backend"] = state.Backend
		inputs["boundInstanceGroups"] = state.BoundInstanceGroups
		inputs["boundLabels"] = state.BoundLabels
		inputs["boundProjects"] = state.BoundProjects
		inputs["boundRegions"] = state.BoundRegions
		inputs["boundServiceAccounts"] = state.BoundServiceAccounts
		inputs["boundZones"] = state.BoundZones
		inputs["maxJwtExp"] = state.MaxJwtExp
		inputs["maxTtl"] = state.MaxTtl
		inputs["period"] = state.Period
		inputs["policies"] = state.Policies
		inputs["role"] = state.Role
		inputs["tokenBoundCidrs"] = state.TokenBoundCidrs
		inputs["tokenExplicitMaxTtl"] = state.TokenExplicitMaxTtl
		inputs["tokenMaxTtl"] = state.TokenMaxTtl
		inputs["tokenNoDefaultPolicy"] = state.TokenNoDefaultPolicy
		inputs["tokenNumUses"] = state.TokenNumUses
		inputs["tokenPeriod"] = state.TokenPeriod
		inputs["tokenPolicies"] = state.TokenPolicies
		inputs["tokenTtl"] = state.TokenTtl
		inputs["tokenType"] = state.TokenType
		inputs["ttl"] = state.Ttl
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("vault:gcp/authBackendRole:AuthBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendRole) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendRole) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *AuthBackendRole) AddGroupAliases() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addGroupAliases"])
}

// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
func (r *AuthBackendRole) AllowGceInference() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["allowGceInference"])
}

// Path to the mounted GCP auth backend
func (r *AuthBackendRole) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
func (r *AuthBackendRole) BoundInstanceGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundInstanceGroups"])
}

// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
func (r *AuthBackendRole) BoundLabels() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundLabels"])
}

// GCP Projects that the role exists within
func (r *AuthBackendRole) BoundProjects() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundProjects"])
}

// The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
func (r *AuthBackendRole) BoundRegions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundRegions"])
}

// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
func (r *AuthBackendRole) BoundServiceAccounts() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundServiceAccounts"])
}

// The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
func (r *AuthBackendRole) BoundZones() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["boundZones"])
}

// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
func (r *AuthBackendRole) MaxJwtExp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["maxJwtExp"])
}

// The maximum allowed lifetime of tokens
// issued using this role, provided as a number of seconds.
func (r *AuthBackendRole) MaxTtl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["maxTtl"])
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. Specified in seconds.
func (r *AuthBackendRole) Period() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["period"])
}

// An array of strings
// specifying the policies to be set on tokens issued using this role.
func (r *AuthBackendRole) Policies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["policies"])
}

// Name of the GCP role
func (r *AuthBackendRole) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (r *AuthBackendRole) TokenBoundCidrs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenBoundCidrs"])
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (r *AuthBackendRole) TokenExplicitMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenExplicitMaxTtl"])
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenMaxTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenMaxTtl"])
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (r *AuthBackendRole) TokenNoDefaultPolicy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["tokenNoDefaultPolicy"])
}

// The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (r *AuthBackendRole) TokenNumUses() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenNumUses"])
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. Specified in seconds.
func (r *AuthBackendRole) TokenPeriod() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenPeriod"])
}

// List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (r *AuthBackendRole) TokenPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tokenPolicies"])
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackendRole) TokenTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["tokenTtl"])
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (r *AuthBackendRole) TokenType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenType"])
}

// The TTL period of tokens issued
// using this role, provided as a number of seconds.
func (r *AuthBackendRole) Ttl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ttl"])
}

// Type of GCP authentication role (either `gce` or `iam`)
func (r *AuthBackendRole) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type AuthBackendRoleState struct {
	AddGroupAliases interface{}
	// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
	AllowGceInference interface{}
	// Path to the mounted GCP auth backend
	Backend interface{}
	// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
	BoundInstanceGroups interface{}
	// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
	BoundLabels interface{}
	// GCP Projects that the role exists within
	BoundProjects interface{}
	// The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
	BoundRegions interface{}
	// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
	BoundServiceAccounts interface{}
	// The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
	BoundZones interface{}
	// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
	MaxJwtExp interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// Name of the GCP role
	Role interface{}
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs interface{}
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl interface{}
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl interface{}
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy interface{}
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod interface{}
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies interface{}
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl interface{}
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType interface{}
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl interface{}
	// Type of GCP authentication role (either `gce` or `iam`)
	Type interface{}
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	AddGroupAliases interface{}
	// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
	AllowGceInference interface{}
	// Path to the mounted GCP auth backend
	Backend interface{}
	// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
	BoundInstanceGroups interface{}
	// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
	BoundLabels interface{}
	// GCP Projects that the role exists within
	BoundProjects interface{}
	// The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
	BoundRegions interface{}
	// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
	BoundServiceAccounts interface{}
	// The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
	BoundZones interface{}
	// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
	MaxJwtExp interface{}
	// The maximum allowed lifetime of tokens
	// issued using this role, provided as a number of seconds.
	MaxTtl interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	Period interface{}
	// An array of strings
	// specifying the policies to be set on tokens issued using this role.
	Policies interface{}
	// Name of the GCP role
	Role interface{}
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs interface{}
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl interface{}
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl interface{}
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy interface{}
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses interface{}
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod interface{}
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies interface{}
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl interface{}
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType interface{}
	// The TTL period of tokens issued
	// using this role, provided as a number of seconds.
	Ttl interface{}
	// Type of GCP authentication role (either `gce` or `iam`)
	Type interface{}
}
