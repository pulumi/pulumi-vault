// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an JWT/OIDC auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/jwt.html) for more
 * information.
 *
 * ## Example Usage
 *
 * Role for JWT backend:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const jwt = new vault.jwt.AuthBackend("jwt", {path: "jwt"});
 * const example = new vault.jwt.AuthBackendRole("example", {
 *     backend: jwt.path,
 *     roleName: "test-role",
 *     tokenPolicies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     boundAudiences: ["https://myco.test"],
 *     boundClaims: {
 *         color: "red,green,blue",
 *     },
 *     userClaim: "https://vault/user",
 *     roleType: "jwt",
 * });
 * ```
 *
 * Role for OIDC backend:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const oidc = new vault.jwt.AuthBackend("oidc", {
 *     path: "oidc",
 *     defaultRole: "test-role",
 * });
 * const example = new vault.jwt.AuthBackendRole("example", {
 *     backend: oidc.path,
 *     roleName: "test-role",
 *     tokenPolicies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     userClaim: "https://vault/user",
 *     roleType: "oidc",
 *     allowedRedirectUris: ["http://localhost:8200/ui/vault/auth/oidc/oidc/callback"],
 * });
 * ```
 *
 * ## Import
 *
 * JWT authentication backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:jwt/authBackendRole:AuthBackendRole example auth/jwt/role/test-role
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
    public static readonly __pulumiType = 'vault:jwt/authBackendRole:AuthBackendRole';

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
     * The list of allowed values for redirectUri during OIDC logins.
     * Required for OIDC roles
     */
    public readonly allowedRedirectUris!: pulumi.Output<string[] | undefined>;
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `jwt`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * (For "jwt" roles, at least one of `boundAudiences`, `boundSubject`, `boundClaims`
     * or `tokenBoundCidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
     * Any match is sufficient.
     */
    public readonly boundAudiences!: pulumi.Output<string[] | undefined>;
    /**
     * If set, a map of claims to values to match against.
     * A claim's value must be a string, which may contain one value or multiple
     * comma-separated values, e.g. `"red"` or `"red,green,blue"`.
     */
    public readonly boundClaims!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * How to interpret values in the claims/values
     * map (`boundClaims`): can be either `string` (exact match) or `glob` (wildcard
     * match). Requires Vault 1.4.0 or above.
     */
    public readonly boundClaimsType!: pulumi.Output<string>;
    /**
     * If set, requires that the `sub` claim matches
     * this value.
     */
    public readonly boundSubject!: pulumi.Output<string | undefined>;
    /**
     * If set, a map of claims (keys) to be copied
     * to specified metadata fields (values).
     */
    public readonly claimMappings!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The amount of leeway to add to all claims to account for clock skew, in
     * seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    public readonly clockSkewLeeway!: pulumi.Output<number | undefined>;
    /**
     * Disable bound claim value parsing. Useful when values contain commas.
     */
    public readonly disableBoundClaimsParsing!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of leeway to add to expiration (`exp`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    public readonly expirationLeeway!: pulumi.Output<number | undefined>;
    /**
     * The claim to use to uniquely identify
     * the set of groups to which the user belongs; this will be used as the names
     * for the Identity group aliases created due to a successful login. The claim
     * value must be a list of strings.
     */
    public readonly groupsClaim!: pulumi.Output<string | undefined>;
    /**
     * Specifies the allowable elapsed time in seconds since the last time 
     * the user was actively authenticated with the OIDC provider.
     */
    public readonly maxAge!: pulumi.Output<number | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The amount of leeway to add to not before (`nbf`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    public readonly notBeforeLeeway!: pulumi.Output<number | undefined>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role.
     * The standard scope "openid" is automatically included and need not be specified.
     */
    public readonly oidcScopes!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * Type of role, either "oidc" (default) or "jwt".
     */
    public readonly roleType!: pulumi.Output<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime of the generated token
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Period
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token to generate, service or batch
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    public readonly userClaim!: pulumi.Output<string>;
    /**
     * Specifies if the `userClaim` value uses
     * [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
     * syntax for referencing claims. By default, the `userClaim` value will not use JSON pointer.
     * Requires Vault 1.11+.
     */
    public readonly userClaimJsonPointer!: pulumi.Output<boolean | undefined>;
    /**
     * Log received OIDC tokens and claims when debug-level
     * logging is active. Not recommended in production since sensitive information may be present
     * in OIDC responses.
     */
    public readonly verboseOidcLogging!: pulumi.Output<boolean | undefined>;

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
            resourceInputs["allowedRedirectUris"] = state ? state.allowedRedirectUris : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["boundAudiences"] = state ? state.boundAudiences : undefined;
            resourceInputs["boundClaims"] = state ? state.boundClaims : undefined;
            resourceInputs["boundClaimsType"] = state ? state.boundClaimsType : undefined;
            resourceInputs["boundSubject"] = state ? state.boundSubject : undefined;
            resourceInputs["claimMappings"] = state ? state.claimMappings : undefined;
            resourceInputs["clockSkewLeeway"] = state ? state.clockSkewLeeway : undefined;
            resourceInputs["disableBoundClaimsParsing"] = state ? state.disableBoundClaimsParsing : undefined;
            resourceInputs["expirationLeeway"] = state ? state.expirationLeeway : undefined;
            resourceInputs["groupsClaim"] = state ? state.groupsClaim : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["notBeforeLeeway"] = state ? state.notBeforeLeeway : undefined;
            resourceInputs["oidcScopes"] = state ? state.oidcScopes : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["roleType"] = state ? state.roleType : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
            resourceInputs["userClaim"] = state ? state.userClaim : undefined;
            resourceInputs["userClaimJsonPointer"] = state ? state.userClaimJsonPointer : undefined;
            resourceInputs["verboseOidcLogging"] = state ? state.verboseOidcLogging : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            if ((!args || args.userClaim === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userClaim'");
            }
            resourceInputs["allowedRedirectUris"] = args ? args.allowedRedirectUris : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["boundAudiences"] = args ? args.boundAudiences : undefined;
            resourceInputs["boundClaims"] = args ? args.boundClaims : undefined;
            resourceInputs["boundClaimsType"] = args ? args.boundClaimsType : undefined;
            resourceInputs["boundSubject"] = args ? args.boundSubject : undefined;
            resourceInputs["claimMappings"] = args ? args.claimMappings : undefined;
            resourceInputs["clockSkewLeeway"] = args ? args.clockSkewLeeway : undefined;
            resourceInputs["disableBoundClaimsParsing"] = args ? args.disableBoundClaimsParsing : undefined;
            resourceInputs["expirationLeeway"] = args ? args.expirationLeeway : undefined;
            resourceInputs["groupsClaim"] = args ? args.groupsClaim : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["notBeforeLeeway"] = args ? args.notBeforeLeeway : undefined;
            resourceInputs["oidcScopes"] = args ? args.oidcScopes : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["roleType"] = args ? args.roleType : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
            resourceInputs["userClaim"] = args ? args.userClaim : undefined;
            resourceInputs["userClaimJsonPointer"] = args ? args.userClaimJsonPointer : undefined;
            resourceInputs["verboseOidcLogging"] = args ? args.verboseOidcLogging : undefined;
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
     * The list of allowed values for redirectUri during OIDC logins.
     * Required for OIDC roles
     */
    allowedRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `jwt`.
     */
    backend?: pulumi.Input<string>;
    /**
     * (For "jwt" roles, at least one of `boundAudiences`, `boundSubject`, `boundClaims`
     * or `tokenBoundCidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
     * Any match is sufficient.
     */
    boundAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a map of claims to values to match against.
     * A claim's value must be a string, which may contain one value or multiple
     * comma-separated values, e.g. `"red"` or `"red,green,blue"`.
     */
    boundClaims?: pulumi.Input<{[key: string]: any}>;
    /**
     * How to interpret values in the claims/values
     * map (`boundClaims`): can be either `string` (exact match) or `glob` (wildcard
     * match). Requires Vault 1.4.0 or above.
     */
    boundClaimsType?: pulumi.Input<string>;
    /**
     * If set, requires that the `sub` claim matches
     * this value.
     */
    boundSubject?: pulumi.Input<string>;
    /**
     * If set, a map of claims (keys) to be copied
     * to specified metadata fields (values).
     */
    claimMappings?: pulumi.Input<{[key: string]: any}>;
    /**
     * The amount of leeway to add to all claims to account for clock skew, in
     * seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    clockSkewLeeway?: pulumi.Input<number>;
    /**
     * Disable bound claim value parsing. Useful when values contain commas.
     */
    disableBoundClaimsParsing?: pulumi.Input<boolean>;
    /**
     * The amount of leeway to add to expiration (`exp`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    expirationLeeway?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the set of groups to which the user belongs; this will be used as the names
     * for the Identity group aliases created due to a successful login. The claim
     * value must be a list of strings.
     */
    groupsClaim?: pulumi.Input<string>;
    /**
     * Specifies the allowable elapsed time in seconds since the last time 
     * the user was actively authenticated with the OIDC provider.
     */
    maxAge?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The amount of leeway to add to not before (`nbf`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    notBeforeLeeway?: pulumi.Input<number>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role.
     * The standard scope "openid" is automatically included and need not be specified.
     */
    oidcScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    roleName?: pulumi.Input<string>;
    /**
     * Type of role, either "oidc" (default) or "jwt".
     */
    roleType?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    userClaim?: pulumi.Input<string>;
    /**
     * Specifies if the `userClaim` value uses
     * [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
     * syntax for referencing claims. By default, the `userClaim` value will not use JSON pointer.
     * Requires Vault 1.11+.
     */
    userClaimJsonPointer?: pulumi.Input<boolean>;
    /**
     * Log received OIDC tokens and claims when debug-level
     * logging is active. Not recommended in production since sensitive information may be present
     * in OIDC responses.
     */
    verboseOidcLogging?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * The list of allowed values for redirectUri during OIDC logins.
     * Required for OIDC roles
     */
    allowedRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `jwt`.
     */
    backend?: pulumi.Input<string>;
    /**
     * (For "jwt" roles, at least one of `boundAudiences`, `boundSubject`, `boundClaims`
     * or `tokenBoundCidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
     * Any match is sufficient.
     */
    boundAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a map of claims to values to match against.
     * A claim's value must be a string, which may contain one value or multiple
     * comma-separated values, e.g. `"red"` or `"red,green,blue"`.
     */
    boundClaims?: pulumi.Input<{[key: string]: any}>;
    /**
     * How to interpret values in the claims/values
     * map (`boundClaims`): can be either `string` (exact match) or `glob` (wildcard
     * match). Requires Vault 1.4.0 or above.
     */
    boundClaimsType?: pulumi.Input<string>;
    /**
     * If set, requires that the `sub` claim matches
     * this value.
     */
    boundSubject?: pulumi.Input<string>;
    /**
     * If set, a map of claims (keys) to be copied
     * to specified metadata fields (values).
     */
    claimMappings?: pulumi.Input<{[key: string]: any}>;
    /**
     * The amount of leeway to add to all claims to account for clock skew, in
     * seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    clockSkewLeeway?: pulumi.Input<number>;
    /**
     * Disable bound claim value parsing. Useful when values contain commas.
     */
    disableBoundClaimsParsing?: pulumi.Input<boolean>;
    /**
     * The amount of leeway to add to expiration (`exp`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    expirationLeeway?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the set of groups to which the user belongs; this will be used as the names
     * for the Identity group aliases created due to a successful login. The claim
     * value must be a list of strings.
     */
    groupsClaim?: pulumi.Input<string>;
    /**
     * Specifies the allowable elapsed time in seconds since the last time 
     * the user was actively authenticated with the OIDC provider.
     */
    maxAge?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The amount of leeway to add to not before (`nbf`) claims to account for
     * clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
     * Only applicable with "jwt" roles.
     */
    notBeforeLeeway?: pulumi.Input<number>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role.
     * The standard scope "openid" is automatically included and need not be specified.
     */
    oidcScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    roleName: pulumi.Input<string>;
    /**
     * Type of role, either "oidc" (default) or "jwt".
     */
    roleType?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    userClaim: pulumi.Input<string>;
    /**
     * Specifies if the `userClaim` value uses
     * [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
     * syntax for referencing claims. By default, the `userClaim` value will not use JSON pointer.
     * Requires Vault 1.11+.
     */
    userClaimJsonPointer?: pulumi.Input<boolean>;
    /**
     * Log received OIDC tokens and claims when debug-level
     * logging is active. Not recommended in production since sensitive information may be present
     * in OIDC responses.
     */
    verboseOidcLogging?: pulumi.Input<boolean>;
}
