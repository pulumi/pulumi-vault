// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates PKI certificate.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const root = new vault.pkisecret.SecretBackendRootSignIntermediate("root", {
 *     backend: vault_mount.root.path,
 *     csr: vault_pki_secret_backend_intermediate_cert_request.intermediate.csr,
 *     commonName: "Intermediate CA",
 *     excludeCnFromSans: true,
 *     ou: "My OU",
 *     organization: "My organization",
 * }, {
 *     dependsOn: [vault_pki_secret_backend_intermediate_cert_request.intermediate],
 * });
 * ```
 */
export class SecretBackendRootSignIntermediate extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRootSignIntermediate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRootSignIntermediateState, opts?: pulumi.CustomResourceOptions): SecretBackendRootSignIntermediate {
        return new SecretBackendRootSignIntermediate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate';

    /**
     * Returns true if the given object is an instance of SecretBackendRootSignIntermediate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRootSignIntermediate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRootSignIntermediate.__pulumiType;
    }

    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * A list of the issuing and intermediate CA certificates in the `format` specified.
     */
    public /*out*/ readonly caChains!: pulumi.Output<string[]>;
    /**
     * The intermediate CA certificate in the `format` specified.
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded). 
     * Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
     */
    public /*out*/ readonly certificateBundle!: pulumi.Output<string>;
    /**
     * CN of intermediate to create
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * The country
     */
    public readonly country!: pulumi.Output<string | undefined>;
    /**
     * The CSR
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * Flag to exclude CN from SANs
     */
    public readonly excludeCnFromSans!: pulumi.Output<boolean | undefined>;
    /**
     * The format of data
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * List of alternative IPs
     */
    public readonly ipSans!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    public readonly issuerRef!: pulumi.Output<string | undefined>;
    /**
     * The issuing CA certificate in the `format` specified.
     */
    public /*out*/ readonly issuingCa!: pulumi.Output<string>;
    /**
     * The locality
     */
    public readonly locality!: pulumi.Output<string | undefined>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    public readonly maxPathLength!: pulumi.Output<number | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The organization
     */
    public readonly organization!: pulumi.Output<string | undefined>;
    /**
     * List of other SANs
     */
    public readonly otherSans!: pulumi.Output<string[] | undefined>;
    /**
     * The organization unit
     */
    public readonly ou!: pulumi.Output<string | undefined>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    public readonly permittedDnsDomains!: pulumi.Output<string[] | undefined>;
    /**
     * The postal code
     */
    public readonly postalCode!: pulumi.Output<string | undefined>;
    /**
     * The province
     */
    public readonly province!: pulumi.Output<string | undefined>;
    /**
     * If set to `true`, the certificate will be revoked on resource destruction.
     */
    public readonly revoke!: pulumi.Output<boolean | undefined>;
    /**
     * The certificate's serial number, hex formatted.
     */
    public /*out*/ readonly serialNumber!: pulumi.Output<string>;
    /**
     * The street address
     */
    public readonly streetAddress!: pulumi.Output<string | undefined>;
    /**
     * Time to live
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * List of alternative URIs
     */
    public readonly uriSans!: pulumi.Output<string[] | undefined>;
    /**
     * Preserve CSR values
     */
    public readonly useCsrValues!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SecretBackendRootSignIntermediate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRootSignIntermediateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRootSignIntermediateArgs | SecretBackendRootSignIntermediateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRootSignIntermediateState | undefined;
            resourceInputs["altNames"] = state ? state.altNames : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["caChains"] = state ? state.caChains : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateBundle"] = state ? state.certificateBundle : undefined;
            resourceInputs["commonName"] = state ? state.commonName : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["ipSans"] = state ? state.ipSans : undefined;
            resourceInputs["issuerRef"] = state ? state.issuerRef : undefined;
            resourceInputs["issuingCa"] = state ? state.issuingCa : undefined;
            resourceInputs["locality"] = state ? state.locality : undefined;
            resourceInputs["maxPathLength"] = state ? state.maxPathLength : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["otherSans"] = state ? state.otherSans : undefined;
            resourceInputs["ou"] = state ? state.ou : undefined;
            resourceInputs["permittedDnsDomains"] = state ? state.permittedDnsDomains : undefined;
            resourceInputs["postalCode"] = state ? state.postalCode : undefined;
            resourceInputs["province"] = state ? state.province : undefined;
            resourceInputs["revoke"] = state ? state.revoke : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["streetAddress"] = state ? state.streetAddress : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["uriSans"] = state ? state.uriSans : undefined;
            resourceInputs["useCsrValues"] = state ? state.useCsrValues : undefined;
        } else {
            const args = argsOrState as SecretBackendRootSignIntermediateArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.commonName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commonName'");
            }
            if ((!args || args.csr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'csr'");
            }
            resourceInputs["altNames"] = args ? args.altNames : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["commonName"] = args ? args.commonName : undefined;
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["csr"] = args ? args.csr : undefined;
            resourceInputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["ipSans"] = args ? args.ipSans : undefined;
            resourceInputs["issuerRef"] = args ? args.issuerRef : undefined;
            resourceInputs["locality"] = args ? args.locality : undefined;
            resourceInputs["maxPathLength"] = args ? args.maxPathLength : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["otherSans"] = args ? args.otherSans : undefined;
            resourceInputs["ou"] = args ? args.ou : undefined;
            resourceInputs["permittedDnsDomains"] = args ? args.permittedDnsDomains : undefined;
            resourceInputs["postalCode"] = args ? args.postalCode : undefined;
            resourceInputs["province"] = args ? args.province : undefined;
            resourceInputs["revoke"] = args ? args.revoke : undefined;
            resourceInputs["streetAddress"] = args ? args.streetAddress : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["uriSans"] = args ? args.uriSans : undefined;
            resourceInputs["useCsrValues"] = args ? args.useCsrValues : undefined;
            resourceInputs["caChains"] = undefined /*out*/;
            resourceInputs["certificate"] = undefined /*out*/;
            resourceInputs["certificateBundle"] = undefined /*out*/;
            resourceInputs["issuingCa"] = undefined /*out*/;
            resourceInputs["serialNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendRootSignIntermediate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRootSignIntermediate resources.
 */
export interface SecretBackendRootSignIntermediateState {
    /**
     * List of alternative names
     */
    altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend?: pulumi.Input<string>;
    /**
     * A list of the issuing and intermediate CA certificates in the `format` specified.
     */
    caChains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The intermediate CA certificate in the `format` specified.
     */
    certificate?: pulumi.Input<string>;
    /**
     * The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded). 
     * Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
     */
    certificateBundle?: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    commonName?: pulumi.Input<string>;
    /**
     * The country
     */
    country?: pulumi.Input<string>;
    /**
     * The CSR
     */
    csr?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * The issuing CA certificate in the `format` specified.
     */
    issuingCa?: pulumi.Input<string>;
    /**
     * The locality
     */
    locality?: pulumi.Input<string>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    maxPathLength?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The organization
     */
    organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    ou?: pulumi.Input<string>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    permittedDnsDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code
     */
    postalCode?: pulumi.Input<string>;
    /**
     * The province
     */
    province?: pulumi.Input<string>;
    /**
     * If set to `true`, the certificate will be revoked on resource destruction.
     */
    revoke?: pulumi.Input<boolean>;
    /**
     * The certificate's serial number, hex formatted.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * The street address
     */
    streetAddress?: pulumi.Input<string>;
    /**
     * Time to live
     */
    ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Preserve CSR values
     */
    useCsrValues?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SecretBackendRootSignIntermediate resource.
 */
export interface SecretBackendRootSignIntermediateArgs {
    /**
     * List of alternative names
     */
    altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    commonName: pulumi.Input<string>;
    /**
     * The country
     */
    country?: pulumi.Input<string>;
    /**
     * The CSR
     */
    csr: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * The locality
     */
    locality?: pulumi.Input<string>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    maxPathLength?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The organization
     */
    organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    ou?: pulumi.Input<string>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    permittedDnsDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code
     */
    postalCode?: pulumi.Input<string>;
    /**
     * The province
     */
    province?: pulumi.Input<string>;
    /**
     * If set to `true`, the certificate will be revoked on resource destruction.
     */
    revoke?: pulumi.Input<boolean>;
    /**
     * The street address
     */
    streetAddress?: pulumi.Input<string>;
    /**
     * Time to live
     */
    ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Preserve CSR values
     */
    useCsrValues?: pulumi.Input<boolean>;
}
