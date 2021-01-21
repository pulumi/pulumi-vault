// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Kubernetes auth backend config in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/kubernetes"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		kubernetes, err := vault.NewAuthBackend(ctx, "kubernetes", &vault.AuthBackendArgs{
// 			Type: pulumi.String("kubernetes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kubernetes.NewAuthBackendConfig(ctx, "example", &kubernetes.AuthBackendConfigArgs{
// 			Backend:              kubernetes.Path,
// 			DisableIssValidation: pulumi.Bool(true),
// 			Issuer:               pulumi.String("api"),
// 			KubernetesCaCert:     pulumi.String(fmt.Sprintf("%v%v%v", "-----BEGIN CERTIFICATE-----\n", "example\n", "-----END CERTIFICATE-----\n")),
// 			KubernetesHost:       pulumi.String("http://example.com:443"),
// 			TokenReviewerJwt:     pulumi.String("ZXhhbXBsZQo="),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Kubernetes authentication backend can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:kubernetes/authBackendConfig:AuthBackendConfig config auth/kubernetes/config
// ```
type AuthBackendConfig struct {
	pulumi.CustomResourceState

	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation pulumi.BoolOutput `pulumi:"disableIssValidation"`
	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt pulumi.BoolOutput `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer pulumi.StringPtrOutput `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrOutput `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringOutput `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayOutput `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrOutput `pulumi:"tokenReviewerJwt"`
}

// NewAuthBackendConfig registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendConfig(ctx *pulumi.Context,
	name string, args *AuthBackendConfigArgs, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesHost == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesHost'")
	}
	var resource AuthBackendConfig
	err := ctx.RegisterResource("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendConfig gets an existing AuthBackendConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendConfigState, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	var resource AuthBackendConfig
	err := ctx.ReadResource("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendConfig resources.
type authBackendConfigState struct {
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation *bool `pulumi:"disableIssValidation"`
	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt *bool `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt *string `pulumi:"tokenReviewerJwt"`
}

type AuthBackendConfigState struct {
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation pulumi.BoolPtrInput
	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt pulumi.BoolPtrInput
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer pulumi.StringPtrInput
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringPtrInput
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrInput
}

func (AuthBackendConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigState)(nil)).Elem()
}

type authBackendConfigArgs struct {
	// Unique name of the kubernetes backend to configure.
	Backend *string `pulumi:"backend"`
	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation *bool `pulumi:"disableIssValidation"`
	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt *bool `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string `pulumi:"kubernetesHost"`
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt *string `pulumi:"tokenReviewerJwt"`
}

// The set of arguments for constructing a AuthBackendConfig resource.
type AuthBackendConfigArgs struct {
	// Unique name of the kubernetes backend to configure.
	Backend pulumi.StringPtrInput
	// Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableIssValidation pulumi.BoolPtrInput
	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
	DisableLocalCaJwt pulumi.BoolPtrInput
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer pulumi.StringPtrInput
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringInput
	// List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput
	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt pulumi.StringPtrInput
}

func (AuthBackendConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigArgs)(nil)).Elem()
}

type AuthBackendConfigInput interface {
	pulumi.Input

	ToAuthBackendConfigOutput() AuthBackendConfigOutput
	ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput
}

func (AuthBackendConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendConfig)(nil)).Elem()
}

func (i AuthBackendConfig) ToAuthBackendConfigOutput() AuthBackendConfigOutput {
	return i.ToAuthBackendConfigOutputWithContext(context.Background())
}

func (i AuthBackendConfig) ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigOutput)
}

type AuthBackendConfigOutput struct {
	*pulumi.OutputState
}

func (AuthBackendConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendConfigOutput)(nil)).Elem()
}

func (o AuthBackendConfigOutput) ToAuthBackendConfigOutput() AuthBackendConfigOutput {
	return o
}

func (o AuthBackendConfigOutput) ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuthBackendConfigOutput{})
}
