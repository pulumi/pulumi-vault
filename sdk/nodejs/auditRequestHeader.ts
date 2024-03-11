// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages additional request headers that appear in audited requests.
 *
 * > **Note**
 * Because of the way the [sys/config/auditing/request-headers API](https://www.vaultproject.io/api-docs/system/config-auditing)
 * is implemented in Vault, this resource will manage existing audited headers with
 * matching names without requiring import.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const xForwardedFor = new vault.AuditRequestHeader("xForwardedFor", {hmac: false});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class AuditRequestHeader extends pulumi.CustomResource {
    /**
     * Get an existing AuditRequestHeader resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuditRequestHeaderState, opts?: pulumi.CustomResourceOptions): AuditRequestHeader {
        return new AuditRequestHeader(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/auditRequestHeader:AuditRequestHeader';

    /**
     * Returns true if the given object is an instance of AuditRequestHeader.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuditRequestHeader {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuditRequestHeader.__pulumiType;
    }

    /**
     * Whether this header's value should be HMAC'd in the audit logs.
     */
    public readonly hmac!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the request header to audit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a AuditRequestHeader resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuditRequestHeaderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuditRequestHeaderArgs | AuditRequestHeaderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuditRequestHeaderState | undefined;
            resourceInputs["hmac"] = state ? state.hmac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as AuditRequestHeaderArgs | undefined;
            resourceInputs["hmac"] = args ? args.hmac : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuditRequestHeader.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuditRequestHeader resources.
 */
export interface AuditRequestHeaderState {
    /**
     * Whether this header's value should be HMAC'd in the audit logs.
     */
    hmac?: pulumi.Input<boolean>;
    /**
     * The name of the request header to audit.
     */
    name?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuditRequestHeader resource.
 */
export interface AuditRequestHeaderArgs {
    /**
     * Whether this header's value should be HMAC'd in the audit logs.
     */
    hmac?: pulumi.Input<boolean>;
    /**
     * The name of the request header to audit.
     */
    name?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
}
