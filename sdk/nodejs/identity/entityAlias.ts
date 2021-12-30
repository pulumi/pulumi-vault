// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Identity entity alias can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:identity/entityAlias:EntityAlias test "3856fb4d-3c91-dcaf-2401-68f446796bfb"
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
     * Create a EntityAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntityAliasArgs | EntityAliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntityAliasState | undefined;
            inputs["canonicalId"] = state ? state.canonicalId : undefined;
            inputs["customMetadata"] = state ? state.customMetadata : undefined;
            inputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as EntityAliasArgs | undefined;
            if ((!args || args.canonicalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'canonicalId'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            inputs["canonicalId"] = args ? args.canonicalId : undefined;
            inputs["customMetadata"] = args ? args.customMetadata : undefined;
            inputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EntityAlias.__pulumiType, name, inputs, opts);
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
}
