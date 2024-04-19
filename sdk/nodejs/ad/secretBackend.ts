// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.ad.SecretBackend("config", {
 *     backend: "ad",
 *     binddn: "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
 *     bindpass: "SuperSecretPassw0rd",
 *     insecureTls: true,
 *     url: "ldaps://ad",
 *     userdn: "CN=Users,DC=corp,DC=example,DC=net",
 * });
 * ```
 *
 * ## Import
 *
 * AD secret backend can be imported using the `backend`, e.g.
 *
 * ```sh
 * $ pulumi import vault:ad/secretBackend:SecretBackend ad ad
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ad/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    public readonly anonymousGroupSearch!: pulumi.Output<boolean | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    public readonly binddn!: pulumi.Output<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    public readonly bindpass!: pulumi.Output<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    public readonly caseSensitiveNames!: pulumi.Output<boolean | undefined>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    public readonly certificate!: pulumi.Output<string | undefined>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    public readonly clientTlsCert!: pulumi.Output<string | undefined>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    public readonly clientTlsKey!: pulumi.Output<string | undefined>;
    /**
     * Default lease duration for secrets in seconds.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    public readonly denyNullBind!: pulumi.Output<boolean | undefined>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    public readonly discoverdn!: pulumi.Output<boolean | undefined>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    public readonly groupattr!: pulumi.Output<string | undefined>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    public readonly groupdn!: pulumi.Output<string | undefined>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    public readonly groupfilter!: pulumi.Output<string | undefined>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    public readonly insecureTls!: pulumi.Output<boolean | undefined>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    public readonly lastRotationTolerance!: pulumi.Output<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    public readonly passwordPolicy!: pulumi.Output<string | undefined>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    public readonly requestTimeout!: pulumi.Output<number | undefined>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    public readonly starttls!: pulumi.Output<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    public readonly tlsMaxVersion!: pulumi.Output<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    public readonly tlsMinVersion!: pulumi.Output<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    public readonly upndomain!: pulumi.Output<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    public readonly usePre111GroupCnBehavior!: pulumi.Output<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    public readonly useTokenGroups!: pulumi.Output<boolean | undefined>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    public readonly userattr!: pulumi.Output<string | undefined>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    public readonly userdn!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["anonymousGroupSearch"] = state ? state.anonymousGroupSearch : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["binddn"] = state ? state.binddn : undefined;
            resourceInputs["bindpass"] = state ? state.bindpass : undefined;
            resourceInputs["caseSensitiveNames"] = state ? state.caseSensitiveNames : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["clientTlsCert"] = state ? state.clientTlsCert : undefined;
            resourceInputs["clientTlsKey"] = state ? state.clientTlsKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["denyNullBind"] = state ? state.denyNullBind : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["discoverdn"] = state ? state.discoverdn : undefined;
            resourceInputs["groupattr"] = state ? state.groupattr : undefined;
            resourceInputs["groupdn"] = state ? state.groupdn : undefined;
            resourceInputs["groupfilter"] = state ? state.groupfilter : undefined;
            resourceInputs["insecureTls"] = state ? state.insecureTls : undefined;
            resourceInputs["lastRotationTolerance"] = state ? state.lastRotationTolerance : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["passwordPolicy"] = state ? state.passwordPolicy : undefined;
            resourceInputs["requestTimeout"] = state ? state.requestTimeout : undefined;
            resourceInputs["starttls"] = state ? state.starttls : undefined;
            resourceInputs["tlsMaxVersion"] = state ? state.tlsMaxVersion : undefined;
            resourceInputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["upndomain"] = state ? state.upndomain : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["usePre111GroupCnBehavior"] = state ? state.usePre111GroupCnBehavior : undefined;
            resourceInputs["useTokenGroups"] = state ? state.useTokenGroups : undefined;
            resourceInputs["userattr"] = state ? state.userattr : undefined;
            resourceInputs["userdn"] = state ? state.userdn : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if ((!args || args.binddn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'binddn'");
            }
            if ((!args || args.bindpass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindpass'");
            }
            resourceInputs["anonymousGroupSearch"] = args ? args.anonymousGroupSearch : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["binddn"] = args ? args.binddn : undefined;
            resourceInputs["bindpass"] = args?.bindpass ? pulumi.secret(args.bindpass) : undefined;
            resourceInputs["caseSensitiveNames"] = args ? args.caseSensitiveNames : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["clientTlsCert"] = args?.clientTlsCert ? pulumi.secret(args.clientTlsCert) : undefined;
            resourceInputs["clientTlsKey"] = args?.clientTlsKey ? pulumi.secret(args.clientTlsKey) : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["denyNullBind"] = args ? args.denyNullBind : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["discoverdn"] = args ? args.discoverdn : undefined;
            resourceInputs["groupattr"] = args ? args.groupattr : undefined;
            resourceInputs["groupdn"] = args ? args.groupdn : undefined;
            resourceInputs["groupfilter"] = args ? args.groupfilter : undefined;
            resourceInputs["insecureTls"] = args ? args.insecureTls : undefined;
            resourceInputs["lastRotationTolerance"] = args ? args.lastRotationTolerance : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["passwordPolicy"] = args ? args.passwordPolicy : undefined;
            resourceInputs["requestTimeout"] = args ? args.requestTimeout : undefined;
            resourceInputs["starttls"] = args ? args.starttls : undefined;
            resourceInputs["tlsMaxVersion"] = args ? args.tlsMaxVersion : undefined;
            resourceInputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["upndomain"] = args ? args.upndomain : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["usePre111GroupCnBehavior"] = args ? args.usePre111GroupCnBehavior : undefined;
            resourceInputs["useTokenGroups"] = args ? args.useTokenGroups : undefined;
            resourceInputs["userattr"] = args ? args.userattr : undefined;
            resourceInputs["userdn"] = args ? args.userdn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["bindpass", "clientTlsCert", "clientTlsKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    anonymousGroupSearch?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    binddn?: pulumi.Input<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    bindpass?: pulumi.Input<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    clientTlsCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    clientTlsKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    denyNullBind?: pulumi.Input<boolean>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    groupattr?: pulumi.Input<string>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    groupdn?: pulumi.Input<string>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    groupfilter?: pulumi.Input<string>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    insecureTls?: pulumi.Input<boolean>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    lastRotationTolerance?: pulumi.Input<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    passwordPolicy?: pulumi.Input<string>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    requestTimeout?: pulumi.Input<number>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    starttls?: pulumi.Input<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    tlsMinVersion?: pulumi.Input<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    ttl?: pulumi.Input<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    upndomain?: pulumi.Input<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    url?: pulumi.Input<string>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    usePre111GroupCnBehavior?: pulumi.Input<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    userattr?: pulumi.Input<string>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    userdn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    anonymousGroupSearch?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    binddn: pulumi.Input<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    bindpass: pulumi.Input<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    clientTlsCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    clientTlsKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    denyNullBind?: pulumi.Input<boolean>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    discoverdn?: pulumi.Input<boolean>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    groupattr?: pulumi.Input<string>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    groupdn?: pulumi.Input<string>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    groupfilter?: pulumi.Input<string>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    insecureTls?: pulumi.Input<boolean>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    lastRotationTolerance?: pulumi.Input<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    passwordPolicy?: pulumi.Input<string>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    requestTimeout?: pulumi.Input<number>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    starttls?: pulumi.Input<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    tlsMinVersion?: pulumi.Input<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    ttl?: pulumi.Input<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    upndomain?: pulumi.Input<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    url?: pulumi.Input<string>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    usePre111GroupCnBehavior?: pulumi.Input<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    userattr?: pulumi.Input<string>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    userdn?: pulumi.Input<string>;
}
