// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages Token auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/token.html) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.tokenauth.AuthBackendRole("example", {
 *     allowedEntityAliases: ["test_entity"],
 *     allowedPolicies: [
 *         "dev",
 *         "test",
 *     ],
 *     disallowedPolicies: ["default"],
 *     orphan: true,
 *     pathSuffix: "path-suffix",
 *     renewable: true,
 *     roleName: "my-role",
 *     tokenExplicitMaxTtl: 115200,
 *     tokenPeriod: 86400,
 * });
 * ```
 *
 * ## Import
 *
 * Token auth backend roles can be imported with `auth/token/roles/` followed by the `role_name`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:tokenauth/authBackendRole:AuthBackendRole example auth/token/roles/my-role
 * ```
 */
export class AuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendRole {
        return new AuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:tokenauth/authBackendRole:AuthBackendRole';

    /**
     * Returns true if the given object is an instance of AuthBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRole.__pulumiType;
    }

    /**
     * List of allowed entity aliases.
     */
    public readonly allowedEntityAliases!: pulumi.Output<string[] | undefined>;
    /**
     * List of allowed policies for given role.
     */
    public readonly allowedPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * Set of allowed policies with glob match for given role.
     */
    public readonly allowedPoliciesGlobs!: pulumi.Output<string[] | undefined>;
    /**
     * List of disallowed policies for given role.
     */
    public readonly disallowedPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * Set of disallowed policies with glob match for given role.
     */
    public readonly disallowedPoliciesGlobs!: pulumi.Output<string[] | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    public readonly orphan!: pulumi.Output<boolean | undefined>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     *
     * > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
     */
    public readonly pathSuffix!: pulumi.Output<string | undefined>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    public readonly renewable!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleArgs | AuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            resourceInputs["allowedEntityAliases"] = state ? state.allowedEntityAliases : undefined;
            resourceInputs["allowedPolicies"] = state ? state.allowedPolicies : undefined;
            resourceInputs["allowedPoliciesGlobs"] = state ? state.allowedPoliciesGlobs : undefined;
            resourceInputs["disallowedPolicies"] = state ? state.disallowedPolicies : undefined;
            resourceInputs["disallowedPoliciesGlobs"] = state ? state.disallowedPoliciesGlobs : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["orphan"] = state ? state.orphan : undefined;
            resourceInputs["pathSuffix"] = state ? state.pathSuffix : undefined;
            resourceInputs["renewable"] = state ? state.renewable : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["allowedEntityAliases"] = args ? args.allowedEntityAliases : undefined;
            resourceInputs["allowedPolicies"] = args ? args.allowedPolicies : undefined;
            resourceInputs["allowedPoliciesGlobs"] = args ? args.allowedPoliciesGlobs : undefined;
            resourceInputs["disallowedPolicies"] = args ? args.disallowedPolicies : undefined;
            resourceInputs["disallowedPoliciesGlobs"] = args ? args.disallowedPoliciesGlobs : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["orphan"] = args ? args.orphan : undefined;
            resourceInputs["pathSuffix"] = args ? args.pathSuffix : undefined;
            resourceInputs["renewable"] = args ? args.renewable : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    /**
     * List of allowed entity aliases.
     */
    allowedEntityAliases?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed policies for given role.
     */
    allowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of allowed policies with glob match for given role.
     */
    allowedPoliciesGlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of disallowed policies with glob match for given role.
     */
    disallowedPoliciesGlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    orphan?: pulumi.Input<boolean>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     *
     * > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
     */
    pathSuffix?: pulumi.Input<string>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    roleName?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * List of allowed entity aliases.
     */
    allowedEntityAliases?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed policies for given role.
     */
    allowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of allowed policies with glob match for given role.
     */
    allowedPoliciesGlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of disallowed policies with glob match for given role.
     */
    disallowedPoliciesGlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    orphan?: pulumi.Input<boolean>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     *
     * > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
     */
    pathSuffix?: pulumi.Input<string>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    roleName: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: pulumi.Input<string>;
}
