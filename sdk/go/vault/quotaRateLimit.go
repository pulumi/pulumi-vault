// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
// A rate limit quota can be created at the root level or defined on a namespace or mount by
// specifying a path when creating the quota.
//
// See [Vault's Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewQuotaRateLimit(ctx, "global", &vault.QuotaRateLimitArgs{
//				Name: pulumi.String("global"),
//				Path: pulumi.String(""),
//				Rate: pulumi.Float64(100),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Rate limit quotas can be imported using their names
//
// ```sh
// $ pulumi import vault:index/quotaRateLimit:QuotaRateLimit global global
// ```
type QuotaRateLimit struct {
	pulumi.CustomResourceState

	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval pulumi.IntPtrOutput `pulumi:"blockInterval"`
	// Attribute used to group requests for rate limiting. Limits are enforced independently for each
	// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
	// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
	// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
	// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
	// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
	// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	GroupBy pulumi.StringOutput `pulumi:"groupBy"`
	// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
	Inheritable pulumi.BoolPtrOutput `pulumi:"inheritable"`
	// The duration in seconds to enforce rate limiting for.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Name of the rate limit quota
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The `rate` must be positive.
	Rate pulumi.Float64Output `pulumi:"rate"`
	// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
	// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
	// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
	SecondaryRate pulumi.Float64Output `pulumi:"secondaryRate"`
}

// NewQuotaRateLimit registers a new resource with the given unique name, arguments, and options.
func NewQuotaRateLimit(ctx *pulumi.Context,
	name string, args *QuotaRateLimitArgs, opts ...pulumi.ResourceOption) (*QuotaRateLimit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rate == nil {
		return nil, errors.New("invalid value for required argument 'Rate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QuotaRateLimit
	err := ctx.RegisterResource("vault:index/quotaRateLimit:QuotaRateLimit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaRateLimit gets an existing QuotaRateLimit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaRateLimit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaRateLimitState, opts ...pulumi.ResourceOption) (*QuotaRateLimit, error) {
	var resource QuotaRateLimit
	err := ctx.ReadResource("vault:index/quotaRateLimit:QuotaRateLimit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaRateLimit resources.
type quotaRateLimitState struct {
	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval *int `pulumi:"blockInterval"`
	// Attribute used to group requests for rate limiting. Limits are enforced independently for each
	// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
	// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
	// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
	// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
	// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
	// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	GroupBy *string `pulumi:"groupBy"`
	// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
	Inheritable *bool `pulumi:"inheritable"`
	// The duration in seconds to enforce rate limiting for.
	Interval *int `pulumi:"interval"`
	// Name of the rate limit quota
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path *string `pulumi:"path"`
	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The `rate` must be positive.
	Rate *float64 `pulumi:"rate"`
	// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role *string `pulumi:"role"`
	// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
	// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
	// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
	SecondaryRate *float64 `pulumi:"secondaryRate"`
}

type QuotaRateLimitState struct {
	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval pulumi.IntPtrInput
	// Attribute used to group requests for rate limiting. Limits are enforced independently for each
	// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
	// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
	// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
	// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
	// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
	// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	GroupBy pulumi.StringPtrInput
	// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
	Inheritable pulumi.BoolPtrInput
	// The duration in seconds to enforce rate limiting for.
	Interval pulumi.IntPtrInput
	// Name of the rate limit quota
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrInput
	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The `rate` must be positive.
	Rate pulumi.Float64PtrInput
	// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role pulumi.StringPtrInput
	// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
	// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
	// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
	SecondaryRate pulumi.Float64PtrInput
}

func (QuotaRateLimitState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaRateLimitState)(nil)).Elem()
}

type quotaRateLimitArgs struct {
	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval *int `pulumi:"blockInterval"`
	// Attribute used to group requests for rate limiting. Limits are enforced independently for each
	// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
	// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
	// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
	// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
	// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
	// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	GroupBy *string `pulumi:"groupBy"`
	// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
	Inheritable *bool `pulumi:"inheritable"`
	// The duration in seconds to enforce rate limiting for.
	Interval *int `pulumi:"interval"`
	// Name of the rate limit quota
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path *string `pulumi:"path"`
	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The `rate` must be positive.
	Rate float64 `pulumi:"rate"`
	// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role *string `pulumi:"role"`
	// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
	// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
	// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
	SecondaryRate *float64 `pulumi:"secondaryRate"`
}

// The set of arguments for constructing a QuotaRateLimit resource.
type QuotaRateLimitArgs struct {
	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval pulumi.IntPtrInput
	// Attribute used to group requests for rate limiting. Limits are enforced independently for each
	// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
	// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
	// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
	// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
	// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
	// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	GroupBy pulumi.StringPtrInput
	// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
	Inheritable pulumi.BoolPtrInput
	// The duration in seconds to enforce rate limiting for.
	Interval pulumi.IntPtrInput
	// Name of the rate limit quota
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrInput
	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The `rate` must be positive.
	Rate pulumi.Float64Input
	// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role pulumi.StringPtrInput
	// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
	// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
	// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
	SecondaryRate pulumi.Float64PtrInput
}

func (QuotaRateLimitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaRateLimitArgs)(nil)).Elem()
}

type QuotaRateLimitInput interface {
	pulumi.Input

	ToQuotaRateLimitOutput() QuotaRateLimitOutput
	ToQuotaRateLimitOutputWithContext(ctx context.Context) QuotaRateLimitOutput
}

func (*QuotaRateLimit) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaRateLimit)(nil)).Elem()
}

