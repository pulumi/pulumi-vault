// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class UiCustomMessage extends pulumi.CustomResource {
    /**
     * Get an existing UiCustomMessage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UiCustomMessageState, opts?: pulumi.CustomResourceOptions): UiCustomMessage {
        return new UiCustomMessage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:config/uiCustomMessage:UiCustomMessage';

    /**
     * Returns true if the given object is an instance of UiCustomMessage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UiCustomMessage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UiCustomMessage.__pulumiType;
    }

    /**
     * A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
     */
    public readonly authenticated!: pulumi.Output<boolean | undefined>;
    /**
     * The ending time of the active period of the custom message. Can be omitted for non-expiring message
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * A block containing a hyperlink associated with the custom message
     */
    public readonly link!: pulumi.Output<outputs.config.UiCustomMessageLink | undefined>;
    /**
     * The base64-encoded content of the custom message
     */
    public readonly messageBase64!: pulumi.Output<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * A map containing additional options for the custom message
     */
    public readonly options!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The starting time of the active period of the custom message
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The title of the custom message
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The display type of custom message. Allowed values are banner and modal
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a UiCustomMessage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UiCustomMessageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UiCustomMessageArgs | UiCustomMessageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UiCustomMessageState | undefined;
            resourceInputs["authenticated"] = state ? state.authenticated : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["link"] = state ? state.link : undefined;
            resourceInputs["messageBase64"] = state ? state.messageBase64 : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as UiCustomMessageArgs | undefined;
            if ((!args || args.messageBase64 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'messageBase64'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["authenticated"] = args ? args.authenticated : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["link"] = args ? args.link : undefined;
            resourceInputs["messageBase64"] = args ? args.messageBase64 : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UiCustomMessage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UiCustomMessage resources.
 */
export interface UiCustomMessageState {
    /**
     * A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
     */
    authenticated?: pulumi.Input<boolean>;
    /**
     * The ending time of the active period of the custom message. Can be omitted for non-expiring message
     */
    endTime?: pulumi.Input<string>;
    /**
     * A block containing a hyperlink associated with the custom message
     */
    link?: pulumi.Input<inputs.config.UiCustomMessageLink>;
    /**
     * The base64-encoded content of the custom message
     */
    messageBase64?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * A map containing additional options for the custom message
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * The starting time of the active period of the custom message
     */
    startTime?: pulumi.Input<string>;
    /**
     * The title of the custom message
     */
    title?: pulumi.Input<string>;
    /**
     * The display type of custom message. Allowed values are banner and modal
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UiCustomMessage resource.
 */
export interface UiCustomMessageArgs {
    /**
     * A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
     */
    authenticated?: pulumi.Input<boolean>;
    /**
     * The ending time of the active period of the custom message. Can be omitted for non-expiring message
     */
    endTime?: pulumi.Input<string>;
    /**
     * A block containing a hyperlink associated with the custom message
     */
    link?: pulumi.Input<inputs.config.UiCustomMessageLink>;
    /**
     * The base64-encoded content of the custom message
     */
    messageBase64: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * A map containing additional options for the custom message
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * The starting time of the active period of the custom message
     */
    startTime: pulumi.Input<string>;
    /**
     * The title of the custom message
     */
    title: pulumi.Input<string>;
    /**
     * The display type of custom message. Allowed values are banner and modal
     */
    type?: pulumi.Input<string>;
}