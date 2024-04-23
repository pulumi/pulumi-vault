// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const cert = new vault.AuthBackend("cert", {
 *     path: "cert",
 *     type: "cert",
 * });
 * const certCertAuthBackendRole = new vault.CertAuthBackendRole("cert", {
 *     name: "foo",
 *     certificate: std.file({
 *         input: "/path/to/certs/ca-cert.pem",
 *     }).then(invoke => invoke.result),
 *     backend: cert.path,
 *     allowedNames: [
 *         "foo.example.org",
 *         "baz.example.org",
 *     ],
 *     tokenTtl: 300,
 *     tokenMaxTtl: 600,
 *     tokenPolicies: ["foo"],
 * });
 * ```
 */
export class CertAuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing CertAuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertAuthBackendRoleState, opts?: pulumi.CustomResourceOptions): CertAuthBackendRole {
        return new CertAuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/certAuthBackendRole:CertAuthBackendRole';

    /**
     * Returns true if the given object is an instance of CertAuthBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertAuthBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertAuthBackendRole.__pulumiType;
    }

    /**
     * Allowed the common names for authenticated client certificates
     */
    public readonly allowedCommonNames!: pulumi.Output<string[]>;
    /**
     * Allowed alternative dns names for authenticated client certificates
     */
    public readonly allowedDnsSans!: pulumi.Output<string[]>;
    /**
     * Allowed emails for authenticated client certificates
     */
    public readonly allowedEmailSans!: pulumi.Output<string[]>;
    /**
     * DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
     */
    public readonly allowedNames!: pulumi.Output<string[]>;
    /**
     * Allowed organization units for authenticated client certificates.
     */
    public readonly allowedOrganizationalUnits!: pulumi.Output<string[] | undefined>;
    /**
     * Allowed URIs for authenticated client certificates
     */
    public readonly allowedUriSans!: pulumi.Output<string[]>;
    /**
     * Path to the mounted Cert auth backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * CA certificate used to validate client certificates
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * The name to display on tokens issued under this role.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Name of the role
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
     * Any additional CA certificates
     * needed to verify OCSP responses. Provided as base64 encoded PEM data.
     * Requires Vault version 1.13+.
     */
    public readonly ocspCaCertificates!: pulumi.Output<string | undefined>;
    /**
     * If enabled, validate certificates'
     * revocation status using OCSP. Requires Vault version 1.13+.
     */
    public readonly ocspEnabled!: pulumi.Output<boolean>;
    /**
     * If true and an OCSP response cannot
     * be fetched or is of an unknown status, the login will proceed as if the
     * certificate has not been revoked.
     * Requires Vault version 1.13+.
     */
    public readonly ocspFailOpen!: pulumi.Output<boolean>;
    /**
     * If set to true, rather than
     * accepting the first successful OCSP response, query all servers and consider
     * the certificate valid only if all servers agree.
     * Requires Vault version 1.13+.
     */
    public readonly ocspQueryAllServers!: pulumi.Output<boolean>;
    /**
     * : A comma-separated list of OCSP
     * server addresses. If unset, the OCSP server is determined from the
     * AuthorityInformationAccess extension on the certificate being inspected.
     * Requires Vault version 1.13+.
     */
    public readonly ocspServersOverrides!: pulumi.Output<string[] | undefined>;
    /**
     * TLS extensions required on
     * client certificates
     */
    public readonly requiredExtensions!: pulumi.Output<string[]>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime of the generated token
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Period
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token to generate, service or batch
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;

    /**
     * Create a CertAuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertAuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertAuthBackendRoleArgs | CertAuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertAuthBackendRoleState | undefined;
            resourceInputs["allowedCommonNames"] = state ? state.allowedCommonNames : undefined;
            resourceInputs["allowedDnsSans"] = state ? state.allowedDnsSans : undefined;
            resourceInputs["allowedEmailSans"] = state ? state.allowedEmailSans : undefined;
            resourceInputs["allowedNames"] = state ? state.allowedNames : undefined;
            resourceInputs["allowedOrganizationalUnits"] = state ? state.allowedOrganizationalUnits : undefined;
            resourceInputs["allowedUriSans"] = state ? state.allowedUriSans : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["ocspCaCertificates"] = state ? state.ocspCaCertificates : undefined;
            resourceInputs["ocspEnabled"] = state ? state.ocspEnabled : undefined;
            resourceInputs["ocspFailOpen"] = state ? state.ocspFailOpen : undefined;
            resourceInputs["ocspQueryAllServers"] = state ? state.ocspQueryAllServers : undefined;
            resourceInputs["ocspServersOverrides"] = state ? state.ocspServersOverrides : undefined;
            resourceInputs["requiredExtensions"] = state ? state.requiredExtensions : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
        } else {
            const args = argsOrState as CertAuthBackendRoleArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            resourceInputs["allowedCommonNames"] = args ? args.allowedCommonNames : undefined;
            resourceInputs["allowedDnsSans"] = args ? args.allowedDnsSans : undefined;
            resourceInputs["allowedEmailSans"] = args ? args.allowedEmailSans : undefined;
            resourceInputs["allowedNames"] = args ? args.allowedNames : undefined;
            resourceInputs["allowedOrganizationalUnits"] = args ? args.allowedOrganizationalUnits : undefined;
            resourceInputs["allowedUriSans"] = args ? args.allowedUriSans : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["ocspCaCertificates"] = args ? args.ocspCaCertificates : undefined;
            resourceInputs["ocspEnabled"] = args ? args.ocspEnabled : undefined;
            resourceInputs["ocspFailOpen"] = args ? args.ocspFailOpen : undefined;
            resourceInputs["ocspQueryAllServers"] = args ? args.ocspQueryAllServers : undefined;
            resourceInputs["ocspServersOverrides"] = args ? args.ocspServersOverrides : undefined;
            resourceInputs["requiredExtensions"] = args ? args.requiredExtensions : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CertAuthBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertAuthBackendRole resources.
 */
export interface CertAuthBackendRoleState {
    /**
     * Allowed the common names for authenticated client certificates
     */
    allowedCommonNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed alternative dns names for authenticated client certificates
     */
    allowedDnsSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed emails for authenticated client certificates
     */
    allowedEmailSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
     */
    allowedNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed organization units for authenticated client certificates.
     */
    allowedOrganizationalUnits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed URIs for authenticated client certificates
     */
    allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Path to the mounted Cert auth backend
     */
    backend?: pulumi.Input<string>;
    /**
     * CA certificate used to validate client certificates
     */
    certificate?: pulumi.Input<string>;
    /**
     * The name to display on tokens issued under this role.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Name of the role
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
     * Any additional CA certificates
     * needed to verify OCSP responses. Provided as base64 encoded PEM data.
     * Requires Vault version 1.13+.
     */
    ocspCaCertificates?: pulumi.Input<string>;
    /**
     * If enabled, validate certificates'
     * revocation status using OCSP. Requires Vault version 1.13+.
     */
    ocspEnabled?: pulumi.Input<boolean>;
    /**
     * If true and an OCSP response cannot
     * be fetched or is of an unknown status, the login will proceed as if the
     * certificate has not been revoked.
     * Requires Vault version 1.13+.
     */
    ocspFailOpen?: pulumi.Input<boolean>;
    /**
     * If set to true, rather than
     * accepting the first successful OCSP response, query all servers and consider
     * the certificate valid only if all servers agree.
     * Requires Vault version 1.13+.
     */
    ocspQueryAllServers?: pulumi.Input<boolean>;
    /**
     * : A comma-separated list of OCSP
     * server addresses. If unset, the OCSP server is determined from the
     * AuthorityInformationAccess extension on the certificate being inspected.
     * Requires Vault version 1.13+.
     */
    ocspServersOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * TLS extensions required on
     * client certificates
     */
    requiredExtensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertAuthBackendRole resource.
 */
export interface CertAuthBackendRoleArgs {
    /**
     * Allowed the common names for authenticated client certificates
     */
    allowedCommonNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed alternative dns names for authenticated client certificates
     */
    allowedDnsSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed emails for authenticated client certificates
     */
    allowedEmailSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
     */
    allowedNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed organization units for authenticated client certificates.
     */
    allowedOrganizationalUnits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed URIs for authenticated client certificates
     */
    allowedUriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Path to the mounted Cert auth backend
     */
    backend?: pulumi.Input<string>;
    /**
     * CA certificate used to validate client certificates
     */
    certificate: pulumi.Input<string>;
    /**
     * The name to display on tokens issued under this role.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Name of the role
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
     * Any additional CA certificates
     * needed to verify OCSP responses. Provided as base64 encoded PEM data.
     * Requires Vault version 1.13+.
     */
    ocspCaCertificates?: pulumi.Input<string>;
    /**
     * If enabled, validate certificates'
     * revocation status using OCSP. Requires Vault version 1.13+.
     */
    ocspEnabled?: pulumi.Input<boolean>;
    /**
     * If true and an OCSP response cannot
     * be fetched or is of an unknown status, the login will proceed as if the
     * certificate has not been revoked.
     * Requires Vault version 1.13+.
     */
    ocspFailOpen?: pulumi.Input<boolean>;
    /**
     * If set to true, rather than
     * accepting the first successful OCSP response, query all servers and consider
     * the certificate valid only if all servers agree.
     * Requires Vault version 1.13+.
     */
    ocspQueryAllServers?: pulumi.Input<boolean>;
    /**
     * : A comma-separated list of OCSP
     * server addresses. If unset, the OCSP server is determined from the
     * AuthorityInformationAccess extension on the certificate being inspected.
     * Requires Vault version 1.13+.
     */
    ocspServersOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * TLS extensions required on
     * client certificates
     */
    requiredExtensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
}
