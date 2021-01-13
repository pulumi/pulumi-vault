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
 * const pki = new vault.pkiSecret.SecretBackend("pki", {
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 *     path: "%s",
 * });
 * const role = new vault.pkiSecret.SecretBackendRole("role", {
 *     backend: pki.path,
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
     * List of allowed domains for certificates
     */
    public readonly allowedDomains!: pulumi.Output<string[] | undefined>;
    /**
     * Defines allowed custom SANs
     */
    public readonly allowedOtherSans!: pulumi.Output<string[] | undefined>;
    /**
     * Defines allowed URI SANs
     */
    public readonly allowedUriSans!: pulumi.Output<string[] | undefined>;
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
     * The number of bits of generated keys
     */
    public readonly keyBits!: pulumi.Output<number | undefined>;
    /**
     * The type of generated keys
     */
    public readonly keyType!: pulumi.Output<string | undefined>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    public readonly keyUsages!: pulumi.Output<string[] | undefined>;
    /**
     * The locality of generated certificates
     */
    public readonly localities!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum TTL
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    public readonly name!: pulumi.Output<string>;
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
     * Specify the list of allowed policies IODs
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
     * The TTL
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            inputs["allowAnyName"] = state ? state.allowAnyName : undefined;
            inputs["allowBareDomains"] = state ? state.allowBareDomains : undefined;
            inputs["allowGlobDomains"] = state ? state.allowGlobDomains : undefined;
            inputs["allowIpSans"] = state ? state.allowIpSans : undefined;
            inputs["allowLocalhost"] = state ? state.allowLocalhost : undefined;
            inputs["allowSubdomains"] = state ? state.allowSubdomains : undefined;
            inputs["allowedDomains"] = state ? state.allowedDomains : undefined;
            inputs["allowedOtherSans"] = state ? state.allowedOtherSans : undefined;
            inputs["allowedUriSans"] = state ? state.allowedUriSans : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["basicConstraintsValidForNonCa"] = state ? state.basicConstraintsValidForNonCa : undefined;
            inputs["clientFlag"] = state ? state.clientFlag : undefined;
            inputs["codeSigningFlag"] = state ? state.codeSigningFlag : undefined;
            inputs["countries"] = state ? state.countries : undefined;
            inputs["emailProtectionFlag"] = state ? state.emailProtectionFlag : undefined;
            inputs["enforceHostnames"] = state ? state.enforceHostnames : undefined;
            inputs["extKeyUsages"] = state ? state.extKeyUsages : undefined;
            inputs["generateLease"] = state ? state.generateLease : undefined;
            inputs["keyBits"] = state ? state.keyBits : undefined;
            inputs["keyType"] = state ? state.keyType : undefined;
            inputs["keyUsages"] = state ? state.keyUsages : undefined;
            inputs["localities"] = state ? state.localities : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["noStore"] = state ? state.noStore : undefined;
            inputs["notBeforeDuration"] = state ? state.notBeforeDuration : undefined;
            inputs["organizationUnit"] = state ? state.organizationUnit : undefined;
            inputs["organizations"] = state ? state.organizations : undefined;
            inputs["policyIdentifiers"] = state ? state.policyIdentifiers : undefined;
            inputs["postalCodes"] = state ? state.postalCodes : undefined;
            inputs["provinces"] = state ? state.provinces : undefined;
            inputs["requireCn"] = state ? state.requireCn : undefined;
            inputs["serverFlag"] = state ? state.serverFlag : undefined;
            inputs["streetAddresses"] = state ? state.streetAddresses : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["useCsrCommonName"] = state ? state.useCsrCommonName : undefined;
            inputs["useCsrSans"] = state ? state.useCsrSans : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'backend'");
            }
            inputs["allowAnyName"] = args ? args.allowAnyName : undefined;
            inputs["allowBareDomains"] = args ? args.allowBareDomains : undefined;
            inputs["allowGlobDomains"] = args ? args.allowGlobDomains : undefined;
            inputs["allowIpSans"] = args ? args.allowIpSans : undefined;
            inputs["allowLocalhost"] = args ? args.allowLocalhost : undefined;
            inputs["allowSubdomains"] = args ? args.allowSubdomains : undefined;
            inputs["allowedDomains"] = args ? args.allowedDomains : undefined;
            inputs["allowedOtherSans"] = args ? args.allowedOtherSans : undefined;
            inputs["allowedUriSans"] = args ? args.allowedUriSans : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["basicConstraintsValidForNonCa"] = args ? args.basicConstraintsValidForNonCa : undefined;
            inputs["clientFlag"] = args ? args.clientFlag : undefined;
            inputs["codeSigningFlag"] = args ? args.codeSigningFlag : undefined;
            inputs["countries"] = args ? args.countries : undefined;
            inputs["emailProtectionFlag"] = args ? args.emailProtectionFlag : undefined;
            inputs["enforceHostnames"] = args ? args.enforceHostnames : undefined;
            inputs["extKeyUsages"] = args ? args.extKeyUsages : undefined;
            inputs["generateLease"] = args ? args.generateLease : undefined;
            inputs["keyBits"] = args ? args.keyBits : undefined;
            inputs["keyType"] = args ? args.keyType : undefined;
            inputs["keyUsages"] = args ? args.keyUsages : undefined;
            inputs["localities"] = args ? args.localities : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noStore"] = args ? args.noStore : undefined;
            inputs["notBeforeDuration"] = args ? args.notBeforeDuration : undefined;
            inputs["organizationUnit"] = args ? args.organizationUnit : undefined;
            inputs["organizations"] = args ? args.organizations : undefined;
            inputs["policyIdentifiers"] = args ? args.policyIdentifiers : undefined;
            inputs["postalCodes"] = args ? args.postalCodes : undefined;
            inputs["provinces"] = args ? args.provinces : undefined;
            inputs["requireCn"] = args ? args.requireCn : undefined;
            inputs["serverFlag"] = args ? args.serverFlag : undefined;
            inputs["streetAddresses"] = args ? args.streetAddresses : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["useCsrCommonName"] = args ? args.useCsrCommonName : undefined;
            inputs["useCsrSans"] = args ? args.useCsrSans : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * Flag to allow any name
     */
    readonly allowAnyName?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching the actual domain
     */
    readonly allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow names containing glob patterns.
     */
    readonly allowGlobDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow IP SANs
     */
    readonly allowIpSans?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates for localhost
     */
    readonly allowLocalhost?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching subdomains
     */
    readonly allowSubdomains?: pulumi.Input<boolean>;
    /**
     * List of allowed domains for certificates
     */
    readonly allowedDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed custom SANs
     */
    readonly allowedOtherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed URI SANs
     */
    readonly allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Flag to mark basic constraints valid when issuing non-CA certificates
     */
    readonly basicConstraintsValidForNonCa?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for client use
     */
    readonly clientFlag?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for code signing use
     */
    readonly codeSigningFlag?: pulumi.Input<boolean>;
    /**
     * The country of generated certificates
     */
    readonly countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to specify certificates for email protection use
     */
    readonly emailProtectionFlag?: pulumi.Input<boolean>;
    /**
     * Flag to allow only valid host names
     */
    readonly enforceHostnames?: pulumi.Input<boolean>;
    /**
     * Specify the allowed extended key usage constraint on issued certificates
     */
    readonly extKeyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to generate leases with certificates
     */
    readonly generateLease?: pulumi.Input<boolean>;
    /**
     * The number of bits of generated keys
     */
    readonly keyBits?: pulumi.Input<number>;
    /**
     * The type of generated keys
     */
    readonly keyType?: pulumi.Input<string>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    readonly keyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locality of generated certificates
     */
    readonly localities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum TTL
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Flag to not store certificates in the storage backend
     */
    readonly noStore?: pulumi.Input<boolean>;
    /**
     * Specifies the duration by which to backdate the NotBefore property.
     */
    readonly notBeforeDuration?: pulumi.Input<string>;
    /**
     * The organization unit of generated certificates
     */
    readonly organizationUnit?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization of generated certificates
     */
    readonly organizations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the list of allowed policies IODs
     */
    readonly policyIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code of generated certificates
     */
    readonly postalCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The province of generated certificates
     */
    readonly provinces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to force CN usage
     */
    readonly requireCn?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for server use
     */
    readonly serverFlag?: pulumi.Input<boolean>;
    /**
     * The street address of generated certificates
     */
    readonly streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The TTL
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Flag to use the CN in the CSR
     */
    readonly useCsrCommonName?: pulumi.Input<boolean>;
    /**
     * Flag to use the SANs in the CSR
     */
    readonly useCsrSans?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * Flag to allow any name
     */
    readonly allowAnyName?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching the actual domain
     */
    readonly allowBareDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow names containing glob patterns.
     */
    readonly allowGlobDomains?: pulumi.Input<boolean>;
    /**
     * Flag to allow IP SANs
     */
    readonly allowIpSans?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates for localhost
     */
    readonly allowLocalhost?: pulumi.Input<boolean>;
    /**
     * Flag to allow certificates matching subdomains
     */
    readonly allowSubdomains?: pulumi.Input<boolean>;
    /**
     * List of allowed domains for certificates
     */
    readonly allowedDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed custom SANs
     */
    readonly allowedOtherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines allowed URI SANs
     */
    readonly allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Flag to mark basic constraints valid when issuing non-CA certificates
     */
    readonly basicConstraintsValidForNonCa?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for client use
     */
    readonly clientFlag?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for code signing use
     */
    readonly codeSigningFlag?: pulumi.Input<boolean>;
    /**
     * The country of generated certificates
     */
    readonly countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to specify certificates for email protection use
     */
    readonly emailProtectionFlag?: pulumi.Input<boolean>;
    /**
     * Flag to allow only valid host names
     */
    readonly enforceHostnames?: pulumi.Input<boolean>;
    /**
     * Specify the allowed extended key usage constraint on issued certificates
     */
    readonly extKeyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to generate leases with certificates
     */
    readonly generateLease?: pulumi.Input<boolean>;
    /**
     * The number of bits of generated keys
     */
    readonly keyBits?: pulumi.Input<number>;
    /**
     * The type of generated keys
     */
    readonly keyType?: pulumi.Input<string>;
    /**
     * Specify the allowed key usage constraint on issued certificates
     */
    readonly keyUsages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locality of generated certificates
     */
    readonly localities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum TTL
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * The name to identify this role within the backend. Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Flag to not store certificates in the storage backend
     */
    readonly noStore?: pulumi.Input<boolean>;
    /**
     * Specifies the duration by which to backdate the NotBefore property.
     */
    readonly notBeforeDuration?: pulumi.Input<string>;
    /**
     * The organization unit of generated certificates
     */
    readonly organizationUnit?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization of generated certificates
     */
    readonly organizations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the list of allowed policies IODs
     */
    readonly policyIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code of generated certificates
     */
    readonly postalCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The province of generated certificates
     */
    readonly provinces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to force CN usage
     */
    readonly requireCn?: pulumi.Input<boolean>;
    /**
     * Flag to specify certificates for server use
     */
    readonly serverFlag?: pulumi.Input<boolean>;
    /**
     * The street address of generated certificates
     */
    readonly streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The TTL
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Flag to use the CN in the CSR
     */
    readonly useCsrCommonName?: pulumi.Input<boolean>;
    /**
     * Flag to use the SANs in the CSR
     */
    readonly useCsrSans?: pulumi.Input<boolean>;
}
