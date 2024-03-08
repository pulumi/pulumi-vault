// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const mongo = new vault.Mount("mongo", {
 *     description: "MongoDB Atlas secret engine mount",
 *     path: "mongodbatlas",
 *     type: "mongodbatlas",
 * });
 * const config = new vault.mongodbatlas.SecretBackend("config", {
 *     mount: "vault_mount.mongo.path",
 *     privateKey: "privateKey",
 *     publicKey: "publicKey",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * MongoDB Atlas secret backends can be imported using the `${mount}/config`, e.g.
 *
 * ```sh
 * $ pulumi import vault:mongodbatlas/secretBackend:SecretBackend config mongodbatlas/config
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:mongodbatlas/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     */
    public readonly mount!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Path where MongoDB Atlas configuration is located
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     */
    public readonly publicKey!: pulumi.Output<string>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["mount"] = state ? state.mount : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if ((!args || args.mount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mount'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["mount"] = args ? args.mount : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["path"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     */
    mount?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Path where MongoDB Atlas configuration is located
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     */
    publicKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     */
    mount: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     */
    privateKey: pulumi.Input<string>;
    /**
     * Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     */
    publicKey: pulumi.Input<string>;
}
