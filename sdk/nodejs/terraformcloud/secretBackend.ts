// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.terraformcloud.SecretBackend("test", {
 *     backend: "terraform",
 *     description: "Manages the Terraform Cloud backend",
 *     token: "V0idfhi2iksSDU234ucdbi2nidsi...",
 * });
 * ```
 *
 * ## Import
 *
 * Terraform Cloud secret backends can be imported using the `backend`, e.g.
 *
 * ```sh
 * $ pulumi import vault:terraformcloud/secretBackend:SecretBackend example terraform
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
    public static readonly __pulumiType = 'vault:terraformcloud/secretBackend:SecretBackend';

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

    public readonly address!: pulumi.Output<string | undefined>;
    public readonly backend!: pulumi.Output<string | undefined>;
    public readonly basePath!: pulumi.Output<string | undefined>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * A human-friendly description for this backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    public readonly token!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["basePath"] = state ? state.basePath : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["basePath"] = args ? args.basePath : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    address?: pulumi.Input<string>;
    backend?: pulumi.Input<string>;
    basePath?: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    address?: pulumi.Input<string>;
    backend?: pulumi.Input<string>;
    basePath?: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
}
