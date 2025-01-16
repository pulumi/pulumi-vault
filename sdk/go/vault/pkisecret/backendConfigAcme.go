// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows setting the ACME server configuration used by specified mount.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkisecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pki, err := vault.NewMount(ctx, "pki", &vault.MountArgs{
//				Path:                   pulumi.String("pki"),
//				Type:                   pulumi.String("pki"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkisecret.NewBackendConfigCluster(ctx, "pki_config_cluster", &pkisecret.BackendConfigClusterArgs{
//				Backend: pki.Path,
//				Path:    pulumi.String("http://127.0.0.1:8200/v1/pki"),
//				AiaPath: pulumi.String("http://127.0.0.1:8200/v1/pki"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkisecret.NewBackendConfigAcme(ctx, "example", &pkisecret.BackendConfigAcmeArgs{
//				Backend: pki.Path,
//				Enabled: pulumi.Bool(true),
//				AllowedIssuers: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				AllowedRoles: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				AllowRoleExtKeyUsage:   pulumi.Bool(false),
//				DefaultDirectoryPolicy: pulumi.String("sign-verbatim"),
//				DnsResolver:            pulumi.String(""),
//				EabPolicy:              pulumi.String("not-required"),
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
// The ACME configuration can be imported using the resource's `id`.
// In the case of the example above the `id` would be `pki/config/acme`,
// where the `pki` component is the resource's `backend`, e.g.
//
// ```sh
// $ pulumi import vault:pkiSecret/backendConfigAcme:BackendConfigAcme example pki/config/acme
// ```
type BackendConfigAcme struct {
	pulumi.CustomResourceState

	// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
	AllowRoleExtKeyUsage pulumi.BoolPtrOutput `pulumi:"allowRoleExtKeyUsage"`
	// Specifies which issuers are allowed for use with ACME.
	AllowedIssuers pulumi.StringArrayOutput `pulumi:"allowedIssuers"`
	// Specifies which roles are allowed for use with ACME.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the policy to be used for non-role-qualified ACME requests.
	// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
	DefaultDirectoryPolicy pulumi.StringOutput `pulumi:"defaultDirectoryPolicy"`
	// DNS resolver to use for domain resolution on this mount.
	// Must be in the format `<host>:<port>`, with both parts mandatory.
	DnsResolver pulumi.StringPtrOutput `pulumi:"dnsResolver"`
	// Specifies the policy to use for external account binding behaviour.
	// Allowed values are `not-required`, `new-account-required` or `always-required`.
	EabPolicy pulumi.StringOutput `pulumi:"eabPolicy"`
	// Specifies whether ACME is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewBackendConfigAcme registers a new resource with the given unique name, arguments, and options.
func NewBackendConfigAcme(ctx *pulumi.Context,
	name string, args *BackendConfigAcmeArgs, opts ...pulumi.ResourceOption) (*BackendConfigAcme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackendConfigAcme
	err := ctx.RegisterResource("vault:pkiSecret/backendConfigAcme:BackendConfigAcme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendConfigAcme gets an existing BackendConfigAcme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendConfigAcme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendConfigAcmeState, opts ...pulumi.ResourceOption) (*BackendConfigAcme, error) {
	var resource BackendConfigAcme
	err := ctx.ReadResource("vault:pkiSecret/backendConfigAcme:BackendConfigAcme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendConfigAcme resources.
type backendConfigAcmeState struct {
	// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
	AllowRoleExtKeyUsage *bool `pulumi:"allowRoleExtKeyUsage"`
	// Specifies which issuers are allowed for use with ACME.
	AllowedIssuers []string `pulumi:"allowedIssuers"`
	// Specifies which roles are allowed for use with ACME.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the policy to be used for non-role-qualified ACME requests.
	// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
	DefaultDirectoryPolicy *string `pulumi:"defaultDirectoryPolicy"`
	// DNS resolver to use for domain resolution on this mount.
	// Must be in the format `<host>:<port>`, with both parts mandatory.
	DnsResolver *string `pulumi:"dnsResolver"`
	// Specifies the policy to use for external account binding behaviour.
	// Allowed values are `not-required`, `new-account-required` or `always-required`.
	EabPolicy *string `pulumi:"eabPolicy"`
	// Specifies whether ACME is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type BackendConfigAcmeState struct {
	// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
	AllowRoleExtKeyUsage pulumi.BoolPtrInput
	// Specifies which issuers are allowed for use with ACME.
	AllowedIssuers pulumi.StringArrayInput
	// Specifies which roles are allowed for use with ACME.
	AllowedRoles pulumi.StringArrayInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the policy to be used for non-role-qualified ACME requests.
	// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
	DefaultDirectoryPolicy pulumi.StringPtrInput
	// DNS resolver to use for domain resolution on this mount.
	// Must be in the format `<host>:<port>`, with both parts mandatory.
	DnsResolver pulumi.StringPtrInput
	// Specifies the policy to use for external account binding behaviour.
	// Allowed values are `not-required`, `new-account-required` or `always-required`.
	EabPolicy pulumi.StringPtrInput
	// Specifies whether ACME is enabled.
	Enabled pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (BackendConfigAcmeState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendConfigAcmeState)(nil)).Elem()
}

type backendConfigAcmeArgs struct {
	// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
	AllowRoleExtKeyUsage *bool `pulumi:"allowRoleExtKeyUsage"`
	// Specifies which issuers are allowed for use with ACME.
	AllowedIssuers []string `pulumi:"allowedIssuers"`
	// Specifies which roles are allowed for use with ACME.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the policy to be used for non-role-qualified ACME requests.
	// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
	DefaultDirectoryPolicy *string `pulumi:"defaultDirectoryPolicy"`
	// DNS resolver to use for domain resolution on this mount.
	// Must be in the format `<host>:<port>`, with both parts mandatory.
	DnsResolver *string `pulumi:"dnsResolver"`
	// Specifies the policy to use for external account binding behaviour.
	// Allowed values are `not-required`, `new-account-required` or `always-required`.
	EabPolicy *string `pulumi:"eabPolicy"`
	// Specifies whether ACME is enabled.
	Enabled bool `pulumi:"enabled"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a BackendConfigAcme resource.
type BackendConfigAcmeArgs struct {
	// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
	AllowRoleExtKeyUsage pulumi.BoolPtrInput
	// Specifies which issuers are allowed for use with ACME.
	AllowedIssuers pulumi.StringArrayInput
	// Specifies which roles are allowed for use with ACME.
	AllowedRoles pulumi.StringArrayInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the policy to be used for non-role-qualified ACME requests.
	// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
	DefaultDirectoryPolicy pulumi.StringPtrInput
	// DNS resolver to use for domain resolution on this mount.
	// Must be in the format `<host>:<port>`, with both parts mandatory.
	DnsResolver pulumi.StringPtrInput
	// Specifies the policy to use for external account binding behaviour.
	// Allowed values are `not-required`, `new-account-required` or `always-required`.
	EabPolicy pulumi.StringPtrInput
	// Specifies whether ACME is enabled.
	Enabled pulumi.BoolInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (BackendConfigAcmeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendConfigAcmeArgs)(nil)).Elem()
}

type BackendConfigAcmeInput interface {
	pulumi.Input

	ToBackendConfigAcmeOutput() BackendConfigAcmeOutput
	ToBackendConfigAcmeOutputWithContext(ctx context.Context) BackendConfigAcmeOutput
}

func (*BackendConfigAcme) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendConfigAcme)(nil)).Elem()
}

func (i *BackendConfigAcme) ToBackendConfigAcmeOutput() BackendConfigAcmeOutput {
	return i.ToBackendConfigAcmeOutputWithContext(context.Background())
}

func (i *BackendConfigAcme) ToBackendConfigAcmeOutputWithContext(ctx context.Context) BackendConfigAcmeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendConfigAcmeOutput)
}

// BackendConfigAcmeArrayInput is an input type that accepts BackendConfigAcmeArray and BackendConfigAcmeArrayOutput values.
// You can construct a concrete instance of `BackendConfigAcmeArrayInput` via:
//
//	BackendConfigAcmeArray{ BackendConfigAcmeArgs{...} }
type BackendConfigAcmeArrayInput interface {
	pulumi.Input

	ToBackendConfigAcmeArrayOutput() BackendConfigAcmeArrayOutput
	ToBackendConfigAcmeArrayOutputWithContext(context.Context) BackendConfigAcmeArrayOutput
}

type BackendConfigAcmeArray []BackendConfigAcmeInput

func (BackendConfigAcmeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendConfigAcme)(nil)).Elem()
}

func (i BackendConfigAcmeArray) ToBackendConfigAcmeArrayOutput() BackendConfigAcmeArrayOutput {
	return i.ToBackendConfigAcmeArrayOutputWithContext(context.Background())
}

func (i BackendConfigAcmeArray) ToBackendConfigAcmeArrayOutputWithContext(ctx context.Context) BackendConfigAcmeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendConfigAcmeArrayOutput)
}

// BackendConfigAcmeMapInput is an input type that accepts BackendConfigAcmeMap and BackendConfigAcmeMapOutput values.
// You can construct a concrete instance of `BackendConfigAcmeMapInput` via:
//
//	BackendConfigAcmeMap{ "key": BackendConfigAcmeArgs{...} }
type BackendConfigAcmeMapInput interface {
	pulumi.Input

	ToBackendConfigAcmeMapOutput() BackendConfigAcmeMapOutput
	ToBackendConfigAcmeMapOutputWithContext(context.Context) BackendConfigAcmeMapOutput
}

type BackendConfigAcmeMap map[string]BackendConfigAcmeInput

func (BackendConfigAcmeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendConfigAcme)(nil)).Elem()
}

func (i BackendConfigAcmeMap) ToBackendConfigAcmeMapOutput() BackendConfigAcmeMapOutput {
	return i.ToBackendConfigAcmeMapOutputWithContext(context.Background())
}

func (i BackendConfigAcmeMap) ToBackendConfigAcmeMapOutputWithContext(ctx context.Context) BackendConfigAcmeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendConfigAcmeMapOutput)
}

type BackendConfigAcmeOutput struct{ *pulumi.OutputState }

func (BackendConfigAcmeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendConfigAcme)(nil)).Elem()
}

func (o BackendConfigAcmeOutput) ToBackendConfigAcmeOutput() BackendConfigAcmeOutput {
	return o
}

func (o BackendConfigAcmeOutput) ToBackendConfigAcmeOutputWithContext(ctx context.Context) BackendConfigAcmeOutput {
	return o
}

// Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
func (o BackendConfigAcmeOutput) AllowRoleExtKeyUsage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.BoolPtrOutput { return v.AllowRoleExtKeyUsage }).(pulumi.BoolPtrOutput)
}

// Specifies which issuers are allowed for use with ACME.
func (o BackendConfigAcmeOutput) AllowedIssuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringArrayOutput { return v.AllowedIssuers }).(pulumi.StringArrayOutput)
}

// Specifies which roles are allowed for use with ACME.
func (o BackendConfigAcmeOutput) AllowedRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringArrayOutput { return v.AllowedRoles }).(pulumi.StringArrayOutput)
}

// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
func (o BackendConfigAcmeOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Specifies the policy to be used for non-role-qualified ACME requests.
// Allowed values are `forbid`, `sign-verbatim`, `role:<role_name>`, `external-policy` or `external-policy:<policy>`.
func (o BackendConfigAcmeOutput) DefaultDirectoryPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringOutput { return v.DefaultDirectoryPolicy }).(pulumi.StringOutput)
}

// DNS resolver to use for domain resolution on this mount.
// Must be in the format `<host>:<port>`, with both parts mandatory.
func (o BackendConfigAcmeOutput) DnsResolver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringPtrOutput { return v.DnsResolver }).(pulumi.StringPtrOutput)
}

// Specifies the policy to use for external account binding behaviour.
// Allowed values are `not-required`, `new-account-required` or `always-required`.
func (o BackendConfigAcmeOutput) EabPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringOutput { return v.EabPolicy }).(pulumi.StringOutput)
}

// Specifies whether ACME is enabled.
func (o BackendConfigAcmeOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o BackendConfigAcmeOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendConfigAcme) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type BackendConfigAcmeArrayOutput struct{ *pulumi.OutputState }

func (BackendConfigAcmeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendConfigAcme)(nil)).Elem()
}

func (o BackendConfigAcmeArrayOutput) ToBackendConfigAcmeArrayOutput() BackendConfigAcmeArrayOutput {
	return o
}

func (o BackendConfigAcmeArrayOutput) ToBackendConfigAcmeArrayOutputWithContext(ctx context.Context) BackendConfigAcmeArrayOutput {
	return o
}

func (o BackendConfigAcmeArrayOutput) Index(i pulumi.IntInput) BackendConfigAcmeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendConfigAcme {
		return vs[0].([]*BackendConfigAcme)[vs[1].(int)]
	}).(BackendConfigAcmeOutput)
}

type BackendConfigAcmeMapOutput struct{ *pulumi.OutputState }

func (BackendConfigAcmeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendConfigAcme)(nil)).Elem()
}

func (o BackendConfigAcmeMapOutput) ToBackendConfigAcmeMapOutput() BackendConfigAcmeMapOutput {
	return o
}

func (o BackendConfigAcmeMapOutput) ToBackendConfigAcmeMapOutputWithContext(ctx context.Context) BackendConfigAcmeMapOutput {
	return o
}

func (o BackendConfigAcmeMapOutput) MapIndex(k pulumi.StringInput) BackendConfigAcmeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendConfigAcme {
		return vs[0].(map[string]*BackendConfigAcme)[vs[1].(string)]
	}).(BackendConfigAcmeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendConfigAcmeInput)(nil)).Elem(), &BackendConfigAcme{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendConfigAcmeArrayInput)(nil)).Elem(), BackendConfigAcmeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendConfigAcmeMapInput)(nil)).Elem(), BackendConfigAcmeMap{})
	pulumi.RegisterOutputType(BackendConfigAcmeOutput{})
	pulumi.RegisterOutputType(BackendConfigAcmeArrayOutput{})
	pulumi.RegisterOutputType(BackendConfigAcmeMapOutput{})
}