func (i *QuotaRateLimit) ToQuotaRateLimitOutput() QuotaRateLimitOutput {
	return i.ToQuotaRateLimitOutputWithContext(context.Background())
}

func (i *QuotaRateLimit) ToQuotaRateLimitOutputWithContext(ctx context.Context) QuotaRateLimitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaRateLimitOutput)
}

// QuotaRateLimitArrayInput is an input type that accepts QuotaRateLimitArray and QuotaRateLimitArrayOutput values.
// You can construct a concrete instance of `QuotaRateLimitArrayInput` via:
//
//	QuotaRateLimitArray{ QuotaRateLimitArgs{...} }
type QuotaRateLimitArrayInput interface {
	pulumi.Input

	ToQuotaRateLimitArrayOutput() QuotaRateLimitArrayOutput
	ToQuotaRateLimitArrayOutputWithContext(context.Context) QuotaRateLimitArrayOutput
}

type QuotaRateLimitArray []QuotaRateLimitInput

func (QuotaRateLimitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaRateLimit)(nil)).Elem()
}

func (i QuotaRateLimitArray) ToQuotaRateLimitArrayOutput() QuotaRateLimitArrayOutput {
	return i.ToQuotaRateLimitArrayOutputWithContext(context.Background())
}

func (i QuotaRateLimitArray) ToQuotaRateLimitArrayOutputWithContext(ctx context.Context) QuotaRateLimitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaRateLimitArrayOutput)
}

// QuotaRateLimitMapInput is an input type that accepts QuotaRateLimitMap and QuotaRateLimitMapOutput values.
// You can construct a concrete instance of `QuotaRateLimitMapInput` via:
//
//	QuotaRateLimitMap{ "key": QuotaRateLimitArgs{...} }
type QuotaRateLimitMapInput interface {
	pulumi.Input

	ToQuotaRateLimitMapOutput() QuotaRateLimitMapOutput
	ToQuotaRateLimitMapOutputWithContext(context.Context) QuotaRateLimitMapOutput
}

type QuotaRateLimitMap map[string]QuotaRateLimitInput

func (QuotaRateLimitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaRateLimit)(nil)).Elem()
}

func (i QuotaRateLimitMap) ToQuotaRateLimitMapOutput() QuotaRateLimitMapOutput {
	return i.ToQuotaRateLimitMapOutputWithContext(context.Background())
}

func (i QuotaRateLimitMap) ToQuotaRateLimitMapOutputWithContext(ctx context.Context) QuotaRateLimitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaRateLimitMapOutput)
}

type QuotaRateLimitOutput struct{ *pulumi.OutputState }

func (QuotaRateLimitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaRateLimit)(nil)).Elem()
}

func (o QuotaRateLimitOutput) ToQuotaRateLimitOutput() QuotaRateLimitOutput {
	return o
}

func (o QuotaRateLimitOutput) ToQuotaRateLimitOutputWithContext(ctx context.Context) QuotaRateLimitOutput {
	return o
}

