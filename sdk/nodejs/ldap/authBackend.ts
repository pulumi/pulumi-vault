// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
 *
 * ## Example Usage
 *
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
 *
 * ## Import
 *
 * LDAP authentication backends can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:ldap/authBackend:AuthBackend ldap ldap
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
     * Trusted CA to validate TLS certificate
     */
    public readonly certificate!: pulumi.Output<string>;
    public readonly denyNullBind!: pulumi.Output<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    public readonly description!: pulumi.Output<string>;
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
     * The userPrincipalDomain used to construct UPN string
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
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["binddn"] = state ? state.binddn : undefined;
            inputs["bindpass"] = state ? state.bindpass : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["denyNullBind"] = state ? state.denyNullBind : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["discoverdn"] = state ? state.discoverdn : undefined;
            inputs["groupattr"] = state ? state.groupattr : undefined;
            inputs["groupdn"] = state ? state.groupdn : undefined;
            inputs["groupfilter"] = state ? state.groupfilter : undefined;
            inputs["insecureTls"] = state ? state.insecureTls : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["starttls"] = state ? state.starttls : undefined;
            inputs["tlsMaxVersion"] = state ? state.tlsMaxVersion : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
            inputs["upndomain"] = state ? state.upndomain : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["useTokenGroups"] = state ? state.useTokenGroups : undefined;
            inputs["userattr"] = state ? state.userattr : undefined;
            inputs["userdn"] = state ? state.userdn : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.url === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'url'");
            }
            inputs["binddn"] = args ? args.binddn : undefined;
            inputs["bindpass"] = args ? args.bindpass : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["denyNullBind"] = args ? args.denyNullBind : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["discoverdn"] = args ? args.discoverdn : undefined;
            inputs["groupattr"] = args ? args.groupattr : undefined;
            inputs["groupdn"] = args ? args.groupdn : undefined;
            inputs["groupfilter"] = args ? args.groupfilter : undefined;
            inputs["insecureTls"] = args ? args.insecureTls : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["starttls"] = args ? args.starttls : undefined;
            inputs["tlsMaxVersion"] = args ? args.tlsMaxVersion : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
            inputs["upndomain"] = args ? args.upndomain : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["useTokenGroups"] = args ? args.useTokenGroups : undefined;
            inputs["userattr"] = args ? args.userattr : undefined;
            inputs["userdn"] = args ? args.userdn : undefined;
            inputs["accessor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The accessor for this auth mount.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * DN of object to bind when performing user search
     */
    readonly binddn?: pulumi.Input<string>;
    /**
     * Password to use with `binddn` when performing user search
     */
    readonly bindpass?: pulumi.Input<string>;
    /**
     * Trusted CA to validate TLS certificate
     */
    readonly certificate?: pulumi.Input<string>;
    readonly denyNullBind?: pulumi.Input<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    readonly description?: pulumi.Input<string>;
    readonly discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     */
    readonly groupattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform group search
     */
    readonly groupdn?: pulumi.Input<string>;
    /**
     * Go template used to construct group membership query
     */
    readonly groupfilter?: pulumi.Input<string>;
    /**
     * Control whether or TLS certificates must be validated
     */
    readonly insecureTls?: pulumi.Input<boolean>;
    /**
     * Path to mount the LDAP auth backend under
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Control use of TLS when conecting to LDAP
     */
    readonly starttls?: pulumi.Input<boolean>;
    /**
     * Maximum acceptable version of TLS
     */
    readonly tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum acceptable version of TLS
     */
    readonly tlsMinVersion?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The userPrincipalDomain used to construct UPN string
     */
    readonly upndomain?: pulumi.Input<string>;
    /**
     * The URL of the LDAP server
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     */
    readonly useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute on user object matching username passed in
     */
    readonly userattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform user search
     */
    readonly userdn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * DN of object to bind when performing user search
     */
    readonly binddn?: pulumi.Input<string>;
    /**
     * Password to use with `binddn` when performing user search
     */
    readonly bindpass?: pulumi.Input<string>;
    /**
     * Trusted CA to validate TLS certificate
     */
    readonly certificate?: pulumi.Input<string>;
    readonly denyNullBind?: pulumi.Input<boolean>;
    /**
     * Description for the LDAP auth backend mount
     */
    readonly description?: pulumi.Input<string>;
    readonly discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     */
    readonly groupattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform group search
     */
    readonly groupdn?: pulumi.Input<string>;
    /**
     * Go template used to construct group membership query
     */
    readonly groupfilter?: pulumi.Input<string>;
    /**
     * Control whether or TLS certificates must be validated
     */
    readonly insecureTls?: pulumi.Input<boolean>;
    /**
     * Path to mount the LDAP auth backend under
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Control use of TLS when conecting to LDAP
     */
    readonly starttls?: pulumi.Input<boolean>;
    /**
     * Maximum acceptable version of TLS
     */
    readonly tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum acceptable version of TLS
     */
    readonly tlsMinVersion?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The userPrincipalDomain used to construct UPN string
     */
    readonly upndomain?: pulumi.Input<string>;
    /**
     * The URL of the LDAP server
     */
    readonly url: pulumi.Input<string>;
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     */
    readonly useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute on user object matching username passed in
     */
    readonly userattr?: pulumi.Input<string>;
    /**
     * Base DN under which to perform user search
     */
    readonly userdn?: pulumi.Input<string>;
}
