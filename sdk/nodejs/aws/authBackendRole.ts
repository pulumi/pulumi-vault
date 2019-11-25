// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS auth backend role in a Vault server. Roles constrain the
 * instances or principals that can perform the login operation against the
 * backend. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
 * information.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_role.html.markdown.
 */
export class AuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendRole {
        return new AuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendRole:AuthBackendRole';

    /**
     * Returns true if the given object is an instance of AuthBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRole.__pulumiType;
    }

    /**
     * If set to `true`, allows migration of
     * the underlying instance where the client resides.
     */
    public readonly allowInstanceMigration!: pulumi.Output<boolean | undefined>;
    /**
     * The auth type permitted for this role. Valid choices
     * are `ec2` and `iam`. Defaults to `iam`.
     */
    public readonly authType!: pulumi.Output<string | undefined>;
    /**
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they should be using the
     * account ID specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    public readonly boundAccountIds!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they should be using the AMI ID
     * specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    public readonly boundAmiIds!: pulumi.Output<string[] | undefined>;
    /**
     * Only EC2 instances that match this instance ID will be permitted to log in.
     */
    public readonly boundEc2InstanceIds!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on
     * the EC2 instances that can perform the login operation that they must be
     * associated with an IAM instance profile ARN which has a prefix that matches
     * the value specified by this field. The value is prefix-matched as though it
     * were a glob ending in `*`. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    public readonly boundIamInstanceProfileArns!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines the IAM principal that
     * must be authenticated when `authType` is set to `iam`. Wildcards are
     * supported at the end of the ARN.
     */
    public readonly boundIamPrincipalArns!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they must match the IAM
     * role ARN specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    public readonly boundIamRoleArns!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that the region in their identity
     * document must match the one specified by this field. `authType` must be set
     * to `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    public readonly boundRegions!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they be associated with
     * the subnet ID that matches the value specified by this field. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    public readonly boundSubnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they be associated with the VPC ID
     * that matches the value specified by this field. `authType` must be set to
     * `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    public readonly boundVpcIds!: pulumi.Output<string[] | undefined>;
    /**
     * IF set to `true`, only allows a
     * single token to be granted per instance ID. This can only be set when
     * `authType` is set to `ec2`.
     */
    public readonly disallowReauthentication!: pulumi.Output<boolean | undefined>;
    /**
     * When `inferredEntityType` is set, this
     * is the region to search for the inferred entities. Required if
     * `inferredEntityType` is set. This only applies when `authType` is set to
     * `iam`.
     */
    public readonly inferredAwsRegion!: pulumi.Output<string | undefined>;
    /**
     * If set, instructs Vault to turn on
     * inferencing. The only valid value is `ec2Instance`, which instructs Vault to
     * infer that the role comes from an EC2 instance in an IAM instance profile.
     * This only applies when `authType` is set to `iam`.
     */
    public readonly inferredEntityType!: pulumi.Output<string | undefined>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * If set to `true`, the
     * `boundIamPrincipalArns` are resolved to [AWS Unique
     * IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids)
     * for the bound principal ARN. This field is ignored when a
     * `boundIamPrincipalArn` ends in a wildcard. Resolving to unique IDs more
     * closely mimics the behavior of AWS services in that if an IAM user or role is
     * deleted and a new one is recreated with the same name, those new users or
     * roles won't get access to roles in Vault that were permissioned to the prior
     * principals of the same name. Defaults to `true`.
     * Once set to `true`, this cannot be changed to `false` without recreating the role.
     */
    public readonly resolveAwsUniqueIds!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * If set, enable role tags for this role. The value set
     * for this field should be the key of the tag on the EC2 instance. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    public readonly roleTag!: pulumi.Output<string | undefined>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;

