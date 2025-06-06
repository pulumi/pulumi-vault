// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.Token("example", {
 *     roleName: "app",
 *     policies: [
 *         "policy1",
 *         "policy2",
 *     ],
 *     renewable: true,
 *     ttl: "24h",
 *     renewMinLease: 43200,
 *     renewIncrement: 86400,
 *     metadata: {
 *         purpose: "service-account",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Tokens can be imported using its `id` as accessor id, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/token:Token example <accessor_id>
 * ```
 */
export class Token extends pulumi.CustomResource {
    /**
     * Get an existing Token resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TokenState, opts?: pulumi.CustomResourceOptions): Token {
        return new Token(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/token:Token';

    /**
     * Returns true if the given object is an instance of Token.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Token {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Token.__pulumiType;
    }

    /**
     * String containing the client token if stored in present file
     */
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * String containing the token display name
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    public readonly explicitMaxTtl!: pulumi.Output<string | undefined>;
    /**
     * String containing the token lease duration if present in state file
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<number>;
    /**
     * String containing the token lease started time if present in state file
     */
    public /*out*/ readonly leaseStarted!: pulumi.Output<string>;
    /**
     * Metadata to be set on this token
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Flag to not attach the default policy to this token
     */
    public readonly noDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to create a token without parent
     */
    public readonly noParent!: pulumi.Output<boolean>;
    /**
     * The number of allowed uses of this token
     */
    public readonly numUses!: pulumi.Output<number>;
    /**
     * The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * List of policies to attach to this token
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * The renew increment. This is specified in seconds
     */
    public readonly renewIncrement!: pulumi.Output<number | undefined>;
    /**
     * The minimal lease to renew this token
     */
    public readonly renewMinLease!: pulumi.Output<number | undefined>;
    /**
     * Flag to allow to renew this token
     */
    public readonly renewable!: pulumi.Output<boolean>;
    /**
     * The token role name
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * The client wrapped token.
     */
    public /*out*/ readonly wrappedToken!: pulumi.Output<string>;
    /**
     * The client wrapping accessor.
     */
    public /*out*/ readonly wrappingAccessor!: pulumi.Output<string>;
    /**
     * The TTL period of the wrapped token.
     */
    public readonly wrappingTtl!: pulumi.Output<string | undefined>;

    /**
     * Create a Token resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenArgs | TokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TokenState | undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["explicitMaxTtl"] = state ? state.explicitMaxTtl : undefined;
            resourceInputs["leaseDuration"] = state ? state.leaseDuration : undefined;
            resourceInputs["leaseStarted"] = state ? state.leaseStarted : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["noDefaultPolicy"] = state ? state.noDefaultPolicy : undefined;
            resourceInputs["noParent"] = state ? state.noParent : undefined;
            resourceInputs["numUses"] = state ? state.numUses : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["renewIncrement"] = state ? state.renewIncrement : undefined;
            resourceInputs["renewMinLease"] = state ? state.renewMinLease : undefined;
            resourceInputs["renewable"] = state ? state.renewable : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["wrappedToken"] = state ? state.wrappedToken : undefined;
            resourceInputs["wrappingAccessor"] = state ? state.wrappingAccessor : undefined;
            resourceInputs["wrappingTtl"] = state ? state.wrappingTtl : undefined;
        } else {
            const args = argsOrState as TokenArgs | undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["explicitMaxTtl"] = args ? args.explicitMaxTtl : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["noDefaultPolicy"] = args ? args.noDefaultPolicy : undefined;
            resourceInputs["noParent"] = args ? args.noParent : undefined;
            resourceInputs["numUses"] = args ? args.numUses : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["renewIncrement"] = args ? args.renewIncrement : undefined;
            resourceInputs["renewMinLease"] = args ? args.renewMinLease : undefined;
            resourceInputs["renewable"] = args ? args.renewable : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["wrappingTtl"] = args ? args.wrappingTtl : undefined;
            resourceInputs["clientToken"] = undefined /*out*/;
            resourceInputs["leaseDuration"] = undefined /*out*/;
            resourceInputs["leaseStarted"] = undefined /*out*/;
            resourceInputs["wrappedToken"] = undefined /*out*/;
            resourceInputs["wrappingAccessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientToken", "wrappedToken", "wrappingAccessor"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Token.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Token resources.
 */
export interface TokenState {
    /**
     * String containing the client token if stored in present file
     */
    clientToken?: pulumi.Input<string>;
    /**
     * String containing the token display name
     */
    displayName?: pulumi.Input<string>;
    /**
     * The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    explicitMaxTtl?: pulumi.Input<string>;
    /**
     * String containing the token lease duration if present in state file
     */
    leaseDuration?: pulumi.Input<number>;
    /**
     * String containing the token lease started time if present in state file
     */
    leaseStarted?: pulumi.Input<string>;
    /**
     * Metadata to be set on this token
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Flag to not attach the default policy to this token
     */
    noDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * Flag to create a token without parent
     */
    noParent?: pulumi.Input<boolean>;
    /**
     * The number of allowed uses of this token
     */
    numUses?: pulumi.Input<number>;
    /**
     * The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    period?: pulumi.Input<string>;
    /**
     * List of policies to attach to this token
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The renew increment. This is specified in seconds
     */
    renewIncrement?: pulumi.Input<number>;
    /**
     * The minimal lease to renew this token
     */
    renewMinLease?: pulumi.Input<number>;
    /**
     * Flag to allow to renew this token
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * The token role name
     */
    roleName?: pulumi.Input<string>;
    /**
     * The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    ttl?: pulumi.Input<string>;
    /**
     * The client wrapped token.
     */
    wrappedToken?: pulumi.Input<string>;
    /**
     * The client wrapping accessor.
     */
    wrappingAccessor?: pulumi.Input<string>;
    /**
     * The TTL period of the wrapped token.
     */
    wrappingTtl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Token resource.
 */
export interface TokenArgs {
    /**
     * String containing the token display name
     */
    displayName?: pulumi.Input<string>;
    /**
     * The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    explicitMaxTtl?: pulumi.Input<string>;
    /**
     * Metadata to be set on this token
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Flag to not attach the default policy to this token
     */
    noDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * Flag to create a token without parent
     */
    noParent?: pulumi.Input<boolean>;
    /**
     * The number of allowed uses of this token
     */
    numUses?: pulumi.Input<number>;
    /**
     * The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    period?: pulumi.Input<string>;
    /**
     * List of policies to attach to this token
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The renew increment. This is specified in seconds
     */
    renewIncrement?: pulumi.Input<number>;
    /**
     * The minimal lease to renew this token
     */
    renewMinLease?: pulumi.Input<number>;
    /**
     * Flag to allow to renew this token
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * The token role name
     */
    roleName?: pulumi.Input<string>;
    /**
     * The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
     */
    ttl?: pulumi.Input<string>;
    /**
     * The TTL period of the wrapped token.
     */
    wrappingTtl?: pulumi.Input<string>;
}
