// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.NomadSecretBackend("config", {
 *     backend: "nomad",
 *     description: "test description",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 7200,
 *     address: "https://127.0.0.1:4646",
 *     token: "ae20ceaa-...",
 * });
 * const test = new vault.NomadSecretRole("test", {
 *     backend: config.backend,
 *     role: "test",
 *     type: "client",
 *     policies: ["readonly"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Nomad secret role can be imported using the `backend`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
 * ```
 */
export class NomadSecretRole extends pulumi.CustomResource {
    /**
     * Get an existing NomadSecretRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NomadSecretRoleState, opts?: pulumi.CustomResourceOptions): NomadSecretRole {
        return new NomadSecretRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/nomadSecretRole:NomadSecretRole';

    /**
     * Returns true if the given object is an instance of NomadSecretRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NomadSecretRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NomadSecretRole.__pulumiType;
    }

    /**
     * The unique path this backend should be mounted at.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies if the generated token should be global. Defaults to 
     * false.
     */
    public readonly global!: pulumi.Output<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * List of policies attached to the generated token. This setting is only used 
     * when `type` is 'client'.
     */
    public readonly policies!: pulumi.Output<string[]>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Specifies the type of token to create when using this role. Valid 
     * settings are 'client' and 'management'. Defaults to 'client'.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a NomadSecretRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NomadSecretRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NomadSecretRoleArgs | NomadSecretRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NomadSecretRoleState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["global"] = state ? state.global : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NomadSecretRoleArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["global"] = args ? args.global : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NomadSecretRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NomadSecretRole resources.
 */
export interface NomadSecretRoleState {
    /**
     * The unique path this backend should be mounted at.
     */
    backend?: pulumi.Input<string>;
    /**
     * Specifies if the generated token should be global. Defaults to 
     * false.
     */
    global?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies attached to the generated token. This setting is only used 
     * when `type` is 'client'.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    role?: pulumi.Input<string>;
    /**
     * Specifies the type of token to create when using this role. Valid 
     * settings are 'client' and 'management'. Defaults to 'client'.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NomadSecretRole resource.
 */
export interface NomadSecretRoleArgs {
    /**
     * The unique path this backend should be mounted at.
     */
    backend: pulumi.Input<string>;
    /**
     * Specifies if the generated token should be global. Defaults to 
     * false.
     */
    global?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies attached to the generated token. This setting is only used 
     * when `type` is 'client'.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    role: pulumi.Input<string>;
    /**
     * Specifies the type of token to create when using this role. Valid 
     * settings are 'client' and 'management'. Defaults to 'client'.
     */
    type?: pulumi.Input<string>;
}
