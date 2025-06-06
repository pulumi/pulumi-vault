// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages policy mappings for Github Users authenticated via Github. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/github/) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.github.AuthBackend("example", {organization: "myorg"});
 * const tfUser = new vault.github.User("tf_user", {
 *     backend: example.id,
 *     user: "john.doe",
 *     policies: [
 *         "developer",
 *         "read-only",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Github user mappings can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:github/user:User tf_user auth/github/map/users/john.doe
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:github/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * GitHub user name.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    backend?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GitHub user name.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    backend?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GitHub user name.
     */
    user: pulumi.Input<string>;
}
