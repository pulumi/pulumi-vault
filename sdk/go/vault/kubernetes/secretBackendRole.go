// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Example using `serviceAccountName` mode:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/kubernetes"
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
//			_, err = kubernetes.NewSecretBackendRole(ctx, "sa-example", &kubernetes.SecretBackendRoleArgs{
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
//			return nil
//		})
//	}
//
// ```
//
// Example using `kubernetesRoleName` mode:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/kubernetes"
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
//			_, err = kubernetes.NewSecretBackendRole(ctx, "name-example", &kubernetes.SecretBackendRoleArgs{
//				Backend: config.Path,
//				Name:    pulumi.String("service-account-name-role"),
//				AllowedKubernetesNamespaces: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				TokenMaxTtl:        pulumi.Int(43200),
//				TokenDefaultTtl:    pulumi.Int(21600),
//				KubernetesRoleName: pulumi.String("vault-k8s-secrets-role"),
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
//			return nil
//		})
//	}
//
// ```
//
// Example using `generatedRoleRules` mode:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/kubernetes"
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
//			_, err = kubernetes.NewSecretBackendRole(ctx, "rules-example", &kubernetes.SecretBackendRoleArgs{
//				Backend: config.Path,
//				Name:    pulumi.String("service-account-name-role"),
//				AllowedKubernetesNamespaces: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				TokenMaxTtl:        pulumi.Int(43200),
//				TokenDefaultTtl:    pulumi.Int(21600),
//				KubernetesRoleType: pulumi.String("Role"),
//				GeneratedRoleRules: pulumi.String("rules:\n- apiGroups: [\"\"]\n  resources: [\"pods\"]\n  verbs: [\"list\"]\n"),
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
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # The Kubernetes secret backend role can be imported using the full path to the role
//
// of the form: `<backend_path>/roles/<role_name>` e.g.
//
// ```sh
// $ pulumi import vault:kubernetes/secretBackendRole:SecretBackendRole example kubernetes kubernetes/roles/example-role
// ```
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// A label selector for Kubernetes namespaces
	// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
	// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
	// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
	AllowedKubernetesNamespaceSelector pulumi.StringPtrOutput `pulumi:"allowedKubernetesNamespaceSelector"`
	// The list of Kubernetes namespaces this role
	// can generate credentials for. If set to `*` all namespaces are allowed. If set with
	// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
	AllowedKubernetesNamespaces pulumi.StringArrayOutput `pulumi:"allowedKubernetesNamespaces"`
	// The path of the Kubernetes Secrets Engine backend mount to create
	// the role in.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Additional annotations to apply to all generated
	// Kubernetes objects.
	ExtraAnnotations pulumi.StringMapOutput `pulumi:"extraAnnotations"`
	// Additional labels to apply to all generated Kubernetes
	// objects.
	//
	// This resource also directly accepts all Mount fields.
	ExtraLabels pulumi.StringMapOutput `pulumi:"extraLabels"`
	// The Role or ClusterRole rules to use when generating
	// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
	// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
	// when credentials are requested.
	GeneratedRoleRules pulumi.StringPtrOutput `pulumi:"generatedRoleRules"`
	// The pre-existing Role or ClusterRole to bind a
	// generated service account to. Mutually exclusive with `serviceAccountName` and
	// `generatedRoleRules`. If set, Kubernetes token, service account, and role
	// binding objects will be created when credentials are requested.
	KubernetesRoleName pulumi.StringPtrOutput `pulumi:"kubernetesRoleName"`
	// Specifies whether the Kubernetes role is a Role or
	// ClusterRole.
	KubernetesRoleType pulumi.StringPtrOutput `pulumi:"kubernetesRoleType"`
	// The name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name template to use when generating service accounts,
	// roles and role bindings. If unset, a default template is used.
	NameTemplate pulumi.StringPtrOutput `pulumi:"nameTemplate"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The pre-existing service account to generate tokens for.
	// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
	// Kubernetes token will be created when credentials are requested.
	ServiceAccountName pulumi.StringPtrOutput `pulumi:"serviceAccountName"`
	// The default TTL for generated Kubernetes tokens in seconds.
	TokenDefaultTtl pulumi.IntPtrOutput `pulumi:"tokenDefaultTtl"`
	// The maximum TTL for generated Kubernetes tokens in seconds.
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:kubernetes/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:kubernetes/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// A label selector for Kubernetes namespaces
	// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
	// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
	// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
	AllowedKubernetesNamespaceSelector *string `pulumi:"allowedKubernetesNamespaceSelector"`
	// The list of Kubernetes namespaces this role
	// can generate credentials for. If set to `*` all namespaces are allowed. If set with
	// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
	AllowedKubernetesNamespaces []string `pulumi:"allowedKubernetesNamespaces"`
	// The path of the Kubernetes Secrets Engine backend mount to create
	// the role in.
	Backend *string `pulumi:"backend"`
	// Additional annotations to apply to all generated
	// Kubernetes objects.
	ExtraAnnotations map[string]string `pulumi:"extraAnnotations"`
	// Additional labels to apply to all generated Kubernetes
	// objects.
	//
	// This resource also directly accepts all Mount fields.
	ExtraLabels map[string]string `pulumi:"extraLabels"`
	// The Role or ClusterRole rules to use when generating
	// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
	// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
	// when credentials are requested.
	GeneratedRoleRules *string `pulumi:"generatedRoleRules"`
	// The pre-existing Role or ClusterRole to bind a
	// generated service account to. Mutually exclusive with `serviceAccountName` and
	// `generatedRoleRules`. If set, Kubernetes token, service account, and role
	// binding objects will be created when credentials are requested.
	KubernetesRoleName *string `pulumi:"kubernetesRoleName"`
	// Specifies whether the Kubernetes role is a Role or
	// ClusterRole.
	KubernetesRoleType *string `pulumi:"kubernetesRoleType"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The name template to use when generating service accounts,
	// roles and role bindings. If unset, a default template is used.
	NameTemplate *string `pulumi:"nameTemplate"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The pre-existing service account to generate tokens for.
	// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
	// Kubernetes token will be created when credentials are requested.
	ServiceAccountName *string `pulumi:"serviceAccountName"`
	// The default TTL for generated Kubernetes tokens in seconds.
	TokenDefaultTtl *int `pulumi:"tokenDefaultTtl"`
	// The maximum TTL for generated Kubernetes tokens in seconds.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
}

type SecretBackendRoleState struct {
	// A label selector for Kubernetes namespaces
	// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
	// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
	// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
	AllowedKubernetesNamespaceSelector pulumi.StringPtrInput
	// The list of Kubernetes namespaces this role
	// can generate credentials for. If set to `*` all namespaces are allowed. If set with
	// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
	AllowedKubernetesNamespaces pulumi.StringArrayInput
	// The path of the Kubernetes Secrets Engine backend mount to create
	// the role in.
	Backend pulumi.StringPtrInput
	// Additional annotations to apply to all generated
	// Kubernetes objects.
	ExtraAnnotations pulumi.StringMapInput
	// Additional labels to apply to all generated Kubernetes
	// objects.
	//
	// This resource also directly accepts all Mount fields.
	ExtraLabels pulumi.StringMapInput
	// The Role or ClusterRole rules to use when generating
	// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
	// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
	// when credentials are requested.
	GeneratedRoleRules pulumi.StringPtrInput
	// The pre-existing Role or ClusterRole to bind a
	// generated service account to. Mutually exclusive with `serviceAccountName` and
	// `generatedRoleRules`. If set, Kubernetes token, service account, and role
	// binding objects will be created when credentials are requested.
	KubernetesRoleName pulumi.StringPtrInput
	// Specifies whether the Kubernetes role is a Role or
	// ClusterRole.
	KubernetesRoleType pulumi.StringPtrInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The name template to use when generating service accounts,
	// roles and role bindings. If unset, a default template is used.
	NameTemplate pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The pre-existing service account to generate tokens for.
	// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
	// Kubernetes token will be created when credentials are requested.
	ServiceAccountName pulumi.StringPtrInput
	// The default TTL for generated Kubernetes tokens in seconds.
	TokenDefaultTtl pulumi.IntPtrInput
	// The maximum TTL for generated Kubernetes tokens in seconds.
	TokenMaxTtl pulumi.IntPtrInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// A label selector for Kubernetes namespaces
	// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
	// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
	// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
	AllowedKubernetesNamespaceSelector *string `pulumi:"allowedKubernetesNamespaceSelector"`
	// The list of Kubernetes namespaces this role
	// can generate credentials for. If set to `*` all namespaces are allowed. If set with
	// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
	AllowedKubernetesNamespaces []string `pulumi:"allowedKubernetesNamespaces"`
	// The path of the Kubernetes Secrets Engine backend mount to create
	// the role in.
	Backend string `pulumi:"backend"`
	// Additional annotations to apply to all generated
	// Kubernetes objects.
	ExtraAnnotations map[string]string `pulumi:"extraAnnotations"`
	// Additional labels to apply to all generated Kubernetes
	// objects.
	//
	// This resource also directly accepts all Mount fields.
	ExtraLabels map[string]string `pulumi:"extraLabels"`
	// The Role or ClusterRole rules to use when generating
	// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
	// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
	// when credentials are requested.
	GeneratedRoleRules *string `pulumi:"generatedRoleRules"`
	// The pre-existing Role or ClusterRole to bind a
	// generated service account to. Mutually exclusive with `serviceAccountName` and
	// `generatedRoleRules`. If set, Kubernetes token, service account, and role
	// binding objects will be created when credentials are requested.
	KubernetesRoleName *string `pulumi:"kubernetesRoleName"`
	// Specifies whether the Kubernetes role is a Role or
	// ClusterRole.
	KubernetesRoleType *string `pulumi:"kubernetesRoleType"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The name template to use when generating service accounts,
	// roles and role bindings. If unset, a default template is used.
	NameTemplate *string `pulumi:"nameTemplate"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The pre-existing service account to generate tokens for.
	// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
	// Kubernetes token will be created when credentials are requested.
	ServiceAccountName *string `pulumi:"serviceAccountName"`
	// The default TTL for generated Kubernetes tokens in seconds.
	TokenDefaultTtl *int `pulumi:"tokenDefaultTtl"`
	// The maximum TTL for generated Kubernetes tokens in seconds.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// A label selector for Kubernetes namespaces
	// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
	// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
	// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
	AllowedKubernetesNamespaceSelector pulumi.StringPtrInput
	// The list of Kubernetes namespaces this role
	// can generate credentials for. If set to `*` all namespaces are allowed. If set with
	// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
	AllowedKubernetesNamespaces pulumi.StringArrayInput
	// The path of the Kubernetes Secrets Engine backend mount to create
	// the role in.
	Backend pulumi.StringInput
	// Additional annotations to apply to all generated
	// Kubernetes objects.
	ExtraAnnotations pulumi.StringMapInput
	// Additional labels to apply to all generated Kubernetes
	// objects.
	//
	// This resource also directly accepts all Mount fields.
	ExtraLabels pulumi.StringMapInput
	// The Role or ClusterRole rules to use when generating
	// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
	// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
	// when credentials are requested.
	GeneratedRoleRules pulumi.StringPtrInput
	// The pre-existing Role or ClusterRole to bind a
	// generated service account to. Mutually exclusive with `serviceAccountName` and
	// `generatedRoleRules`. If set, Kubernetes token, service account, and role
	// binding objects will be created when credentials are requested.
	KubernetesRoleName pulumi.StringPtrInput
	// Specifies whether the Kubernetes role is a Role or
	// ClusterRole.
	KubernetesRoleType pulumi.StringPtrInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The name template to use when generating service accounts,
	// roles and role bindings. If unset, a default template is used.
	NameTemplate pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The pre-existing service account to generate tokens for.
	// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
	// Kubernetes token will be created when credentials are requested.
	ServiceAccountName pulumi.StringPtrInput
	// The default TTL for generated Kubernetes tokens in seconds.
	TokenDefaultTtl pulumi.IntPtrInput
	// The maximum TTL for generated Kubernetes tokens in seconds.
	TokenMaxTtl pulumi.IntPtrInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}

type SecretBackendRoleInput interface {
	pulumi.Input

	ToSecretBackendRoleOutput() SecretBackendRoleOutput
	ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput
}

func (*SecretBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (i *SecretBackendRole) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return i.ToSecretBackendRoleOutputWithContext(context.Background())
}

func (i *SecretBackendRole) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleOutput)
}

// SecretBackendRoleArrayInput is an input type that accepts SecretBackendRoleArray and SecretBackendRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleArrayInput` via:
//
//	SecretBackendRoleArray{ SecretBackendRoleArgs{...} }
type SecretBackendRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput
	ToSecretBackendRoleArrayOutputWithContext(context.Context) SecretBackendRoleArrayOutput
}

