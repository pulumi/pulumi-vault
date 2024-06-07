// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const github_apps = new vault.secrets.SyncGithubApps("github-apps", {
 *     name: "gh-apps",
 *     appId: appId,
 *     privateKey: std.file({
 *         input: privatekeyFile,
 *     }).then(invoke => invoke.result),
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Apps Secrets sync configuration endpoint can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:secrets/syncGithubApps:SyncGithubApps gh github-apps
 * ```
 */
export class SyncGithubApps extends pulumi.CustomResource {
    /**
     * Get an existing SyncGithubApps resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncGithubAppsState, opts?: pulumi.CustomResourceOptions): SyncGithubApps {
        return new SyncGithubApps(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:secrets/syncGithubApps:SyncGithubApps';

    /**
     * Returns true if the given object is an instance of SyncGithubApps.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncGithubApps {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncGithubApps.__pulumiType;
    }

    /**
     * The GitHub application ID.
     */
    public readonly appId!: pulumi.Output<number>;
    /**
     * A fingerprint of a private key.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The user-defined name of the GitHub App configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The content of a PEM formatted private key generated on GitHub for the app.
     */
    public readonly privateKey!: pulumi.Output<string>;

    /**
     * Create a SyncGithubApps resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncGithubAppsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncGithubAppsArgs | SyncGithubAppsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncGithubAppsState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
        } else {
            const args = argsOrState as SyncGithubAppsArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SyncGithubApps.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncGithubApps resources.
 */
export interface SyncGithubAppsState {
    /**
     * The GitHub application ID.
     */
    appId?: pulumi.Input<number>;
    /**
     * A fingerprint of a private key.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The user-defined name of the GitHub App configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * The content of a PEM formatted private key generated on GitHub for the app.
     */
    privateKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncGithubApps resource.
 */
export interface SyncGithubAppsArgs {
    /**
     * The GitHub application ID.
     */
    appId: pulumi.Input<number>;
    /**
     * The user-defined name of the GitHub App configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * The content of a PEM formatted private key generated on GitHub for the app.
     */
    privateKey: pulumi.Input<string>;
}
