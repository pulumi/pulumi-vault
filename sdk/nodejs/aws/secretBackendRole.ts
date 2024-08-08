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
 * const aws = new vault.aws.SecretBackend("aws", {
 *     accessKey: "AKIA.....",
 *     secretKey: "AWS secret key",
 * });
 * const role = new vault.aws.SecretBackendRole("role", {
 *     backend: aws.path,
 *     name: "deploy",
 *     credentialType: "iam_user",
 *     policyDocument: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Action": "iam:*",
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * AWS secret backend roles can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:aws/secretBackendRole:SecretBackendRole role aws/roles/deploy
 * ```
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    public readonly credentialType!: pulumi.Output<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    public readonly defaultStsTtl!: pulumi.Output<number>;
    /**
     * External ID to set for assume role creds. 
     * Valid only when `credentialType` is set to `assumedRole`.
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * A list of IAM group names. IAM users generated
     * against this vault role will be added to these IAM Groups. For a credential
     * type of `assumedRole` or `federationToken`, the policies sent to the
     * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
     * policies from each group in `iamGroups` combined with the `policyDocument`
     * and `policyArns` parameters.
     */
    public readonly iamGroups!: pulumi.Output<string[] | undefined>;
    /**
     * A map of strings representing key/value pairs
     * to be used as tags for any IAM user that is created by this role.
     */
    public readonly iamTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    public readonly maxStsTtl!: pulumi.Output<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
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
     * The ARN of the AWS Permissions 
     * Boundary to attach to IAM users created in the role. Valid only when
     * `credentialType` is `iamUser`. If not specified, then no permissions boundary
     * policy will be attached.
     */
    public readonly permissionsBoundaryArn!: pulumi.Output<string | undefined>;
    /**
     * Specifies a list of AWS managed policy ARNs. The
     * behavior depends on the credential type. With `iamUser`, the policies will be
     * attached to IAM users when they are requested. With `assumedRole` and
     * `federationToken`, the policy ARNs will act as a filter on what the credentials
     * can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
     * `federationToken`, at least one of `policyDocument` or `policyArns` must
     * be specified.
     */
    public readonly policyArns!: pulumi.Output<string[] | undefined>;
    /**
     * The IAM policy document for the role. The
     * behavior depends on the credential type. With `iamUser`, the policy document
     * will be attached to the IAM user generated and augment the permissions the IAM
     * user has. With `assumedRole` and `federationToken`, the policy document will
     * act as a filter on what the credentials can do, similar to `policyArns`.
     */
    public readonly policyDocument!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    public readonly roleArns!: pulumi.Output<string[] | undefined>;
    /**
     * A map of strings representing key/value pairs to be set
     * during assume role creds creation. Valid only when `credentialType` is set to
     * `assumedRole`.
     */
    public readonly sessionTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The path for the user name. Valid only when 
     * `credentialType` is `iamUser`. Default is `/`.
     */
    public readonly userPath!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["credentialType"] = state ? state.credentialType : undefined;
            resourceInputs["defaultStsTtl"] = state ? state.defaultStsTtl : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["iamGroups"] = state ? state.iamGroups : undefined;
            resourceInputs["iamTags"] = state ? state.iamTags : undefined;
            resourceInputs["maxStsTtl"] = state ? state.maxStsTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["permissionsBoundaryArn"] = state ? state.permissionsBoundaryArn : undefined;
            resourceInputs["policyArns"] = state ? state.policyArns : undefined;
            resourceInputs["policyDocument"] = state ? state.policyDocument : undefined;
            resourceInputs["roleArns"] = state ? state.roleArns : undefined;
            resourceInputs["sessionTags"] = state ? state.sessionTags : undefined;
            resourceInputs["userPath"] = state ? state.userPath : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.credentialType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'credentialType'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["credentialType"] = args ? args.credentialType : undefined;
            resourceInputs["defaultStsTtl"] = args ? args.defaultStsTtl : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["iamGroups"] = args ? args.iamGroups : undefined;
            resourceInputs["iamTags"] = args ? args.iamTags : undefined;
            resourceInputs["maxStsTtl"] = args ? args.maxStsTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["permissionsBoundaryArn"] = args ? args.permissionsBoundaryArn : undefined;
            resourceInputs["policyArns"] = args ? args.policyArns : undefined;
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["roleArns"] = args ? args.roleArns : undefined;
            resourceInputs["sessionTags"] = args ? args.sessionTags : undefined;
            resourceInputs["userPath"] = args ? args.userPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    credentialType?: pulumi.Input<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    defaultStsTtl?: pulumi.Input<number>;
    /**
     * External ID to set for assume role creds. 
     * Valid only when `credentialType` is set to `assumedRole`.
     */
    externalId?: pulumi.Input<string>;
    /**
     * A list of IAM group names. IAM users generated
     * against this vault role will be added to these IAM Groups. For a credential
     * type of `assumedRole` or `federationToken`, the policies sent to the
     * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
     * policies from each group in `iamGroups` combined with the `policyDocument`
     * and `policyArns` parameters.
     */
    iamGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of strings representing key/value pairs
     * to be used as tags for any IAM user that is created by this role.
     */
    iamTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    maxStsTtl?: pulumi.Input<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
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
     * The ARN of the AWS Permissions 
     * Boundary to attach to IAM users created in the role. Valid only when
     * `credentialType` is `iamUser`. If not specified, then no permissions boundary
     * policy will be attached.
     */
    permissionsBoundaryArn?: pulumi.Input<string>;
    /**
     * Specifies a list of AWS managed policy ARNs. The
     * behavior depends on the credential type. With `iamUser`, the policies will be
     * attached to IAM users when they are requested. With `assumedRole` and
     * `federationToken`, the policy ARNs will act as a filter on what the credentials
     * can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
     * `federationToken`, at least one of `policyDocument` or `policyArns` must
     * be specified.
     */
    policyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IAM policy document for the role. The
     * behavior depends on the credential type. With `iamUser`, the policy document
     * will be attached to the IAM user generated and augment the permissions the IAM
     * user has. With `assumedRole` and `federationToken`, the policy document will
     * act as a filter on what the credentials can do, similar to `policyArns`.
     */
    policyDocument?: pulumi.Input<string>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    roleArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of strings representing key/value pairs to be set
     * during assume role creds creation. Valid only when `credentialType` is set to
     * `assumedRole`.
     */
    sessionTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path for the user name. Valid only when 
     * `credentialType` is `iamUser`. Default is `/`.
     */
    userPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    credentialType: pulumi.Input<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    defaultStsTtl?: pulumi.Input<number>;
    /**
     * External ID to set for assume role creds. 
     * Valid only when `credentialType` is set to `assumedRole`.
     */
    externalId?: pulumi.Input<string>;
    /**
     * A list of IAM group names. IAM users generated
     * against this vault role will be added to these IAM Groups. For a credential
     * type of `assumedRole` or `federationToken`, the policies sent to the
     * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
     * policies from each group in `iamGroups` combined with the `policyDocument`
     * and `policyArns` parameters.
     */
    iamGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of strings representing key/value pairs
     * to be used as tags for any IAM user that is created by this role.
     */
    iamTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    maxStsTtl?: pulumi.Input<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
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
     * The ARN of the AWS Permissions 
     * Boundary to attach to IAM users created in the role. Valid only when
     * `credentialType` is `iamUser`. If not specified, then no permissions boundary
     * policy will be attached.
     */
    permissionsBoundaryArn?: pulumi.Input<string>;
    /**
     * Specifies a list of AWS managed policy ARNs. The
     * behavior depends on the credential type. With `iamUser`, the policies will be
     * attached to IAM users when they are requested. With `assumedRole` and
     * `federationToken`, the policy ARNs will act as a filter on what the credentials
     * can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
     * `federationToken`, at least one of `policyDocument` or `policyArns` must
     * be specified.
     */
    policyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IAM policy document for the role. The
     * behavior depends on the credential type. With `iamUser`, the policy document
     * will be attached to the IAM user generated and augment the permissions the IAM
     * user has. With `assumedRole` and `federationToken`, the policy document will
     * act as a filter on what the credentials can do, similar to `policyArns`.
     */
    policyDocument?: pulumi.Input<string>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    roleArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of strings representing key/value pairs to be set
     * during assume role creds creation. Valid only when `credentialType` is set to
     * `assumedRole`.
     */
    sessionTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path for the user name. Valid only when 
     * `credentialType` is `iamUser`. Default is `/`.
     */
    userPath?: pulumi.Input<string>;
}
