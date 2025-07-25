// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/kubernetes"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "/path/to/cert",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			invokeFile1, err := std.File(ctx, &std.FileArgs{
//				Input: "/path/to/token",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			config, err := kubernetes.NewSecretBackend(ctx, "config", &kubernetes.SecretBackendArgs{
//				Path:              pulumi.String("kubernetes"),
//				Description:       pulumi.String("kubernetes secrets engine description"),
//				KubernetesHost:    pulumi.String("https://127.0.0.1:61233"),
//				KubernetesCaCert:  pulumi.String(invokeFile.Result),
//				ServiceAccountJwt: pulumi.String(invokeFile1.Result),
//				DisableLocalCaJwt: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			role, err := kubernetes.NewSecretBackendRole(ctx, "role", &kubernetes.SecretBackendRoleArgs{
//				Backend: config.Path,
//				Name:    pulumi.String("service-account-name-role"),
//				AllowedKubernetesNamespaces: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				TokenMaxTtl:        pulumi.Int(43200),
//				TokenDefaultTtl:    pulumi.Int(21600),
//				ServiceAccountName: pulumi.String("test-service-account-with-generated-token"),
//				ExtraLabels: pulumi.StringMap{
//					"id":   pulumi.String("abc123"),
//					"name": pulumi.String("some_name"),
//				},
//				ExtraAnnotations: pulumi.StringMap{
//					"env":      pulumi.String("development"),
//					"location": pulumi.String("earth"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = kubernetes.GetServiceAccountTokenOutput(ctx, kubernetes.GetServiceAccountTokenOutputArgs{
//				Backend:             config.Path,
//				Role:                role.Name,
//				KubernetesNamespace: pulumi.String("test"),
//				ClusterRoleBinding:  pulumi.Bool(false),
//				Ttl:                 pulumi.String("1h"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetServiceAccountToken(ctx *pulumi.Context, args *GetServiceAccountTokenArgs, opts ...pulumi.InvokeOption) (*GetServiceAccountTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceAccountTokenResult
	err := ctx.Invoke("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccountToken.
type GetServiceAccountTokenArgs struct {
	// The Kubernetes secret backend to generate service account
	// tokens from.
	Backend string `pulumi:"backend"`
	// If true, generate a ClusterRoleBinding to grant
	// permissions across the whole cluster instead of within a namespace.
	ClusterRoleBinding *bool `pulumi:"clusterRoleBinding"`
	// The name of the Kubernetes namespace in which to
	// generate the credentials.
	KubernetesNamespace string `pulumi:"kubernetesNamespace"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the Kubernetes secret backend role to generate service
	// account tokens from.
	Role string `pulumi:"role"`
	// The TTL of the generated Kubernetes service account token, specified in
	// seconds or as a Go duration format string.
	Ttl *string `pulumi:"ttl"`
}

// A collection of values returned by getServiceAccountToken.
type GetServiceAccountTokenResult struct {
	Backend            string `pulumi:"backend"`
	ClusterRoleBinding *bool  `pulumi:"clusterRoleBinding"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	KubernetesNamespace string `pulumi:"kubernetesNamespace"`
	// The duration of the lease in seconds.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault.
	LeaseId string `pulumi:"leaseId"`
	// True if the duration of this lease can be extended through renewal.
	LeaseRenewable bool    `pulumi:"leaseRenewable"`
	Namespace      *string `pulumi:"namespace"`
	Role           string  `pulumi:"role"`
	// The name of the service account associated with the token.
	ServiceAccountName string `pulumi:"serviceAccountName"`
	// The Kubernetes namespace that the service account resides in.
	ServiceAccountNamespace string `pulumi:"serviceAccountNamespace"`
	// The Kubernetes service account token.
	ServiceAccountToken string  `pulumi:"serviceAccountToken"`
	Ttl                 *string `pulumi:"ttl"`
}

func GetServiceAccountTokenOutput(ctx *pulumi.Context, args GetServiceAccountTokenOutputArgs, opts ...pulumi.InvokeOption) GetServiceAccountTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServiceAccountTokenResultOutput, error) {
			args := v.(GetServiceAccountTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", args, GetServiceAccountTokenResultOutput{}, options).(GetServiceAccountTokenResultOutput), nil
		}).(GetServiceAccountTokenResultOutput)
}

// A collection of arguments for invoking getServiceAccountToken.
type GetServiceAccountTokenOutputArgs struct {
	// The Kubernetes secret backend to generate service account
	// tokens from.
	Backend pulumi.StringInput `pulumi:"backend"`
	// If true, generate a ClusterRoleBinding to grant
	// permissions across the whole cluster instead of within a namespace.
	ClusterRoleBinding pulumi.BoolPtrInput `pulumi:"clusterRoleBinding"`
	// The name of the Kubernetes namespace in which to
	// generate the credentials.
	KubernetesNamespace pulumi.StringInput `pulumi:"kubernetesNamespace"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// The name of the Kubernetes secret backend role to generate service
	// account tokens from.
	Role pulumi.StringInput `pulumi:"role"`
	// The TTL of the generated Kubernetes service account token, specified in
	// seconds or as a Go duration format string.
	Ttl pulumi.StringPtrInput `pulumi:"ttl"`
}

func (GetServiceAccountTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountTokenArgs)(nil)).Elem()
}

// A collection of values returned by getServiceAccountToken.
type GetServiceAccountTokenResultOutput struct{ *pulumi.OutputState }

func (GetServiceAccountTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountTokenResult)(nil)).Elem()
}

func (o GetServiceAccountTokenResultOutput) ToGetServiceAccountTokenResultOutput() GetServiceAccountTokenResultOutput {
	return o
}

func (o GetServiceAccountTokenResultOutput) ToGetServiceAccountTokenResultOutputWithContext(ctx context.Context) GetServiceAccountTokenResultOutput {
	return o
}

func (o GetServiceAccountTokenResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.Backend }).(pulumi.StringOutput)
}

func (o GetServiceAccountTokenResultOutput) ClusterRoleBinding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) *bool { return v.ClusterRoleBinding }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceAccountTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceAccountTokenResultOutput) KubernetesNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.KubernetesNamespace }).(pulumi.StringOutput)
}

// The duration of the lease in seconds.
func (o GetServiceAccountTokenResultOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) int { return v.LeaseDuration }).(pulumi.IntOutput)
}

// The lease identifier assigned by Vault.
func (o GetServiceAccountTokenResultOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.LeaseId }).(pulumi.StringOutput)
}

// True if the duration of this lease can be extended through renewal.
func (o GetServiceAccountTokenResultOutput) LeaseRenewable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) bool { return v.LeaseRenewable }).(pulumi.BoolOutput)
}

func (o GetServiceAccountTokenResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetServiceAccountTokenResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.Role }).(pulumi.StringOutput)
}

// The name of the service account associated with the token.
func (o GetServiceAccountTokenResultOutput) ServiceAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.ServiceAccountName }).(pulumi.StringOutput)
}

// The Kubernetes namespace that the service account resides in.
func (o GetServiceAccountTokenResultOutput) ServiceAccountNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.ServiceAccountNamespace }).(pulumi.StringOutput)
}

// The Kubernetes service account token.
func (o GetServiceAccountTokenResultOutput) ServiceAccountToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) string { return v.ServiceAccountToken }).(pulumi.StringOutput)
}

func (o GetServiceAccountTokenResultOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceAccountTokenResult) *string { return v.Ttl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceAccountTokenResultOutput{})
}
