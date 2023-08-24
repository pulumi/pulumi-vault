// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS auth backend identity configuration in a Vault server. This configuration defines how Vault interacts
 * with the identity store. See the [Vault documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const aws = new vault.AuthBackend("aws", {type: "aws"});
 * const example = new vault.aws.AuthBackendConfigIdentity("example", {
 *     backend: aws.path,
 *     iamAlias: "full_arn",
 *     iamMetadatas: [
 *         "canonical_arn",
 *         "account_id",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * AWS auth backend identity config can be imported using `auth/`, the `backend` path, and `/config/identity` e.g.
 *
 * ```sh
 *  $ pulumi import vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity example auth/aws/config/identity
 * ```
 */
export class AuthBackendConfigIdentity extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendConfigIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendConfigIdentityState, opts?: pulumi.CustomResourceOptions): AuthBackendConfigIdentity {
        return new AuthBackendConfigIdentity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity';

    /**
     * Returns true if the given object is an instance of AuthBackendConfigIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendConfigIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendConfigIdentity.__pulumiType;
    }

    /**
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * How to generate the identity alias when using the ec2 auth method. Valid choices are
     * `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
     */
    public readonly ec2Alias!: pulumi.Output<string | undefined>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `ec2Alias`
     */
    public readonly ec2Metadatas!: pulumi.Output<string[] | undefined>;
    /**
     * How to generate the identity alias when using the iam auth method. Valid choices are
     * `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
     */
    public readonly iamAlias!: pulumi.Output<string | undefined>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `iamAlias`
     */
    public readonly iamMetadatas!: pulumi.Output<string[] | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendConfigIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendConfigIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendConfigIdentityArgs | AuthBackendConfigIdentityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendConfigIdentityState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["ec2Alias"] = state ? state.ec2Alias : undefined;
            resourceInputs["ec2Metadatas"] = state ? state.ec2Metadatas : undefined;
            resourceInputs["iamAlias"] = state ? state.iamAlias : undefined;
            resourceInputs["iamMetadatas"] = state ? state.iamMetadatas : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as AuthBackendConfigIdentityArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["ec2Alias"] = args ? args.ec2Alias : undefined;
            resourceInputs["ec2Metadatas"] = args ? args.ec2Metadatas : undefined;
            resourceInputs["iamAlias"] = args ? args.iamAlias : undefined;
            resourceInputs["iamMetadatas"] = args ? args.iamMetadatas : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendConfigIdentity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendConfigIdentity resources.
 */
export interface AuthBackendConfigIdentityState {
    /**
     * Unique name of the auth backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * How to generate the identity alias when using the ec2 auth method. Valid choices are
     * `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
     */
    ec2Alias?: pulumi.Input<string>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `ec2Alias`
     */
    ec2Metadatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How to generate the identity alias when using the iam auth method. Valid choices are
     * `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
     */
    iamAlias?: pulumi.Input<string>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `iamAlias`
     */
    iamMetadatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendConfigIdentity resource.
 */
export interface AuthBackendConfigIdentityArgs {
    /**
     * Unique name of the auth backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * How to generate the identity alias when using the ec2 auth method. Valid choices are
     * `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
     */
    ec2Alias?: pulumi.Input<string>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `ec2Alias`
     */
    ec2Metadatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How to generate the identity alias when using the iam auth method. Valid choices are
     * `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
     */
    iamAlias?: pulumi.Input<string>;
    /**
     * The metadata to include on the token returned by the `login` endpoint. This metadata will be
     * added to both audit logs, and on the `iamAlias`
     */
    iamMetadatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
