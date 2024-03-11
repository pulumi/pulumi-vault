// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Namespaces can be imported using its `name` as accessor id
 *
 * ```sh
 * $ pulumi import vault:index/namespace:Namespace example <name>
 * ```
 *
 * If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path.
 *
 * hcl
 *
 * provider "vault" {
 *
 * # Configuration options
 *
 *   namespace = "example"
 *
 *   alias     = "example"
 *
 * }
 *
 * resource "vault_namespace" "example2" {
 *
 *   provider = vault.example
 *
 *   path     = "example2"
 *
 * }
 *
 * ```sh
 * $ pulumi import vault:index/namespace:Namespace example2 example2
 * ```
 *
 * $ terraform state show vault_namespace.example2
 *
 * vault_namespace.example2:
 *
 * resource "vault_namespace" "example2" {
 *
 *     id           = "example/example2/"
 *     
 *     namespace_id = <known after import>
 *     
 *     path         = "example2"
 *     
 *     path_fq      = "example2"
 *
 * }
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Custom metadata describing this namespace. Value type
     * is `map[string]string`. Requires Vault version 1.12+.
     */
    public readonly customMetadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Vault server's internal ID of the namespace.
     */
    public /*out*/ readonly namespaceId!: pulumi.Output<string>;
    /**
     * The path of the namespace. Must not have a trailing `/`.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
     * The path is relative to the provider's `namespace` argument.
     */
    public readonly pathFq!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["customMetadata"] = state ? state.customMetadata : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["pathFq"] = state ? state.pathFq : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["customMetadata"] = args ? args.customMetadata : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["pathFq"] = args ? args.pathFq : undefined;
            resourceInputs["namespaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Custom metadata describing this namespace. Value type
     * is `map[string]string`. Requires Vault version 1.12+.
     */
    customMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Vault server's internal ID of the namespace.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The path of the namespace. Must not have a trailing `/`.
     */
    path?: pulumi.Input<string>;
    /**
     * The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
     * The path is relative to the provider's `namespace` argument.
     */
    pathFq?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Custom metadata describing this namespace. Value type
     * is `map[string]string`. Requires Vault version 1.12+.
     */
    customMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The path of the namespace. Must not have a trailing `/`.
     */
    path: pulumi.Input<string>;
    /**
     * The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
     * The path is relative to the provider's `namespace` argument.
     */
    pathFq?: pulumi.Input<string>;
}
