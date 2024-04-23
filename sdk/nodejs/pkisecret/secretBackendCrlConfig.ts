// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const pki = new vault.Mount("pki", {
 *     path: "%s",
 *     type: "pki",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 * });
 * const crlConfig = new vault.pkisecret.SecretBackendCrlConfig("crl_config", {
 *     backend: pki.path,
 *     expiry: "72h",
 *     disable: false,
 * });
 * ```
 */
export class SecretBackendCrlConfig extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendCrlConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendCrlConfigState, opts?: pulumi.CustomResourceOptions): SecretBackendCrlConfig {
        return new SecretBackendCrlConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig';

    /**
     * Returns true if the given object is an instance of SecretBackendCrlConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendCrlConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendCrlConfig.__pulumiType;
    }

    /**
     * Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
     */
    public readonly autoRebuild!: pulumi.Output<boolean | undefined>;
    /**
     * Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
     */
    public readonly autoRebuildGracePeriod!: pulumi.Output<string>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Enable cross-cluster revocation request queues. **Vault 1.13+**
     */
    public readonly crossClusterRevocation!: pulumi.Output<boolean>;
    /**
     * Interval to check for new revocations on, to regenerate the delta CRL.
     */
    public readonly deltaRebuildInterval!: pulumi.Output<string>;
    /**
     * Disables or enables CRL building.
     */
    public readonly disable!: pulumi.Output<boolean | undefined>;
    /**
     * Enables building of delta CRLs with up-to-date revocation information, 
     * augmenting the last complete CRL.  **Vault 1.12+**
     */
    public readonly enableDelta!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the time until expiration.
     */
    public readonly expiry!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Disables the OCSP responder in Vault. **Vault 1.12+**
     */
    public readonly ocspDisable!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of time an OCSP response can be cached for, useful for OCSP stapling 
     * refresh durations. **Vault 1.12+**
     */
    public readonly ocspExpiry!: pulumi.Output<string>;
    /**
     * Enables unified CRL and OCSP building. **Vault 1.13+**
     */
    public readonly unifiedCrl!: pulumi.Output<boolean>;
    /**
     * Enables serving the unified CRL and OCSP on the existing, previously
     * cluster-local paths. **Vault 1.13+**
     */
    public readonly unifiedCrlOnExistingPaths!: pulumi.Output<boolean>;

    /**
     * Create a SecretBackendCrlConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendCrlConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendCrlConfigArgs | SecretBackendCrlConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendCrlConfigState | undefined;
            resourceInputs["autoRebuild"] = state ? state.autoRebuild : undefined;
            resourceInputs["autoRebuildGracePeriod"] = state ? state.autoRebuildGracePeriod : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["crossClusterRevocation"] = state ? state.crossClusterRevocation : undefined;
            resourceInputs["deltaRebuildInterval"] = state ? state.deltaRebuildInterval : undefined;
            resourceInputs["disable"] = state ? state.disable : undefined;
            resourceInputs["enableDelta"] = state ? state.enableDelta : undefined;
            resourceInputs["expiry"] = state ? state.expiry : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["ocspDisable"] = state ? state.ocspDisable : undefined;
            resourceInputs["ocspExpiry"] = state ? state.ocspExpiry : undefined;
            resourceInputs["unifiedCrl"] = state ? state.unifiedCrl : undefined;
            resourceInputs["unifiedCrlOnExistingPaths"] = state ? state.unifiedCrlOnExistingPaths : undefined;
        } else {
            const args = argsOrState as SecretBackendCrlConfigArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            resourceInputs["autoRebuild"] = args ? args.autoRebuild : undefined;
            resourceInputs["autoRebuildGracePeriod"] = args ? args.autoRebuildGracePeriod : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["crossClusterRevocation"] = args ? args.crossClusterRevocation : undefined;
            resourceInputs["deltaRebuildInterval"] = args ? args.deltaRebuildInterval : undefined;
            resourceInputs["disable"] = args ? args.disable : undefined;
            resourceInputs["enableDelta"] = args ? args.enableDelta : undefined;
            resourceInputs["expiry"] = args ? args.expiry : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["ocspDisable"] = args ? args.ocspDisable : undefined;
            resourceInputs["ocspExpiry"] = args ? args.ocspExpiry : undefined;
            resourceInputs["unifiedCrl"] = args ? args.unifiedCrl : undefined;
            resourceInputs["unifiedCrlOnExistingPaths"] = args ? args.unifiedCrlOnExistingPaths : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendCrlConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendCrlConfig resources.
 */
export interface SecretBackendCrlConfigState {
    /**
     * Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
     */
    autoRebuild?: pulumi.Input<boolean>;
    /**
     * Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
     */
    autoRebuildGracePeriod?: pulumi.Input<string>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * Enable cross-cluster revocation request queues. **Vault 1.13+**
     */
    crossClusterRevocation?: pulumi.Input<boolean>;
    /**
     * Interval to check for new revocations on, to regenerate the delta CRL.
     */
    deltaRebuildInterval?: pulumi.Input<string>;
    /**
     * Disables or enables CRL building.
     */
    disable?: pulumi.Input<boolean>;
    /**
     * Enables building of delta CRLs with up-to-date revocation information, 
     * augmenting the last complete CRL.  **Vault 1.12+**
     */
    enableDelta?: pulumi.Input<boolean>;
    /**
     * Specifies the time until expiration.
     */
    expiry?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Disables the OCSP responder in Vault. **Vault 1.12+**
     */
    ocspDisable?: pulumi.Input<boolean>;
    /**
     * The amount of time an OCSP response can be cached for, useful for OCSP stapling 
     * refresh durations. **Vault 1.12+**
     */
    ocspExpiry?: pulumi.Input<string>;
    /**
     * Enables unified CRL and OCSP building. **Vault 1.13+**
     */
    unifiedCrl?: pulumi.Input<boolean>;
    /**
     * Enables serving the unified CRL and OCSP on the existing, previously
     * cluster-local paths. **Vault 1.13+**
     */
    unifiedCrlOnExistingPaths?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SecretBackendCrlConfig resource.
 */
export interface SecretBackendCrlConfigArgs {
    /**
     * Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
     */
    autoRebuild?: pulumi.Input<boolean>;
    /**
     * Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
     */
    autoRebuildGracePeriod?: pulumi.Input<string>;
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Enable cross-cluster revocation request queues. **Vault 1.13+**
     */
    crossClusterRevocation?: pulumi.Input<boolean>;
    /**
     * Interval to check for new revocations on, to regenerate the delta CRL.
     */
    deltaRebuildInterval?: pulumi.Input<string>;
    /**
     * Disables or enables CRL building.
     */
    disable?: pulumi.Input<boolean>;
    /**
     * Enables building of delta CRLs with up-to-date revocation information, 
     * augmenting the last complete CRL.  **Vault 1.12+**
     */
    enableDelta?: pulumi.Input<boolean>;
    /**
     * Specifies the time until expiration.
     */
    expiry?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Disables the OCSP responder in Vault. **Vault 1.12+**
     */
    ocspDisable?: pulumi.Input<boolean>;
    /**
     * The amount of time an OCSP response can be cached for, useful for OCSP stapling 
     * refresh durations. **Vault 1.12+**
     */
    ocspExpiry?: pulumi.Input<string>;
    /**
     * Enables unified CRL and OCSP building. **Vault 1.13+**
     */
    unifiedCrl?: pulumi.Input<boolean>;
    /**
     * Enables serving the unified CRL and OCSP on the existing, previously
     * cluster-local paths. **Vault 1.13+**
     */
    unifiedCrlOnExistingPaths?: pulumi.Input<boolean>;
}
