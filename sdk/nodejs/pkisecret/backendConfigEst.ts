// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows setting the EST configuration on a PKI Secret Backend
 *
 * ## Import
 *
 * The PKI config cluster can be imported using the resource's `id`.
 * In the case of the example above the `id` would be `pki-root/config/est`,
 * where the `pki-root` component is the resource's `backend`, e.g.
 *
 * ```sh
 * $ pulumi import vault:pkiSecret/backendConfigEst:BackendConfigEst example pki-root/config/est
 * ```
 */
export class BackendConfigEst extends pulumi.CustomResource {
    /**
     * Get an existing BackendConfigEst resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendConfigEstState, opts?: pulumi.CustomResourceOptions): BackendConfigEst {
        return new BackendConfigEst(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/backendConfigEst:BackendConfigEst';

    /**
     * Returns true if the given object is an instance of BackendConfigEst.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendConfigEst {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendConfigEst.__pulumiType;
    }

    /**
     * Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     *
     * <a id="nestedatt--authenticators"></a>
     */
    public readonly auditFields!: pulumi.Output<string[]>;
    /**
     * Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     */
    public readonly authenticators!: pulumi.Output<outputs.pkiSecret.BackendConfigEstAuthenticators>;
    /**
     * The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     */
    public readonly defaultMount!: pulumi.Output<boolean | undefined>;
    /**
     * Required to be set if defaultMount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:<role_name>.
     */
    public readonly defaultPathPolicy!: pulumi.Output<string | undefined>;
    /**
     * If set, parse out fields from the provided CSR making them available for Sentinel policies.
     */
    public readonly enableSentinelParsing!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether EST is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:<role_name>. Labels must be unique across Vault cluster, and will register .well-known/est/<label> URL paths.
     */
    public readonly labelToPathPolicy!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A read-only timestamp representing the last time the configuration was updated.
     */
    public /*out*/ readonly lastUpdated!: pulumi.Output<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a BackendConfigEst resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendConfigEstArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendConfigEstArgs | BackendConfigEstState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendConfigEstState | undefined;
            resourceInputs["auditFields"] = state ? state.auditFields : undefined;
            resourceInputs["authenticators"] = state ? state.authenticators : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["defaultMount"] = state ? state.defaultMount : undefined;
            resourceInputs["defaultPathPolicy"] = state ? state.defaultPathPolicy : undefined;
            resourceInputs["enableSentinelParsing"] = state ? state.enableSentinelParsing : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["labelToPathPolicy"] = state ? state.labelToPathPolicy : undefined;
            resourceInputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as BackendConfigEstArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            resourceInputs["auditFields"] = args ? args.auditFields : undefined;
            resourceInputs["authenticators"] = args ? args.authenticators : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["defaultMount"] = args ? args.defaultMount : undefined;
            resourceInputs["defaultPathPolicy"] = args ? args.defaultPathPolicy : undefined;
            resourceInputs["enableSentinelParsing"] = args ? args.enableSentinelParsing : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["labelToPathPolicy"] = args ? args.labelToPathPolicy : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["lastUpdated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackendConfigEst.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendConfigEst resources.
 */
export interface BackendConfigEstState {
    /**
     * Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     *
     * <a id="nestedatt--authenticators"></a>
     */
    auditFields?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     */
    authenticators?: pulumi.Input<inputs.pkiSecret.BackendConfigEstAuthenticators>;
    /**
     * The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     */
    defaultMount?: pulumi.Input<boolean>;
    /**
     * Required to be set if defaultMount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:<role_name>.
     */
    defaultPathPolicy?: pulumi.Input<string>;
    /**
     * If set, parse out fields from the provided CSR making them available for Sentinel policies.
     */
    enableSentinelParsing?: pulumi.Input<boolean>;
    /**
     * Specifies whether EST is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:<role_name>. Labels must be unique across Vault cluster, and will register .well-known/est/<label> URL paths.
     */
    labelToPathPolicy?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A read-only timestamp representing the last time the configuration was updated.
     */
    lastUpdated?: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendConfigEst resource.
 */
export interface BackendConfigEstArgs {
    /**
     * Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     *
     * <a id="nestedatt--authenticators"></a>
     */
    auditFields?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     */
    authenticators?: pulumi.Input<inputs.pkiSecret.BackendConfigEstAuthenticators>;
    /**
     * The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     */
    defaultMount?: pulumi.Input<boolean>;
    /**
     * Required to be set if defaultMount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:<role_name>.
     */
    defaultPathPolicy?: pulumi.Input<string>;
    /**
     * If set, parse out fields from the provided CSR making them available for Sentinel policies.
     */
    enableSentinelParsing?: pulumi.Input<boolean>;
    /**
     * Specifies whether EST is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:<role_name>. Labels must be unique across Vault cluster, and will register .well-known/est/<label> URL paths.
     */
    labelToPathPolicy?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
