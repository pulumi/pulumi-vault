// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
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
 * const root = new vault.pkisecret.SecretBackendRootCert("root", {
 *     backend: pki.path,
 *     type: "internal",
 *     commonName: "test",
 *     ttl: "86400",
 * });
 * const example = new vault.pkisecret.SecretBackendIssuer("example", {
 *     backend: root.backend,
 *     issuerRef: root.issuerId,
 *     issuerName: "example-issuer",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * PKI secret backend issuer can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer example pki/issuer/bf9b0d48-d0dd-652c-30be-77d04fc7e94d
 * ```
 */
export class SecretBackendIssuer extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendIssuer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendIssuerState, opts?: pulumi.CustomResourceOptions): SecretBackendIssuer {
        return new SecretBackendIssuer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer';

    /**
     * Returns true if the given object is an instance of SecretBackendIssuer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendIssuer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendIssuer.__pulumiType;
    }

    /**
     * The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies the URL values for the CRL
     * Distribution Points field.
     */
    public readonly crlDistributionPoints!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies that the AIA URL values should
     * be templated.
     */
    public readonly enableAiaUrlTemplating!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the issuer.
     */
    public /*out*/ readonly issuerId!: pulumi.Output<string>;
    /**
     * Name of the issuer.
     */
    public readonly issuerName!: pulumi.Output<string | undefined>;
    /**
     * Reference to an existing issuer.
     */
    public readonly issuerRef!: pulumi.Output<string>;
    /**
     * Specifies the URL values for the Issuing
     * Certificate field.
     */
    public readonly issuingCertificates!: pulumi.Output<string[] | undefined>;
    /**
     * Behavior of a leaf's NotAfter field during
     * issuance.
     */
    public readonly leafNotAfterBehavior!: pulumi.Output<string>;
    /**
     * Chain of issuer references to build this issuer's
     * computed CAChain field from, when non-empty.
     */
    public readonly manualChains!: pulumi.Output<string[] | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Specifies the URL values for the OCSP Servers field.
     */
    public readonly ocspServers!: pulumi.Output<string[] | undefined>;
    /**
     * Which signature algorithm to use
     * when building CRLs.
     */
    public readonly revocationSignatureAlgorithm!: pulumi.Output<string>;
    /**
     * Allowed usages for this issuer.
     */
    public readonly usage!: pulumi.Output<string>;

    /**
     * Create a SecretBackendIssuer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendIssuerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendIssuerArgs | SecretBackendIssuerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendIssuerState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["crlDistributionPoints"] = state ? state.crlDistributionPoints : undefined;
            resourceInputs["enableAiaUrlTemplating"] = state ? state.enableAiaUrlTemplating : undefined;
            resourceInputs["issuerId"] = state ? state.issuerId : undefined;
            resourceInputs["issuerName"] = state ? state.issuerName : undefined;
            resourceInputs["issuerRef"] = state ? state.issuerRef : undefined;
            resourceInputs["issuingCertificates"] = state ? state.issuingCertificates : undefined;
            resourceInputs["leafNotAfterBehavior"] = state ? state.leafNotAfterBehavior : undefined;
            resourceInputs["manualChains"] = state ? state.manualChains : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["ocspServers"] = state ? state.ocspServers : undefined;
            resourceInputs["revocationSignatureAlgorithm"] = state ? state.revocationSignatureAlgorithm : undefined;
            resourceInputs["usage"] = state ? state.usage : undefined;
        } else {
            const args = argsOrState as SecretBackendIssuerArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.issuerRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuerRef'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["crlDistributionPoints"] = args ? args.crlDistributionPoints : undefined;
            resourceInputs["enableAiaUrlTemplating"] = args ? args.enableAiaUrlTemplating : undefined;
            resourceInputs["issuerName"] = args ? args.issuerName : undefined;
            resourceInputs["issuerRef"] = args ? args.issuerRef : undefined;
            resourceInputs["issuingCertificates"] = args ? args.issuingCertificates : undefined;
            resourceInputs["leafNotAfterBehavior"] = args ? args.leafNotAfterBehavior : undefined;
            resourceInputs["manualChains"] = args ? args.manualChains : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["ocspServers"] = args ? args.ocspServers : undefined;
            resourceInputs["revocationSignatureAlgorithm"] = args ? args.revocationSignatureAlgorithm : undefined;
            resourceInputs["usage"] = args ? args.usage : undefined;
            resourceInputs["issuerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendIssuer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendIssuer resources.
 */
export interface SecretBackendIssuerState {
    /**
     * The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * Specifies the URL values for the CRL
     * Distribution Points field.
     */
    crlDistributionPoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies that the AIA URL values should
     * be templated.
     */
    enableAiaUrlTemplating?: pulumi.Input<boolean>;
    /**
     * ID of the issuer.
     */
    issuerId?: pulumi.Input<string>;
    /**
     * Name of the issuer.
     */
    issuerName?: pulumi.Input<string>;
    /**
     * Reference to an existing issuer.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * Specifies the URL values for the Issuing
     * Certificate field.
     */
    issuingCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Behavior of a leaf's NotAfter field during
     * issuance.
     */
    leafNotAfterBehavior?: pulumi.Input<string>;
    /**
     * Chain of issuer references to build this issuer's
     * computed CAChain field from, when non-empty.
     */
    manualChains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the URL values for the OCSP Servers field.
     */
    ocspServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Which signature algorithm to use
     * when building CRLs.
     */
    revocationSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Allowed usages for this issuer.
     */
    usage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendIssuer resource.
 */
export interface SecretBackendIssuerArgs {
    /**
     * The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Specifies the URL values for the CRL
     * Distribution Points field.
     */
    crlDistributionPoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies that the AIA URL values should
     * be templated.
     */
    enableAiaUrlTemplating?: pulumi.Input<boolean>;
    /**
     * Name of the issuer.
     */
    issuerName?: pulumi.Input<string>;
    /**
     * Reference to an existing issuer.
     */
    issuerRef: pulumi.Input<string>;
    /**
     * Specifies the URL values for the Issuing
     * Certificate field.
     */
    issuingCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Behavior of a leaf's NotAfter field during
     * issuance.
     */
    leafNotAfterBehavior?: pulumi.Input<string>;
    /**
     * Chain of issuer references to build this issuer's
     * computed CAChain field from, when non-empty.
     */
    manualChains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the URL values for the OCSP Servers field.
     */
    ocspServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Which signature algorithm to use
     * when building CRLs.
     */
    revocationSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Allowed usages for this issuer.
     */
    usage?: pulumi.Input<string>;
}
