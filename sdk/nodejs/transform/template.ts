// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports the `/transform/template/{name}` Vault endpoint.
 *
 * It creates or updates a template with the given name. If a template with the name does not exist,
 * it will be created. If the template exists, it will be updated with the new attributes.
 *
 * > Requires _Vault Enterprise with the Advanced Data Protection Transform Module_.
 * See [Transform Secrets Engine](https://www.vaultproject.io/docs/secrets/transform)
 * for more information.
 *
 * ## Example Usage
 *
 * Please note that the `pattern` below holds a regex. The regex shown
 * is identical to the one in our [Setup](https://www.vaultproject.io/docs/secrets/transform#setup)
 * docs, `(\d{4})-(\d{4})-(\d{4})-(\d{4})`. However, due to HCL, the
 * backslashes must be escaped to appear correctly in Vault. For further
 * assistance escaping your own custom regex, see String Literals.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transform = new vault.Mount("transform", {
 *     path: "transform",
 *     type: "transform",
 * });
 * const numerics = new vault.transform.Alphabet("numerics", {
 *     path: transform.path,
 *     name: "numerics",
 *     alphabet: "0123456789",
 * });
 * const test = new vault.transform.Template("test", {
 *     path: numerics.path,
 *     name: "ccn",
 *     type: "regex",
 *     pattern: "(\\d{4})[- ](\\d{4})[- ](\\d{4})[- ](\\d{4})",
 *     alphabet: "numerics",
 *     encodeFormat: "$1-$2-$3-$4",
 *     decodeFormats: {
 *         "last-four-digits": "$4",
 *     },
 * });
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:transform/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     */
    public readonly alphabet!: pulumi.Output<string | undefined>;
    /**
     * Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     */
    public readonly decodeFormats!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The regular expression template used to format encoded values.
     * (requires Vault Enterprise 1.9+)
     */
    public readonly encodeFormat!: pulumi.Output<string | undefined>;
    /**
     * The name of the template.
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
     * Path to where the back-end is mounted within Vault.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     */
    public readonly pattern!: pulumi.Output<string | undefined>;
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateState | undefined;
            resourceInputs["alphabet"] = state ? state.alphabet : undefined;
            resourceInputs["decodeFormats"] = state ? state.decodeFormats : undefined;
            resourceInputs["encodeFormat"] = state ? state.encodeFormat : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["alphabet"] = args ? args.alphabet : undefined;
            resourceInputs["decodeFormats"] = args ? args.decodeFormats : undefined;
            resourceInputs["encodeFormat"] = args ? args.encodeFormat : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     */
    alphabet?: pulumi.Input<string>;
    /**
     * Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     */
    decodeFormats?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The regular expression template used to format encoded values.
     * (requires Vault Enterprise 1.9+)
     */
    encodeFormat?: pulumi.Input<string>;
    /**
     * The name of the template.
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
     * Path to where the back-end is mounted within Vault.
     */
    path?: pulumi.Input<string>;
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     */
    pattern?: pulumi.Input<string>;
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     */
    alphabet?: pulumi.Input<string>;
    /**
     * Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     */
    decodeFormats?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The regular expression template used to format encoded values.
     * (requires Vault Enterprise 1.9+)
     */
    encodeFormat?: pulumi.Input<string>;
    /**
     * The name of the template.
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
     * Path to where the back-end is mounted within Vault.
     */
    path: pulumi.Input<string>;
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     */
    pattern?: pulumi.Input<string>;
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     */
    type?: pulumi.Input<string>;
}
