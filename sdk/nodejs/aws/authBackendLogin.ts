// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Logs into a Vault server using an AWS auth backend. Login can be
 * accomplished using a signed identity request from IAM or using ec2
 * instance metadata. For more information, see the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/aws.html).
 */
export class AuthBackendLogin extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendLogin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendLoginState, opts?: pulumi.CustomResourceOptions): AuthBackendLogin {
        return new AuthBackendLogin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendLogin:AuthBackendLogin';

    /**
     * Returns true if the given object is an instance of AuthBackendLogin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendLogin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendLogin.__pulumiType;
    }

    /**
     * The token's accessor.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The authentication type used to generate this token.
     */
    public /*out*/ readonly authType!: pulumi.Output<string>;
    /**
     * The unique name of the AWS auth backend. Defaults to
     * 'aws'.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The token returned by Vault.
     */
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * The HTTP method used in the signed IAM
     * request.
     */
    public readonly iamHttpRequestMethod!: pulumi.Output<string | undefined>;
    /**
     * The base64-encoded body of the signed
     * request.
     */
    public readonly iamRequestBody!: pulumi.Output<string | undefined>;
    /**
     * The base64-encoded, JSON serialized
     * representation of the GetCallerIdentity HTTP request headers.
     */
    public readonly iamRequestHeaders!: pulumi.Output<string | undefined>;
    /**
     * The base64-encoded HTTP URL used in the signed
     * request.
     */
    public readonly iamRequestUrl!: pulumi.Output<string | undefined>;
    /**
     * The base64-encoded EC2 instance identity document to
     * authenticate with. Can be retrieved from the EC2 metadata server.
     */
    public readonly identity!: pulumi.Output<string | undefined>;
    /**
     * The duration in seconds the token will be valid, relative
     * to the time in `leaseStartTime`.
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<number>;
    public /*out*/ readonly leaseStartTime!: pulumi.Output<string>;
    /**
     * A map of information returned by the Vault server about the
     * authentication used to generate this token.
     */
    public /*out*/ readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The unique nonce to be used for login requests. Can be
     * set to a user-specified value, or will contain the server-generated value
     * once a token is issued. EC2 instances can only acquire a single token until
     * the whitelist is tidied again unless they keep track of this nonce.
     */
    public readonly nonce!: pulumi.Output<string>;
    /**
     * The PKCS#7 signature of the identity document to
     * authenticate with, with all newline characters removed. Can be retrieved from
     * the EC2 metadata server.
     */
    public readonly pkcs7!: pulumi.Output<string | undefined>;
    /**
     * The Vault policies assigned to this token.
     */
    public /*out*/ readonly policies!: pulumi.Output<string[]>;
    /**
     * Set to true if the token can be extended through renewal.
     */
    public /*out*/ readonly renewable!: pulumi.Output<boolean>;
    /**
     * The name of the AWS auth backend role to create tokens
     * against.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The base64-encoded SHA256 RSA signature of the
     * instance identity document to authenticate with, with all newline characters
     * removed. Can be retrieved from the EC2 metadata server.
     */
    public readonly signature!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendLogin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendLoginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendLoginArgs | AuthBackendLoginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendLoginState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["iamHttpRequestMethod"] = state ? state.iamHttpRequestMethod : undefined;
            resourceInputs["iamRequestBody"] = state ? state.iamRequestBody : undefined;
            resourceInputs["iamRequestHeaders"] = state ? state.iamRequestHeaders : undefined;
            resourceInputs["iamRequestUrl"] = state ? state.iamRequestUrl : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["leaseDuration"] = state ? state.leaseDuration : undefined;
            resourceInputs["leaseStartTime"] = state ? state.leaseStartTime : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["nonce"] = state ? state.nonce : undefined;
            resourceInputs["pkcs7"] = state ? state.pkcs7 : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["renewable"] = state ? state.renewable : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
        } else {
            const args = argsOrState as AuthBackendLoginArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["iamHttpRequestMethod"] = args ? args.iamHttpRequestMethod : undefined;
            resourceInputs["iamRequestBody"] = args ? args.iamRequestBody : undefined;
            resourceInputs["iamRequestHeaders"] = args ? args.iamRequestHeaders : undefined;
            resourceInputs["iamRequestUrl"] = args ? args.iamRequestUrl : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["nonce"] = args ? args.nonce : undefined;
            resourceInputs["pkcs7"] = args ? args.pkcs7 : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
            resourceInputs["authType"] = undefined /*out*/;
            resourceInputs["clientToken"] = undefined /*out*/;
            resourceInputs["leaseDuration"] = undefined /*out*/;
            resourceInputs["leaseStartTime"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["policies"] = undefined /*out*/;
            resourceInputs["renewable"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackendLogin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendLogin resources.
 */
export interface AuthBackendLoginState {
    /**
     * The token's accessor.
     */
    accessor?: pulumi.Input<string>;
    /**
     * The authentication type used to generate this token.
     */
    authType?: pulumi.Input<string>;
    /**
     * The unique name of the AWS auth backend. Defaults to
     * 'aws'.
     */
    backend?: pulumi.Input<string>;
    /**
     * The token returned by Vault.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The HTTP method used in the signed IAM
     * request.
     */
    iamHttpRequestMethod?: pulumi.Input<string>;
    /**
     * The base64-encoded body of the signed
     * request.
     */
    iamRequestBody?: pulumi.Input<string>;
    /**
     * The base64-encoded, JSON serialized
     * representation of the GetCallerIdentity HTTP request headers.
     */
    iamRequestHeaders?: pulumi.Input<string>;
    /**
     * The base64-encoded HTTP URL used in the signed
     * request.
     */
    iamRequestUrl?: pulumi.Input<string>;
    /**
     * The base64-encoded EC2 instance identity document to
     * authenticate with. Can be retrieved from the EC2 metadata server.
     */
    identity?: pulumi.Input<string>;
    /**
     * The duration in seconds the token will be valid, relative
     * to the time in `leaseStartTime`.
     */
    leaseDuration?: pulumi.Input<number>;
    leaseStartTime?: pulumi.Input<string>;
    /**
     * A map of information returned by the Vault server about the
     * authentication used to generate this token.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique nonce to be used for login requests. Can be
     * set to a user-specified value, or will contain the server-generated value
     * once a token is issued. EC2 instances can only acquire a single token until
     * the whitelist is tidied again unless they keep track of this nonce.
     */
    nonce?: pulumi.Input<string>;
    /**
     * The PKCS#7 signature of the identity document to
     * authenticate with, with all newline characters removed. Can be retrieved from
     * the EC2 metadata server.
     */
    pkcs7?: pulumi.Input<string>;
    /**
     * The Vault policies assigned to this token.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to true if the token can be extended through renewal.
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * The name of the AWS auth backend role to create tokens
     * against.
     */
    role?: pulumi.Input<string>;
    /**
     * The base64-encoded SHA256 RSA signature of the
     * instance identity document to authenticate with, with all newline characters
     * removed. Can be retrieved from the EC2 metadata server.
     */
    signature?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendLogin resource.
 */
export interface AuthBackendLoginArgs {
    /**
     * The unique name of the AWS auth backend. Defaults to
     * 'aws'.
     */
    backend?: pulumi.Input<string>;
    /**
     * The HTTP method used in the signed IAM
     * request.
     */
    iamHttpRequestMethod?: pulumi.Input<string>;
    /**
     * The base64-encoded body of the signed
     * request.
     */
    iamRequestBody?: pulumi.Input<string>;
    /**
     * The base64-encoded, JSON serialized
     * representation of the GetCallerIdentity HTTP request headers.
     */
    iamRequestHeaders?: pulumi.Input<string>;
    /**
     * The base64-encoded HTTP URL used in the signed
     * request.
     */
    iamRequestUrl?: pulumi.Input<string>;
    /**
     * The base64-encoded EC2 instance identity document to
     * authenticate with. Can be retrieved from the EC2 metadata server.
     */
    identity?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique nonce to be used for login requests. Can be
     * set to a user-specified value, or will contain the server-generated value
     * once a token is issued. EC2 instances can only acquire a single token until
     * the whitelist is tidied again unless they keep track of this nonce.
     */
    nonce?: pulumi.Input<string>;
    /**
     * The PKCS#7 signature of the identity document to
     * authenticate with, with all newline characters removed. Can be retrieved from
     * the EC2 metadata server.
     */
    pkcs7?: pulumi.Input<string>;
    /**
     * The name of the AWS auth backend role to create tokens
     * against.
     */
    role?: pulumi.Input<string>;
    /**
     * The base64-encoded SHA256 RSA signature of the
     * instance identity document to authenticate with, with all newline characters
     * removed. Can be retrieved from the EC2 metadata server.
     */
    signature?: pulumi.Input<string>;
}
