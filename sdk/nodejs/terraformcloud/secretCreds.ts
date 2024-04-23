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
 * const test = new vault.terraformcloud.SecretBackend("test", {
 *     backend: "terraform",
 *     description: "Manages the Terraform Cloud backend",
 *     token: "V0idfhi2iksSDU234ucdbi2nidsi...",
 * });
 * const example = new vault.terraformcloud.SecretRole("example", {
 *     backend: test.backend,
 *     name: "test-role",
 *     organization: "example-organization-name",
 *     teamId: "team-ieF4isC...",
 * });
 * const token = new vault.terraformcloud.SecretCreds("token", {
 *     backend: test.backend,
 *     role: example.name,
 * });
 * ```
 */
export class SecretCreds extends pulumi.CustomResource {
    /**
     * Get an existing SecretCreds resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretCredsState, opts?: pulumi.CustomResourceOptions): SecretCreds {
        return new SecretCreds(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:terraformcloud/secretCreds:SecretCreds';

    /**
     * Returns true if the given object is an instance of SecretCreds.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretCreds {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretCreds.__pulumiType;
    }

    /**
     * Terraform Cloud secret backend to generate tokens from
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The lease associated with the token. Only user tokens will have a 
     * Vault lease associated with them.
     */
    public /*out*/ readonly leaseId!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The organization associated with the token provided.
     */
    public /*out*/ readonly organization!: pulumi.Output<string>;
    /**
     * Name of the role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The team id associated with the token provided.
     */
    public /*out*/ readonly teamId!: pulumi.Output<string>;
    /**
     * The actual token that was generated and can be used with API calls
     * to identify the user of the call.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The public identifier for a specific token. It can be used 
     * to look up information about a token or to revoke a token.
     */
    public /*out*/ readonly tokenId!: pulumi.Output<string>;

    /**
     * Create a SecretCreds resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretCredsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretCredsArgs | SecretCredsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretCredsState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["leaseId"] = state ? state.leaseId : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["tokenId"] = state ? state.tokenId : undefined;
        } else {
            const args = argsOrState as SecretCredsArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["leaseId"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["teamId"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["tokenId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["leaseId", "token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretCreds.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretCreds resources.
 */
export interface SecretCredsState {
    /**
     * Terraform Cloud secret backend to generate tokens from
     */
    backend?: pulumi.Input<string>;
    /**
     * The lease associated with the token. Only user tokens will have a 
     * Vault lease associated with them.
     */
    leaseId?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The organization associated with the token provided.
     */
    organization?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    role?: pulumi.Input<string>;
    /**
     * The team id associated with the token provided.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The actual token that was generated and can be used with API calls
     * to identify the user of the call.
     */
    token?: pulumi.Input<string>;
    /**
     * The public identifier for a specific token. It can be used 
     * to look up information about a token or to revoke a token.
     */
    tokenId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretCreds resource.
 */
export interface SecretCredsArgs {
    /**
     * Terraform Cloud secret backend to generate tokens from
     */
    backend: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    role: pulumi.Input<string>;
}