type SecretBackendRoleArray []SecretBackendRoleInput

func (SecretBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return i.ToSecretBackendRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleArrayOutput)
}

// SecretBackendRoleMapInput is an input type that accepts SecretBackendRoleMap and SecretBackendRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendRoleMapInput` via:
//
//	SecretBackendRoleMap{ "key": SecretBackendRoleArgs{...} }
type SecretBackendRoleMapInput interface {
	pulumi.Input

	ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput
	ToSecretBackendRoleMapOutputWithContext(context.Context) SecretBackendRoleMapOutput
}

type SecretBackendRoleMap map[string]SecretBackendRoleInput

func (SecretBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return i.ToSecretBackendRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleMapOutput)
}

type SecretBackendRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return o
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return o
}

// A label selector for Kubernetes namespaces
// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
// If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
func (o SecretBackendRoleOutput) AllowedKubernetesNamespaceSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.AllowedKubernetesNamespaceSelector }).(pulumi.StringPtrOutput)
}

// The list of Kubernetes namespaces this role
// can generate credentials for. If set to `*` all namespaces are allowed. If set with
// `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
func (o SecretBackendRoleOutput) AllowedKubernetesNamespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.AllowedKubernetesNamespaces }).(pulumi.StringArrayOutput)
}

// The path of the Kubernetes Secrets Engine backend mount to create
// the role in.
func (o SecretBackendRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Additional annotations to apply to all generated
// Kubernetes objects.
func (o SecretBackendRoleOutput) ExtraAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringMapOutput { return v.ExtraAnnotations }).(pulumi.StringMapOutput)
}

