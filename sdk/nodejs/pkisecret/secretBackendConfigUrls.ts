// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.
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
 * const configUrls = new vault.pkiSecret.SecretBackendConfigUrls("config_urls", {
 *     backend: pki.path,
 *     issuingCertificates: "http://127.0.0.1:8200/v1/pki/ca",
 * });
 * ```
 */
export class SecretBackendConfigUrls extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendConfigUrls resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendConfigUrlsState, opts?: pulumi.CustomResourceOptions): SecretBackendConfigUrls {
        return new SecretBackendConfigUrls(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls';

    /**
     * Returns true if the given object is an instance of SecretBackendConfigUrls.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendConfigUrls {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendConfigUrls.__pulumiType;
    }

    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies the URL values for the CRL Distribution Points field. Comma-separated string if multiple.
     */
    public readonly crlDistributionPoints!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the URL values for the Issuing Certificate field. Comma-separated string if multiple.
     */
    public readonly issuingCertificates!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the URL values for the OCSP Servers field. Comma-separated string if multiple.
     */
    public readonly ocspServers!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendConfigUrls resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendConfigUrlsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendConfigUrlsArgs | SecretBackendConfigUrlsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendConfigUrlsState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["crlDistributionPoints"] = state ? state.crlDistributionPoints : undefined;
            inputs["issuingCertificates"] = state ? state.issuingCertificates : undefined;
            inputs["ocspServers"] = state ? state.ocspServers : undefined;
        } else {
            const args = argsOrState as SecretBackendConfigUrlsArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["crlDistributionPoints"] = args ? args.crlDistributionPoints : undefined;
            inputs["issuingCertificates"] = args ? args.issuingCertificates : undefined;
            inputs["ocspServers"] = args ? args.ocspServers : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendConfigUrls.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendConfigUrls resources.
 */
export interface SecretBackendConfigUrlsState {
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Specifies the URL values for the CRL Distribution Points field. Comma-separated string if multiple.
     */
    readonly crlDistributionPoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the URL values for the Issuing Certificate field. Comma-separated string if multiple.
     */
    readonly issuingCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the URL values for the OCSP Servers field. Comma-separated string if multiple.
     */
    readonly ocspServers?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendConfigUrls resource.
 */
export interface SecretBackendConfigUrlsArgs {
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Specifies the URL values for the CRL Distribution Points field. Comma-separated string if multiple.
     */
    readonly crlDistributionPoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the URL values for the Issuing Certificate field. Comma-separated string if multiple.
     */
    readonly issuingCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the URL values for the OCSP Servers field. Comma-separated string if multiple.
     */
    readonly ocspServers?: pulumi.Input<pulumi.Input<string>[]>;
}
