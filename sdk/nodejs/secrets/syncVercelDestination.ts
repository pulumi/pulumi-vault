// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const vercel = new vault.secrets.SyncVercelDestination("vercel", {
 *     name: "vercel-dest",
 *     accessToken: accessToken,
 *     projectId: projectId,
 *     deploymentEnvironments: [
 *         "development",
 *         "preview",
 *         "production",
 *     ],
 *     secretNameTemplate: "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Secrets sync destinations can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:secrets/syncVercelDestination:SyncVercelDestination vercel vercel-dest
 * ```
 */
export class SyncVercelDestination extends pulumi.CustomResource {
    /**
     * Get an existing SyncVercelDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncVercelDestinationState, opts?: pulumi.CustomResourceOptions): SyncVercelDestination {
        return new SyncVercelDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:secrets/syncVercelDestination:SyncVercelDestination';

    /**
     * Returns true if the given object is an instance of SyncVercelDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncVercelDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncVercelDestination.__pulumiType;
    }

    /**
     * Vercel API access token with the permissions to manage environment
     * variables.
     */
    public readonly accessToken!: pulumi.Output<string>;
    /**
     * Deployment environments where the environment variables
     * are available. Accepts `development`, `preview` and `production`.
     */
    public readonly deploymentEnvironments!: pulumi.Output<string[]>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    public readonly granularity!: pulumi.Output<string | undefined>;
    /**
     * Unique name of the GitHub destination.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Project ID where to manage environment variables.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    public readonly secretNameTemplate!: pulumi.Output<string>;
    /**
     * Team ID where to manage environment variables.
     */
    public readonly teamId!: pulumi.Output<string | undefined>;
    /**
     * The type of the secrets destination (`vercel-project`).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SyncVercelDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncVercelDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncVercelDestinationArgs | SyncVercelDestinationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncVercelDestinationState | undefined;
            resourceInputs["accessToken"] = state ? state.accessToken : undefined;
            resourceInputs["deploymentEnvironments"] = state ? state.deploymentEnvironments : undefined;
            resourceInputs["granularity"] = state ? state.granularity : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["secretNameTemplate"] = state ? state.secretNameTemplate : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SyncVercelDestinationArgs | undefined;
            if ((!args || args.accessToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessToken'");
            }
            if ((!args || args.deploymentEnvironments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentEnvironments'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["accessToken"] = args?.accessToken ? pulumi.secret(args.accessToken) : undefined;
            resourceInputs["deploymentEnvironments"] = args ? args.deploymentEnvironments : undefined;
            resourceInputs["granularity"] = args ? args.granularity : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["secretNameTemplate"] = args ? args.secretNameTemplate : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SyncVercelDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncVercelDestination resources.
 */
export interface SyncVercelDestinationState {
    /**
     * Vercel API access token with the permissions to manage environment
     * variables.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * Deployment environments where the environment variables
     * are available. Accepts `development`, `preview` and `production`.
     */
    deploymentEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * Unique name of the GitHub destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Project ID where to manage environment variables.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
    /**
     * Team ID where to manage environment variables.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The type of the secrets destination (`vercel-project`).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncVercelDestination resource.
 */
export interface SyncVercelDestinationArgs {
    /**
     * Vercel API access token with the permissions to manage environment
     * variables.
     */
    accessToken: pulumi.Input<string>;
    /**
     * Deployment environments where the environment variables
     * are available. Accepts `development`, `preview` and `production`.
     */
    deploymentEnvironments: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * Unique name of the GitHub destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Project ID where to manage environment variables.
     */
    projectId: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
    /**
     * Team ID where to manage environment variables.
     */
    teamId?: pulumi.Input<string>;
}
