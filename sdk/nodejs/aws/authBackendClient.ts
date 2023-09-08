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
 * const exampleAuthBackend = new vault.AuthBackend("exampleAuthBackend", {type: "aws"});
 * const exampleAuthBackendClient = new vault.aws.AuthBackendClient("exampleAuthBackendClient", {
 *     backend: exampleAuthBackend.path,
 *     accessKey: "INSERT_AWS_ACCESS_KEY",
 *     secretKey: "INSERT_AWS_SECRET_KEY",
 * });
 * ```
 *
 * ## Import
 *
 * AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.
 *
 * ```sh
 *  $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
 * ```
 */
export class AuthBackendClient extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendClientState, opts?: pulumi.CustomResourceOptions): AuthBackendClient {
        return new AuthBackendClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendClient:AuthBackendClient';

    /**
     * Returns true if the given object is an instance of AuthBackendClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendClient.__pulumiType;
    }

    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    public readonly ec2Endpoint!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    public readonly iamEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    public readonly iamServerIdHeaderValue!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    public readonly stsEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Override the default region when making STS API 
     * calls. The `stsEndpoint` argument must be set when using `stsRegion`.
     */
    public readonly stsRegion!: pulumi.Output<string | undefined>;
    /**
     * Available in Vault v1.15+. If set, 
     * overrides both `stsEndpoint` and `stsRegion` to instead use the region
     * specified in the client request headers for IAM-based authentication.
     * This can be useful when you have client requests coming from different
     * regions and want flexibility in which regional STS API is used.
     */
    public readonly useStsRegionFromClient!: pulumi.Output<boolean>;

    /**
     * Create a AuthBackendClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendClientArgs | AuthBackendClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendClientState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["ec2Endpoint"] = state ? state.ec2Endpoint : undefined;
            resourceInputs["iamEndpoint"] = state ? state.iamEndpoint : undefined;
            resourceInputs["iamServerIdHeaderValue"] = state ? state.iamServerIdHeaderValue : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["stsEndpoint"] = state ? state.stsEndpoint : undefined;
            resourceInputs["stsRegion"] = state ? state.stsRegion : undefined;
            resourceInputs["useStsRegionFromClient"] = state ? state.useStsRegionFromClient : undefined;
        } else {
            const args = argsOrState as AuthBackendClientArgs | undefined;
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["ec2Endpoint"] = args ? args.ec2Endpoint : undefined;
            resourceInputs["iamEndpoint"] = args ? args.iamEndpoint : undefined;
            resourceInputs["iamServerIdHeaderValue"] = args ? args.iamServerIdHeaderValue : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["stsEndpoint"] = args ? args.stsEndpoint : undefined;
            resourceInputs["stsRegion"] = args ? args.stsRegion : undefined;
            resourceInputs["useStsRegionFromClient"] = args ? args.useStsRegionFromClient : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackendClient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendClient resources.
 */
export interface AuthBackendClientState {
    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    ec2Endpoint?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    iamEndpoint?: pulumi.Input<string>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    iamServerIdHeaderValue?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Override the default region when making STS API 
     * calls. The `stsEndpoint` argument must be set when using `stsRegion`.
     */
    stsRegion?: pulumi.Input<string>;
    /**
     * Available in Vault v1.15+. If set, 
     * overrides both `stsEndpoint` and `stsRegion` to instead use the region
     * specified in the client request headers for IAM-based authentication.
     * This can be useful when you have client requests coming from different
     * regions and want flexibility in which regional STS API is used.
     */
    useStsRegionFromClient?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AuthBackendClient resource.
 */
export interface AuthBackendClientArgs {
    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    backend?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    ec2Endpoint?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    iamEndpoint?: pulumi.Input<string>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    iamServerIdHeaderValue?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Override the default region when making STS API 
     * calls. The `stsEndpoint` argument must be set when using `stsRegion`.
     */
    stsRegion?: pulumi.Input<string>;
    /**
     * Available in Vault v1.15+. If set, 
     * overrides both `stsEndpoint` and `stsRegion` to instead use the region
     * specified in the client request headers for IAM-based authentication.
     * This can be useful when you have client requests coming from different
     * regions and want flexibility in which regional STS API is used.
     */
    useStsRegionFromClient?: pulumi.Input<boolean>;
}
