// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const example = new vault.Mount("example", {
 *     type: "ssh",
 * });
 * const foo = new vault.ssh.SecretBackendRole("foo", {
 *     allowUserCertificates: true,
 *     backend: example.path,
 *     keyType: "ca",
 * });
 * const bar = new vault.ssh.SecretBackendRole("bar", {
 *     allowedUsers: "default,baz",
 *     backend: example.path,
 *     cidrList: "0.0.0.0/0",
 *     defaultUser: "default",
 *     keyType: "otp",
 * });
 * ```
 *
 * ## Import
 *
 * SSH secret backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:ssh/secretBackendRole:SecretBackendRole foo ssh/roles/my-role
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
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    public readonly allowedExtensions!: pulumi.Output<string | undefined>;
    /**
     * Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
     */
    public readonly allowedUserKeyLengths!: pulumi.Output<{[key: string]: any} | undefined>;
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
    public readonly defaultCriticalOptions!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    public readonly defaultExtensions!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    public readonly defaultUser!: pulumi.Output<string | undefined>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            inputs["algorithmSigner"] = state ? state.algorithmSigner : undefined;
            inputs["allowBareDomains"] = state ? state.allowBareDomains : undefined;
            inputs["allowHostCertificates"] = state ? state.allowHostCertificates : undefined;
            inputs["allowSubdomains"] = state ? state.allowSubdomains : undefined;
            inputs["allowUserCertificates"] = state ? state.allowUserCertificates : undefined;
            inputs["allowUserKeyIds"] = state ? state.allowUserKeyIds : undefined;
            inputs["allowedCriticalOptions"] = state ? state.allowedCriticalOptions : undefined;
            inputs["allowedDomains"] = state ? state.allowedDomains : undefined;
            inputs["allowedExtensions"] = state ? state.allowedExtensions : undefined;
            inputs["allowedUserKeyLengths"] = state ? state.allowedUserKeyLengths : undefined;
            inputs["allowedUsers"] = state ? state.allowedUsers : undefined;
            inputs["allowedUsersTemplate"] = state ? state.allowedUsersTemplate : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["cidrList"] = state ? state.cidrList : undefined;
            inputs["defaultCriticalOptions"] = state ? state.defaultCriticalOptions : undefined;
            inputs["defaultExtensions"] = state ? state.defaultExtensions : undefined;
            inputs["defaultUser"] = state ? state.defaultUser : undefined;
            inputs["keyIdFormat"] = state ? state.keyIdFormat : undefined;
            inputs["keyType"] = state ? state.keyType : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.keyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyType'");
            }
            inputs["algorithmSigner"] = args ? args.algorithmSigner : undefined;
            inputs["allowBareDomains"] = args ? args.allowBareDomains : undefined;
            inputs["allowHostCertificates"] = args ? args.allowHostCertificates : undefined;
            inputs["allowSubdomains"] = args ? args.allowSubdomains : undefined;
            inputs["allowUserCertificates"] = args ? args.allowUserCertificates : undefined;
            inputs["allowUserKeyIds"] = args ? args.allowUserKeyIds : undefined;
            inputs["allowedCriticalOptions"] = args ? args.allowedCriticalOptions : undefined;
            inputs["allowedDomains"] = args ? args.allowedDomains : undefined;
            inputs["allowedExtensions"] = args ? args.allowedExtensions : undefined;
            inputs["allowedUserKeyLengths"] = args ? args.allowedUserKeyLengths : undefined;
            inputs["allowedUsers"] = args ? args.allowedUsers : undefined;
            inputs["allowedUsersTemplate"] = args ? args.allowedUsersTemplate : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["cidrList"] = args ? args.cidrList : undefined;
            inputs["defaultCriticalOptions"] = args ? args.defaultCriticalOptions : undefined;
            inputs["defaultExtensions"] = args ? args.defaultExtensions : undefined;
            inputs["defaultUser"] = args ? args.defaultUser : undefined;
            inputs["keyIdFormat"] = args ? args.keyIdFormat : undefined;
            inputs["keyType"] = args ? args.keyType : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecretBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     */
    readonly algorithmSigner?: pulumi.Input<string>;
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
     */
    readonly allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'host'.
     */
    readonly allowHostCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
     */
    readonly allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'user'.
     */
    readonly allowUserCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if users can override the key ID for a signed certificate with the `keyId` field.
     */
    readonly allowUserKeyIds?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     */
    readonly allowedCriticalOptions?: pulumi.Input<string>;
    /**
     * The list of domains for which a client can request a host certificate.
     */
    readonly allowedDomains?: pulumi.Input<string>;
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    readonly allowedExtensions?: pulumi.Input<string>;
    /**
     * Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
     */
    readonly allowedUserKeyLengths?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     */
    readonly allowedUsers?: pulumi.Input<string>;
    /**
     * Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
     */
    readonly allowedUsersTemplate?: pulumi.Input<boolean>;
    /**
     * The path where the SSH secret backend is mounted.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     */
    readonly cidrList?: pulumi.Input<string>;
    /**
     * Specifies a map of critical options that certificates have when signed.
     */
    readonly defaultCriticalOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    readonly defaultExtensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    readonly defaultUser?: pulumi.Input<string>;
    /**
     * Specifies a custom format for the key id of a signed certificate.
     */
    readonly keyIdFormat?: pulumi.Input<string>;
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     */
    readonly keyType?: pulumi.Input<string>;
    /**
     * Specifies the maximum Time To Live value.
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * Specifies the name of the role to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Time To Live value.
     */
    readonly ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     */
    readonly algorithmSigner?: pulumi.Input<string>;
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
     */
    readonly allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'host'.
     */
    readonly allowHostCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
     */
    readonly allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Specifies if certificates are allowed to be signed for use as a 'user'.
     */
    readonly allowUserCertificates?: pulumi.Input<boolean>;
    /**
     * Specifies if users can override the key ID for a signed certificate with the `keyId` field.
     */
    readonly allowUserKeyIds?: pulumi.Input<boolean>;
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     */
    readonly allowedCriticalOptions?: pulumi.Input<string>;
    /**
     * The list of domains for which a client can request a host certificate.
     */
    readonly allowedDomains?: pulumi.Input<string>;
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     */
    readonly allowedExtensions?: pulumi.Input<string>;
    /**
     * Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
     */
    readonly allowedUserKeyLengths?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     */
    readonly allowedUsers?: pulumi.Input<string>;
    /**
     * Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
     */
    readonly allowedUsersTemplate?: pulumi.Input<boolean>;
    /**
     * The path where the SSH secret backend is mounted.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     */
    readonly cidrList?: pulumi.Input<string>;
    /**
     * Specifies a map of critical options that certificates have when signed.
     */
    readonly defaultCriticalOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies a map of extensions that certificates have when signed.
     */
    readonly defaultExtensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the default username for which a credential will be generated.
     */
    readonly defaultUser?: pulumi.Input<string>;
    /**
     * Specifies a custom format for the key id of a signed certificate.
     */
    readonly keyIdFormat?: pulumi.Input<string>;
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     */
    readonly keyType: pulumi.Input<string>;
    /**
     * Specifies the maximum Time To Live value.
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * Specifies the name of the role to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Time To Live value.
     */
    readonly ttl?: pulumi.Input<string>;
}
