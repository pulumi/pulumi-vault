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
 * const test = new vault.pkisecret.SecretBackendSign("test", {
 *     backend: vault_mount.pki.path,
 *     csr: `-----BEGIN CERTIFICATE REQUEST-----
 * MIIEqDCCApACAQAwYzELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
 * ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEcMBoGA1UEAwwTY2Vy
 * dC50ZXN0Lm15LmRvbWFpbjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
 * AJupYCQ8UVCWII1Zof1c6YcSSaM9hEaDU78cfKP5RoSeH10BvrWRfT+mzCONVpNP
 * CW9Iabtvk6hm0ot6ilnndEyVJbc0g7hdDLBX5BM25D+DGZGJRKUz1V+uBrWmXtIt
 * Vonj7JTDTe7ViH0GDsB7CvqXFGXO2a2cDYBchLkL6vQiFPshxvUsLtwxuy/qdYgy
 * X6ya+AUoZcoQGy1XxNjfH6cPtWSWQGEp1oPR6vL9hU3laTZb3C+VV4jZem+he8/0
 * V+qV6fLG92WTXm2hmf8nrtUqqJ+C7mW/RJod+TviviBadIX0OHXW7k5HVsZood01
 * te8vMRUNJNiZfa9EMIK5oncbQn0LcM3Wo9VrjpL7jREb/4HCS2gswYGv7hzk9cCS
 * kVY4rDucchKbApuI3kfzmO7GFOF5eiSkYZpY/czNn7VVM3WCu6dpOX4+3rhgrZQw
 * kY14L930DaLVRUgve/zKVP2D2GHdEOs+MbV7s96UgigT9pXly/yHPj+1sSYqmnaD
 * 5b7jSeJusmzO/nrwXVGLsnezR87VzHl9Ux9g5s6zh+R+PrZuVxYsLvoUpaasH47O
 * gIcBzSb/6pSGZKAUizmYsHsR1k88dAvsQ+FsUDaNokdi9VndEB4QPmiFmjyLV+0I
 * 1TFoXop4sW11NPz1YCq+IxnYrEaIN3PyhY0GvBJDFY1/AgMBAAGgADANBgkqhkiG
 * 9w0BAQsFAAOCAgEActuqnqS8Y9UF7e08w7tR3FPzGecWreuvxILrlFEZJxiLPFqL
 * It7uJvtypCVQvz6UQzKdBYO7tMpRaWViB8DrWzXNZjLMrg+QHcpveg8C0Ett4scG
 * fnvLk6fTDFYrnGvwHTqiHos5i0y3bFLyS1BGwSpdLAykGtvC+VM8mRyw/Y7CPcKN
 * 77kebY/9xduW1g2uxWLr0x90RuQDv9psPojT+59tRLGSp5Kt0IeD3QtnAZEFE4aN
 * vt+Pd69eg3BgZ8ZeDgoqAw3yppvOkpAFiE5pw2qPZaM4SRphl4d2Lek2zNIMyZqv
 * do5zh356HOgXtDaSg0POnRGrN/Ua+LMCRTg6GEPUnx9uQb/zt8Zu0hIexDGyykp1
 * OGqtWlv/Nc8UYuS38v0BeB6bMPeoqQUjkqs8nHlAEFn0KlgYdtDC+7SdQx6wS4te
 * dBKRNDfC4lS3jYJgs55jHqonZgkpSi3bamlxpfpW0ukGBcmq91wRe4bOw/4uD/vf
 * UwqMWOdCYcU3mdYNjTWy22ORW3SGFQxMBwpUEURCSoeqWr6aJeQ7KAYkx1PrB5T8
 * OTEc13lWf+B0PU9UJuGTsmpIuImPDVd0EVDayr3mT5dDbqTVDbe8ppf2IswABmf0
 * o3DybUeUmknYjl109rdSf+76nuREICHatxXgN3xCMFuBaN4WLO+ksd6Y1Ys=
 * -----END CERTIFICATE REQUEST-----
 * `,
 *     commonName: "test.my.domain",
 * }, {
 *     dependsOn: [vault_pki_secret_backend_role.admin],
 * });
 * ```
 */
export class SecretBackendSign extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendSign resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendSignState, opts?: pulumi.CustomResourceOptions): SecretBackendSign {
        return new SecretBackendSign(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendSign:SecretBackendSign';

    /**
     * Returns true if the given object is an instance of SecretBackendSign.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendSign {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendSign.__pulumiType;
    }

    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The CA chain
     */
    public /*out*/ readonly caChains!: pulumi.Output<string[]>;
    /**
     * The certificate
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * CN of certificate to create
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * The CSR
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * Flag to exclude CN from SANs
     */
    public readonly excludeCnFromSans!: pulumi.Output<boolean | undefined>;
    /**
     * The expiration date of the certificate in unix epoch format
     */
    public /*out*/ readonly expiration!: pulumi.Output<number>;
    /**
     * The format of data
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * List of alternative IPs
     */
    public readonly ipSans!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    public readonly issuerRef!: pulumi.Output<string | undefined>;
    /**
     * The issuing CA
     */
    public /*out*/ readonly issuingCa!: pulumi.Output<string>;
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    public readonly minSecondsRemaining!: pulumi.Output<number | undefined>;
    /**
     * Name of the role to create the certificate against
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
     * List of other SANs
     */
    public readonly otherSans!: pulumi.Output<string[] | undefined>;
    /**
     * `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     */
    public /*out*/ readonly renewPending!: pulumi.Output<boolean>;
    /**
     * The certificate's serial number, hex formatted.
     */
    public /*out*/ readonly serialNumber!: pulumi.Output<string>;
    /**
     * Time to live
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * List of alternative URIs
     */
    public readonly uriSans!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendSign resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendSignArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendSignArgs | SecretBackendSignState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendSignState | undefined;
            resourceInputs["altNames"] = state ? state.altNames : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["caChains"] = state ? state.caChains : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["commonName"] = state ? state.commonName : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            resourceInputs["expiration"] = state ? state.expiration : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["ipSans"] = state ? state.ipSans : undefined;
            resourceInputs["issuerRef"] = state ? state.issuerRef : undefined;
            resourceInputs["issuingCa"] = state ? state.issuingCa : undefined;
            resourceInputs["minSecondsRemaining"] = state ? state.minSecondsRemaining : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["otherSans"] = state ? state.otherSans : undefined;
            resourceInputs["renewPending"] = state ? state.renewPending : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["uriSans"] = state ? state.uriSans : undefined;
        } else {
            const args = argsOrState as SecretBackendSignArgs | undefined;
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
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["commonName"] = args ? args.commonName : undefined;
            resourceInputs["csr"] = args ? args.csr : undefined;
            resourceInputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["ipSans"] = args ? args.ipSans : undefined;
            resourceInputs["issuerRef"] = args ? args.issuerRef : undefined;
            resourceInputs["minSecondsRemaining"] = args ? args.minSecondsRemaining : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["otherSans"] = args ? args.otherSans : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["uriSans"] = args ? args.uriSans : undefined;
            resourceInputs["caChains"] = undefined /*out*/;
            resourceInputs["certificate"] = undefined /*out*/;
            resourceInputs["expiration"] = undefined /*out*/;
            resourceInputs["issuingCa"] = undefined /*out*/;
            resourceInputs["renewPending"] = undefined /*out*/;
            resourceInputs["serialNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendSign.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendSign resources.
 */
export interface SecretBackendSignState {
    /**
     * List of alternative names
     */
    altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend?: pulumi.Input<string>;
    /**
     * The CA chain
     */
    caChains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The certificate
     */
    certificate?: pulumi.Input<string>;
    /**
     * CN of certificate to create
     */
    commonName?: pulumi.Input<string>;
    /**
     * The CSR
     */
    csr?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The expiration date of the certificate in unix epoch format
     */
    expiration?: pulumi.Input<number>;
    /**
     * The format of data
     */
    format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * The issuing CA
     */
    issuingCa?: pulumi.Input<string>;
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    minSecondsRemaining?: pulumi.Input<number>;
    /**
     * Name of the role to create the certificate against
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
     * List of other SANs
     */
    otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     */
    renewPending?: pulumi.Input<boolean>;
    /**
     * The certificate's serial number, hex formatted.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * Time to live
     */
    ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendSign resource.
 */
export interface SecretBackendSignArgs {
    /**
     * List of alternative names
     */
    altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    backend: pulumi.Input<string>;
    /**
     * CN of certificate to create
     */
    commonName: pulumi.Input<string>;
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
     * Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role's `issuerRef` value.
     */
    issuerRef?: pulumi.Input<string>;
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    minSecondsRemaining?: pulumi.Input<number>;
    /**
     * Name of the role to create the certificate against
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
     * List of other SANs
     */
    otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Time to live
     */
    ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}