// If set, when a client reaches a rate limit threshold, the client will
// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
func (o QuotaRateLimitOutput) BlockInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.IntPtrOutput { return v.BlockInterval }).(pulumi.IntPtrOutput)
}

// Attribute used to group requests for rate limiting. Limits are enforced independently for each
// group. Valid `groupBy` modes are: 1) `ip` that groups requests by their source IP address (**`groupBy` defaults to
// `ip` if unset, which is the only supported mode in community edition**); 2) `none` that groups together all requests
// that match the rate limit quota rule; 3) `entityThenIp` that groups requests by their entity ID for authenticated
// requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not
// connected to an entity); and 4) `entityThenNone` which also groups requests by their entity ID when available, but
// the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
func (o QuotaRateLimitOutput) GroupBy() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.StringOutput { return v.GroupBy }).(pulumi.StringOutput)
}

// If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
func (o QuotaRateLimitOutput) Inheritable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.BoolPtrOutput { return v.Inheritable }).(pulumi.BoolPtrOutput)
}

// The duration in seconds to enforce rate limiting for.
func (o QuotaRateLimitOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

// Name of the rate limit quota
func (o QuotaRateLimitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured namespace.
// *Available only for Vault Enterprise*.
func (o QuotaRateLimitOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Path of the mount or namespace to apply the quota. A blank path configures a
// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
// Updating this field on an existing quota can have "moving" effects. For example, updating
// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
func (o QuotaRateLimitOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The maximum number of requests at any given second to be allowed by the quota
// rule. The `rate` must be positive.
func (o QuotaRateLimitOutput) Rate() pulumi.Float64Output {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.Float64Output { return v.Rate }).(pulumi.Float64Output)
}

// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
func (o QuotaRateLimitOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// Can only be set for the `groupBy` modes `entityThenIp` or `entityThenNone`. This is
// the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests
// that contain an entity ID are subject to the `rate` field instead. Defaults to the same value as `rate`.
func (o QuotaRateLimitOutput) SecondaryRate() pulumi.Float64Output {
	return o.ApplyT(func(v *QuotaRateLimit) pulumi.Float64Output { return v.SecondaryRate }).(pulumi.Float64Output)
}

type QuotaRateLimitArrayOutput struct{ *pulumi.OutputState }

func (QuotaRateLimitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaRateLimit)(nil)).Elem()
}

func (o QuotaRateLimitArrayOutput) ToQuotaRateLimitArrayOutput() QuotaRateLimitArrayOutput {
	return o
}

func (o QuotaRateLimitArrayOutput) ToQuotaRateLimitArrayOutputWithContext(ctx context.Context) QuotaRateLimitArrayOutput {
	return o
}

func (o QuotaRateLimitArrayOutput) Index(i pulumi.IntInput) QuotaRateLimitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuotaRateLimit {
		return vs[0].([]*QuotaRateLimit)[vs[1].(int)]
	}).(QuotaRateLimitOutput)
}

type QuotaRateLimitMapOutput struct{ *pulumi.OutputState }

func (QuotaRateLimitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaRateLimit)(nil)).Elem()
}

func (o QuotaRateLimitMapOutput) ToQuotaRateLimitMapOutput() QuotaRateLimitMapOutput {
	return o
}

func (o QuotaRateLimitMapOutput) ToQuotaRateLimitMapOutputWithContext(ctx context.Context) QuotaRateLimitMapOutput {
	return o
}

func (o QuotaRateLimitMapOutput) MapIndex(k pulumi.StringInput) QuotaRateLimitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuotaRateLimit {
		return vs[0].(map[string]*QuotaRateLimit)[vs[1].(string)]
	}).(QuotaRateLimitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaRateLimitInput)(nil)).Elem(), &QuotaRateLimit{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaRateLimitArrayInput)(nil)).Elem(), QuotaRateLimitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaRateLimitMapInput)(nil)).Elem(), QuotaRateLimitMap{})
	pulumi.RegisterOutputType(QuotaRateLimitOutput{})
	pulumi.RegisterOutputType(QuotaRateLimitArrayOutput{})
	pulumi.RegisterOutputType(QuotaRateLimitMapOutput{})
}
