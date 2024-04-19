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
 * const test = new vault.identity.EntityAlias("test", {
 *     canonicalId: "49877D63-07AD-4B85-BDA8-B61626C477E8",
 *     mountAccessor: "token_1f2bd5",
 * });
 * ```
 *
 * ## Import
 *
 * Identity entity alias can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import vault:identity/entityAlias:EntityAlias test "3856fb4d-3c91-dcaf-2401-68f446796bfb"
 * ```
 */
export class EntityAlias extends pulumi.CustomResource {
    /**
     * Get an existing EntityAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntityAliasState, opts?: pulumi.CustomResourceOptions): EntityAlias {
        return new EntityAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/entityAlias:EntityAlias';

    /**
     * Returns true if the given object is an instance of EntityAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityAlias.__pulumiType;
    }

    /**
     * Entity ID to which this alias belongs to.
     */
    public readonly canonicalId!: pulumi.Output<string>;
    /**
     * Custom metadata to be associated with this alias.
     */
    public readonly customMetadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Accessor of the mount to which the alias should belong to.
     */
    public readonly mountAccessor!: pulumi.Output<string>;
    /**
     * Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
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
     * Create a EntityAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntityAliasArgs | EntityAliasState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntityAliasState | undefined;
            resourceInputs["canonicalId"] = state ? state.canonicalId : undefined;
            resourceInputs["customMetadata"] = state ? state.customMetadata : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as EntityAliasArgs | undefined;
            if ((!args || args.canonicalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'canonicalId'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            resourceInputs["canonicalId"] = args ? args.canonicalId : undefined;
            resourceInputs["customMetadata"] = args ? args.customMetadata : undefined;
            resourceInputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntityAlias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EntityAlias resources.
 */
export interface EntityAliasState {
    /**
     * Entity ID to which this alias belongs to.
     */
    canonicalId?: pulumi.Input<string>;
    /**
     * Custom metadata to be associated with this alias.
     */
    customMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Accessor of the mount to which the alias should belong to.
     */
    mountAccessor?: pulumi.Input<string>;
    /**
     * Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EntityAlias resource.
 */
export interface EntityAliasArgs {
    /**
     * Entity ID to which this alias belongs to.
     */
    canonicalId: pulumi.Input<string>;
    /**
     * Custom metadata to be associated with this alias.
     */
    customMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Accessor of the mount to which the alias should belong to.
     */
    mountAccessor: pulumi.Input<string>;
    /**
     * Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
