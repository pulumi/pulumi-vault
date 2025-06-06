// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kubernetes
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// Example using `service_account_name` mode:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Std = Pulumi.Std;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Vault.Kubernetes.SecretBackend("config", new()
    ///     {
    ///         Path = "kubernetes",
    ///         Description = "kubernetes secrets engine description",
    ///         KubernetesHost = "https://127.0.0.1:61233",
    ///         KubernetesCaCert = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/cert",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         ServiceAccountJwt = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/token",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         DisableLocalCaJwt = false,
    ///     });
    /// 
    ///     var sa_example = new Vault.Kubernetes.SecretBackendRole("sa-example", new()
    ///     {
    ///         Backend = config.Path,
    ///         Name = "service-account-name-role",
    ///         AllowedKubernetesNamespaces = new[]
    ///         {
    ///             "*",
    ///         },
    ///         TokenMaxTtl = 43200,
    ///         TokenDefaultTtl = 21600,
    ///         ServiceAccountName = "test-service-account-with-generated-token",
    ///         ExtraLabels = 
    ///         {
    ///             { "id", "abc123" },
    ///             { "name", "some_name" },
    ///         },
    ///         ExtraAnnotations = 
    ///         {
    ///             { "env", "development" },
    ///             { "location", "earth" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Example using `kubernetes_role_name` mode:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Std = Pulumi.Std;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Vault.Kubernetes.SecretBackend("config", new()
    ///     {
    ///         Path = "kubernetes",
    ///         Description = "kubernetes secrets engine description",
    ///         KubernetesHost = "https://127.0.0.1:61233",
    ///         KubernetesCaCert = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/cert",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         ServiceAccountJwt = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/token",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         DisableLocalCaJwt = false,
    ///     });
    /// 
    ///     var name_example = new Vault.Kubernetes.SecretBackendRole("name-example", new()
    ///     {
    ///         Backend = config.Path,
    ///         Name = "service-account-name-role",
    ///         AllowedKubernetesNamespaces = new[]
    ///         {
    ///             "*",
    ///         },
    ///         TokenMaxTtl = 43200,
    ///         TokenDefaultTtl = 21600,
    ///         KubernetesRoleName = "vault-k8s-secrets-role",
    ///         ExtraLabels = 
    ///         {
    ///             { "id", "abc123" },
    ///             { "name", "some_name" },
    ///         },
    ///         ExtraAnnotations = 
    ///         {
    ///             { "env", "development" },
    ///             { "location", "earth" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Example using `generated_role_rules` mode:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Std = Pulumi.Std;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Vault.Kubernetes.SecretBackend("config", new()
    ///     {
    ///         Path = "kubernetes",
    ///         Description = "kubernetes secrets engine description",
    ///         KubernetesHost = "https://127.0.0.1:61233",
    ///         KubernetesCaCert = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/cert",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         ServiceAccountJwt = Std.File.Invoke(new()
    ///         {
    ///             Input = "/path/to/token",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         DisableLocalCaJwt = false,
    ///     });
    /// 
    ///     var rules_example = new Vault.Kubernetes.SecretBackendRole("rules-example", new()
    ///     {
    ///         Backend = config.Path,
    ///         Name = "service-account-name-role",
    ///         AllowedKubernetesNamespaces = new[]
    ///         {
    ///             "*",
    ///         },
    ///         TokenMaxTtl = 43200,
    ///         TokenDefaultTtl = 21600,
    ///         KubernetesRoleType = "Role",
    ///         GeneratedRoleRules = @"rules:
    /// - apiGroups: [""""]
    ///   resources: [""pods""]
    ///   verbs: [""list""]
    /// ",
    ///         ExtraLabels = 
    ///         {
    ///             { "id", "abc123" },
    ///             { "name", "some_name" },
    ///         },
    ///         ExtraAnnotations = 
    ///         {
    ///             { "env", "development" },
    ///             { "location", "earth" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The Kubernetes secret backend role can be imported using the full path to the role
    /// 
    /// of the form: `&lt;backend_path&gt;/roles/&lt;role_name&gt;` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:kubernetes/secretBackendRole:SecretBackendRole example kubernetes kubernetes/roles/example-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:kubernetes/secretBackendRole:SecretBackendRole")]
    public partial class SecretBackendRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A label selector for Kubernetes namespaces 
        /// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
        /// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
        /// If set with `allowed_kubernetes_namespace`, the conditions are `OR`ed.
        /// </summary>
        [Output("allowedKubernetesNamespaceSelector")]
        public Output<string?> AllowedKubernetesNamespaceSelector { get; private set; } = null!;

        /// <summary>
        /// The list of Kubernetes namespaces this role 
        /// can generate credentials for. If set to `*` all namespaces are allowed. If set with
        /// `allowed_kubernetes_namespace_selector`, the conditions are `OR`ed.
        /// </summary>
        [Output("allowedKubernetesNamespaces")]
        public Output<ImmutableArray<string>> AllowedKubernetesNamespaces { get; private set; } = null!;

        /// <summary>
        /// The path of the Kubernetes Secrets Engine backend mount to create
        /// the role in.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Additional annotations to apply to all generated 
        /// Kubernetes objects.
        /// </summary>
        [Output("extraAnnotations")]
        public Output<ImmutableDictionary<string, string>?> ExtraAnnotations { get; private set; } = null!;

        /// <summary>
        /// Additional labels to apply to all generated Kubernetes 
        /// objects.
        /// 
        /// This resource also directly accepts all vault.Mount fields.
        /// </summary>
        [Output("extraLabels")]
        public Output<ImmutableDictionary<string, string>?> ExtraLabels { get; private set; } = null!;

        /// <summary>
        /// The Role or ClusterRole rules to use when generating 
        /// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `service_account_name`
        /// and `kubernetes_role_name`. If set, the entire chain of Kubernetes objects will be generated
        /// when credentials are requested.
        /// </summary>
        [Output("generatedRoleRules")]
        public Output<string?> GeneratedRoleRules { get; private set; } = null!;

        /// <summary>
        /// The pre-existing Role or ClusterRole to bind a 
        /// generated service account to. Mutually exclusive with `service_account_name` and
        /// `generated_role_rules`. If set, Kubernetes token, service account, and role
        /// binding objects will be created when credentials are requested.
        /// </summary>
        [Output("kubernetesRoleName")]
        public Output<string?> KubernetesRoleName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Kubernetes role is a Role or 
        /// ClusterRole.
        /// </summary>
        [Output("kubernetesRoleType")]
        public Output<string?> KubernetesRoleType { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name template to use when generating service accounts, 
        /// roles and role bindings. If unset, a default template is used.
        /// </summary>
        [Output("nameTemplate")]
        public Output<string?> NameTemplate { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The pre-existing service account to generate tokens for.
        /// Mutually exclusive with `kubernetes_role_name` and `generated_role_rules`. If set, only a
        /// Kubernetes token will be created when credentials are requested.
        /// </summary>
        [Output("serviceAccountName")]
        public Output<string?> ServiceAccountName { get; private set; } = null!;

        /// <summary>
        /// The default TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Output("tokenDefaultTtl")]
        public Output<int?> TokenDefaultTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:kubernetes/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:kubernetes/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A label selector for Kubernetes namespaces 
        /// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
        /// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
        /// If set with `allowed_kubernetes_namespace`, the conditions are `OR`ed.
        /// </summary>
        [Input("allowedKubernetesNamespaceSelector")]
        public Input<string>? AllowedKubernetesNamespaceSelector { get; set; }

        [Input("allowedKubernetesNamespaces")]
        private InputList<string>? _allowedKubernetesNamespaces;

        /// <summary>
        /// The list of Kubernetes namespaces this role 
        /// can generate credentials for. If set to `*` all namespaces are allowed. If set with
        /// `allowed_kubernetes_namespace_selector`, the conditions are `OR`ed.
        /// </summary>
        public InputList<string> AllowedKubernetesNamespaces
        {
            get => _allowedKubernetesNamespaces ?? (_allowedKubernetesNamespaces = new InputList<string>());
            set => _allowedKubernetesNamespaces = value;
        }

        /// <summary>
        /// The path of the Kubernetes Secrets Engine backend mount to create
        /// the role in.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("extraAnnotations")]
        private InputMap<string>? _extraAnnotations;

        /// <summary>
        /// Additional annotations to apply to all generated 
        /// Kubernetes objects.
        /// </summary>
        public InputMap<string> ExtraAnnotations
        {
            get => _extraAnnotations ?? (_extraAnnotations = new InputMap<string>());
            set => _extraAnnotations = value;
        }

        [Input("extraLabels")]
        private InputMap<string>? _extraLabels;

        /// <summary>
        /// Additional labels to apply to all generated Kubernetes 
        /// objects.
        /// 
        /// This resource also directly accepts all vault.Mount fields.
        /// </summary>
        public InputMap<string> ExtraLabels
        {
            get => _extraLabels ?? (_extraLabels = new InputMap<string>());
            set => _extraLabels = value;
        }

        /// <summary>
        /// The Role or ClusterRole rules to use when generating 
        /// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `service_account_name`
        /// and `kubernetes_role_name`. If set, the entire chain of Kubernetes objects will be generated
        /// when credentials are requested.
        /// </summary>
        [Input("generatedRoleRules")]
        public Input<string>? GeneratedRoleRules { get; set; }

        /// <summary>
        /// The pre-existing Role or ClusterRole to bind a 
        /// generated service account to. Mutually exclusive with `service_account_name` and
        /// `generated_role_rules`. If set, Kubernetes token, service account, and role
        /// binding objects will be created when credentials are requested.
        /// </summary>
        [Input("kubernetesRoleName")]
        public Input<string>? KubernetesRoleName { get; set; }

        /// <summary>
        /// Specifies whether the Kubernetes role is a Role or 
        /// ClusterRole.
        /// </summary>
        [Input("kubernetesRoleType")]
        public Input<string>? KubernetesRoleType { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name template to use when generating service accounts, 
        /// roles and role bindings. If unset, a default template is used.
        /// </summary>
        [Input("nameTemplate")]
        public Input<string>? NameTemplate { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The pre-existing service account to generate tokens for.
        /// Mutually exclusive with `kubernetes_role_name` and `generated_role_rules`. If set, only a
        /// Kubernetes token will be created when credentials are requested.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        /// <summary>
        /// The default TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Input("tokenDefaultTtl")]
        public Input<int>? TokenDefaultTtl { get; set; }

        /// <summary>
        /// The maximum TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        public SecretBackendRoleArgs()
        {
        }
        public static new SecretBackendRoleArgs Empty => new SecretBackendRoleArgs();
    }

    public sealed class SecretBackendRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A label selector for Kubernetes namespaces 
        /// in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
        /// of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
        /// If set with `allowed_kubernetes_namespace`, the conditions are `OR`ed.
        /// </summary>
        [Input("allowedKubernetesNamespaceSelector")]
        public Input<string>? AllowedKubernetesNamespaceSelector { get; set; }

        [Input("allowedKubernetesNamespaces")]
        private InputList<string>? _allowedKubernetesNamespaces;

        /// <summary>
        /// The list of Kubernetes namespaces this role 
        /// can generate credentials for. If set to `*` all namespaces are allowed. If set with
        /// `allowed_kubernetes_namespace_selector`, the conditions are `OR`ed.
        /// </summary>
        public InputList<string> AllowedKubernetesNamespaces
        {
            get => _allowedKubernetesNamespaces ?? (_allowedKubernetesNamespaces = new InputList<string>());
            set => _allowedKubernetesNamespaces = value;
        }

        /// <summary>
        /// The path of the Kubernetes Secrets Engine backend mount to create
        /// the role in.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("extraAnnotations")]
        private InputMap<string>? _extraAnnotations;

        /// <summary>
        /// Additional annotations to apply to all generated 
        /// Kubernetes objects.
        /// </summary>
        public InputMap<string> ExtraAnnotations
        {
            get => _extraAnnotations ?? (_extraAnnotations = new InputMap<string>());
            set => _extraAnnotations = value;
        }

        [Input("extraLabels")]
        private InputMap<string>? _extraLabels;

        /// <summary>
        /// Additional labels to apply to all generated Kubernetes 
        /// objects.
        /// 
        /// This resource also directly accepts all vault.Mount fields.
        /// </summary>
        public InputMap<string> ExtraLabels
        {
            get => _extraLabels ?? (_extraLabels = new InputMap<string>());
            set => _extraLabels = value;
        }

        /// <summary>
        /// The Role or ClusterRole rules to use when generating 
        /// a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `service_account_name`
        /// and `kubernetes_role_name`. If set, the entire chain of Kubernetes objects will be generated
        /// when credentials are requested.
        /// </summary>
        [Input("generatedRoleRules")]
        public Input<string>? GeneratedRoleRules { get; set; }

        /// <summary>
        /// The pre-existing Role or ClusterRole to bind a 
        /// generated service account to. Mutually exclusive with `service_account_name` and
        /// `generated_role_rules`. If set, Kubernetes token, service account, and role
        /// binding objects will be created when credentials are requested.
        /// </summary>
        [Input("kubernetesRoleName")]
        public Input<string>? KubernetesRoleName { get; set; }

        /// <summary>
        /// Specifies whether the Kubernetes role is a Role or 
        /// ClusterRole.
        /// </summary>
        [Input("kubernetesRoleType")]
        public Input<string>? KubernetesRoleType { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name template to use when generating service accounts, 
        /// roles and role bindings. If unset, a default template is used.
        /// </summary>
        [Input("nameTemplate")]
        public Input<string>? NameTemplate { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The pre-existing service account to generate tokens for.
        /// Mutually exclusive with `kubernetes_role_name` and `generated_role_rules`. If set, only a
        /// Kubernetes token will be created when credentials are requested.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        /// <summary>
        /// The default TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Input("tokenDefaultTtl")]
        public Input<int>? TokenDefaultTtl { get; set; }

        /// <summary>
        /// The maximum TTL for generated Kubernetes tokens in seconds.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        public SecretBackendRoleState()
        {
        }
        public static new SecretBackendRoleState Empty => new SecretBackendRoleState();
    }
}
