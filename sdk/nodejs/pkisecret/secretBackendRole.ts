// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a role on an PKI Secret Backend for Vault.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const pki = new vault.Mount("pki", {
 *     path: "pki",
 *     type: "pki",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 * });
 * const role = new vault.pkisecret.SecretBackendRole("role", {
 *     backend: pki.path,
 *     ttl: "3600",
 *     allowIpSans: true,
 *     keyType: "rsa",
 *     keyBits: 4096,
 *     allowedDomains: [
 *         "example.com",
 *         "my.domain",
 *     ],
 *     allowSubdomains: true,
 * });
 * ```
 *
 * ## Import
 *
 * PKI secret backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:pkiSecret/secretBackendRole:SecretBackendRole role pki/roles/my_role
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
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendRole:SecretBackendRole';

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
     * Flag to allow any name
     */
    public readonly allowAnyName!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow certificates matching the actual domain
     */
    public readonly allowBareDomains!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow names containing glob patterns.
     */
    public readonly allowGlobDomains!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow IP SANs
     */
    public readonly allowIpSans!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow certificates for localhost
     */
    public readonly allowLocalhost!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow certificates matching subdomains
     */
    public readonly allowSubdomains!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow wildcard certificates.
     */
    public readonly allowWildcardCertificates!: pulumi.Output<boolean | undefined>;
    /**
     * List of allowed domains for certificates
     */
    public readonly allowedDomains!: pulumi.Output<string[] | undefined>;
    /**
     * Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    public readonly allowedDomainsTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * Defines allowed custom SANs
     */
    public readonly allowedOtherSans!: pulumi.Output<string[] | undefined>;
    /**
     * An array of allowed serial numbers to put in Subject
     */
    public readonly allowedSerialNumbers!: pulumi.Output<string[] | undefined>;
    /**
     * Defines allowed URI SANs
     */
    public readonly allowedUriSans!: pulumi.Output<string[] | undefined>;
    /**
     * Flag, if set, `allowedUriSans` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    public readonly allowedUriSansTemplate!: pulumi.Output<boolean>;
    /**
     * Defines allowed User IDs
     */
    public readonly allowedUserIds!: pulumi.Output<string[] | undefined>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Flag to mark basic constraints valid when issuing non-CA certificates
     */
    public readonly basicConstraintsValidForNonCa!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to specify certificates for client use
     */
    public readonly clientFlag!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to specify certificates for code signing use
     */
    public readonly codeSigningFlag!: pulumi.Output<boolean | undefined>;
    /**
     * The country of generated certificates
     */
    public readonly countries!: pulumi.Output<string[] | undefined>;
    /**
     * Flag to specify certificates for email protection use
     */
    public readonly emailProtectionFlag!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to allow only valid host names
     */
    public readonly enforceHostnames!: pulumi.Output<boolean | undefined>;
    /**
     * Specify the allowed extended key usage constraint on issued certificates
     */
    public readonly extKeyUsages!: pulumi.Output<string[] | undefined>;
    /**
     * Flag to generate leases with certificates
     */
    public readonly generateLease!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    public readonly issuerRef!: pulumi.Output<string>;
    /**
     * The number of bits of generated keys
     */
    public readonly keyBits!: pulumi.Output<number | undefined>;
    /**
     * The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
     * Defaults to `rsa`
     */
    public readonly keyType!: pulumi.Output<string | undefined>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    public readonly keyUsages!: pulumi.Output<string[]>;
    /**
     * The locality of generated certificates
     */
    public readonly localities!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum lease TTL, in seconds, for the role.
     */
    public readonly maxTtl!: pulumi.Output<string>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Flag to not store certificates in the storage backend
     */
    public readonly noStore!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the duration by which to backdate the NotBefore property.
     */
    public readonly notBeforeDuration!: pulumi.Output<string>;
    /**
     * The organization unit of generated certificates
     */
    public readonly organizationUnit!: pulumi.Output<string[] | undefined>;
    /**
     * The organization of generated certificates
     */
    public readonly organizations!: pulumi.Output<string[] | undefined>;
    /**
     * Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policyIdentifier` blocks instead
     */
    public readonly policyIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * The postal code of generated certificates
     */
    public readonly postalCodes!: pulumi.Output<string[] | undefined>;
    /**
     * The province of generated certificates
     */
    public readonly provinces!: pulumi.Output<string[] | undefined>;
    /**
     * Flag to force CN usage
     */
    public readonly requireCn!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to specify certificates for server use
     */
    public readonly serverFlag!: pulumi.Output<boolean | undefined>;
    /**
     * The street address of generated certificates
     */
    public readonly streetAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * The TTL, in seconds, for any certificate issued against this role.
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * Flag to use the CN in the CSR
     */
    public readonly useCsrCommonName!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to use the SANs in the CSR
     */
    public readonly useCsrSans!: pulumi.Output<boolean | undefined>;

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
            resourceInputs["allowAnyName"] = state ? state.allowAnyName : undefined;
            resourceInputs["allowBareDomains"] = state ? state.allowBareDomains : undefined;
            resourceInputs["allowGlobDomains"] = state ? state.allowGlobDomains : undefined;
            resourceInputs["allowIpSans"] = state ? state.allowIpSans : undefined;
            resourceInputs["allowLocalhost"] = state ? state.allowLocalhost : undefined;
            resourceInputs["allowSubdomains"] = state ? state.allowSubdomains : undefined;
            resourceInputs["allowWildcardCertificates"] = state ? state.allowWildcardCertificates : undefined;
            resourceInputs["allowedDomains"] = state ? state.allowedDomains : undefined;
            resourceInputs["allowedDomainsTemplate"] = state ? state.allowedDomainsTemplate : undefined;
            resourceInputs["allowedOtherSans"] = state ? state.allowedOtherSans : undefined;
            resourceInputs["allowedSerialNumbers"] = state ? state.allowedSerialNumbers : undefined;
            resourceInputs["allowedUriSans"] = state ? state.allowedUriSans : undefined;
            resourceInputs["allowedUriSansTemplate"] = state ? state.allowedUriSansTemplate : undefined;
            resourceInputs["allowedUserIds"] = state ? state.allowedUserIds : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["basicConstraintsValidForNonCa"] = state ? state.basicConstraintsValidForNonCa : undefined;
            resourceInputs["clientFlag"] = state ? state.clientFlag : undefined;
            resourceInputs["codeSigningFlag"] = state ? state.codeSigningFlag : undefined;
            resourceInputs["countries"] = state ? state.countries : undefined;
            resourceInputs["emailProtectionFlag"] = state ? state.emailProtectionFlag : undefined;
            resourceInputs["enforceHostnames"] = state ? state.enforceHostnames : undefined;
            resourceInputs["extKeyUsages"] = state ? state.extKeyUsages : undefined;
            resourceInputs["generateLease"] = state ? state.generateLease : undefined;
            resourceInputs["issuerRef"] = state ? state.issuerRef : undefined;
            resourceInputs["keyBits"] = state ? state.keyBits : undefined;
            resourceInputs["keyType"] = state ? state.keyType : undefined;
            resourceInputs["keyUsages"] = state ? state.keyUsages : undefined;
            resourceInputs["localities"] = state ? state.localities : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["noStore"] = state ? state.noStore : undefined;
            resourceInputs["notBeforeDuration"] = state ? state.notBeforeDuration : undefined;
            resourceInputs["organizationUnit"] = state ? state.organizationUnit : undefined;
            resourceInputs["organizations"] = state ? state.organizations : undefined;
            resourceInputs["policyIdentifiers"] = state ? state.policyIdentifiers : undefined;
            resourceInputs["postalCodes"] = state ? state.postalCodes : undefined;
            resourceInputs["provinces"] = state ? state.provinces : undefined;
            resourceInputs["requireCn"] = state ? state.requireCn : undefined;
            resourceInputs["serverFlag"] = state ? state.serverFlag : undefined;
            resourceInputs["streetAddresses"] = state ? state.streetAddresses : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["useCsrCommonName"] = state ? state.useCsrCommonName : undefined;
            resourceInputs["useCsrSans"] = state ? state.useCsrSans : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            resourceInputs["allowAnyName"] = args ? args.allowAnyName : undefined;
            resourceInputs["allowBareDomains"] = args ? args.allowBareDomains : undefined;
            resourceInputs["allowGlobDomains"] = args ? args.allowGlobDomains : undefined;
            resourceInputs["allowIpSans"] = args ? args.allowIpSans : undefined;
            resourceInputs["allowLocalhost"] = args ? args.allowLocalhost : undefined;
            resourceInputs["allowSubdomains"] = args ? args.allowSubdomains : undefined;
            resourceInputs["allowWildcardCertificates"] = args ? args.allowWildcardCertificates : undefined;
            resourceInputs["allowedDomains"] = args ? args.allowedDomains : undefined;
            resourceInputs["allowedDomainsTemplate"] = args ? args.allowedDomainsTemplate : undefined;
            resourceInputs["allowedOtherSans"] = args ? args.allowedOtherSans : undefined;
            resourceInputs["allowedSerialNumbers"] = args ? args.allowedSerialNumbers : undefined;
            resourceInputs["allowedUriSans"] = args ? args.allowedUriSans : undefined;
            resourceInputs["allowedUriSansTemplate"] = args ? args.allowedUriSansTemplate : undefined;
            resourceInputs["allowedUserIds"] = args ? args.allowedUserIds : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["basicConstraintsValidForNonCa"] = args ? args.basicConstraintsValidForNonCa : undefined;
            resourceInputs["clientFlag"] = args ? args.clientFlag : undefined;
            resourceInputs["codeSigningFlag"] = args ? args.codeSigningFlag : undefined;
            resourceInputs["countries"] = args ? args.countries : undefined;
            resourceInputs["emailProtectionFlag"] = args ? args.emailProtectionFlag : undefined;
            resourceInputs["enforceHostnames"] = args ? args.enforceHostnames : undefined;
            resourceInputs["extKeyUsages"] = args ? args.extKeyUsages : undefined;
            resourceInputs["generateLease"] = args ? args.generateLease : undefined;
            resourceInputs["issuerRef"] = args ? args.issuerRef : undefined;
            resourceInputs["keyBits"] = args ? args.keyBits : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["keyUsages"] = args ? args.keyUsages : undefined;
            resourceInputs["localities"] = args ? args.localities : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["noStore"] = args ? args.noStore : undefined;
            resourceInputs["notBeforeDuration"] = args ? args.notBeforeDuration : undefined;
            resourceInputs["organizationUnit"] = args ? args.organizationUnit : undefined;
            resourceInputs["organizations"] = args ? args.organizations : undefined;
            resourceInputs["policyIdentifiers"] = args ? args.policyIdentifiers : undefined;
            resourceInputs["postalCodes"] = args ? args.postalCodes : undefined;
            resourceInputs["provinces"] = args ? args.provinces : undefined;
            resourceInputs["requireCn"] = args ? args.requireCn : undefined;
            resourceInputs["serverFlag"] = args ? args.serverFlag : undefined;
            resourceInputs["streetAddresses"] = args ? args.streetAddresses : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["useCsrCommonName"] = args ? args.useCsrCommonName : undefined;
            resourceInputs["useCsrSans"] = args ? args.useCsrSans : undefined;
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
     * Flag to allow any name
     */
    allowAnyName?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching the actual domain
     */
    allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow names containing glob patterns.
     */
    allowGlobDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow IP SANs
     */
    allowIpSans?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates for localhost
     */
    allowLocalhost?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching subdomains
     */
    allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow wildcard certificates.
     */
    allowWildcardCertificates?: pulumi.Input<boolean>;
    /**
     * List of allowed domains for certificates
     */
    allowedDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    allowedDomainsTemplate?: pulumi.Input<boolean>;
    /**
     * Defines allowed custom SANs
     */
    allowedOtherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of allowed serial numbers to put in Subject
     */
    allowedSerialNumbers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed URI SANs
     */
    allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag, if set, `allowedUriSans` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    allowedUriSansTemplate?: pulumi.Input<boolean>;
    /**
     * Defines allowed User IDs
     */
    allowedUserIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * Flag to mark basic constraints valid when issuing non-CA certificates
     */
    basicConstraintsValidForNonCa?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for client use
     */
    clientFlag?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for code signing use
     */
    codeSigningFlag?: pulumi.Input<boolean>;
    /**
     * The country of generated certificates
     */
    countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to specify certificates for email protection use
     */
    emailProtectionFlag?: pulumi.Input<boolean>;
    /**
     * Flag to allow only valid host names
     */
    enforceHostnames?: pulumi.Input<boolean>;
    /**
     * Specify the allowed extended key usage constraint on issued certificates
     */
    extKeyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to generate leases with certificates
     */
    generateLease?: pulumi.Input<boolean>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * The number of bits of generated keys
     */
    keyBits?: pulumi.Input<number>;
    /**
     * The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
     * Defaults to `rsa`
     */
    keyType?: pulumi.Input<string>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    keyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locality of generated certificates
     */
    localities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum lease TTL, in seconds, for the role.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Flag to not store certificates in the storage backend
     */
    noStore?: pulumi.Input<boolean>;
    /**
     * Specifies the duration by which to backdate the NotBefore property.
     */
    notBeforeDuration?: pulumi.Input<string>;
    /**
     * The organization unit of generated certificates
     */
    organizationUnit?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization of generated certificates
     */
    organizations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policyIdentifier` blocks instead
     */
    policyIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code of generated certificates
     */
    postalCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The province of generated certificates
     */
    provinces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to force CN usage
     */
    requireCn?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for server use
     */
    serverFlag?: pulumi.Input<boolean>;
    /**
     * The street address of generated certificates
     */
    streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The TTL, in seconds, for any certificate issued against this role.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Flag to use the CN in the CSR
     */
    useCsrCommonName?: pulumi.Input<boolean>;
    /**
     * Flag to use the SANs in the CSR
     */
    useCsrSans?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * Flag to allow any name
     */
    allowAnyName?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching the actual domain
     */
    allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow names containing glob patterns.
     */
    allowGlobDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow IP SANs
     */
    allowIpSans?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates for localhost
     */
    allowLocalhost?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching subdomains
     */
    allowSubdomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow wildcard certificates.
     */
    allowWildcardCertificates?: pulumi.Input<boolean>;
    /**
     * List of allowed domains for certificates
     */
    allowedDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    allowedDomainsTemplate?: pulumi.Input<boolean>;
    /**
     * Defines allowed custom SANs
     */
    allowedOtherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of allowed serial numbers to put in Subject
     */
    allowedSerialNumbers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed URI SANs
     */
    allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag, if set, `allowedUriSans` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
     */
    allowedUriSansTemplate?: pulumi.Input<boolean>;
    /**
     * Defines allowed User IDs
     */
    allowedUserIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Flag to mark basic constraints valid when issuing non-CA certificates
     */
    basicConstraintsValidForNonCa?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for client use
     */
    clientFlag?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for code signing use
     */
    codeSigningFlag?: pulumi.Input<boolean>;
    /**
     * The country of generated certificates
     */
    countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to specify certificates for email protection use
     */
    emailProtectionFlag?: pulumi.Input<boolean>;
    /**
     * Flag to allow only valid host names
     */
    enforceHostnames?: pulumi.Input<boolean>;
    /**
     * Specify the allowed extended key usage constraint on issued certificates
     */
    extKeyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to generate leases with certificates
     */
    generateLease?: pulumi.Input<boolean>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * The number of bits of generated keys
     */
    keyBits?: pulumi.Input<number>;
    /**
     * The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
     * Defaults to `rsa`
     */
    keyType?: pulumi.Input<string>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    keyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locality of generated certificates
     */
    localities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum lease TTL, in seconds, for the role.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Flag to not store certificates in the storage backend
     */
    noStore?: pulumi.Input<boolean>;
    /**
     * Specifies the duration by which to backdate the NotBefore property.
     */
    notBeforeDuration?: pulumi.Input<string>;
    /**
     * The organization unit of generated certificates
     */
    organizationUnit?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization of generated certificates
     */
    organizations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policyIdentifier` blocks instead
     */
    policyIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code of generated certificates
     */
    postalCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The province of generated certificates
     */
    provinces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to force CN usage
     */
    requireCn?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for server use
     */
    serverFlag?: pulumi.Input<boolean>;
    /**
     * The street address of generated certificates
     */
    streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The TTL, in seconds, for any certificate issued against this role.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Flag to use the CN in the CSR
     */
    useCsrCommonName?: pulumi.Input<boolean>;
    /**
     * Flag to use the SANs in the CSR
     */
    useCsrSans?: pulumi.Input<boolean>;
}
