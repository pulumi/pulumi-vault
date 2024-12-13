// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the Role of an Kubernetes from a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
// information.
func LookupAuthBackendConfig(ctx *pulumi.Context, args *LookupAuthBackendConfigArgs, opts ...pulumi.InvokeOption) (*LookupAuthBackendConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthBackendConfigResult
	err := ctx.Invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackendConfig.
type LookupAuthBackendConfigArgs struct {
	// The unique name for the Kubernetes backend the config to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend *string `pulumi:"backend"`
	// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation *bool `pulumi:"disableIssValidation"`
	// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt *bool `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `pulumi:"kubernetesHost"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
	UseAnnotationsAsAliasMetadata *bool `pulumi:"useAnnotationsAsAliasMetadata"`
}

// A collection of values returned by getAuthBackendConfig.
type LookupAuthBackendConfigResult struct {
	Backend *string `pulumi:"backend"`
	// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation bool `pulumi:"disableIssValidation"`
	// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt bool `pulumi:"disableLocalCaJwt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string  `pulumi:"kubernetesHost"`
	Namespace      *string `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
	UseAnnotationsAsAliasMetadata bool `pulumi:"useAnnotationsAsAliasMetadata"`
}

func LookupAuthBackendConfigOutput(ctx *pulumi.Context, args LookupAuthBackendConfigOutputArgs, opts ...pulumi.InvokeOption) LookupAuthBackendConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAuthBackendConfigResultOutput, error) {
			args := v.(LookupAuthBackendConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args, LookupAuthBackendConfigResultOutput{}, options).(LookupAuthBackendConfigResultOutput), nil
		}).(LookupAuthBackendConfigResultOutput)
}

// A collection of arguments for invoking getAuthBackendConfig.
type LookupAuthBackendConfigOutputArgs struct {
	// The unique name for the Kubernetes backend the config to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend pulumi.StringPtrInput `pulumi:"backend"`
	// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation pulumi.BoolPtrInput `pulumi:"disableIssValidation"`
	// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt pulumi.BoolPtrInput `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer pulumi.StringPtrInput `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringPtrInput `pulumi:"kubernetesHost"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput `pulumi:"pemKeys"`
	// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
	UseAnnotationsAsAliasMetadata pulumi.BoolPtrInput `pulumi:"useAnnotationsAsAliasMetadata"`
}

func (LookupAuthBackendConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendConfigArgs)(nil)).Elem()
}

// A collection of values returned by getAuthBackendConfig.
type LookupAuthBackendConfigResultOutput struct{ *pulumi.OutputState }

func (LookupAuthBackendConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendConfigResult)(nil)).Elem()
}

func (o LookupAuthBackendConfigResultOutput) ToLookupAuthBackendConfigResultOutput() LookupAuthBackendConfigResultOutput {
	return o
}

func (o LookupAuthBackendConfigResultOutput) ToLookupAuthBackendConfigResultOutputWithContext(ctx context.Context) LookupAuthBackendConfigResultOutput {
	return o
}

func (o LookupAuthBackendConfigResultOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) *string { return v.Backend }).(pulumi.StringPtrOutput)
}

// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
func (o LookupAuthBackendConfigResultOutput) DisableIssValidation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) bool { return v.DisableIssValidation }).(pulumi.BoolOutput)
}

// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
func (o LookupAuthBackendConfigResultOutput) DisableLocalCaJwt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) bool { return v.DisableLocalCaJwt }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAuthBackendConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
func (o LookupAuthBackendConfigResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
func (o LookupAuthBackendConfigResultOutput) KubernetesCaCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.KubernetesCaCert }).(pulumi.StringOutput)
}

// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
func (o LookupAuthBackendConfigResultOutput) KubernetesHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.KubernetesHost }).(pulumi.StringOutput)
}

func (o LookupAuthBackendConfigResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
func (o LookupAuthBackendConfigResultOutput) PemKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) []string { return v.PemKeys }).(pulumi.StringArrayOutput)
}

// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
func (o LookupAuthBackendConfigResultOutput) UseAnnotationsAsAliasMetadata() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) bool { return v.UseAnnotationsAsAliasMetadata }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthBackendConfigResultOutput{})
}
