// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage roles in an SSH secret backend
 * [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.Mount("example", {type: "ssh"});
 * const foo = new vault.ssh.SecretBackendRole("foo", {
 *     name: "my-role",
 *     backend: example.path,
 *     keyType: "ca",
 *     allowUserCertificates: true,
 * });
 * const bar = new vault.ssh.SecretBackendRole("bar", {
 *     name: "otp-role",
 *     backend: example.path,
 *     keyType: "otp",
 *     defaultUser: "default",
 *     allowedUsers: "default,baz",
 *     cidrList: "0.0.0.0/0",
 * });
 * ```
 *
 * ## Import
 *
 * SSH secret backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:ssh/secretBackendRole:SecretBackendRole foo ssh/roles/my-role
 * ```
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ssh/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     */
    public readonly algorithmSigner!: pulumi.Output<string>;
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
     */
    public readonly allowBareDomains!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'host'.
     */
    public readonly allowHostCertificates!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
     */
    public readonly allowSubdomains!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'user'.
     */
    public readonly allowUserCertificates!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if users can override the key ID for a signed certificate with the `keyId` field.
     */
    public readonly allowUserKeyIds!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     */
    public readonly allowedCriticalOptions!: pulumi.Output<string | undefined>;
    /**
     * The list of domains for which a client can request a host certificate.
     */
    public readonly allowedDomains!: pulumi.Output<string | undefined>;
    /**
     * Specifies if `allowedDomains` can be declared using
     * identity template policies. Non-templated domains are also permitted.
     */
    public readonly allowedDomainsTemplate!: pulumi.Output<boolean>;
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    public readonly allowedExtensions!: pulumi.Output<string | undefined>;
    /**
     * Set of configuration blocks to define allowed  
     * user key configuration, like key type and their lengths. Can be specified multiple times.
     * *See Configuration-Options for more info*
     */
    public readonly allowedUserKeyConfigs!: pulumi.Output<outputs.ssh.SecretBackendRoleAllowedUserKeyConfig[] | undefined>;
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     */
    public readonly allowedUsers!: pulumi.Output<string | undefined>;
    /**
     * Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
     */
    public readonly allowedUsersTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * The path where the SSH secret backend is mounted.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     */
    public readonly cidrList!: pulumi.Output<string | undefined>;
    /**
     * Specifies a map of critical options that certificates have when signed.
     */
    public readonly defaultCriticalOptions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    public readonly defaultExtensions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    public readonly defaultUser!: pulumi.Output<string | undefined>;
    /**
     * If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
     */
    public readonly defaultUserTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a custom format for the key id of a signed certificate.
     */
    public readonly keyIdFormat!: pulumi.Output<string | undefined>;
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     */
    public readonly keyType!: pulumi.Output<string>;
    /**
     * Specifies the maximum Time To Live value.
     */
    public readonly maxTtl!: pulumi.Output<string>;
    /**
     * Specifies the name of the role to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
     */
    public readonly notBeforeDuration!: pulumi.Output<string>;
    /**
     * Specifies the Time To Live value.
     */
    public readonly ttl!: pulumi.Output<string>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            resourceInputs["algorithmSigner"] = state ? state.algorithmSigner : undefined;
            resourceInputs["allowBareDomains"] = state ? state.allowBareDomains : undefined;
            resourceInputs["allowHostCertificates"] = state ? state.allowHostCertificates : undefined;
            resourceInputs["allowSubdomains"] = state ? state.allowSubdomains : undefined;
            resourceInputs["allowUserCertificates"] = state ? state.allowUserCertificates : undefined;
            resourceInputs["allowUserKeyIds"] = state ? state.allowUserKeyIds : undefined;
            resourceInputs["allowedCriticalOptions"] = state ? state.allowedCriticalOptions : undefined;
            resourceInputs["allowedDomains"] = state ? state.allowedDomains : undefined;
            resourceInputs["allowedDomainsTemplate"] = state ? state.allowedDomainsTemplate : undefined;
            resourceInputs["allowedExtensions"] = state ? state.allowedExtensions : undefined;
            resourceInputs["allowedUserKeyConfigs"] = state ? state.allowedUserKeyConfigs : undefined;
            resourceInputs["allowedUsers"] = state ? state.allowedUsers : undefined;
            resourceInputs["allowedUsersTemplate"] = state ? state.allowedUsersTemplate : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["cidrList"] = state ? state.cidrList : undefined;
            resourceInputs["defaultCriticalOptions"] = state ? state.defaultCriticalOptions : undefined;
            resourceInputs["defaultExtensions"] = state ? state.defaultExtensions : undefined;
            resourceInputs["defaultUser"] = state ? state.defaultUser : undefined;
            resourceInputs["defaultUserTemplate"] = state ? state.defaultUserTemplate : undefined;
            resourceInputs["keyIdFormat"] = state ? state.keyIdFormat : undefined;
            resourceInputs["keyType"] = state ? state.keyType : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["notBeforeDuration"] = state ? state.notBeforeDuration : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.keyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyType'");
            }
            resourceInputs["algorithmSigner"] = args ? args.algorithmSigner : undefined;
            resourceInputs["allowBareDomains"] = args ? args.allowBareDomains : undefined;
            resourceInputs["allowHostCertificates"] = args ? args.allowHostCertificates : undefined;
            resourceInputs["allowSubdomains"] = args ? args.allowSubdomains : undefined;
            resourceInputs["allowUserCertificates"] = args ? args.allowUserCertificates : undefined;
            resourceInputs["allowUserKeyIds"] = args ? args.allowUserKeyIds : undefined;
            resourceInputs["allowedCriticalOptions"] = args ? args.allowedCriticalOptions : undefined;
            resourceInputs["allowedDomains"] = args ? args.allowedDomains : undefined;
            resourceInputs["allowedDomainsTemplate"] = args ? args.allowedDomainsTemplate : undefined;
            resourceInputs["allowedExtensions"] = args ? args.allowedExtensions : undefined;
            resourceInputs["allowedUserKeyConfigs"] = args ? args.allowedUserKeyConfigs : undefined;
            resourceInputs["allowedUsers"] = args ? args.allowedUsers : undefined;
            resourceInputs["allowedUsersTemplate"] = args ? args.allowedUsersTemplate : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["cidrList"] = args ? args.cidrList : undefined;
            resourceInputs["defaultCriticalOptions"] = args ? args.defaultCriticalOptions : undefined;
            resourceInputs["defaultExtensions"] = args ? args.defaultExtensions : undefined;
            resourceInputs["defaultUser"] = args ? args.defaultUser : undefined;
            resourceInputs["defaultUserTemplate"] = args ? args.defaultUserTemplate : undefined;
            resourceInputs["keyIdFormat"] = args ? args.keyIdFormat : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["notBeforeDuration"] = args ? args.notBeforeDuration : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     */
    algorithmSigner?: pulumi.Input<string>;
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
     */
    allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'host'.
     */
    allowHostCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
     */
    allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'user'.
     */
    allowUserCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if users can override the key ID for a signed certificate with the `keyId` field.
     */
    allowUserKeyIds?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     */
    allowedCriticalOptions?: pulumi.Input<string>;
    /**
     * The list of domains for which a client can request a host certificate.
     */
    allowedDomains?: pulumi.Input<string>;
    /**
     * Specifies if `allowedDomains` can be declared using
     * identity template policies. Non-templated domains are also permitted.
     */
    allowedDomainsTemplate?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    allowedExtensions?: pulumi.Input<string>;
    /**
     * Set of configuration blocks to define allowed  
     * user key configuration, like key type and their lengths. Can be specified multiple times.
     * *See Configuration-Options for more info*
     */
    allowedUserKeyConfigs?: pulumi.Input<pulumi.Input<inputs.ssh.SecretBackendRoleAllowedUserKeyConfig>[]>;
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     */
    allowedUsers?: pulumi.Input<string>;
    /**
     * Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
     */
    allowedUsersTemplate?: pulumi.Input<boolean>;
    /**
     * The path where the SSH secret backend is mounted.
     */
    backend?: pulumi.Input<string>;
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     */
    cidrList?: pulumi.Input<string>;
    /**
     * Specifies a map of critical options that certificates have when signed.
     */
    defaultCriticalOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    defaultExtensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    defaultUser?: pulumi.Input<string>;
    /**
     * If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
     */
    defaultUserTemplate?: pulumi.Input<boolean>;
    /**
     * Specifies a custom format for the key id of a signed certificate.
     */
    keyIdFormat?: pulumi.Input<string>;
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     */
    keyType?: pulumi.Input<string>;
    /**
     * Specifies the maximum Time To Live value.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * Specifies the name of the role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
     */
    notBeforeDuration?: pulumi.Input<string>;
    /**
     * Specifies the Time To Live value.
     */
    ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     */
    algorithmSigner?: pulumi.Input<string>;
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
     */
    allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'host'.
     */
    allowHostCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
     */
    allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'user'.
     */
    allowUserCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if users can override the key ID for a signed certificate with the `keyId` field.
     */
    allowUserKeyIds?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     */
    allowedCriticalOptions?: pulumi.Input<string>;
    /**
     * The list of domains for which a client can request a host certificate.
     */
    allowedDomains?: pulumi.Input<string>;
    /**
     * Specifies if `allowedDomains` can be declared using
     * identity template policies. Non-templated domains are also permitted.
     */
    allowedDomainsTemplate?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    allowedExtensions?: pulumi.Input<string>;
    /**
     * Set of configuration blocks to define allowed  
     * user key configuration, like key type and their lengths. Can be specified multiple times.
     * *See Configuration-Options for more info*
     */
    allowedUserKeyConfigs?: pulumi.Input<pulumi.Input<inputs.ssh.SecretBackendRoleAllowedUserKeyConfig>[]>;
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     */
    allowedUsers?: pulumi.Input<string>;
    /**
     * Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
     */
    allowedUsersTemplate?: pulumi.Input<boolean>;
    /**
     * The path where the SSH secret backend is mounted.
     */
    backend: pulumi.Input<string>;
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     */
    cidrList?: pulumi.Input<string>;
    /**
     * Specifies a map of critical options that certificates have when signed.
     */
    defaultCriticalOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    defaultExtensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    defaultUser?: pulumi.Input<string>;
    /**
     * If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
     */
    defaultUserTemplate?: pulumi.Input<boolean>;
    /**
     * Specifies a custom format for the key id of a signed certificate.
     */
    keyIdFormat?: pulumi.Input<string>;
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     */
    keyType: pulumi.Input<string>;
    /**
     * Specifies the maximum Time To Live value.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * Specifies the name of the role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
     */
    notBeforeDuration?: pulumi.Input<string>;
    /**
     * Specifies the Time To Live value.
     */
    ttl?: pulumi.Input<string>;
}
