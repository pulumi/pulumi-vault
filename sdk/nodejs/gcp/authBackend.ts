// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure the [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
 *
 * ## Example Usage
 *
 * You can setup the GCP auth backend with Workload Identity Federation (WIF) for a secret-less configuration:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const gcp = new vault.gcp.AuthBackend("gcp", {
 *     identityTokenKey: "example-key",
 *     identityTokenTtl: 1800,
 *     identityTokenAudience: "<TOKEN_AUDIENCE>",
 *     serviceAccountEmail: "<SERVICE_ACCOUNT_EMAIL>",
 * });
 * ```
 *
 * ## Import
 *
 * GCP authentication backends can be imported using the backend name, e.g.
 *
 * ```sh
 * $ pulumi import vault:gcp/authBackend:AuthBackend gcp gcp
 * ```
 */
export class AuthBackend extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendState, opts?: pulumi.CustomResourceOptions): AuthBackend {
        return new AuthBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/authBackend:AuthBackend';

    /**
     * Returns true if the given object is an instance of AuthBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackend.__pulumiType;
    }

    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The clients email associated with the credentials
     */
    public readonly clientEmail!: pulumi.Output<string>;
    /**
     * The Client ID of the credentials
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     *
     * Overrides are set at the subdomain level using the following keys:
     */
    public readonly customEndpoint!: pulumi.Output<outputs.gcp.AuthBackendCustomEndpoint | undefined>;
    /**
     * A description of the auth method.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    public readonly identityTokenAudience!: pulumi.Output<string | undefined>;
    /**
     * The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    public readonly identityTokenKey!: pulumi.Output<string | undefined>;
    /**
     * The TTL of generated tokens.
     */
    public readonly identityTokenTtl!: pulumi.Output<number | undefined>;
    /**
     * Specifies if the auth method is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The ID of the private key from the credentials
     */
    public readonly privateKeyId!: pulumi.Output<string>;
    /**
     * The GCP Project ID
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Service Account to impersonate for plugin workload identity federation.
     * Required with `identityTokenAudience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string | undefined>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    public readonly tune!: pulumi.Output<outputs.gcp.AuthBackendTune>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["clientEmail"] = state ? state.clientEmail : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["customEndpoint"] = state ? state.customEndpoint : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["identityTokenAudience"] = state ? state.identityTokenAudience : undefined;
            resourceInputs["identityTokenKey"] = state ? state.identityTokenKey : undefined;
            resourceInputs["identityTokenTtl"] = state ? state.identityTokenTtl : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["privateKeyId"] = state ? state.privateKeyId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["tune"] = state ? state.tune : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            resourceInputs["clientEmail"] = args ? args.clientEmail : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["credentials"] = args?.credentials ? pulumi.secret(args.credentials) : undefined;
            resourceInputs["customEndpoint"] = args ? args.customEndpoint : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["identityTokenAudience"] = args ? args.identityTokenAudience : undefined;
            resourceInputs["identityTokenKey"] = args ? args.identityTokenKey : undefined;
            resourceInputs["identityTokenTtl"] = args ? args.identityTokenTtl : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["privateKeyId"] = args ? args.privateKeyId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            resourceInputs["tune"] = args ? args.tune : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["credentials"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     */
    accessor?: pulumi.Input<string>;
    /**
     * The clients email associated with the credentials
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * The Client ID of the credentials
     */
    clientId?: pulumi.Input<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    credentials?: pulumi.Input<string>;
    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     *
     * Overrides are set at the subdomain level using the following keys:
     */
    customEndpoint?: pulumi.Input<inputs.gcp.AuthBackendCustomEndpoint>;
    /**
     * A description of the auth method.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    identityTokenKey?: pulumi.Input<string>;
    /**
     * The TTL of generated tokens.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the private key from the credentials
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * The GCP Project ID
     */
    projectId?: pulumi.Input<string>;
    /**
     * Service Account to impersonate for plugin workload identity federation.
     * Required with `identityTokenAudience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    tune?: pulumi.Input<inputs.gcp.AuthBackendTune>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The clients email associated with the credentials
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * The Client ID of the credentials
     */
    clientId?: pulumi.Input<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    credentials?: pulumi.Input<string>;
    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     *
     * Overrides are set at the subdomain level using the following keys:
     */
    customEndpoint?: pulumi.Input<inputs.gcp.AuthBackendCustomEndpoint>;
    /**
     * A description of the auth method.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    identityTokenKey?: pulumi.Input<string>;
    /**
     * The TTL of generated tokens.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the private key from the credentials
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * The GCP Project ID
     */
    projectId?: pulumi.Input<string>;
    /**
     * Service Account to impersonate for plugin workload identity federation.
     * Required with `identityTokenAudience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    tune?: pulumi.Input<inputs.gcp.AuthBackendTune>;
}
