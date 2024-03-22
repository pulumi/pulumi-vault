// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
 *
 * Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as vault from "@pulumi/vault";
 *
 * const project = "my-awesome-project";
 * const gcp = new vault.gcp.SecretBackend("gcp", {
 *     path: "gcp",
 *     credentials: fs.readFileSync("credentials.json", "utf8"),
 * });
 * const roleset = new vault.gcp.SecretRoleset("roleset", {
 *     backend: gcp.path,
 *     roleset: "project_viewer",
 *     secretType: "access_token",
 *     project: project,
 *     tokenScopes: ["https://www.googleapis.com/auth/cloud-platform"],
 *     bindings: [{
 *         resource: `//cloudresourcemanager.googleapis.com/projects/${project}`,
 *         roles: ["roles/viewer"],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * A roleset can be imported using its Vault Path. For example, referencing the example above,
 *
 * ```sh
 * $ pulumi import vault:gcp/secretRoleset:SecretRoleset roleset gcp/roleset/project_viewer
 * ```
 */
export class SecretRoleset extends pulumi.CustomResource {
    /**
     * Get an existing SecretRoleset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretRolesetState, opts?: pulumi.CustomResourceOptions): SecretRoleset {
        return new SecretRoleset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/secretRoleset:SecretRoleset';

    /**
     * Returns true if the given object is an instance of SecretRoleset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretRoleset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretRoleset.__pulumiType;
    }

    /**
     * Path where the GCP Secrets Engine is mounted
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    public readonly bindings!: pulumi.Output<outputs.gcp.SecretRolesetBinding[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Name of the Roleset to create
     */
    public readonly roleset!: pulumi.Output<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    public readonly secretType!: pulumi.Output<string>;
    /**
     * Email of the service account created by Vault for this Roleset.
     */
    public /*out*/ readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    public readonly tokenScopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretRoleset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretRolesetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretRolesetArgs | SecretRolesetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretRolesetState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["bindings"] = state ? state.bindings : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["roleset"] = state ? state.roleset : undefined;
            resourceInputs["secretType"] = state ? state.secretType : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as SecretRolesetArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.bindings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindings'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.roleset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleset'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["bindings"] = args ? args.bindings : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["roleset"] = args ? args.roleset : undefined;
            resourceInputs["secretType"] = args ? args.secretType : undefined;
            resourceInputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            resourceInputs["serviceAccountEmail"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretRoleset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretRoleset resources.
 */
export interface SecretRolesetState {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend?: pulumi.Input<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    bindings?: pulumi.Input<pulumi.Input<inputs.gcp.SecretRolesetBinding>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    project?: pulumi.Input<string>;
    /**
     * Name of the Roleset to create
     */
    roleset?: pulumi.Input<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    secretType?: pulumi.Input<string>;
    /**
     * Email of the service account created by Vault for this Roleset.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretRoleset resource.
 */
export interface SecretRolesetArgs {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend: pulumi.Input<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    bindings: pulumi.Input<pulumi.Input<inputs.gcp.SecretRolesetBinding>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    project: pulumi.Input<string>;
    /**
     * Name of the Roleset to create
     */
    roleset: pulumi.Input<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    secretType?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}
