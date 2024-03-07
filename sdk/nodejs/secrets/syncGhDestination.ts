// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const gh = new vault.secrets.SyncGhDestination("gh", {
 *     accessToken: _var.access_token,
 *     repositoryOwner: _var.repo_owner,
 *     repositoryName: "repo-name-example",
 *     secretNameTemplate: "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GitHub Secrets sync destinations can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:secrets/syncGhDestination:SyncGhDestination gh gh-dest
 * ```
 */
export class SyncGhDestination extends pulumi.CustomResource {
    /**
     * Get an existing SyncGhDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncGhDestinationState, opts?: pulumi.CustomResourceOptions): SyncGhDestination {
        return new SyncGhDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:secrets/syncGhDestination:SyncGhDestination';

    /**
     * Returns true if the given object is an instance of SyncGhDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncGhDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncGhDestination.__pulumiType;
    }

    /**
     * Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     */
    public readonly accessToken!: pulumi.Output<string | undefined>;
    /**
     * Unique name of the GitHub destination.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     */
    public readonly repositoryName!: pulumi.Output<string | undefined>;
    /**
     * GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     */
    public readonly repositoryOwner!: pulumi.Output<string | undefined>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    public readonly secretNameTemplate!: pulumi.Output<string>;
    /**
     * The type of the secrets destination (`gh`).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SyncGhDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SyncGhDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncGhDestinationArgs | SyncGhDestinationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncGhDestinationState | undefined;
            resourceInputs["accessToken"] = state ? state.accessToken : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["repositoryName"] = state ? state.repositoryName : undefined;
            resourceInputs["repositoryOwner"] = state ? state.repositoryOwner : undefined;
            resourceInputs["secretNameTemplate"] = state ? state.secretNameTemplate : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SyncGhDestinationArgs | undefined;
            resourceInputs["accessToken"] = args?.accessToken ? pulumi.secret(args.accessToken) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["repositoryName"] = args ? args.repositoryName : undefined;
            resourceInputs["repositoryOwner"] = args ? args.repositoryOwner : undefined;
            resourceInputs["secretNameTemplate"] = args ? args.secretNameTemplate : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SyncGhDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncGhDestination resources.
 */
export interface SyncGhDestinationState {
    /**
     * Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * Unique name of the GitHub destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     */
    repositoryName?: pulumi.Input<string>;
    /**
     * GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     */
    repositoryOwner?: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
    /**
     * The type of the secrets destination (`gh`).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncGhDestination resource.
 */
export interface SyncGhDestinationArgs {
    /**
     * Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * Unique name of the GitHub destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     */
    repositoryName?: pulumi.Input<string>;
    /**
     * GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     */
    repositoryOwner?: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
}
