// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a GitHub Auth mount in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/github/) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.github.AuthBackend("example", {
 *     organization: "myorg",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub authentication mounts can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:github/authBackend:AuthBackend example github
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
    public static readonly __pulumiType = 'vault:github/authBackend:AuthBackend';

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
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     */
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The organization configured users must be part of.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The ID of the organization users must be part of.
     * Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
     */
    public readonly organizationId!: pulumi.Output<number>;
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * Extra configuration block. Structure is documented below.
     */
    public readonly tune!: pulumi.Output<outputs.github.AuthBackendTune>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["baseUrl"] = state ? state.baseUrl : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
            resourceInputs["tune"] = state ? state.tune : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            resourceInputs["baseUrl"] = args ? args.baseUrl : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
            resourceInputs["tune"] = args ? args.tune : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     */
    description?: pulumi.Input<string>;
    /**
     * The organization configured users must be part of.
     */
    organization?: pulumi.Input<string>;
    /**
     * The ID of the organization users must be part of.
     * Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
     */
    organizationId?: pulumi.Input<number>;
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     */
    path?: pulumi.Input<string>;
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * (Optional) The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    tokenType?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     */
    tune?: pulumi.Input<inputs.github.AuthBackendTune>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     */
    description?: pulumi.Input<string>;
    /**
     * The organization configured users must be part of.
     */
    organization: pulumi.Input<string>;
    /**
     * The ID of the organization users must be part of.
     * Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
     */
    organizationId?: pulumi.Input<number>;
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     */
    path?: pulumi.Input<string>;
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * (Optional) The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    tokenType?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     */
    tune?: pulumi.Input<inputs.github.AuthBackendTune>;
}
