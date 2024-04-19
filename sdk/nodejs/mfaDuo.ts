// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
 *
 * **Note** this feature is available only with Vault Enterprise.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const userpass = new vault.AuthBackend("userpass", {
 *     type: "userpass",
 *     path: "userpass",
 * });
 * const myDuo = new vault.MfaDuo("myDuo", {
 *     mountAccessor: userpass.accessor,
 *     secretKey: "8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz",
 *     integrationKey: "BIACEUEAXI20BNWTEYXT",
 *     apiHostname: "api-2b5c39f5.duosecurity.com",
 * });
 * ```
 *
 * ## Import
 *
 * Mounts can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
 * ```
 */
export class MfaDuo extends pulumi.CustomResource {
    /**
     * Get an existing MfaDuo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaDuoState, opts?: pulumi.CustomResourceOptions): MfaDuo {
        return new MfaDuo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/mfaDuo:MfaDuo';

    /**
     * Returns true if the given object is an instance of MfaDuo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaDuo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaDuo.__pulumiType;
    }

    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    public readonly apiHostname!: pulumi.Output<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    public readonly mountAccessor!: pulumi.Output<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
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
     * `(string)` - Push information for Duo.
     */
    public readonly pushInfo!: pulumi.Output<string | undefined>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    public readonly usernameFormat!: pulumi.Output<string | undefined>;

    /**
     * Create a MfaDuo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaDuoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaDuoArgs | MfaDuoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaDuoState | undefined;
            resourceInputs["apiHostname"] = state ? state.apiHostname : undefined;
            resourceInputs["integrationKey"] = state ? state.integrationKey : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["pushInfo"] = state ? state.pushInfo : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["usernameFormat"] = state ? state.usernameFormat : undefined;
        } else {
            const args = argsOrState as MfaDuoArgs | undefined;
            if ((!args || args.apiHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiHostname'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["apiHostname"] = args ? args.apiHostname : undefined;
            resourceInputs["integrationKey"] = args?.integrationKey ? pulumi.secret(args.integrationKey) : undefined;
            resourceInputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["pushInfo"] = args ? args.pushInfo : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["usernameFormat"] = args ? args.usernameFormat : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["integrationKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MfaDuo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaDuo resources.
 */
export interface MfaDuoState {
    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    apiHostname?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    integrationKey?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    mountAccessor?: pulumi.Input<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * `(string)` - Push information for Duo.
     */
    pushInfo?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    usernameFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MfaDuo resource.
 */
export interface MfaDuoArgs {
    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    apiHostname: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    integrationKey: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    mountAccessor: pulumi.Input<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * `(string)` - Push information for Duo.
     */
    pushInfo?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    secretKey: pulumi.Input<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    usernameFormat?: pulumi.Input<string>;
}
