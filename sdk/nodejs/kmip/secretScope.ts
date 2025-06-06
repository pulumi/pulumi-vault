// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages KMIP Secret Scopes in a Vault server. This feature requires
 * Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
 * for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const _default = new vault.kmip.SecretBackend("default", {
 *     path: "kmip",
 *     description: "Vault KMIP backend",
 * });
 * const dev = new vault.kmip.SecretScope("dev", {
 *     path: _default.path,
 *     scope: "dev",
 *     force: true,
 * });
 * ```
 *
 * ## Import
 *
 * KMIP Secret scope can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:kmip/secretScope:SecretScope dev kmip
 * ```
 */
export class SecretScope extends pulumi.CustomResource {
    /**
     * Get an existing SecretScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretScopeState, opts?: pulumi.CustomResourceOptions): SecretScope {
        return new SecretScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kmip/secretScope:SecretScope';

    /**
     * Returns true if the given object is an instance of SecretScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretScope.__pulumiType;
    }

    /**
     * Boolean field to force deletion even if there are managed objects in the scope.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Name of the scope.
     */
    public readonly scope!: pulumi.Output<string>;

    /**
     * Create a SecretScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretScopeArgs | SecretScopeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretScopeState | undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as SecretScopeArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretScope.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretScope resources.
 */
export interface SecretScopeState {
    /**
     * Boolean field to force deletion even if there are managed objects in the scope.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    path?: pulumi.Input<string>;
    /**
     * Name of the scope.
     */
    scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretScope resource.
 */
export interface SecretScopeArgs {
    /**
     * Boolean field to force deletion even if there are managed objects in the scope.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    path: pulumi.Input<string>;
    /**
     * Name of the scope.
     */
    scope: pulumi.Input<string>;
}
