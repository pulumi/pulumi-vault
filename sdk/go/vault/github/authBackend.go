// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Github Auth mount in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/github.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/github_auth_backend.html.markdown.
type AuthBackend struct {
	s *pulumi.ResourceState
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOpt) (*AuthBackend, error) {
	if args == nil || args.Organization == nil {
		return nil, errors.New("missing required argument 'Organization'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["baseUrl"] = nil
		inputs["description"] = nil
		inputs["maxTtl"] = nil
		inputs["organization"] = nil
		inputs["path"] = nil
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
		inputs["tune"] = nil
	} else {
		inputs["baseUrl"] = args.BaseUrl
		inputs["description"] = args.Description
		inputs["maxTtl"] = args.MaxTtl
		inputs["organization"] = args.Organization
		inputs["path"] = args.Path
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
		inputs["tune"] = args.Tune
	}
	inputs["accessor"] = nil
	s, err := ctx.RegisterResource("vault:github/authBackend:AuthBackend", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackend{s: s}, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendState, opts ...pulumi.ResourceOpt) (*AuthBackend, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessor"] = state.Accessor
		inputs["baseUrl"] = state.BaseUrl
		inputs["description"] = state.Description
		inputs["maxTtl"] = state.MaxTtl
		inputs["organization"] = state.Organization
		inputs["path"] = state.Path
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
		inputs["tune"] = state.Tune
	}
	s, err := ctx.ReadResource("vault:github/authBackend:AuthBackend", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackend{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackend) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackend) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
func (r *AuthBackend) Accessor() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessor"])
}

// The API endpoint to use. Useful if you
// are running GitHub Enterprise or an API-compatible authentication server.
func (r *AuthBackend) BaseUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["baseUrl"])
}

// Specifies the description of the mount.
// This overrides the current stored value, if any.
func (r *AuthBackend) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// (Optional; Deprecated, use `tokenMaxTtl` instead) The maximum allowed lifetime of tokens
// issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
func (r *AuthBackend) MaxTtl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["maxTtl"])
}

// The organization configured users must be part of.
func (r *AuthBackend) Organization() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["organization"])
}

// Path where the auth backend is mounted. Defaults to `auth/github`
// if not specified.
func (r *AuthBackend) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// (Optional) List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (r *AuthBackend) TokenBoundCidrs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tokenBoundCidrs"])
}

// (Optional) If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (r *AuthBackend) TokenExplicitMaxTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenExplicitMaxTtl"])
}

// (Optional) The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackend) TokenMaxTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenMaxTtl"])
}

// (Optional) If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (r *AuthBackend) TokenNoDefaultPolicy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["tokenNoDefaultPolicy"])
}

// (Optional) The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (r *AuthBackend) TokenNumUses() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenNumUses"])
}

// Generated Token's Period
func (r *AuthBackend) TokenPeriod() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenPeriod"])
}

// (Optional) List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (r *AuthBackend) TokenPolicies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tokenPolicies"])
}

// (Optional) The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (r *AuthBackend) TokenTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tokenTtl"])
}

// (Optional) The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (r *AuthBackend) TokenType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tokenType"])
}

// (Optional; Deprecated, use `tokenTtl` isntead) The TTL period of tokens issued
// using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
func (r *AuthBackend) Ttl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ttl"])
}

func (r *AuthBackend) Tune() *pulumi.Output {
	return r.s.State["tune"]
}

// Input properties used for looking up and filtering AuthBackend resources.
type AuthBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor interface{}
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl interface{}
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description interface{}
	// (Optional; Deprecated, use `tokenMaxTtl` instead) The maximum allowed lifetime of tokens
	// issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl interface{}
	// The organization configured users must be part of.
	Organization interface{}
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path interface{}
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs interface{}
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl interface{}
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl interface{}
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy interface{}
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses interface{}
	// Generated Token's Period
	TokenPeriod interface{}
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies interface{}
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl interface{}
	// (Optional) The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType interface{}
	// (Optional; Deprecated, use `tokenTtl` isntead) The TTL period of tokens issued
	// using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
	Ttl interface{}
	Tune interface{}
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl interface{}
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description interface{}
	// (Optional; Deprecated, use `tokenMaxTtl` instead) The maximum allowed lifetime of tokens
	// issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl interface{}
	// The organization configured users must be part of.
	Organization interface{}
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path interface{}
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs interface{}
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl interface{}
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl interface{}
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy interface{}
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses interface{}
	// Generated Token's Period
	TokenPeriod interface{}
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies interface{}
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl interface{}
	// (Optional) The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType interface{}
	// (Optional; Deprecated, use `tokenTtl` isntead) The TTL period of tokens issued
	// using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
	Ttl interface{}
	Tune interface{}
}
