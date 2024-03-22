// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages member entities for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 *
 * ## Example Usage
 *
 * ### Exclusive Member Entities
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const internal = new vault.identity.Group("internal", {
 *     type: "internal",
 *     externalMemberEntityIds: true,
 *     metadata: {
 *         version: "2",
 *     },
 * });
 * const user = new vault.identity.Entity("user", {});
 * const members = new vault.identity.GroupMemberEntityIds("members", {
 *     exclusive: true,
 *     memberEntityIds: [user.id],
 *     groupId: internal.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Non-exclusive Member Entities
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const internal = new vault.identity.Group("internal", {
 *     type: "internal",
 *     externalMemberEntityIds: true,
 *     metadata: {
 *         version: "2",
 *     },
 * });
 * const testUser = new vault.identity.Entity("testUser", {});
 * const secondTestUser = new vault.identity.Entity("secondTestUser", {});
 * const devUser = new vault.identity.Entity("devUser", {});
 * const test = new vault.identity.GroupMemberEntityIds("test", {
 *     memberEntityIds: [
 *         testUser.id,
 *         secondTestUser.id,
 *     ],
 *     exclusive: false,
 *     groupId: internal.id,
 * });
 * const others = new vault.identity.GroupMemberEntityIds("others", {
 *     memberEntityIds: [devUser.id],
 *     exclusive: false,
 *     groupId: internal.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class GroupMemberEntityIds extends pulumi.CustomResource {
    /**
     * Get an existing GroupMemberEntityIds resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMemberEntityIdsState, opts?: pulumi.CustomResourceOptions): GroupMemberEntityIds {
        return new GroupMemberEntityIds(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/groupMemberEntityIds:GroupMemberEntityIds';

    /**
     * Returns true if the given object is an instance of GroupMemberEntityIds.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMemberEntityIds {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMemberEntityIds.__pulumiType;
    }

    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     */
    public readonly exclusive!: pulumi.Output<boolean | undefined>;
    /**
     * Group ID to assign member entities to.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * List of member entities that belong to the group
     */
    public readonly memberEntityIds!: pulumi.Output<string[] | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMemberEntityIds resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMemberEntityIdsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMemberEntityIdsArgs | GroupMemberEntityIdsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMemberEntityIdsState | undefined;
            resourceInputs["exclusive"] = state ? state.exclusive : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["memberEntityIds"] = state ? state.memberEntityIds : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as GroupMemberEntityIdsArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["exclusive"] = args ? args.exclusive : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["memberEntityIds"] = args ? args.memberEntityIds : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupMemberEntityIds.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMemberEntityIds resources.
 */
export interface GroupMemberEntityIdsState {
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * Group ID to assign member entities to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * List of member entities that belong to the group
     */
    memberEntityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMemberEntityIds resource.
 */
export interface GroupMemberEntityIdsArgs {
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * Group ID to assign member entities to.
     */
    groupId: pulumi.Input<string>;
    /**
     * List of member entities that belong to the group
     */
    memberEntityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
