// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Static Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
 *
 * Each [static account](https://www.vaultproject.io/docs/secrets/gcp/index.html#static-accounts) is tied to a separately managed
 * Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings) associated with it.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * import * as vault from "@pulumi/vault";
 *
 * const _this = new gcp.serviceaccount.Account("this", {accountId: "my-awesome-account"});
 * const gcp = new vault.gcp.SecretBackend("gcp", {
 *     path: "gcp",
 *     credentials: fs.readFileSync("credentials.json", "utf8"),
 * });
 * const staticAccount = new vault.gcp.SecretStaticAccount("staticAccount", {
 *     backend: gcp.path,
 *     staticAccount: "project_viewer",
 *     secretType: "access_token",
 *     tokenScopes: ["https://www.googleapis.com/auth/cloud-platform"],
 *     serviceAccountEmail: _this.email,
 *     bindings: [{
 *         resource: pulumi.interpolate`//cloudresourcemanager.googleapis.com/projects/${_this.project}`,
 *         roles: ["roles/viewer"],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * A static account can be imported using its Vault Path. For example, referencing the example above,
 *
 * ```sh
 * $ pulumi import vault:gcp/secretStaticAccount:SecretStaticAccount static_account gcp/static-account/project_viewer
 * ```
 */
export class SecretStaticAccount extends pulumi.CustomResource {
    /**
     * Get an existing SecretStaticAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretStaticAccountState, opts?: pulumi.CustomResourceOptions): SecretStaticAccount {
        return new SecretStaticAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/secretStaticAccount:SecretStaticAccount';

    /**
     * Returns true if the given object is an instance of SecretStaticAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretStaticAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretStaticAccount.__pulumiType;
    }

    /**
     * Path where the GCP Secrets Engine is mounted
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    public readonly bindings!: pulumi.Output<outputs.gcp.SecretStaticAccountBinding[] | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    public readonly secretType!: pulumi.Output<string>;
    /**
     * Email of the GCP service account to manage.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * Project the service account belongs to.
     */
    public /*out*/ readonly serviceAccountProject!: pulumi.Output<string>;
    /**
     * Name of the Static Account to create
     */
    public readonly staticAccount!: pulumi.Output<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
     */
    public readonly tokenScopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretStaticAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretStaticAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretStaticAccountArgs | SecretStaticAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretStaticAccountState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["bindings"] = state ? state.bindings : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["secretType"] = state ? state.secretType : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["serviceAccountProject"] = state ? state.serviceAccountProject : undefined;
            resourceInputs["staticAccount"] = state ? state.staticAccount : undefined;
            resourceInputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as SecretStaticAccountArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.serviceAccountEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountEmail'");
            }
            if ((!args || args.staticAccount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'staticAccount'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["bindings"] = args ? args.bindings : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["secretType"] = args ? args.secretType : undefined;
            resourceInputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            resourceInputs["staticAccount"] = args ? args.staticAccount : undefined;
            resourceInputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            resourceInputs["serviceAccountProject"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretStaticAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretStaticAccount resources.
 */
export interface SecretStaticAccountState {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend?: pulumi.Input<string>;
    /**
     * Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    bindings?: pulumi.Input<pulumi.Input<inputs.gcp.SecretStaticAccountBinding>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    secretType?: pulumi.Input<string>;
    /**
     * Email of the GCP service account to manage.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * Project the service account belongs to.
     */
    serviceAccountProject?: pulumi.Input<string>;
    /**
     * Name of the Static Account to create
     */
    staticAccount?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretStaticAccount resource.
 */
export interface SecretStaticAccountArgs {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend: pulumi.Input<string>;
    /**
     * Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    bindings?: pulumi.Input<pulumi.Input<inputs.gcp.SecretStaticAccountBinding>[]>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    secretType?: pulumi.Input<string>;
    /**
     * Email of the GCP service account to manage.
     */
    serviceAccountEmail: pulumi.Input<string>;
    /**
     * Name of the Static Account to create
     */
    staticAccount: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}
