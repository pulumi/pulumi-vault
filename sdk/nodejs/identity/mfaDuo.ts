// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for configuring the duo MFA method.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.identity.MfaDuo("example", {
 *     apiHostname: "api-xxxxxxxx.duosecurity.com",
 *     integrationKey: "secret-int-key",
 *     secretKey: "secret-key",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Resource can be imported using its `uuid` field, e.g.
 *
 * ```sh
 * $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
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
    public static readonly __pulumiType = 'vault:identity/mfaDuo:MfaDuo';

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
     * API hostname for Duo
     */
    public readonly apiHostname!: pulumi.Output<string>;
    /**
     * Integration key for Duo
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * Method ID.
     */
    public /*out*/ readonly methodId!: pulumi.Output<string>;
    /**
     * Mount accessor.
     */
    public /*out*/ readonly mountAccessor!: pulumi.Output<string>;
    /**
     * Method name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Method's namespace ID.
     */
    public /*out*/ readonly namespaceId!: pulumi.Output<string>;
    /**
     * Method's namespace path.
     */
    public /*out*/ readonly namespacePath!: pulumi.Output<string>;
    /**
     * Push information for Duo.
     */
    public readonly pushInfo!: pulumi.Output<string | undefined>;
    /**
     * Secret key for Duo
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * MFA type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Require passcode upon MFA validation.
     */
    public readonly usePasscode!: pulumi.Output<boolean | undefined>;
    /**
     * A template string for mapping Identity names to MFA methods.
     */
    public readonly usernameFormat!: pulumi.Output<string | undefined>;
    /**
     * Resource UUID.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

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
            resourceInputs["methodId"] = state ? state.methodId : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["namespacePath"] = state ? state.namespacePath : undefined;
            resourceInputs["pushInfo"] = state ? state.pushInfo : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["usePasscode"] = state ? state.usePasscode : undefined;
            resourceInputs["usernameFormat"] = state ? state.usernameFormat : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as MfaDuoArgs | undefined;
            if ((!args || args.apiHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiHostname'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["apiHostname"] = args ? args.apiHostname : undefined;
            resourceInputs["integrationKey"] = args?.integrationKey ? pulumi.secret(args.integrationKey) : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["pushInfo"] = args ? args.pushInfo : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["usePasscode"] = args ? args.usePasscode : undefined;
            resourceInputs["usernameFormat"] = args ? args.usernameFormat : undefined;
            resourceInputs["methodId"] = undefined /*out*/;
            resourceInputs["mountAccessor"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["namespacePath"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
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
     * API hostname for Duo
     */
    apiHostname?: pulumi.Input<string>;
    /**
     * Integration key for Duo
     */
    integrationKey?: pulumi.Input<string>;
    /**
     * Method ID.
     */
    methodId?: pulumi.Input<string>;
    /**
     * Mount accessor.
     */
    mountAccessor?: pulumi.Input<string>;
    /**
     * Method name.
     */
    name?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Method's namespace ID.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * Method's namespace path.
     */
    namespacePath?: pulumi.Input<string>;
    /**
     * Push information for Duo.
     */
    pushInfo?: pulumi.Input<string>;
    /**
     * Secret key for Duo
     */
    secretKey?: pulumi.Input<string>;
    /**
     * MFA type.
     */
    type?: pulumi.Input<string>;
    /**
     * Require passcode upon MFA validation.
     */
    usePasscode?: pulumi.Input<boolean>;
    /**
     * A template string for mapping Identity names to MFA methods.
     */
    usernameFormat?: pulumi.Input<string>;
    /**
     * Resource UUID.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MfaDuo resource.
 */
export interface MfaDuoArgs {
    /**
     * API hostname for Duo
     */
    apiHostname: pulumi.Input<string>;
    /**
     * Integration key for Duo
     */
    integrationKey: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Push information for Duo.
     */
    pushInfo?: pulumi.Input<string>;
    /**
     * Secret key for Duo
     */
    secretKey: pulumi.Input<string>;
    /**
     * Require passcode upon MFA validation.
     */
    usePasscode?: pulumi.Input<boolean>;
    /**
     * A template string for mapping Identity names to MFA methods.
     */
    usernameFormat?: pulumi.Input<string>;
}
