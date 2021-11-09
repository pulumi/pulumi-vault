// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AppRole auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/approle) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const approle = new vault.AuthBackend("approle", {type: "approle"});
 * const example = new vault.approle.AuthBackendRole("example", {
 *     backend: approle.path,
 *     roleName: "test-role",
 *     tokenPolicies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * AppRole authentication backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:appRole/authBackendRole:AuthBackendRole example auth/approle/role/test-role
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
    public static readonly __pulumiType = 'vault:appRole/authBackendRole:AuthBackendRole';

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
     * The unique name of the auth backend to configure.
     * Defaults to `approle`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Whether or not to require `secretId` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     */
    public readonly bindSecretId!: pulumi.Output<boolean | undefined>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     *
     * @deprecated use `secret_id_bound_cidrs` instead
     */
    public readonly boundCidrLists!: pulumi.Output<string[] | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * The RoleID of this role. If not specified, one will be
     * auto-generated.
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * The name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     */
    public readonly secretIdBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     */
    public readonly secretIdNumUses!: pulumi.Output<number | undefined>;
    /**
     * The number of seconds after which any SecretID
     * expires.
     */
    public readonly secretIdTtl!: pulumi.Output<number | undefined>;
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
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
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
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["bindSecretId"] = state ? state.bindSecretId : undefined;
            inputs["boundCidrLists"] = state ? state.boundCidrLists : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["secretIdBoundCidrs"] = state ? state.secretIdBoundCidrs : undefined;
            inputs["secretIdNumUses"] = state ? state.secretIdNumUses : undefined;
            inputs["secretIdTtl"] = state ? state.secretIdTtl : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["bindSecretId"] = args ? args.bindSecretId : undefined;
            inputs["boundCidrLists"] = args ? args.boundCidrLists : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["secretIdBoundCidrs"] = args ? args.secretIdBoundCidrs : undefined;
            inputs["secretIdNumUses"] = args ? args.secretIdNumUses : undefined;
            inputs["secretIdTtl"] = args ? args.secretIdTtl : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `approle`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Whether or not to require `secretId` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     */
    bindSecretId?: pulumi.Input<boolean>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     *
     * @deprecated use `secret_id_bound_cidrs` instead
     */
    boundCidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The RoleID of this role. If not specified, one will be
     * auto-generated.
     */
    roleId?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    roleName?: pulumi.Input<string>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     */
    secretIdBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     */
    secretIdNumUses?: pulumi.Input<number>;
    /**
     * The number of seconds after which any SecretID
     * expires.
     */
    secretIdTtl?: pulumi.Input<number>;
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
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
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
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
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
     * The unique name of the auth backend to configure.
     * Defaults to `approle`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Whether or not to require `secretId` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     */
    bindSecretId?: pulumi.Input<boolean>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     *
     * @deprecated use `secret_id_bound_cidrs` instead
     */
    boundCidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The RoleID of this role. If not specified, one will be
     * auto-generated.
     */
    roleId?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    roleName: pulumi.Input<string>;
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     */
    secretIdBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     */
    secretIdNumUses?: pulumi.Input<number>;
    /**
     * The number of seconds after which any SecretID
     * expires.
     */
    secretIdTtl?: pulumi.Input<number>;
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
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
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
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
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
