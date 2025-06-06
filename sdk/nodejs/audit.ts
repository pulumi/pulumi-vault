// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### File Audit Device)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.Audit("test", {
 *     type: "file",
 *     options: {
 *         file_path: "C:/temp/audit.txt",
 *     },
 * });
 * ```
 *
 * ### Socket Audit Device)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.Audit("test", {
 *     type: "socket",
 *     path: "app_socket",
 *     local: false,
 *     options: {
 *         address: "127.0.0.1:8000",
 *         socket_type: "tcp",
 *         description: "application x socket",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Audit devices can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/audit:Audit test syslog
 * ```
 */
export class Audit extends pulumi.CustomResource {
    /**
     * Get an existing Audit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuditState, opts?: pulumi.CustomResourceOptions): Audit {
        return new Audit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/audit:Audit';

    /**
     * Returns true if the given object is an instance of Audit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Audit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Audit.__pulumiType;
    }

    /**
     * Human-friendly description of the audit device.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Configuration options to pass to the audit device itself.
     *
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     */
    public readonly options!: pulumi.Output<{[key: string]: string}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Audit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuditArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuditArgs | AuditState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuditState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuditArgs | undefined;
            if ((!args || args.options === undefined) && !opts.urn) {
                throw new Error("Missing required property 'options'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Audit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Audit resources.
 */
export interface AuditState {
    /**
     * Human-friendly description of the audit device.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Configuration options to pass to the audit device itself.
     *
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    path?: pulumi.Input<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Audit resource.
 */
export interface AuditArgs {
    /**
     * Human-friendly description of the audit device.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Configuration options to pass to the audit device itself.
     *
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     */
    options: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    path?: pulumi.Input<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    type: pulumi.Input<string>;
}
