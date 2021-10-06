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
     * (Optional; Deprecated, use `tokenMaxTtl` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens
     * issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * The organization configured users must be part of.
     */
    public readonly organization!: pulumi.Output<string>;
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
     * (Optional; Deprecated, use `tokenTtl` instead if you are running Vault >= 1.2) The TTL period of tokens issued
     * using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["baseUrl"] = state ? state.baseUrl : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["organization"] = state ? state.organization : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["tune"] = state ? state.tune : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            inputs["baseUrl"] = args ? args.baseUrl : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["organization"] = args ? args.organization : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["tune"] = args ? args.tune : undefined;
            inputs["accessor"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     */
    readonly baseUrl?: pulumi.Input<string>;
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated, use `tokenMaxTtl` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens
     * issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * The organization configured users must be part of.
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * (Optional) The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated, use `tokenTtl` instead if you are running Vault >= 1.2) The TTL period of tokens issued
     * using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     */
    readonly tune?: pulumi.Input<inputs.github.AuthBackendTune>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     */
    readonly baseUrl?: pulumi.Input<string>;
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated, use `tokenMaxTtl` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens
     * issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * The organization configured users must be part of.
     */
    readonly organization: pulumi.Input<string>;
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * (Optional) The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated, use `tokenTtl` instead if you are running Vault >= 1.2) The TTL period of tokens issued
     * using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     */
    readonly tune?: pulumi.Input<inputs.github.AuthBackendTune>;
}
