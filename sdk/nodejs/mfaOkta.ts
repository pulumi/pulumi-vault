// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).
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
 * const myOkta = new vault.MfaOkta("my_okta", {
 *     name: "my_okta",
 *     mountAccessor: userpass.accessor,
 *     usernameFormat: "user@example.com",
 *     orgName: "hashicorp",
 *     apiToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
 * });
 * ```
 *
 * ## Import
 *
 * Mounts can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
 * ```
 */
export class MfaOkta extends pulumi.CustomResource {
    /**
     * Get an existing MfaOkta resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaOktaState, opts?: pulumi.CustomResourceOptions): MfaOkta {
        return new MfaOkta(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/mfaOkta:MfaOkta';

    /**
     * Returns true if the given object is an instance of MfaOkta.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaOkta {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaOkta.__pulumiType;
    }

    /**
     * `(string: <required>)` - Okta API key.
     */
    public readonly apiToken!: pulumi.Output<string>;
    /**
     * `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
     * `oktapreview.com`, and `okta-emea.com`.
     */
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
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
     * `(string: <required>)` - Name of the organization to be used in the Okta API.
     */
    public readonly orgName!: pulumi.Output<string>;
    /**
     * `(string: <required>)` - If set to true, the username will only match the 
     * primary email for the account.
     */
    public readonly primaryEmail!: pulumi.Output<boolean | undefined>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. 
     * Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
     * If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    public readonly usernameFormat!: pulumi.Output<string | undefined>;

    /**
     * Create a MfaOkta resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaOktaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaOktaArgs | MfaOktaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaOktaState | undefined;
            resourceInputs["apiToken"] = state ? state.apiToken : undefined;
            resourceInputs["baseUrl"] = state ? state.baseUrl : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["orgName"] = state ? state.orgName : undefined;
            resourceInputs["primaryEmail"] = state ? state.primaryEmail : undefined;
            resourceInputs["usernameFormat"] = state ? state.usernameFormat : undefined;
        } else {
            const args = argsOrState as MfaOktaArgs | undefined;
            if ((!args || args.apiToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiToken'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            if ((!args || args.orgName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgName'");
            }
            resourceInputs["apiToken"] = args?.apiToken ? pulumi.secret(args.apiToken) : undefined;
            resourceInputs["baseUrl"] = args ? args.baseUrl : undefined;
            resourceInputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["orgName"] = args ? args.orgName : undefined;
            resourceInputs["primaryEmail"] = args ? args.primaryEmail : undefined;
            resourceInputs["usernameFormat"] = args ? args.usernameFormat : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MfaOkta.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaOkta resources.
 */
export interface MfaOktaState {
    /**
     * `(string: <required>)` - Okta API key.
     */
    apiToken?: pulumi.Input<string>;
    /**
     * `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
     * `oktapreview.com`, and `okta-emea.com`.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
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
     * `(string: <required>)` - Name of the organization to be used in the Okta API.
     */
    orgName?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - If set to true, the username will only match the 
     * primary email for the account.
     */
    primaryEmail?: pulumi.Input<boolean>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. 
     * Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
     * If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    usernameFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MfaOkta resource.
 */
export interface MfaOktaArgs {
    /**
     * `(string: <required>)` - Okta API key.
     */
    apiToken: pulumi.Input<string>;
    /**
     * `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
     * `oktapreview.com`, and `okta-emea.com`.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
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
     * `(string: <required>)` - Name of the organization to be used in the Okta API.
     */
    orgName: pulumi.Input<string>;
    /**
     * `(string: <required>)` - If set to true, the username will only match the 
     * primary email for the account.
     */
    primaryEmail?: pulumi.Input<boolean>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. 
     * Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
     * If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    usernameFormat?: pulumi.Input<string>;
}