// Additional labels to apply to all generated Kubernetes
// objects.
//
// This resource also directly accepts all Mount fields.
func (o SecretBackendRoleOutput) ExtraLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringMapOutput { return v.ExtraLabels }).(pulumi.StringMapOutput)
}

// The Role or ClusterRole rules to use when generating
// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
// and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
// when credentials are requested.
func (o SecretBackendRoleOutput) GeneratedRoleRules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.GeneratedRoleRules }).(pulumi.StringPtrOutput)
}

// The pre-existing Role or ClusterRole to bind a
// generated service account to. Mutually exclusive with `serviceAccountName` and
// `generatedRoleRules`. If set, Kubernetes token, service account, and role
// binding objects will be created when credentials are requested.
func (o SecretBackendRoleOutput) KubernetesRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.KubernetesRoleName }).(pulumi.StringPtrOutput)
}

// Specifies whether the Kubernetes role is a Role or
// ClusterRole.
func (o SecretBackendRoleOutput) KubernetesRoleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.KubernetesRoleType }).(pulumi.StringPtrOutput)
}

// The name of the role.
func (o SecretBackendRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name template to use when generating service accounts,
// roles and role bindings. If unset, a default template is used.
func (o SecretBackendRoleOutput) NameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.NameTemplate }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The pre-existing service account to generate tokens for.
// Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
// Kubernetes token will be created when credentials are requested.
func (o SecretBackendRoleOutput) ServiceAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.ServiceAccountName }).(pulumi.StringPtrOutput)
}

// The default TTL for generated Kubernetes tokens in seconds.
func (o SecretBackendRoleOutput) TokenDefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntPtrOutput { return v.TokenDefaultTtl }).(pulumi.IntPtrOutput)
}

// The maximum TTL for generated Kubernetes tokens in seconds.
func (o SecretBackendRoleOutput) TokenMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntPtrOutput { return v.TokenMaxTtl }).(pulumi.IntPtrOutput)
}

type SecretBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].([]*SecretBackendRole)[vs[1].(int)]
	}).(SecretBackendRoleOutput)
}

type SecretBackendRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].(map[string]*SecretBackendRole)[vs[1].(string)]
	}).(SecretBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleInput)(nil)).Elem(), &SecretBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleArrayInput)(nil)).Elem(), SecretBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleMapInput)(nil)).Elem(), SecretBackendRoleMap{})
	pulumi.RegisterOutputType(SecretBackendRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleMapOutput{})
}
