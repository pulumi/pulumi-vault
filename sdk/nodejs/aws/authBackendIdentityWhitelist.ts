// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configures the periodic tidying operation of the whitelisted identity entries.
 *
 * For more information, see the
 * [Vault docs](https://www.vaultproject.io/api-docs/auth/aws#configure-identity-whitelist-tidy-operation).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const exampleAuthBackend = new vault.AuthBackend("exampleAuthBackend", {type: "aws"});
 * const exampleAuthBackendIdentityWhitelist = new vault.aws.AuthBackendIdentityWhitelist("exampleAuthBackendIdentityWhitelist", {
 *     backend: exampleAuthBackend.path,
 *     safetyBuffer: 3600,
 * });
 * ```
 *
 * ## Import
 *
 * AWS auth backend identity whitelists can be imported using `auth/`, the `backend` path, and `/config/tidy/identity-whitelist` e.g.
 *
 * ```sh
 * $ pulumi import vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist example auth/aws/config/tidy/identity-whitelist
 * ```
 */
export class AuthBackendIdentityWhitelist extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendIdentityWhitelist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendIdentityWhitelistState, opts?: pulumi.CustomResourceOptions): AuthBackendIdentityWhitelist {
        return new AuthBackendIdentityWhitelist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist';

    /**
     * Returns true if the given object is an instance of AuthBackendIdentityWhitelist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendIdentityWhitelist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendIdentityWhitelist.__pulumiType;
    }

    /**
     * The path of the AWS backend being configured.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set to true, disables the periodic
     * tidying of the identity-whitelist entries.
     */
    public readonly disablePeriodicTidy!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The amount of extra time, in minutes, that must
     * have passed beyond the roletag expiration, before it is removed from the
     * backend storage.
     */
    public readonly safetyBuffer!: pulumi.Output<number | undefined>;

    /**
     * Create a AuthBackendIdentityWhitelist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendIdentityWhitelistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendIdentityWhitelistArgs | AuthBackendIdentityWhitelistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendIdentityWhitelistState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["disablePeriodicTidy"] = state ? state.disablePeriodicTidy : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["safetyBuffer"] = state ? state.safetyBuffer : undefined;
        } else {
            const args = argsOrState as AuthBackendIdentityWhitelistArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["disablePeriodicTidy"] = args ? args.disablePeriodicTidy : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["safetyBuffer"] = args ? args.safetyBuffer : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendIdentityWhitelist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendIdentityWhitelist resources.
 */
export interface AuthBackendIdentityWhitelistState {
    /**
     * The path of the AWS backend being configured.
     */
    backend?: pulumi.Input<string>;
    /**
     * If set to true, disables the periodic
     * tidying of the identity-whitelist entries.
     */
    disablePeriodicTidy?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The amount of extra time, in minutes, that must
     * have passed beyond the roletag expiration, before it is removed from the
     * backend storage.
     */
    safetyBuffer?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AuthBackendIdentityWhitelist resource.
 */
export interface AuthBackendIdentityWhitelistArgs {
    /**
     * The path of the AWS backend being configured.
     */
    backend?: pulumi.Input<string>;
    /**
     * If set to true, disables the periodic
     * tidying of the identity-whitelist entries.
     */
    disablePeriodicTidy?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The amount of extra time, in minutes, that must
     * have passed beyond the roletag expiration, before it is removed from the
     * backend storage.
     */
    safetyBuffer?: pulumi.Input<number>;
}
