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
 * const root = new vault.Mount("root", {
 *     path: "pki-root",
 *     type: "pki",
 *     description: "root",
 *     defaultLeaseTtlSeconds: 8640000,
 *     maxLeaseTtlSeconds: 8640000,
 * });
 * const intermediate = new vault.Mount("intermediate", {
 *     path: "pki-int",
 *     type: root.type,
 *     description: "intermediate",
 *     defaultLeaseTtlSeconds: 86400,
 *     maxLeaseTtlSeconds: 86400,
 * });
 * const example = new vault.pkisecret.SecretBackendRootCert("example", {
 *     backend: root.path,
 *     type: "internal",
 *     commonName: "RootOrg Root CA",
 *     ttl: "86400",
 *     format: "pem",
 *     privateKeyFormat: "der",
 *     keyType: "rsa",
 *     keyBits: 4096,
 *     excludeCnFromSans: true,
 *     ou: "Organizational Unit",
 *     organization: "RootOrg",
 *     country: "US",
 *     locality: "San Francisco",
 *     province: "CA",
 * });
 * const exampleSecretBackendIntermediateCertRequest = new vault.pkisecret.SecretBackendIntermediateCertRequest("example", {
 *     backend: intermediate.path,
 *     type: example.type,
 *     commonName: "SubOrg Intermediate CA",
 * });
 * const exampleSecretBackendRootSignIntermediate = new vault.pkisecret.SecretBackendRootSignIntermediate("example", {
 *     backend: root.path,
 *     csr: exampleSecretBackendIntermediateCertRequest.csr,
 *     commonName: "SubOrg Intermediate CA",
 *     excludeCnFromSans: true,
 *     ou: "SubUnit",
 *     organization: "SubOrg",
 *     country: "US",
 *     locality: "San Francisco",
 *     province: "CA",
 *     revoke: true,
 * });
 * const exampleSecretBackendIntermediateSetSigned = new vault.pkisecret.SecretBackendIntermediateSetSigned("example", {
 *     backend: intermediate.path,
 *     certificate: exampleSecretBackendRootSignIntermediate.certificate,
 * });
 * ```
 */
export class SecretBackendIntermediateSetSigned extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendIntermediateSetSigned resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendIntermediateSetSignedState, opts?: pulumi.CustomResourceOptions): SecretBackendIntermediateSetSigned {
        return new SecretBackendIntermediateSetSigned(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned';

    /**
     * Returns true if the given object is an instance of SecretBackendIntermediateSetSigned.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendIntermediateSetSigned {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendIntermediateSetSigned.__pulumiType;
    }

    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies the PEM encoded certificate. May optionally append additional
     * CA certificates to populate the whole chain, which will then enable returning the full chain from
     * issue and sign operations.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * The imported issuers indicating which issuers were created as part of
     * this request.
     */
    public /*out*/ readonly importedIssuers!: pulumi.Output<string[]>;
    /**
     * The imported keys indicating which keys were created as part of this request.
     */
    public /*out*/ readonly importedKeys!: pulumi.Output<string[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackendIntermediateSetSigned resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendIntermediateSetSignedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendIntermediateSetSignedArgs | SecretBackendIntermediateSetSignedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendIntermediateSetSignedState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["importedIssuers"] = state ? state.importedIssuers : undefined;
            resourceInputs["importedKeys"] = state ? state.importedKeys : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as SecretBackendIntermediateSetSignedArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["importedIssuers"] = undefined /*out*/;
            resourceInputs["importedKeys"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendIntermediateSetSigned.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendIntermediateSetSigned resources.
 */
export interface SecretBackendIntermediateSetSignedState {
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend?: pulumi.Input<string>;
    /**
     * Specifies the PEM encoded certificate. May optionally append additional
     * CA certificates to populate the whole chain, which will then enable returning the full chain from
     * issue and sign operations.
     */
    certificate?: pulumi.Input<string>;
    /**
     * The imported issuers indicating which issuers were created as part of
     * this request.
     */
    importedIssuers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The imported keys indicating which keys were created as part of this request.
     */
    importedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendIntermediateSetSigned resource.
 */
export interface SecretBackendIntermediateSetSignedArgs {
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend: pulumi.Input<string>;
    /**
     * Specifies the PEM encoded certificate. May optionally append additional
     * CA certificates to populate the whole chain, which will then enable returning the full chain from
     * issue and sign operations.
     */
    certificate: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
