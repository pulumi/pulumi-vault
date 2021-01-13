// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
     * The name of the template.
     */
    public readonly name!: pulumi.Output<string>;
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TemplateState | undefined;
            inputs["alphabet"] = state ? state.alphabet : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["pattern"] = state ? state.pattern : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.path === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'path'");
            }
            inputs["alphabet"] = args ? args.alphabet : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["pattern"] = args ? args.pattern : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Template.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     */
    readonly alphabet?: pulumi.Input<string>;
    /**
     * The name of the template.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path to where the back-end is mounted within Vault.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     */
    readonly pattern?: pulumi.Input<string>;
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     */
    readonly alphabet?: pulumi.Input<string>;
    /**
     * The name of the template.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path to where the back-end is mounted within Vault.
     */
    readonly path: pulumi.Input<string>;
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     */
    readonly pattern?: pulumi.Input<string>;
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     */
    readonly type?: pulumi.Input<string>;
}
