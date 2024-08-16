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
 * const gcp = new vault.secrets.SyncGcpDestination("gcp", {
 *     name: "gcp-dest",
 *     projectId: "gcp-project-id",
 *     credentials: std.file({
 *         input: credentialsFile,
 *     }).then(invoke => invoke.result),
 *     secretNameTemplate: "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
 *     customTags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * GCP Secrets sync destinations can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:secrets/syncGcpDestination:SyncGcpDestination gcp gcp-dest
 * ```
 */
export class SyncGcpDestination extends pulumi.CustomResource {
    /**
     * Get an existing SyncGcpDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncGcpDestinationState, opts?: pulumi.CustomResourceOptions): SyncGcpDestination {
        return new SyncGcpDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:secrets/syncGcpDestination:SyncGcpDestination';

    /**
     * Returns true if the given object is an instance of SyncGcpDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncGcpDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncGcpDestination.__pulumiType;
    }

    /**
     * JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * Custom tags to set on the secret managed at the destination.
     */
    public readonly customTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    public readonly granularity!: pulumi.Output<string | undefined>;
    /**
     * Unique name of the GCP destination.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The target project to manage secrets in. If set,
     * overrides the project ID derived from the service account JSON credentials or application
     * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
     * to perform Secret Manager actions in the target project.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    public readonly secretNameTemplate!: pulumi.Output<string>;
    /**
     * The type of the secrets destination (`gcp-sm`).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SyncGcpDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SyncGcpDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncGcpDestinationArgs | SyncGcpDestinationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncGcpDestinationState | undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["customTags"] = state ? state.customTags : undefined;
            resourceInputs["granularity"] = state ? state.granularity : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["secretNameTemplate"] = state ? state.secretNameTemplate : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SyncGcpDestinationArgs | undefined;
            resourceInputs["credentials"] = args?.credentials ? pulumi.secret(args.credentials) : undefined;
            resourceInputs["customTags"] = args ? args.customTags : undefined;
            resourceInputs["granularity"] = args ? args.granularity : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["secretNameTemplate"] = args ? args.secretNameTemplate : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["credentials"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SyncGcpDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncGcpDestination resources.
 */
export interface SyncGcpDestinationState {
    /**
     * JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     */
    credentials?: pulumi.Input<string>;
    /**
     * Custom tags to set on the secret managed at the destination.
     */
    customTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * Unique name of the GCP destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * The target project to manage secrets in. If set,
     * overrides the project ID derived from the service account JSON credentials or application
     * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
     * to perform Secret Manager actions in the target project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
    /**
     * The type of the secrets destination (`gcp-sm`).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncGcpDestination resource.
 */
export interface SyncGcpDestinationArgs {
    /**
     * JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     */
    credentials?: pulumi.Input<string>;
    /**
     * Custom tags to set on the secret managed at the destination.
     */
    customTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * Unique name of the GCP destination.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     */
    namespace?: pulumi.Input<string>;
    /**
     * The target project to manage secrets in. If set,
     * overrides the project ID derived from the service account JSON credentials or application
     * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
     * to perform Secret Manager actions in the target project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     */
    secretNameTemplate?: pulumi.Input<string>;
}
