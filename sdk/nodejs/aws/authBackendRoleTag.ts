// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Reads role tag information from an AWS auth backend in Vault.
 */
export class AuthBackendRoleTag extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRoleTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleTagState, opts?: pulumi.CustomResourceOptions): AuthBackendRoleTag {
        return new AuthBackendRoleTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendRoleTag:AuthBackendRoleTag';

    /**
     * Returns true if the given object is an instance of AuthBackendRoleTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRoleTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRoleTag.__pulumiType;
    }

    /**
     * If set, allows migration of the underlying instances where the client resides. Use with caution.
     */
    public readonly allowInstanceMigration!: pulumi.Output<boolean | undefined>;
    /**
     * The path to the AWS auth backend to
     * read role tags from, with no leading or trailing `/`s. Defaults to "aws".
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set, only allows a single token to be granted per instance ID.
     */
    public readonly disallowReauthentication!: pulumi.Output<boolean | undefined>;
    /**
     * Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * The maximum TTL of the tokens issued using this role.
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The policies to be associated with the tag. Must be a subset of the policies associated with the role.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the AWS auth backend role to read
     * role tags from, with no leading or trailing `/`s.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The key of the role tag.
     */
    public /*out*/ readonly tagKey!: pulumi.Output<string>;
    /**
     * The value to set the role key.
     */
    public /*out*/ readonly tagValue!: pulumi.Output<string>;

    /**
     * Create a AuthBackendRoleTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleTagArgs | AuthBackendRoleTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendRoleTagState | undefined;
            resourceInputs["allowInstanceMigration"] = state ? state.allowInstanceMigration : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["disallowReauthentication"] = state ? state.disallowReauthentication : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["tagKey"] = state ? state.tagKey : undefined;
            resourceInputs["tagValue"] = state ? state.tagValue : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleTagArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["allowInstanceMigration"] = args ? args.allowInstanceMigration : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["disallowReauthentication"] = args ? args.disallowReauthentication : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["tagKey"] = undefined /*out*/;
            resourceInputs["tagValue"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendRoleTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRoleTag resources.
 */
export interface AuthBackendRoleTagState {
    /**
     * If set, allows migration of the underlying instances where the client resides. Use with caution.
     */
    allowInstanceMigration?: pulumi.Input<boolean>;
    /**
     * The path to the AWS auth backend to
     * read role tags from, with no leading or trailing `/`s. Defaults to "aws".
     */
    backend?: pulumi.Input<string>;
    /**
     * If set, only allows a single token to be granted per instance ID.
     */
    disallowReauthentication?: pulumi.Input<boolean>;
    /**
     * Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The maximum TTL of the tokens issued using this role.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The policies to be associated with the tag. Must be a subset of the policies associated with the role.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the AWS auth backend role to read
     * role tags from, with no leading or trailing `/`s.
     */
    role?: pulumi.Input<string>;
    /**
     * The key of the role tag.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * The value to set the role key.
     */
    tagValue?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRoleTag resource.
 */
export interface AuthBackendRoleTagArgs {
    /**
     * If set, allows migration of the underlying instances where the client resides. Use with caution.
     */
    allowInstanceMigration?: pulumi.Input<boolean>;
    /**
     * The path to the AWS auth backend to
     * read role tags from, with no leading or trailing `/`s. Defaults to "aws".
     */
    backend?: pulumi.Input<string>;
    /**
     * If set, only allows a single token to be granted per instance ID.
     */
    disallowReauthentication?: pulumi.Input<boolean>;
    /**
     * Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The maximum TTL of the tokens issued using this role.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The policies to be associated with the tag. Must be a subset of the policies associated with the role.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the AWS auth backend role to read
     * role tags from, with no leading or trailing `/`s.
     */
    role: pulumi.Input<string>;
}
