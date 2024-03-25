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
 * const userpass = new vault.AuthBackend("userpass", {type: "userpass"});
 * const u1 = new vault.generic.Endpoint("u1", {
 *     path: "auth/userpass/users/u1",
 *     ignoreAbsentFields: true,
 *     dataJson: `{
 *   "policies": ["p1"],
 *   "password": "changeme"
 * }
 * `,
 * }, {
 *     dependsOn: [userpass],
 * });
 * const u1Token = new vault.generic.Endpoint("u1Token", {
 *     path: "auth/userpass/login/u1",
 *     disableRead: true,
 *     disableDelete: true,
 *     dataJson: `{
 *   "password": "changeme"
 * }
 * `,
 * }, {
 *     dependsOn: [u1],
 * });
 * const u1Entity = new vault.generic.Endpoint("u1Entity", {
 *     disableRead: true,
 *     disableDelete: true,
 *     path: "identity/lookup/entity",
 *     ignoreAbsentFields: true,
 *     writeFields: ["id"],
 *     dataJson: `{
 *   "alias_name": "u1",
 *   "alias_mount_accessor": vault_auth_backend.userpass.accessor
 * }
 * `,
 * }, {
 *     dependsOn: [u1Token],
 * });
 * export const u1Id = u1Entity.writeData.id;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `create` or `update` capability
 * (depending on whether the resource already exists) on the given path. If
 * `disableDelete` is false, the `delete` capability is also required. If
 * `disableRead` is false, the `read` capability is required.
 *
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
     * - (Optional) True/false. Set this to true if your
     * vault authentication is not able to delete the data or if the endpoint
     * does not support the `DELETE` method. Defaults to false.
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
     * - (Optional) True/false. If set to true,
     * ignore any fields present when the endpoint is read but that were not
     * in `dataJson`. Also, if a field that was written is not returned when
     * the endpoint is read, treat that field as being up to date. You should
     * set this to `true` when writing to endpoint that, when read, returns a
     * different set of fields from the ones you wrote, as is common with
     * many configuration endpoints. Defaults to false.
     */
    public readonly ignoreAbsentFields!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * - A map whose keys are the top-level data keys
     * returned from Vault by the write operation and whose values are the
     * corresponding values. This map can only represent string data, so
     * any non-string values returned from Vault are serialized as JSON.
     * Only fields set in `writeFields` are present in the JSON data.
     */
    public /*out*/ readonly writeData!: pulumi.Output<{[key: string]: string}>;
    /**
     * - The JSON data returned by the write operation.
     * Only fields set in `writeFields` are present in the JSON data.
     */
    public /*out*/ readonly writeDataJson!: pulumi.Output<string>;
    /**
     * - (Optional). A list of fields that should be returned
     * in `writeDataJson` and `writeData`. If omitted, data returned by
     * the write operation is not available to the resource or included in
     * state. This helps to avoid accidental storage of sensitive values in
     * state. Some endpoints, such as many dynamic secrets endpoints, return
     * data from writing to an endpoint rather than reading it. You should
     * use `writeFields` if you need information returned in this way.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointState | undefined;
            resourceInputs["dataJson"] = state ? state.dataJson : undefined;
            resourceInputs["disableDelete"] = state ? state.disableDelete : undefined;
            resourceInputs["disableRead"] = state ? state.disableRead : undefined;
            resourceInputs["ignoreAbsentFields"] = state ? state.ignoreAbsentFields : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["writeData"] = state ? state.writeData : undefined;
            resourceInputs["writeDataJson"] = state ? state.writeDataJson : undefined;
            resourceInputs["writeFields"] = state ? state.writeFields : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if ((!args || args.dataJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataJson'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["dataJson"] = args?.dataJson ? pulumi.secret(args.dataJson) : undefined;
            resourceInputs["disableDelete"] = args ? args.disableDelete : undefined;
            resourceInputs["disableRead"] = args ? args.disableRead : undefined;
            resourceInputs["ignoreAbsentFields"] = args ? args.ignoreAbsentFields : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["writeFields"] = args ? args.writeFields : undefined;
            resourceInputs["writeData"] = undefined /*out*/;
            resourceInputs["writeDataJson"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["dataJson"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Endpoint.__pulumiType, name, resourceInputs, opts);
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
     * - (Optional) True/false. Set this to true if your
     * vault authentication is not able to delete the data or if the endpoint
     * does not support the `DELETE` method. Defaults to false.
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
     * - (Optional) True/false. If set to true,
     * ignore any fields present when the endpoint is read but that were not
     * in `dataJson`. Also, if a field that was written is not returned when
     * the endpoint is read, treat that field as being up to date. You should
     * set this to `true` when writing to endpoint that, when read, returns a
     * different set of fields from the ones you wrote, as is common with
     * many configuration endpoints. Defaults to false.
     */
    ignoreAbsentFields?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    path?: pulumi.Input<string>;
    /**
     * - A map whose keys are the top-level data keys
     * returned from Vault by the write operation and whose values are the
     * corresponding values. This map can only represent string data, so
     * any non-string values returned from Vault are serialized as JSON.
     * Only fields set in `writeFields` are present in the JSON data.
     */
    writeData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * - The JSON data returned by the write operation.
     * Only fields set in `writeFields` are present in the JSON data.
     */
    writeDataJson?: pulumi.Input<string>;
    /**
     * - (Optional). A list of fields that should be returned
     * in `writeDataJson` and `writeData`. If omitted, data returned by
     * the write operation is not available to the resource or included in
     * state. This helps to avoid accidental storage of sensitive values in
     * state. Some endpoints, such as many dynamic secrets endpoints, return
     * data from writing to an endpoint rather than reading it. You should
     * use `writeFields` if you need information returned in this way.
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
     * - (Optional) True/false. Set this to true if your
     * vault authentication is not able to delete the data or if the endpoint
     * does not support the `DELETE` method. Defaults to false.
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
     * - (Optional) True/false. If set to true,
     * ignore any fields present when the endpoint is read but that were not
     * in `dataJson`. Also, if a field that was written is not returned when
     * the endpoint is read, treat that field as being up to date. You should
     * set this to `true` when writing to endpoint that, when read, returns a
     * different set of fields from the ones you wrote, as is common with
     * many configuration endpoints. Defaults to false.
     */
    ignoreAbsentFields?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The full logical path at which to write the given
     * data. Consult each backend's documentation to see which endpoints
     * support the `PUT` methods and to determine whether they also support
     * `DELETE` and `GET`.
     */
    path: pulumi.Input<string>;
    /**
     * - (Optional). A list of fields that should be returned
     * in `writeDataJson` and `writeData`. If omitted, data returned by
     * the write operation is not available to the resource or included in
     * state. This helps to avoid accidental storage of sensitive values in
     * state. Some endpoints, such as many dynamic secrets endpoints, return
     * data from writing to an endpoint rather than reading it. You should
     * use `writeFields` if you need information returned in this way.
     */
    writeFields?: pulumi.Input<pulumi.Input<string>[]>;
}
