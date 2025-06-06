// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Kubernetes auth backend config in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kubernetes = new vault.AuthBackend("kubernetes", {type: "kubernetes"});
 * const example = new vault.kubernetes.AuthBackendConfig("example", {
 *     backend: kubernetes.path,
 *     kubernetesHost: "http://example.com:443",
 *     kubernetesCaCert: `-----BEGIN CERTIFICATE-----
 * example
 * -----END CERTIFICATE-----`,
 *     tokenReviewerJwt: "ZXhhbXBsZQo=",
 *     issuer: "api",
 *     disableIssValidation: true,
 * });
 * ```
 *
 * ## Import
 *
 * Kubernetes authentication backend can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:kubernetes/authBackendConfig:AuthBackendConfig config auth/kubernetes/config
 * ```
 */
export class AuthBackendConfig extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendConfigState, opts?: pulumi.CustomResourceOptions): AuthBackendConfig {
        return new AuthBackendConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kubernetes/authBackendConfig:AuthBackendConfig';

    /**
     * Returns true if the given object is an instance of AuthBackendConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendConfig.__pulumiType;
    }

    /**
     * Unique name of the kubernetes backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    public readonly disableIssValidation!: pulumi.Output<boolean>;
    /**
     * Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    public readonly disableLocalCaJwt!: pulumi.Output<boolean>;
    /**
     * JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    public readonly issuer!: pulumi.Output<string | undefined>;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    public readonly kubernetesCaCert!: pulumi.Output<string>;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    public readonly kubernetesHost!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    public readonly pemKeys!: pulumi.Output<string[] | undefined>;
    /**
     * A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     */
    public readonly tokenReviewerJwt!: pulumi.Output<string | undefined>;
    /**
     * Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    public readonly useAnnotationsAsAliasMetadata!: pulumi.Output<boolean>;

    /**
     * Create a AuthBackendConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendConfigArgs | AuthBackendConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendConfigState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["disableIssValidation"] = state ? state.disableIssValidation : undefined;
            resourceInputs["disableLocalCaJwt"] = state ? state.disableLocalCaJwt : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["kubernetesCaCert"] = state ? state.kubernetesCaCert : undefined;
            resourceInputs["kubernetesHost"] = state ? state.kubernetesHost : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["pemKeys"] = state ? state.pemKeys : undefined;
            resourceInputs["tokenReviewerJwt"] = state ? state.tokenReviewerJwt : undefined;
            resourceInputs["useAnnotationsAsAliasMetadata"] = state ? state.useAnnotationsAsAliasMetadata : undefined;
        } else {
            const args = argsOrState as AuthBackendConfigArgs | undefined;
            if ((!args || args.kubernetesHost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesHost'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["disableIssValidation"] = args ? args.disableIssValidation : undefined;
            resourceInputs["disableLocalCaJwt"] = args ? args.disableLocalCaJwt : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["kubernetesCaCert"] = args ? args.kubernetesCaCert : undefined;
            resourceInputs["kubernetesHost"] = args ? args.kubernetesHost : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["pemKeys"] = args ? args.pemKeys : undefined;
            resourceInputs["tokenReviewerJwt"] = args?.tokenReviewerJwt ? pulumi.secret(args.tokenReviewerJwt) : undefined;
            resourceInputs["useAnnotationsAsAliasMetadata"] = args ? args.useAnnotationsAsAliasMetadata : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tokenReviewerJwt"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackendConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendConfig resources.
 */
export interface AuthBackendConfigState {
    /**
     * Unique name of the kubernetes backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableIssValidation?: pulumi.Input<boolean>;
    /**
     * Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableLocalCaJwt?: pulumi.Input<boolean>;
    /**
     * JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    kubernetesHost?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    pemKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     */
    tokenReviewerJwt?: pulumi.Input<string>;
    /**
     * Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    useAnnotationsAsAliasMetadata?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AuthBackendConfig resource.
 */
export interface AuthBackendConfigArgs {
    /**
     * Unique name of the kubernetes backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableIssValidation?: pulumi.Input<boolean>;
    /**
     * Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableLocalCaJwt?: pulumi.Input<boolean>;
    /**
     * JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    kubernetesHost: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    pemKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     */
    tokenReviewerJwt?: pulumi.Input<string>;
    /**
     * Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    useAnnotationsAsAliasMetadata?: pulumi.Input<boolean>;
}