    /**
     * Create a AuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleArgs | AuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            inputs["allowInstanceMigration"] = state ? state.allowInstanceMigration : undefined;
            inputs["authType"] = state ? state.authType : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["boundAccountIds"] = state ? state.boundAccountIds : undefined;
            inputs["boundAmiIds"] = state ? state.boundAmiIds : undefined;
            inputs["boundEc2InstanceIds"] = state ? state.boundEc2InstanceIds : undefined;
            inputs["boundIamInstanceProfileArns"] = state ? state.boundIamInstanceProfileArns : undefined;
            inputs["boundIamPrincipalArns"] = state ? state.boundIamPrincipalArns : undefined;
            inputs["boundIamRoleArns"] = state ? state.boundIamRoleArns : undefined;
            inputs["boundRegions"] = state ? state.boundRegions : undefined;
            inputs["boundSubnetIds"] = state ? state.boundSubnetIds : undefined;
            inputs["boundVpcIds"] = state ? state.boundVpcIds : undefined;
            inputs["disallowReauthentication"] = state ? state.disallowReauthentication : undefined;
            inputs["inferredAwsRegion"] = state ? state.inferredAwsRegion : undefined;
            inputs["inferredEntityType"] = state ? state.inferredEntityType : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["resolveAwsUniqueIds"] = state ? state.resolveAwsUniqueIds : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["roleTag"] = state ? state.roleTag : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["allowInstanceMigration"] = args ? args.allowInstanceMigration : undefined;
            inputs["authType"] = args ? args.authType : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["boundAccountIds"] = args ? args.boundAccountIds : undefined;
            inputs["boundAmiIds"] = args ? args.boundAmiIds : undefined;
            inputs["boundEc2InstanceIds"] = args ? args.boundEc2InstanceIds : undefined;
            inputs["boundIamInstanceProfileArns"] = args ? args.boundIamInstanceProfileArns : undefined;
            inputs["boundIamPrincipalArns"] = args ? args.boundIamPrincipalArns : undefined;
            inputs["boundIamRoleArns"] = args ? args.boundIamRoleArns : undefined;
            inputs["boundRegions"] = args ? args.boundRegions : undefined;
            inputs["boundSubnetIds"] = args ? args.boundSubnetIds : undefined;
            inputs["boundVpcIds"] = args ? args.boundVpcIds : undefined;
            inputs["disallowReauthentication"] = args ? args.disallowReauthentication : undefined;
            inputs["inferredAwsRegion"] = args ? args.inferredAwsRegion : undefined;
            inputs["inferredEntityType"] = args ? args.inferredEntityType : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["resolveAwsUniqueIds"] = args ? args.resolveAwsUniqueIds : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["roleTag"] = args ? args.roleTag : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    /**
     * If set to `true`, allows migration of
     * the underlying instance where the client resides.
     */
    readonly allowInstanceMigration?: pulumi.Input<boolean>;
    /**
     * The auth type permitted for this role. Valid choices
     * are `ec2` and `iam`. Defaults to `iam`.
     */
    readonly authType?: pulumi.Input<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they should be using the
     * account ID specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they should be using the AMI ID
     * specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundAmiIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Only EC2 instances that match this instance ID will be permitted to log in.
     */
    readonly boundEc2InstanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on
     * the EC2 instances that can perform the login operation that they must be
     * associated with an IAM instance profile ARN which has a prefix that matches
     * the value specified by this field. The value is prefix-matched as though it
     * were a glob ending in `*`. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundIamInstanceProfileArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines the IAM principal that
     * must be authenticated when `authType` is set to `iam`. Wildcards are
     * supported at the end of the ARN.
     */
    readonly boundIamPrincipalArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they must match the IAM
     * role ARN specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundIamRoleArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that the region in their identity
     * document must match the one specified by this field. `authType` must be set
     * to `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    readonly boundRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they be associated with
     * the subnet ID that matches the value specified by this field. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    readonly boundSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they be associated with the VPC ID
     * that matches the value specified by this field. `authType` must be set to
     * `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    readonly boundVpcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IF set to `true`, only allows a
     * single token to be granted per instance ID. This can only be set when
     * `authType` is set to `ec2`.
     */
    readonly disallowReauthentication?: pulumi.Input<boolean>;
    /**
     * When `inferredEntityType` is set, this
     * is the region to search for the inferred entities. Required if
     * `inferredEntityType` is set. This only applies when `authType` is set to
     * `iam`.
     */
    readonly inferredAwsRegion?: pulumi.Input<string>;
    /**
     * If set, instructs Vault to turn on
     * inferencing. The only valid value is `ec2Instance`, which instructs Vault to
     * infer that the role comes from an EC2 instance in an IAM instance profile.
     * This only applies when `authType` is set to `iam`.
     */
    readonly inferredEntityType?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, the
     * `boundIamPrincipalArns` are resolved to [AWS Unique
     * IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids)
     * for the bound principal ARN. This field is ignored when a
     * `boundIamPrincipalArn` ends in a wildcard. Resolving to unique IDs more
     * closely mimics the behavior of AWS services in that if an IAM user or role is
     * deleted and a new one is recreated with the same name, those new users or
     * roles won't get access to roles in Vault that were permissioned to the prior
     * principals of the same name. Defaults to `true`.
     * Once set to `true`, this cannot be changed to `false` without recreating the role.
     */
    readonly resolveAwsUniqueIds?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * If set, enable role tags for this role. The value set
     * for this field should be the key of the tag on the EC2 instance. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    readonly roleTag?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     */
    readonly ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * If set to `true`, allows migration of
     * the underlying instance where the client resides.
     */
    readonly allowInstanceMigration?: pulumi.Input<boolean>;
    /**
     * The auth type permitted for this role. Valid choices
     * are `ec2` and `iam`. Defaults to `iam`.
     */
    readonly authType?: pulumi.Input<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they should be using the
     * account ID specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they should be using the AMI ID
     * specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundAmiIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Only EC2 instances that match this instance ID will be permitted to log in.
     */
    readonly boundEc2InstanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on
     * the EC2 instances that can perform the login operation that they must be
     * associated with an IAM instance profile ARN which has a prefix that matches
     * the value specified by this field. The value is prefix-matched as though it
     * were a glob ending in `*`. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundIamInstanceProfileArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines the IAM principal that
     * must be authenticated when `authType` is set to `iam`. Wildcards are
     * supported at the end of the ARN.
     */
    readonly boundIamPrincipalArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they must match the IAM
     * role ARN specified by this field. `authType` must be set to `ec2` or
     * `inferredEntityType` must be set to `ec2Instance` to use this constraint.
     */
    readonly boundIamRoleArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that the region in their identity
     * document must match the one specified by this field. `authType` must be set
     * to `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    readonly boundRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they be associated with
     * the subnet ID that matches the value specified by this field. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    readonly boundSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they be associated with the VPC ID
     * that matches the value specified by this field. `authType` must be set to
     * `ec2` or `inferredEntityType` must be set to `ec2Instance` to use this
     * constraint.
     */
    readonly boundVpcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IF set to `true`, only allows a
     * single token to be granted per instance ID. This can only be set when
     * `authType` is set to `ec2`.
     */
    readonly disallowReauthentication?: pulumi.Input<boolean>;
    /**
     * When `inferredEntityType` is set, this
     * is the region to search for the inferred entities. Required if
     * `inferredEntityType` is set. This only applies when `authType` is set to
     * `iam`.
     */
    readonly inferredAwsRegion?: pulumi.Input<string>;
    /**
     * If set, instructs Vault to turn on
     * inferencing. The only valid value is `ec2Instance`, which instructs Vault to
     * infer that the role comes from an EC2 instance in an IAM instance profile.
     * This only applies when `authType` is set to `iam`.
     */
    readonly inferredEntityType?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, the
     * `boundIamPrincipalArns` are resolved to [AWS Unique
     * IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids)
     * for the bound principal ARN. This field is ignored when a
     * `boundIamPrincipalArn` ends in a wildcard. Resolving to unique IDs more
     * closely mimics the behavior of AWS services in that if an IAM user or role is
     * deleted and a new one is recreated with the same name, those new users or
     * roles won't get access to roles in Vault that were permissioned to the prior
     * principals of the same name. Defaults to `true`.
     * Once set to `true`, this cannot be changed to `false` without recreating the role.
     */
    readonly resolveAwsUniqueIds?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    readonly role: pulumi.Input<string>;
    /**
     * If set, enable role tags for this role. The value set
     * for this field should be the key of the tag on the EC2 instance. `authType`
     * must be set to `ec2` or `inferredEntityType` must be set to `ec2Instance`
     * to use this constraint.
     */
    readonly roleTag?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     */
    readonly ttl?: pulumi.Input<number>;
}
