// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const azure = new vault.azure.Backend("azure", {
 *     subscriptionId: _var.subscription_id,
 *     tenantId: _var.tenant_id,
 *     clientSecret: _var.client_secret,
 *     clientId: _var.client_id,
 * });
 * const generatedRole = new vault.azure.BackendRole("generatedRole", {
 *     backend: azure.path,
 *     role: "generated_role",
 *     ttl: 300,
 *     maxTtl: 600,
 *     azureRoles: [{
 *         roleName: "Reader",
 *         scope: `/subscriptions/${_var.subscription_id}/resourceGroups/azure-vault-group`,
 *     }],
 * });
 * const existingObjectId = new vault.azure.BackendRole("existingObjectId", {
 *     backend: azure.path,
 *     role: "existing_object_id",
 *     applicationObjectId: "11111111-2222-3333-4444-44444444444",
 *     ttl: 300,
 *     maxTtl: 600,
 * });
 * ```
 */
export class BackendRole extends pulumi.CustomResource {
    /**
     * Get an existing BackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendRoleState, opts?: pulumi.CustomResourceOptions): BackendRole {
        return new BackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:azure/backendRole:BackendRole';

    /**
     * Returns true if the given object is an instance of BackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendRole.__pulumiType;
    }

    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    public readonly applicationObjectId!: pulumi.Output<string | undefined>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    public readonly azureGroups!: pulumi.Output<outputs.azure.BackendRoleAzureGroup[] | undefined>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    public readonly azureRoles!: pulumi.Output<outputs.azure.BackendRoleAzureRole[] | undefined>;
    /**
     * Path to the mounted Azure auth backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * Name of the Azure role
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    public readonly ttl!: pulumi.Output<string | undefined>;

    /**
     * Create a BackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendRoleArgs | BackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendRoleState | undefined;
            resourceInputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            resourceInputs["azureGroups"] = state ? state.azureGroups : undefined;
            resourceInputs["azureRoles"] = state ? state.azureRoles : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as BackendRoleArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            resourceInputs["azureGroups"] = args ? args.azureGroups : undefined;
            resourceInputs["azureRoles"] = args ? args.azureRoles : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendRole resources.
 */
export interface BackendRoleState {
    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    azureGroups?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureGroup>[]>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    azureRoles?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureRole>[]>;
    /**
     * Path to the mounted Azure auth backend
     */
    backend?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * Name of the Azure role
     */
    role?: pulumi.Input<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendRole resource.
 */
export interface BackendRoleArgs {
    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    azureGroups?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureGroup>[]>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    azureRoles?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureRole>[]>;
    /**
     * Path to the mounted Azure auth backend
     */
    backend?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * Name of the Azure role
     */
    role: pulumi.Input<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    ttl?: pulumi.Input<string>;
}
