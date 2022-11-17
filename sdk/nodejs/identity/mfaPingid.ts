// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for configuring the pingid MFA method.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.identity.MfaDuo("example", {
 *     settingsFileBase64: "CnVzZV9iYXNlNjR[...]HBtCg==",
 * });
 * ```
 *
 * ## Import
 *
 * Resource can be imported using its `uuid` field, e.g.
 *
 * ```sh
 *  $ pulumi import vault:identity/mfaPingid:MfaPingid example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
 * ```
 */
export class MfaPingid extends pulumi.CustomResource {
    /**
     * Get an existing MfaPingid resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaPingidState, opts?: pulumi.CustomResourceOptions): MfaPingid {
        return new MfaPingid(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/mfaPingid:MfaPingid';

    /**
     * Returns true if the given object is an instance of MfaPingid.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaPingid {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaPingid.__pulumiType;
    }

    /**
     * The admin URL, derived from "settingsFileBase64"
     */
    public /*out*/ readonly adminUrl!: pulumi.Output<string>;
    /**
     * A unique identifier of the organization, derived from "settingsFileBase64"
     */
    public /*out*/ readonly authenticatorUrl!: pulumi.Output<string>;
    /**
     * The IDP URL, derived from "settingsFileBase64"
     */
    public /*out*/ readonly idpUrl!: pulumi.Output<string>;
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
     * The name of the PingID client organization, derived from "settingsFileBase64"
     */
    public /*out*/ readonly orgAlias!: pulumi.Output<string>;
    /**
     * A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
     */
    public readonly settingsFileBase64!: pulumi.Output<string>;
    /**
     * MFA type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Use signature value, derived from "settingsFileBase64"
     */
    public /*out*/ readonly useSignature!: pulumi.Output<boolean>;
    /**
     * A template string for mapping Identity names to MFA methods.
     */
    public readonly usernameFormat!: pulumi.Output<string | undefined>;
    /**
     * Resource UUID.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a MfaPingid resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaPingidArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaPingidArgs | MfaPingidState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaPingidState | undefined;
            resourceInputs["adminUrl"] = state ? state.adminUrl : undefined;
            resourceInputs["authenticatorUrl"] = state ? state.authenticatorUrl : undefined;
            resourceInputs["idpUrl"] = state ? state.idpUrl : undefined;
            resourceInputs["methodId"] = state ? state.methodId : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["namespacePath"] = state ? state.namespacePath : undefined;
            resourceInputs["orgAlias"] = state ? state.orgAlias : undefined;
            resourceInputs["settingsFileBase64"] = state ? state.settingsFileBase64 : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["useSignature"] = state ? state.useSignature : undefined;
            resourceInputs["usernameFormat"] = state ? state.usernameFormat : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as MfaPingidArgs | undefined;
            if ((!args || args.settingsFileBase64 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settingsFileBase64'");
            }
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["settingsFileBase64"] = args ? args.settingsFileBase64 : undefined;
            resourceInputs["usernameFormat"] = args ? args.usernameFormat : undefined;
            resourceInputs["adminUrl"] = undefined /*out*/;
            resourceInputs["authenticatorUrl"] = undefined /*out*/;
            resourceInputs["idpUrl"] = undefined /*out*/;
            resourceInputs["methodId"] = undefined /*out*/;
            resourceInputs["mountAccessor"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["namespacePath"] = undefined /*out*/;
            resourceInputs["orgAlias"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["useSignature"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MfaPingid.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaPingid resources.
 */
export interface MfaPingidState {
    /**
     * The admin URL, derived from "settingsFileBase64"
     */
    adminUrl?: pulumi.Input<string>;
    /**
     * A unique identifier of the organization, derived from "settingsFileBase64"
     */
    authenticatorUrl?: pulumi.Input<string>;
    /**
     * The IDP URL, derived from "settingsFileBase64"
     */
    idpUrl?: pulumi.Input<string>;
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
     * The name of the PingID client organization, derived from "settingsFileBase64"
     */
    orgAlias?: pulumi.Input<string>;
    /**
     * A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
     */
    settingsFileBase64?: pulumi.Input<string>;
    /**
     * MFA type.
     */
    type?: pulumi.Input<string>;
    /**
     * Use signature value, derived from "settingsFileBase64"
     */
    useSignature?: pulumi.Input<boolean>;
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
 * The set of arguments for constructing a MfaPingid resource.
 */
export interface MfaPingidArgs {
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
     */
    settingsFileBase64: pulumi.Input<string>;
    /**
     * A template string for mapping Identity names to MFA methods.
     */
    usernameFormat?: pulumi.Input<string>;
}