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
 * const test = new vault.pkisecret.SecretBackendIntermediateCertRequest("test", {
 *     backend: pki.path,
 *     type: "internal",
 *     commonName: "app.my.domain",
 * }, {
 *     dependsOn: [pki],
 * });
 * ```
 */
export class SecretBackendIntermediateCertRequest extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendIntermediateCertRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendIntermediateCertRequestState, opts?: pulumi.CustomResourceOptions): SecretBackendIntermediateCertRequest {
        return new SecretBackendIntermediateCertRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest';

    /**
     * Returns true if the given object is an instance of SecretBackendIntermediateCertRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendIntermediateCertRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendIntermediateCertRequest.__pulumiType;
    }

    /**
     * Adds a Basic Constraints extension with 'CA: true'.
     * Only needed as a workaround in some compatibility scenarios with Active Directory
     * Certificate Services
     */
    public readonly addBasicConstraints!: pulumi.Output<boolean | undefined>;
    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
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
    public /*out*/ readonly csr!: pulumi.Output<string>;
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
     * The number of bits to use
     */
    public readonly keyBits!: pulumi.Output<number | undefined>;
    /**
     * The ID of the generated key.
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     */
    public readonly keyName!: pulumi.Output<string>;
    /**
     * Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     */
    public readonly keyRef!: pulumi.Output<string>;
    /**
     * The desired key type
     */
    public readonly keyType!: pulumi.Output<string | undefined>;
    /**
     * The locality
     */
    public readonly locality!: pulumi.Output<string | undefined>;
    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managedKeyName`
     */
    public readonly managedKeyId!: pulumi.Output<string | undefined>;
    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managedKeyId`
     */
    public readonly managedKeyName!: pulumi.Output<string | undefined>;
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
     * The postal code
     */
    public readonly postalCode!: pulumi.Output<string | undefined>;
    /**
     * The private key
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * The private key format
     */
    public readonly privateKeyFormat!: pulumi.Output<string | undefined>;
    /**
     * The private key type
     */
    public /*out*/ readonly privateKeyType!: pulumi.Output<string>;
    /**
     * The province
     */
    public readonly province!: pulumi.Output<string | undefined>;
    /**
     * The street address
     */
    public readonly streetAddress!: pulumi.Output<string | undefined>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     * or \"kms\"
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * List of alternative URIs
     */
    public readonly uriSans!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendIntermediateCertRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendIntermediateCertRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendIntermediateCertRequestArgs | SecretBackendIntermediateCertRequestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendIntermediateCertRequestState | undefined;
            resourceInputs["addBasicConstraints"] = state ? state.addBasicConstraints : undefined;
            resourceInputs["altNames"] = state ? state.altNames : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["commonName"] = state ? state.commonName : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["ipSans"] = state ? state.ipSans : undefined;
            resourceInputs["keyBits"] = state ? state.keyBits : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["keyName"] = state ? state.keyName : undefined;
            resourceInputs["keyRef"] = state ? state.keyRef : undefined;
            resourceInputs["keyType"] = state ? state.keyType : undefined;
            resourceInputs["locality"] = state ? state.locality : undefined;
            resourceInputs["managedKeyId"] = state ? state.managedKeyId : undefined;
            resourceInputs["managedKeyName"] = state ? state.managedKeyName : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["otherSans"] = state ? state.otherSans : undefined;
            resourceInputs["ou"] = state ? state.ou : undefined;
            resourceInputs["postalCode"] = state ? state.postalCode : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyFormat"] = state ? state.privateKeyFormat : undefined;
            resourceInputs["privateKeyType"] = state ? state.privateKeyType : undefined;
            resourceInputs["province"] = state ? state.province : undefined;
            resourceInputs["streetAddress"] = state ? state.streetAddress : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uriSans"] = state ? state.uriSans : undefined;
        } else {
            const args = argsOrState as SecretBackendIntermediateCertRequestArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.commonName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commonName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["addBasicConstraints"] = args ? args.addBasicConstraints : undefined;
            resourceInputs["altNames"] = args ? args.altNames : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["commonName"] = args ? args.commonName : undefined;
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["ipSans"] = args ? args.ipSans : undefined;
            resourceInputs["keyBits"] = args ? args.keyBits : undefined;
            resourceInputs["keyName"] = args ? args.keyName : undefined;
            resourceInputs["keyRef"] = args ? args.keyRef : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["locality"] = args ? args.locality : undefined;
            resourceInputs["managedKeyId"] = args ? args.managedKeyId : undefined;
            resourceInputs["managedKeyName"] = args ? args.managedKeyName : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["otherSans"] = args ? args.otherSans : undefined;
            resourceInputs["ou"] = args ? args.ou : undefined;
            resourceInputs["postalCode"] = args ? args.postalCode : undefined;
            resourceInputs["privateKeyFormat"] = args ? args.privateKeyFormat : undefined;
            resourceInputs["province"] = args ? args.province : undefined;
            resourceInputs["streetAddress"] = args ? args.streetAddress : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uriSans"] = args ? args.uriSans : undefined;
            resourceInputs["csr"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["privateKey"] = undefined /*out*/;
            resourceInputs["privateKeyType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackendIntermediateCertRequest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendIntermediateCertRequest resources.
 */
export interface SecretBackendIntermediateCertRequestState {
    /**
     * Adds a Basic Constraints extension with 'CA: true'.
     * Only needed as a workaround in some compatibility scenarios with Active Directory
     * Certificate Services
     */
    addBasicConstraints?: pulumi.Input<boolean>;
    /**
     * List of alternative names
     */
    altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend?: pulumi.Input<string>;
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
     * The number of bits to use
     */
    keyBits?: pulumi.Input<number>;
    /**
     * The ID of the generated key.
     */
    keyId?: pulumi.Input<string>;
    /**
     * When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     */
    keyName?: pulumi.Input<string>;
    /**
     * Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     */
    keyRef?: pulumi.Input<string>;
    /**
     * The desired key type
     */
    keyType?: pulumi.Input<string>;
    /**
     * The locality
     */
    locality?: pulumi.Input<string>;
    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managedKeyName`
     */
    managedKeyId?: pulumi.Input<string>;
    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managedKeyId`
     */
    managedKeyName?: pulumi.Input<string>;
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
     * The postal code
     */
    postalCode?: pulumi.Input<string>;
    /**
     * The private key
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The private key format
     */
    privateKeyFormat?: pulumi.Input<string>;
    /**
     * The private key type
     */
    privateKeyType?: pulumi.Input<string>;
    /**
     * The province
     */
    province?: pulumi.Input<string>;
    /**
     * The street address
     */
    streetAddress?: pulumi.Input<string>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     * or \"kms\"
     */
    type?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendIntermediateCertRequest resource.
 */
export interface SecretBackendIntermediateCertRequestArgs {
    /**
     * Adds a Basic Constraints extension with 'CA: true'.
     * Only needed as a workaround in some compatibility scenarios with Active Directory
     * Certificate Services
     */
    addBasicConstraints?: pulumi.Input<boolean>;
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
     * The number of bits to use
     */
    keyBits?: pulumi.Input<number>;
    /**
     * When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     */
    keyName?: pulumi.Input<string>;
    /**
     * Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     */
    keyRef?: pulumi.Input<string>;
    /**
     * The desired key type
     */
    keyType?: pulumi.Input<string>;
    /**
     * The locality
     */
    locality?: pulumi.Input<string>;
    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managedKeyName`
     */
    managedKeyId?: pulumi.Input<string>;
    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managedKeyId`
     */
    managedKeyName?: pulumi.Input<string>;
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
     * The postal code
     */
    postalCode?: pulumi.Input<string>;
    /**
     * The private key format
     */
    privateKeyFormat?: pulumi.Input<string>;
    /**
     * The province
     */
    province?: pulumi.Input<string>;
    /**
     * The street address
     */
    streetAddress?: pulumi.Input<string>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     * or \"kms\"
     */
    type: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}
