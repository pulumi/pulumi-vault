// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Example using `serviceAccountName` mode:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * const sa_example = new vault.kubernetes.SecretBackendRole("sa-example", {
 *     backend: config.path,
 *     name: "service-account-name-role",
 *     allowedKubernetesNamespaces: ["*"],
 *     tokenMaxTtl: 43200,
 *     tokenDefaultTtl: 21600,
 *     serviceAccountName: "test-service-account-with-generated-token",
 *     extraLabels: {
 *         id: "abc123",
 *         name: "some_name",
 *     },
 *     extraAnnotations: {
 *         env: "development",
 *         location: "earth",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Example using `kubernetesRoleName` mode:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * const name_example = new vault.kubernetes.SecretBackendRole("name-example", {
 *     backend: config.path,
 *     name: "service-account-name-role",
 *     allowedKubernetesNamespaces: ["*"],
 *     tokenMaxTtl: 43200,
 *     tokenDefaultTtl: 21600,
 *     kubernetesRoleName: "vault-k8s-secrets-role",
 *     extraLabels: {
 *         id: "abc123",
 *         name: "some_name",
 *     },
 *     extraAnnotations: {
 *         env: "development",
 *         location: "earth",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Example using `generatedRoleRules` mode:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * const rules_example = new vault.kubernetes.SecretBackendRole("rules-example", {
 *     backend: config.path,
 *     name: "service-account-name-role",
 *     allowedKubernetesNamespaces: ["*"],
 *     tokenMaxTtl: 43200,
 *     tokenDefaultTtl: 21600,
 *     kubernetesRoleType: "Role",
 *     generatedRoleRules: `rules:
 * - apiGroups: [""]
 *   resources: ["pods"]
 *   verbs: ["list"]
 * `,
 *     extraLabels: {
 *         id: "abc123",
 *         name: "some_name",
 *     },
 *     extraAnnotations: {
 *         env: "development",
 *         location: "earth",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * The Kubernetes secret backend role can be imported using the full path to the role
 *
 * of the form: `<backend_path>/roles/<role_name>` e.g.
 *
 * ```sh
 * $ pulumi import vault:kubernetes/secretBackendRole:SecretBackendRole example kubernetes kubernetes/roles/example-role
 * ```
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kubernetes/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * A label selector for Kubernetes namespaces 
     * in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
     * of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
     * If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
     */
    public readonly allowedKubernetesNamespaceSelector!: pulumi.Output<string | undefined>;
    /**
     * The list of Kubernetes namespaces this role 
     * can generate credentials for. If set to `*` all namespaces are allowed. If set with
     * `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
     */
    public readonly allowedKubernetesNamespaces!: pulumi.Output<string[] | undefined>;
    /**
     * The path of the Kubernetes Secrets Engine backend mount to create
     * the role in.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Additional annotations to apply to all generated 
     * Kubernetes objects.
     */
    public readonly extraAnnotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Additional labels to apply to all generated Kubernetes 
     * objects.
     *
     * This resource also directly accepts all vault.Mount fields.
     */
    public readonly extraLabels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Role or ClusterRole rules to use when generating 
     * a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
     * and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
     * when credentials are requested.
     */
    public readonly generatedRoleRules!: pulumi.Output<string | undefined>;
    /**
     * The pre-existing Role or ClusterRole to bind a 
     * generated service account to. Mutually exclusive with `serviceAccountName` and
     * `generatedRoleRules`. If set, Kubernetes token, service account, and role
     * binding objects will be created when credentials are requested.
     */
    public readonly kubernetesRoleName!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the Kubernetes role is a Role or 
     * ClusterRole.
     */
    public readonly kubernetesRoleType!: pulumi.Output<string | undefined>;
    /**
     * The name of the role.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name template to use when generating service accounts, 
     * roles and role bindings. If unset, a default template is used.
     */
    public readonly nameTemplate!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The pre-existing service account to generate tokens for.
     * Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
     * Kubernetes token will be created when credentials are requested.
     */
    public readonly serviceAccountName!: pulumi.Output<string | undefined>;
    /**
     * The default TTL for generated Kubernetes tokens in seconds.
     */
    public readonly tokenDefaultTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum TTL for generated Kubernetes tokens in seconds.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            resourceInputs["allowedKubernetesNamespaceSelector"] = state ? state.allowedKubernetesNamespaceSelector : undefined;
            resourceInputs["allowedKubernetesNamespaces"] = state ? state.allowedKubernetesNamespaces : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["extraAnnotations"] = state ? state.extraAnnotations : undefined;
            resourceInputs["extraLabels"] = state ? state.extraLabels : undefined;
            resourceInputs["generatedRoleRules"] = state ? state.generatedRoleRules : undefined;
            resourceInputs["kubernetesRoleName"] = state ? state.kubernetesRoleName : undefined;
            resourceInputs["kubernetesRoleType"] = state ? state.kubernetesRoleType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameTemplate"] = state ? state.nameTemplate : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["serviceAccountName"] = state ? state.serviceAccountName : undefined;
            resourceInputs["tokenDefaultTtl"] = state ? state.tokenDefaultTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            resourceInputs["allowedKubernetesNamespaceSelector"] = args ? args.allowedKubernetesNamespaceSelector : undefined;
            resourceInputs["allowedKubernetesNamespaces"] = args ? args.allowedKubernetesNamespaces : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["extraAnnotations"] = args ? args.extraAnnotations : undefined;
            resourceInputs["extraLabels"] = args ? args.extraLabels : undefined;
            resourceInputs["generatedRoleRules"] = args ? args.generatedRoleRules : undefined;
            resourceInputs["kubernetesRoleName"] = args ? args.kubernetesRoleName : undefined;
            resourceInputs["kubernetesRoleType"] = args ? args.kubernetesRoleType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nameTemplate"] = args ? args.nameTemplate : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["serviceAccountName"] = args ? args.serviceAccountName : undefined;
            resourceInputs["tokenDefaultTtl"] = args ? args.tokenDefaultTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * A label selector for Kubernetes namespaces 
     * in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
     * of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
     * If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
     */
    allowedKubernetesNamespaceSelector?: pulumi.Input<string>;
    /**
     * The list of Kubernetes namespaces this role 
     * can generate credentials for. If set to `*` all namespaces are allowed. If set with
     * `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
     */
    allowedKubernetesNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path of the Kubernetes Secrets Engine backend mount to create
     * the role in.
     */
    backend?: pulumi.Input<string>;
    /**
     * Additional annotations to apply to all generated 
     * Kubernetes objects.
     */
    extraAnnotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Additional labels to apply to all generated Kubernetes 
     * objects.
     *
     * This resource also directly accepts all vault.Mount fields.
     */
    extraLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Role or ClusterRole rules to use when generating 
     * a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
     * and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
     * when credentials are requested.
     */
    generatedRoleRules?: pulumi.Input<string>;
    /**
     * The pre-existing Role or ClusterRole to bind a 
     * generated service account to. Mutually exclusive with `serviceAccountName` and
     * `generatedRoleRules`. If set, Kubernetes token, service account, and role
     * binding objects will be created when credentials are requested.
     */
    kubernetesRoleName?: pulumi.Input<string>;
    /**
     * Specifies whether the Kubernetes role is a Role or 
     * ClusterRole.
     */
    kubernetesRoleType?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The name template to use when generating service accounts, 
     * roles and role bindings. If unset, a default template is used.
     */
    nameTemplate?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The pre-existing service account to generate tokens for.
     * Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
     * Kubernetes token will be created when credentials are requested.
     */
    serviceAccountName?: pulumi.Input<string>;
    /**
     * The default TTL for generated Kubernetes tokens in seconds.
     */
    tokenDefaultTtl?: pulumi.Input<number>;
    /**
     * The maximum TTL for generated Kubernetes tokens in seconds.
     */
    tokenMaxTtl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * A label selector for Kubernetes namespaces 
     * in which credentials can be generated. Accepts either a JSON or YAML object. The value should be
     * of type [LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#labelselector-v1-meta).
     * If set with `allowedKubernetesNamespace`, the conditions are `OR`ed.
     */
    allowedKubernetesNamespaceSelector?: pulumi.Input<string>;
    /**
     * The list of Kubernetes namespaces this role 
     * can generate credentials for. If set to `*` all namespaces are allowed. If set with
     * `allowedKubernetesNamespaceSelector`, the conditions are `OR`ed.
     */
    allowedKubernetesNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path of the Kubernetes Secrets Engine backend mount to create
     * the role in.
     */
    backend: pulumi.Input<string>;
    /**
     * Additional annotations to apply to all generated 
     * Kubernetes objects.
     */
    extraAnnotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Additional labels to apply to all generated Kubernetes 
     * objects.
     *
     * This resource also directly accepts all vault.Mount fields.
     */
    extraLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Role or ClusterRole rules to use when generating 
     * a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with `serviceAccountName`
     * and `kubernetesRoleName`. If set, the entire chain of Kubernetes objects will be generated
     * when credentials are requested.
     */
    generatedRoleRules?: pulumi.Input<string>;
    /**
     * The pre-existing Role or ClusterRole to bind a 
     * generated service account to. Mutually exclusive with `serviceAccountName` and
     * `generatedRoleRules`. If set, Kubernetes token, service account, and role
     * binding objects will be created when credentials are requested.
     */
    kubernetesRoleName?: pulumi.Input<string>;
    /**
     * Specifies whether the Kubernetes role is a Role or 
     * ClusterRole.
     */
    kubernetesRoleType?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The name template to use when generating service accounts, 
     * roles and role bindings. If unset, a default template is used.
     */
    nameTemplate?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The pre-existing service account to generate tokens for.
     * Mutually exclusive with `kubernetesRoleName` and `generatedRoleRules`. If set, only a
     * Kubernetes token will be created when credentials are requested.
     */
    serviceAccountName?: pulumi.Input<string>;
    /**
     * The default TTL for generated Kubernetes tokens in seconds.
     */
    tokenDefaultTtl?: pulumi.Input<number>;
    /**
     * The maximum TTL for generated Kubernetes tokens in seconds.
     */
    tokenMaxTtl?: pulumi.Input<number>;
}
