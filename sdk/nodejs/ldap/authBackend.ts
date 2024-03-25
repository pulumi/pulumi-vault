// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const ldap = new vault.ldap.AuthBackend("ldap", {
 *     discoverdn: false,
 *     groupdn: "OU=Groups,DC=example,DC=org",
 *     groupfilter: "(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
 *     path: "ldap",
 *     upndomain: "EXAMPLE.ORG",
 *     url: "ldaps://dc-01.example.org",
 *     userattr: "sAMAccountName",
 *     userdn: "OU=Users,OU=Accounts,DC=example,DC=org",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * LDAP authentication backends can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:ldap/authBackend:AuthBackend ldap ldap
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
    public static readonly __pulumiType = 'vault:ldap/authBackend:AuthBackend';

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
     * The accessor for this auth mount.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * DN of object to bind when performing user search
     */
    public readonly binddn!: pulumi.Output<string>;
    /**
     * Password to use with `binddn` when performing user search
     */
    public readonly bindpass!: pulumi.Output<string>;
    /**
     * Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     */
    public readonly caseSensitiveNames!: pulumi.Output<boolean>;
    /**
     * Trusted CA to validate TLS certificate
     */
    public readonly certificate!: pulumi.Output<string>;
    public readonly clientTlsCert!: pulumi.Output<string>;
    public readonly clientTlsKey!: pulumi.Output<string>;
    /**
     * Prevents users from bypassing authentication when providing an empty password.
     */
    public readonly denyNullBind!: pulumi.Output<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Use anonymous bind to discover the bind DN of a user.
     */
    public readonly discoverdn!: pulumi.Output<boolean>;
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     */
    public readonly groupattr!: pulumi.Output<string>;
    /**
     * Base DN under which to perform group search
     */
    public readonly groupdn!: pulumi.Output<string>;
    /**
     * Go template used to construct group membership query
     */
    public readonly groupfilter!: pulumi.Output<string>;
    /**
     * Control whether or TLS certificates must be validated
     */
    public readonly insecureTls!: pulumi.Output<boolean>;
    /**
     * Specifies if the auth method is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Sets the max page size for LDAP lookups, by default it's set to -1.
     * *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
     */
    public readonly maxPageSize!: pulumi.Output<number | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Path to mount the LDAP auth backend under
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * Control use of TLS when conecting to LDAP
     */
    public readonly starttls!: pulumi.Output<boolean>;
    /**
     * Maximum acceptable version of TLS
     */
    public readonly tlsMaxVersion!: pulumi.Output<string>;
    /**
     * Minimum acceptable version of TLS
     */
    public readonly tlsMinVersion!: pulumi.Output<string>;
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
     * The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
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
     * The type of token to generate, service or batch
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
     */
    public readonly upndomain!: pulumi.Output<string>;
    /**
     * The URL of the LDAP server
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     */
    public readonly useTokenGroups!: pulumi.Output<boolean>;
    /**
     * Attribute on user object matching username passed in
     */
    public readonly userattr!: pulumi.Output<string>;
    /**
     * Base DN under which to perform user search
     */
    public readonly userdn!: pulumi.Output<string>;
    /**
     * LDAP user search filter
     */
    public readonly userfilter!: pulumi.Output<string>;
    /**
     * Force the auth method to use the username passed by the user as the alias name.
     */
    public readonly usernameAsAlias!: pulumi.Output<boolean>;

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
            resourceInputs["binddn"] = state ? state.binddn : undefined;
            resourceInputs["bindpass"] = state ? state.bindpass : undefined;
            resourceInputs["caseSensitiveNames"] = state ? state.caseSensitiveNames : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["clientTlsCert"] = state ? state.clientTlsCert : undefined;
            resourceInputs["clientTlsKey"] = state ? state.clientTlsKey : undefined;
            resourceInputs["denyNullBind"] = state ? state.denyNullBind : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["discoverdn"] = state ? state.discoverdn : undefined;
            resourceInputs["groupattr"] = state ? state.groupattr : undefined;
            resourceInputs["groupdn"] = state ? state.groupdn : undefined;
            resourceInputs["groupfilter"] = state ? state.groupfilter : undefined;
            resourceInputs["insecureTls"] = state ? state.insecureTls : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxPageSize"] = state ? state.maxPageSize : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["starttls"] = state ? state.starttls : undefined;
            resourceInputs["tlsMaxVersion"] = state ? state.tlsMaxVersion : undefined;
            resourceInputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
            resourceInputs["upndomain"] = state ? state.upndomain : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["useTokenGroups"] = state ? state.useTokenGroups : undefined;
            resourceInputs["userattr"] = state ? state.userattr : undefined;
            resourceInputs["userdn"] = state ? state.userdn : undefined;
            resourceInputs["userfilter"] = state ? state.userfilter : undefined;
            resourceInputs["usernameAsAlias"] = state ? state.usernameAsAlias : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["binddn"] = args ? args.binddn : undefined;
            resourceInputs["bindpass"] = args?.bindpass ? pulumi.secret(args.bindpass) : undefined;
            resourceInputs["caseSensitiveNames"] = args ? args.caseSensitiveNames : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["clientTlsCert"] = args ? args.clientTlsCert : undefined;
            resourceInputs["clientTlsKey"] = args?.clientTlsKey ? pulumi.secret(args.clientTlsKey) : undefined;
            resourceInputs["denyNullBind"] = args ? args.denyNullBind : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["discoverdn"] = args ? args.discoverdn : undefined;
            resourceInputs["groupattr"] = args ? args.groupattr : undefined;
            resourceInputs["groupdn"] = args ? args.groupdn : undefined;
            resourceInputs["groupfilter"] = args ? args.groupfilter : undefined;
            resourceInputs["insecureTls"] = args ? args.insecureTls : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxPageSize"] = args ? args.maxPageSize : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["starttls"] = args ? args.starttls : undefined;
            resourceInputs["tlsMaxVersion"] = args ? args.tlsMaxVersion : undefined;
            resourceInputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
            resourceInputs["upndomain"] = args ? args.upndomain : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["useTokenGroups"] = args ? args.useTokenGroups : undefined;
            resourceInputs["userattr"] = args ? args.userattr : undefined;
            resourceInputs["userdn"] = args ? args.userdn : undefined;
            resourceInputs["userfilter"] = args ? args.userfilter : undefined;
            resourceInputs["usernameAsAlias"] = args ? args.usernameAsAlias : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["bindpass", "clientTlsKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The accessor for this auth mount.
     */
    accessor?: pulumi.Input<string>;
    /**
     * DN of object to bind when performing user search
     */
    binddn?: pulumi.Input<string>;
    /**
     * Password to use with `binddn` when performing user search
     */
    bindpass?: pulumi.Input<string>;
    /**
     * Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     */
    caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * Trusted CA to validate TLS certificate
     */
    certificate?: pulumi.Input<string>;
    clientTlsCert?: pulumi.Input<string>;
    clientTlsKey?: pulumi.Input<string>;
    /**
     * Prevents users from bypassing authentication when providing an empty password.
     */
    denyNullBind?: pulumi.Input<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Use anonymous bind to discover the bind DN of a user.
     */
    discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     */
    groupattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform group search
     */
    groupdn?: pulumi.Input<string>;
    /**
     * Go template used to construct group membership query
     */
    groupfilter?: pulumi.Input<string>;
    /**
     * Control whether or TLS certificates must be validated
     */
    insecureTls?: pulumi.Input<boolean>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Sets the max page size for LDAP lookups, by default it's set to -1.
     * *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
     */
    maxPageSize?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Path to mount the LDAP auth backend under
     */
    path?: pulumi.Input<string>;
    /**
     * Control use of TLS when conecting to LDAP
     */
    starttls?: pulumi.Input<boolean>;
    /**
     * Maximum acceptable version of TLS
     */
    tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum acceptable version of TLS
     */
    tlsMinVersion?: pulumi.Input<string>;
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
     * The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
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
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
     */
    upndomain?: pulumi.Input<string>;
    /**
     * The URL of the LDAP server
     */
    url?: pulumi.Input<string>;
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     */
    useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute on user object matching username passed in
     */
    userattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform user search
     */
    userdn?: pulumi.Input<string>;
    /**
     * LDAP user search filter
     */
    userfilter?: pulumi.Input<string>;
    /**
     * Force the auth method to use the username passed by the user as the alias name.
     */
    usernameAsAlias?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * DN of object to bind when performing user search
     */
    binddn?: pulumi.Input<string>;
    /**
     * Password to use with `binddn` when performing user search
     */
    bindpass?: pulumi.Input<string>;
    /**
     * Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     */
    caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * Trusted CA to validate TLS certificate
     */
    certificate?: pulumi.Input<string>;
    clientTlsCert?: pulumi.Input<string>;
    clientTlsKey?: pulumi.Input<string>;
    /**
     * Prevents users from bypassing authentication when providing an empty password.
     */
    denyNullBind?: pulumi.Input<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Use anonymous bind to discover the bind DN of a user.
     */
    discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     */
    groupattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform group search
     */
    groupdn?: pulumi.Input<string>;
    /**
     * Go template used to construct group membership query
     */
    groupfilter?: pulumi.Input<string>;
    /**
     * Control whether or TLS certificates must be validated
     */
    insecureTls?: pulumi.Input<boolean>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Sets the max page size for LDAP lookups, by default it's set to -1.
     * *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
     */
    maxPageSize?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Path to mount the LDAP auth backend under
     */
    path?: pulumi.Input<string>;
    /**
     * Control use of TLS when conecting to LDAP
     */
    starttls?: pulumi.Input<boolean>;
    /**
     * Maximum acceptable version of TLS
     */
    tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum acceptable version of TLS
     */
    tlsMinVersion?: pulumi.Input<string>;
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
     * The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
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
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
     */
    upndomain?: pulumi.Input<string>;
    /**
     * The URL of the LDAP server
     */
    url: pulumi.Input<string>;
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     */
    useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute on user object matching username passed in
     */
    userattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform user search
     */
    userdn?: pulumi.Input<string>;
    /**
     * LDAP user search filter
     */
    userfilter?: pulumi.Input<string>;
    /**
     * Force the auth method to use the username passed by the user as the alias name.
     */
    usernameAsAlias?: pulumi.Input<boolean>;
}
