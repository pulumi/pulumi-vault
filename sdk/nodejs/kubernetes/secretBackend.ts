// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     defaultLeaseTtlSeconds: 43200,
 *     maxLeaseTtlSeconds: 86400,
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * ```
 *
 * ## Import
 *
 * The Kubernetes secret backend can be imported using its `path` e.g.
 *
 * ```sh
 * $ pulumi import vault:kubernetes/secretBackend:SecretBackend config kubernetes
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kubernetes/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * Accessor of the mount
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * List of managed key registry entry names that the mount in question is allowed to access
     */
    public readonly allowedManagedKeys!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    public readonly auditNonHmacRequestKeys!: pulumi.Output<string[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    public readonly auditNonHmacResponseKeys!: pulumi.Output<string[]>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Human-friendly description of the mount
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Disable defaulting to the local CA certificate and 
     * service account JWT when Vault is running in a Kubernetes pod.
     */
    public readonly disableLocalCaJwt!: pulumi.Output<boolean | undefined>;
    /**
     * Enable the secrets engine to access Vault's external entropy source
     */
    public readonly externalEntropyAccess!: pulumi.Output<boolean | undefined>;
    /**
     * A PEM-encoded CA certificate used by the 
     * secrets engine to verify the Kubernetes API server certificate. Defaults to the local
     * pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
     * Vault is running.
     */
    public readonly kubernetesCaCert!: pulumi.Output<string | undefined>;
    /**
     * The Kubernetes API URL to connect to. Required if the 
     * standard pod environment variables `KUBERNETES_SERVICE_HOST` or `KUBERNETES_SERVICE_PORT`
     * are not set on the host that Vault is running on.
     */
    public readonly kubernetesHost!: pulumi.Output<string | undefined>;
    /**
     * Local mount flag that can be explicitly set to true to enforce local mount in HA environment
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    public readonly options!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Where the secret backend will be mounted
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    public readonly sealWrap!: pulumi.Output<boolean>;
    /**
     * The JSON web token of the service account used by the
     * secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
     * is running in Kubernetes.
     */
    public readonly serviceAccountJwt!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["allowedManagedKeys"] = state ? state.allowedManagedKeys : undefined;
            resourceInputs["auditNonHmacRequestKeys"] = state ? state.auditNonHmacRequestKeys : undefined;
            resourceInputs["auditNonHmacResponseKeys"] = state ? state.auditNonHmacResponseKeys : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableLocalCaJwt"] = state ? state.disableLocalCaJwt : undefined;
            resourceInputs["externalEntropyAccess"] = state ? state.externalEntropyAccess : undefined;
            resourceInputs["kubernetesCaCert"] = state ? state.kubernetesCaCert : undefined;
            resourceInputs["kubernetesHost"] = state ? state.kubernetesHost : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["sealWrap"] = state ? state.sealWrap : undefined;
            resourceInputs["serviceAccountJwt"] = state ? state.serviceAccountJwt : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["allowedManagedKeys"] = args ? args.allowedManagedKeys : undefined;
            resourceInputs["auditNonHmacRequestKeys"] = args ? args.auditNonHmacRequestKeys : undefined;
            resourceInputs["auditNonHmacResponseKeys"] = args ? args.auditNonHmacResponseKeys : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableLocalCaJwt"] = args ? args.disableLocalCaJwt : undefined;
            resourceInputs["externalEntropyAccess"] = args ? args.externalEntropyAccess : undefined;
            resourceInputs["kubernetesCaCert"] = args ? args.kubernetesCaCert : undefined;
            resourceInputs["kubernetesHost"] = args ? args.kubernetesHost : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["sealWrap"] = args ? args.sealWrap : undefined;
            resourceInputs["serviceAccountJwt"] = args?.serviceAccountJwt ? pulumi.secret(args.serviceAccountJwt) : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["serviceAccountJwt"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Accessor of the mount
     */
    accessor?: pulumi.Input<string>;
    /**
     * List of managed key registry entry names that the mount in question is allowed to access
     */
    allowedManagedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    description?: pulumi.Input<string>;
    /**
     * Disable defaulting to the local CA certificate and 
     * service account JWT when Vault is running in a Kubernetes pod.
     */
    disableLocalCaJwt?: pulumi.Input<boolean>;
    /**
     * Enable the secrets engine to access Vault's external entropy source
     */
    externalEntropyAccess?: pulumi.Input<boolean>;
    /**
     * A PEM-encoded CA certificate used by the 
     * secrets engine to verify the Kubernetes API server certificate. Defaults to the local
     * pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
     * Vault is running.
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * The Kubernetes API URL to connect to. Required if the 
     * standard pod environment variables `KUBERNETES_SERVICE_HOST` or `KUBERNETES_SERVICE_PORT`
     * are not set on the host that Vault is running on.
     */
    kubernetesHost?: pulumi.Input<string>;
    /**
     * Local mount flag that can be explicitly set to true to enforce local mount in HA environment
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * Where the secret backend will be mounted
     */
    path?: pulumi.Input<string>;
    /**
     * Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    sealWrap?: pulumi.Input<boolean>;
    /**
     * The JSON web token of the service account used by the
     * secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
     * is running in Kubernetes.
     */
    serviceAccountJwt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * List of managed key registry entry names that the mount in question is allowed to access
     */
    allowedManagedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    description?: pulumi.Input<string>;
    /**
     * Disable defaulting to the local CA certificate and 
     * service account JWT when Vault is running in a Kubernetes pod.
     */
    disableLocalCaJwt?: pulumi.Input<boolean>;
    /**
     * Enable the secrets engine to access Vault's external entropy source
     */
    externalEntropyAccess?: pulumi.Input<boolean>;
    /**
     * A PEM-encoded CA certificate used by the 
     * secrets engine to verify the Kubernetes API server certificate. Defaults to the local
     * pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
     * Vault is running.
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * The Kubernetes API URL to connect to. Required if the 
     * standard pod environment variables `KUBERNETES_SERVICE_HOST` or `KUBERNETES_SERVICE_PORT`
     * are not set on the host that Vault is running on.
     */
    kubernetesHost?: pulumi.Input<string>;
    /**
     * Local mount flag that can be explicitly set to true to enforce local mount in HA environment
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * Where the secret backend will be mounted
     */
    path: pulumi.Input<string>;
    /**
     * Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    sealWrap?: pulumi.Input<boolean>;
    /**
     * The JSON web token of the service account used by the
     * secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
     * is running in Kubernetes.
     */
    serviceAccountJwt?: pulumi.Input<string>;
}
