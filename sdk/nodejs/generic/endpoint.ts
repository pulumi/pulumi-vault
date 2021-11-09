// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Import is not supported for this resource.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:generic/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * String containing a JSON-encoded object that will be
     * written to the given path as the secret data.
     */
    public readonly dataJson!: pulumi.Output<string>;
    /**
     * Don't attempt to delete the path from Vault if true
     */
    public readonly disableDelete!: pulumi.Output<boolean | undefined>;
    /**
     * True/false. Set this to true if your vault
     * authentication is not able to read the data or if the endpoint does
     * not support the `GET` method. Setting this to `true` will break drift
     * detection. You should set this to `true` for endpoints that are
     * write-only. Defaults to false.
     */
    public readonly disableRead!: pulumi.Output<boolean | undefined>;
    /**
     * When reading, disregard fields not present in data_json
     */
    public readonly ignoreAbsentFields!: pulumi.Output<boolean | undefined>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Map of strings returned by write operation
     */
    public /*out*/ readonly writeData!: pulumi.Output<{[key: string]: string}>;
    /**
     * JSON data returned by write operation
     */
    public /*out*/ readonly writeDataJson!: pulumi.Output<string>;
    /**
     * Top-level fields returned by write to persist in state
     */
    public readonly writeFields!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["dataJson"] = state ? state.dataJson : undefined;
            inputs["disableDelete"] = state ? state.disableDelete : undefined;
            inputs["disableRead"] = state ? state.disableRead : undefined;
            inputs["ignoreAbsentFields"] = state ? state.ignoreAbsentFields : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["writeData"] = state ? state.writeData : undefined;
            inputs["writeDataJson"] = state ? state.writeDataJson : undefined;
            inputs["writeFields"] = state ? state.writeFields : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if ((!args || args.dataJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataJson'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            inputs["dataJson"] = args ? args.dataJson : undefined;
            inputs["disableDelete"] = args ? args.disableDelete : undefined;
            inputs["disableRead"] = args ? args.disableRead : undefined;
            inputs["ignoreAbsentFields"] = args ? args.ignoreAbsentFields : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["writeFields"] = args ? args.writeFields : undefined;
            inputs["writeData"] = undefined /*out*/;
            inputs["writeDataJson"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * String containing a JSON-encoded object that will be
     * written to the given path as the secret data.
     */
    dataJson?: pulumi.Input<string>;
    /**
     * Don't attempt to delete the path from Vault if true
     */
    disableDelete?: pulumi.Input<boolean>;
    /**
     * True/false. Set this to true if your vault
     * authentication is not able to read the data or if the endpoint does
     * not support the `GET` method. Setting this to `true` will break drift
     * detection. You should set this to `true` for endpoints that are
     * write-only. Defaults to false.
     */
    disableRead?: pulumi.Input<boolean>;
    /**
     * When reading, disregard fields not present in data_json
     */
    ignoreAbsentFields?: pulumi.Input<boolean>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    path?: pulumi.Input<string>;
    /**
     * Map of strings returned by write operation
     */
    writeData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * JSON data returned by write operation
     */
    writeDataJson?: pulumi.Input<string>;
    /**
     * Top-level fields returned by write to persist in state
     */
    writeFields?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * String containing a JSON-encoded object that will be
     * written to the given path as the secret data.
     */
    dataJson: pulumi.Input<string>;
    /**
     * Don't attempt to delete the path from Vault if true
     */
    disableDelete?: pulumi.Input<boolean>;
    /**
     * True/false. Set this to true if your vault
     * authentication is not able to read the data or if the endpoint does
     * not support the `GET` method. Setting this to `true` will break drift
     * detection. You should set this to `true` for endpoints that are
     * write-only. Defaults to false.
     */
    disableRead?: pulumi.Input<boolean>;
    /**
     * When reading, disregard fields not present in data_json
     */
    ignoreAbsentFields?: pulumi.Input<boolean>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    path: pulumi.Input<string>;
    /**
     * Top-level fields returned by write to persist in state
     */
    writeFields?: pulumi.Input<pulumi.Input<string>[]>;
}
