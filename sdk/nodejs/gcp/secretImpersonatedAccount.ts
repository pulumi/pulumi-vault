// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
 *
 * Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
 * Service Account.
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
 * const impersonatedAccount = new vault.gcp.SecretImpersonatedAccount("impersonatedAccount", {
 *     backend: gcp.path,
 *     impersonatedAccount: "this",
 *     serviceAccountEmail: _this.email,
 *     tokenScopes: ["https://www.googleapis.com/auth/cloud-platform"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * A impersonated account can be imported using its Vault Path. For example, referencing the example above,
 *
 * ```sh
 * $ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
 * ```
 */
export class SecretImpersonatedAccount extends pulumi.CustomResource {
    /**
     * Get an existing SecretImpersonatedAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretImpersonatedAccountState, opts?: pulumi.CustomResourceOptions): SecretImpersonatedAccount {
        return new SecretImpersonatedAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount';

    /**
     * Returns true if the given object is an instance of SecretImpersonatedAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretImpersonatedAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretImpersonatedAccount.__pulumiType;
    }

    /**
     * Path where the GCP Secrets Engine is mounted
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Name of the Impersonated Account to create
     */
    public readonly impersonatedAccount!: pulumi.Output<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Email of the GCP service account to impersonate.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * Project the service account belongs to.
     */
    public /*out*/ readonly serviceAccountProject!: pulumi.Output<string>;
    /**
     * List of OAuth scopes to assign to access tokens generated under this impersonated account.
     */
    public readonly tokenScopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretImpersonatedAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretImpersonatedAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretImpersonatedAccountArgs | SecretImpersonatedAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretImpersonatedAccountState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["impersonatedAccount"] = state ? state.impersonatedAccount : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["serviceAccountProject"] = state ? state.serviceAccountProject : undefined;
            resourceInputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as SecretImpersonatedAccountArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.impersonatedAccount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'impersonatedAccount'");
            }
            if ((!args || args.serviceAccountEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountEmail'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["impersonatedAccount"] = args ? args.impersonatedAccount : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            resourceInputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            resourceInputs["serviceAccountProject"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretImpersonatedAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretImpersonatedAccount resources.
 */
export interface SecretImpersonatedAccountState {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend?: pulumi.Input<string>;
    /**
     * Name of the Impersonated Account to create
     */
    impersonatedAccount?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Email of the GCP service account to impersonate.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * Project the service account belongs to.
     */
    serviceAccountProject?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to access tokens generated under this impersonated account.
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretImpersonatedAccount resource.
 */
export interface SecretImpersonatedAccountArgs {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    backend: pulumi.Input<string>;
    /**
     * Name of the Impersonated Account to create
     */
    impersonatedAccount: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Email of the GCP service account to impersonate.
     */
    serviceAccountEmail: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to access tokens generated under this impersonated account.
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}
