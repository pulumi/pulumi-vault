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
 * const jwt = new vault.jwt.AuthBackend("jwt", {
 *     path: "jwt",
 * });
 * const example = new vault.jwt.AuthBackendRole("example", {
 *     backend: jwt.path,
 *     boundAudiences: ["https://myco.test"],
 *     policies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     roleName: "test-role",
 *     roleType: "jwt",
 *     userClaim: "https://vault/user",
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
 *     defaultRole: "test-role",
 *     path: "oidc",
 * });
 * const example = new vault.jwt.AuthBackendRole("example", {
 *     allowedRedirectUris: ["http://localhost:8200/ui/vault/auth/oidc/oidc/callback"],
 *     backend: oidc.path,
 *     boundAudiences: ["https://myco.test"],
 *     policies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     roleName: "test-role",
 *     roleType: "oidc",
 *     userClaim: "https://vault/user",
 * });
 * ```
 *
 * ## Import
 *
 * JWT authentication backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:jwt/authBackendRole:AuthBackendRole example auth/jwt/role/test-role
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
     * List of `aud` claims to match
     * against. Any match is sufficient.
     */
    public readonly boundAudiences!: pulumi.Output<string[] | undefined>;
    /**
     * If set, a list of CIDRs valid as the source
     * address for login requests. This value is also encoded into any resulting
     * token.
     *
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    public readonly boundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * If set, a map of claims/values to match against. 
     * The expected value may be a single string or a list of strings.
     */
    public readonly boundClaims!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * How to interpret values in the claims/values map: can be either "string" (exact match) or "glob" (wildcard match).
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
     * The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and
     * can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    public readonly clockSkewLeeway!: pulumi.Output<number | undefined>;
    /**
     * The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 60 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
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
     * (Optional; Deprecated. This field has been
     * removed since Vault 1.1. If the groups claim is not at the top level, it can
     * now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
     * A pattern of delimiters
     * used to allow the groupsClaim to live outside of the top-level JWT structure.
     * For instance, a groupsClaim of meta/user.name/groups with this field
     * set to // will expect nested structures named meta, user.name, and groups.
     * If this field was set to /./ the groups information would expect to be
     * via nested structures of meta, user, name, and groups.
     *
     * @deprecated `groups_claim_delimiter_pattern` has been removed since Vault 1.1. If the groups claim is not at the top level, it can now be specified as a JSONPointer.
     */
    public readonly groupsClaimDelimiterPattern!: pulumi.Output<string | undefined>;
    /**
     * The maximum allowed lifetime of tokens issued using
     * this role, in seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    /**
     * The amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds. Defaults to 150 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    public readonly notBeforeLeeway!: pulumi.Output<number | undefined>;
    /**
     * If set, puts a use-count limitation on the issued
     * token.
     *
     * @deprecated use `token_num_uses` instead if you are running Vault >= 1.2
     */
    public readonly numUses!: pulumi.Output<number | undefined>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role. 
     * The standard scope "openid" is automatically included and need not be specified.
     */
    public readonly oidcScopes!: pulumi.Output<string[] | undefined>;
    /**
     * If set, indicates that the token generated
     * using this role should never expire, but instead always use the value set
     * here as the TTL for every renewal.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
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
     * The initial/renewal TTL of tokens issued using this role,
     * in seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    public readonly userClaim!: pulumi.Output<string>;
    /**
     * Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive
     * information may be present in OIDC responses.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            inputs["allowedRedirectUris"] = state ? state.allowedRedirectUris : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["boundAudiences"] = state ? state.boundAudiences : undefined;
            inputs["boundCidrs"] = state ? state.boundCidrs : undefined;
            inputs["boundClaims"] = state ? state.boundClaims : undefined;
            inputs["boundClaimsType"] = state ? state.boundClaimsType : undefined;
            inputs["boundSubject"] = state ? state.boundSubject : undefined;
            inputs["claimMappings"] = state ? state.claimMappings : undefined;
            inputs["clockSkewLeeway"] = state ? state.clockSkewLeeway : undefined;
            inputs["expirationLeeway"] = state ? state.expirationLeeway : undefined;
            inputs["groupsClaim"] = state ? state.groupsClaim : undefined;
            inputs["groupsClaimDelimiterPattern"] = state ? state.groupsClaimDelimiterPattern : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["notBeforeLeeway"] = state ? state.notBeforeLeeway : undefined;
            inputs["numUses"] = state ? state.numUses : undefined;
            inputs["oidcScopes"] = state ? state.oidcScopes : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["roleType"] = state ? state.roleType : undefined;
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
            inputs["userClaim"] = state ? state.userClaim : undefined;
            inputs["verboseOidcLogging"] = state ? state.verboseOidcLogging : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if (!args || args.roleName === undefined) {
                throw new Error("Missing required property 'roleName'");
            }
            if (!args || args.userClaim === undefined) {
                throw new Error("Missing required property 'userClaim'");
            }
            inputs["allowedRedirectUris"] = args ? args.allowedRedirectUris : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["boundAudiences"] = args ? args.boundAudiences : undefined;
            inputs["boundCidrs"] = args ? args.boundCidrs : undefined;
            inputs["boundClaims"] = args ? args.boundClaims : undefined;
            inputs["boundClaimsType"] = args ? args.boundClaimsType : undefined;
            inputs["boundSubject"] = args ? args.boundSubject : undefined;
            inputs["claimMappings"] = args ? args.claimMappings : undefined;
            inputs["clockSkewLeeway"] = args ? args.clockSkewLeeway : undefined;
            inputs["expirationLeeway"] = args ? args.expirationLeeway : undefined;
            inputs["groupsClaim"] = args ? args.groupsClaim : undefined;
            inputs["groupsClaimDelimiterPattern"] = args ? args.groupsClaimDelimiterPattern : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["notBeforeLeeway"] = args ? args.notBeforeLeeway : undefined;
            inputs["numUses"] = args ? args.numUses : undefined;
            inputs["oidcScopes"] = args ? args.oidcScopes : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["roleType"] = args ? args.roleType : undefined;
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
            inputs["userClaim"] = args ? args.userClaim : undefined;
            inputs["verboseOidcLogging"] = args ? args.verboseOidcLogging : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendRole.__pulumiType, name, inputs, opts);
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
    readonly allowedRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `jwt`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * List of `aud` claims to match
     * against. Any match is sufficient.
     */
    readonly boundAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a list of CIDRs valid as the source
     * address for login requests. This value is also encoded into any resulting
     * token.
     *
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a map of claims/values to match against. 
     * The expected value may be a single string or a list of strings.
     */
    readonly boundClaims?: pulumi.Input<{[key: string]: any}>;
    /**
     * How to interpret values in the claims/values map: can be either "string" (exact match) or "glob" (wildcard match).
     */
    readonly boundClaimsType?: pulumi.Input<string>;
    /**
     * If set, requires that the `sub` claim matches
     * this value.
     */
    readonly boundSubject?: pulumi.Input<string>;
    /**
     * If set, a map of claims (keys) to be copied 
     * to specified metadata fields (values).
     */
    readonly claimMappings?: pulumi.Input<{[key: string]: any}>;
    /**
     * The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and
     * can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly clockSkewLeeway?: pulumi.Input<number>;
    /**
     * The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 60 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly expirationLeeway?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the set of groups to which the user belongs; this will be used as the names
     * for the Identity group aliases created due to a successful login. The claim
     * value must be a list of strings.
     */
    readonly groupsClaim?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated. This field has been
     * removed since Vault 1.1. If the groups claim is not at the top level, it can
     * now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
     * A pattern of delimiters
     * used to allow the groupsClaim to live outside of the top-level JWT structure.
     * For instance, a groupsClaim of meta/user.name/groups with this field
     * set to // will expect nested structures named meta, user.name, and groups.
     * If this field was set to /./ the groups information would expect to be
     * via nested structures of meta, user, name, and groups.
     *
     * @deprecated `groups_claim_delimiter_pattern` has been removed since Vault 1.1. If the groups claim is not at the top level, it can now be specified as a JSONPointer.
     */
    readonly groupsClaimDelimiterPattern?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens issued using
     * this role, in seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * The amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds. Defaults to 150 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly notBeforeLeeway?: pulumi.Input<number>;
    /**
     * If set, puts a use-count limitation on the issued
     * token.
     *
     * @deprecated use `token_num_uses` instead if you are running Vault >= 1.2
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role. 
     * The standard scope "openid" is automatically included and need not be specified.
     */
    readonly oidcScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, indicates that the token generated
     * using this role should never expire, but instead always use the value set
     * here as the TTL for every renewal.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * Type of role, either "oidc" (default) or "jwt".
     */
    readonly roleType?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The initial/renewal TTL of tokens issued using this role,
     * in seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    readonly userClaim?: pulumi.Input<string>;
    /**
     * Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive
     * information may be present in OIDC responses.
     */
    readonly verboseOidcLogging?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * The list of allowed values for redirectUri during OIDC logins.
     * Required for OIDC roles
     */
    readonly allowedRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `jwt`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * List of `aud` claims to match
     * against. Any match is sufficient.
     */
    readonly boundAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a list of CIDRs valid as the source
     * address for login requests. This value is also encoded into any resulting
     * token.
     *
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a map of claims/values to match against. 
     * The expected value may be a single string or a list of strings.
     */
    readonly boundClaims?: pulumi.Input<{[key: string]: any}>;
    /**
     * How to interpret values in the claims/values map: can be either "string" (exact match) or "glob" (wildcard match).
     */
    readonly boundClaimsType?: pulumi.Input<string>;
    /**
     * If set, requires that the `sub` claim matches
     * this value.
     */
    readonly boundSubject?: pulumi.Input<string>;
    /**
     * If set, a map of claims (keys) to be copied 
     * to specified metadata fields (values).
     */
    readonly claimMappings?: pulumi.Input<{[key: string]: any}>;
    /**
     * The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and
     * can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly clockSkewLeeway?: pulumi.Input<number>;
    /**
     * The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 60 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly expirationLeeway?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the set of groups to which the user belongs; this will be used as the names
     * for the Identity group aliases created due to a successful login. The claim
     * value must be a list of strings.
     */
    readonly groupsClaim?: pulumi.Input<string>;
    /**
     * (Optional; Deprecated. This field has been
     * removed since Vault 1.1. If the groups claim is not at the top level, it can
     * now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
     * A pattern of delimiters
     * used to allow the groupsClaim to live outside of the top-level JWT structure.
     * For instance, a groupsClaim of meta/user.name/groups with this field
     * set to // will expect nested structures named meta, user.name, and groups.
     * If this field was set to /./ the groups information would expect to be
     * via nested structures of meta, user, name, and groups.
     *
     * @deprecated `groups_claim_delimiter_pattern` has been removed since Vault 1.1. If the groups claim is not at the top level, it can now be specified as a JSONPointer.
     */
    readonly groupsClaimDelimiterPattern?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens issued using
     * this role, in seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * The amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds. Defaults to 150 seconds if
     * set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
     */
    readonly notBeforeLeeway?: pulumi.Input<number>;
    /**
     * If set, puts a use-count limitation on the issued
     * token.
     *
     * @deprecated use `token_num_uses` instead if you are running Vault >= 1.2
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * If set, a list of OIDC scopes to be used with an OIDC role. 
     * The standard scope "openid" is automatically included and need not be specified.
     */
    readonly oidcScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, indicates that the token generated
     * using this role should never expire, but instead always use the value set
     * here as the TTL for every renewal.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    readonly roleName: pulumi.Input<string>;
    /**
     * Type of role, either "oidc" (default) or "jwt".
     */
    readonly roleType?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The initial/renewal TTL of tokens issued using this role,
     * in seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The claim to use to uniquely identify
     * the user; this will be used as the name for the Identity entity alias created
     * due to a successful login.
     */
    readonly userClaim: pulumi.Input<string>;
    /**
     * Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive
     * information may be present in OIDC responses.
     */
    readonly verboseOidcLogging?: pulumi.Input<boolean>;
}
