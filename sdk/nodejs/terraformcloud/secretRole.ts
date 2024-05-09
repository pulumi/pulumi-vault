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
 * ```
 *
 * ## Import
 *
 * Terraform Cloud secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
 *
 * ```sh
 * $ pulumi import vault:terraformcloud/secretRole:SecretRole example terraform/roles/my-role
 * ```
 */
export class SecretRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretRoleState, opts?: pulumi.CustomResourceOptions): SecretRole {
        return new SecretRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:terraformcloud/secretRole:SecretRole';

    /**
     * Returns true if the given object is an instance of SecretRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretRole.__pulumiType;
    }

    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    public readonly organization!: pulumi.Output<string | undefined>;
    public readonly teamId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the TTL for this role.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretRoleArgs | SecretRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretRoleState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as SecretRoleArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretRole resources.
 */
export interface SecretRoleState {
    backend?: pulumi.Input<string>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
    /**
     * Specifies the TTL for this role.
     */
    ttl?: pulumi.Input<number>;
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretRole resource.
 */
export interface SecretRoleArgs {
    backend?: pulumi.Input<string>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
    /**
     * Specifies the TTL for this role.
     */
    ttl?: pulumi.Input<number>;
    userId?: pulumi.Input<string>;
}
