// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure the [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as vault from "@pulumi/vault";
 *
 * const gcp = new vault.gcp.AuthBackend("gcp", {
 *     credentials: fs.readFileSync("vault-gcp-credentials.json", "utf-8"),
 * });
 * ```
 *
 * ## Import
 *
 * GCP authentication backends can be imported using the backend name, e.g.
 *
 * ```sh
 *  $ pulumi import vault:gcp/authBackend:AuthBackend gcp gcp
 * ```
 */
export class AuthBackend extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendState, opts?: pulumi.CustomResourceOptions): AuthBackend {
        return new AuthBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/authBackend:AuthBackend';

    /**
     * Returns true if the given object is an instance of AuthBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackend.__pulumiType;
    }

    /**
     * The clients email associated with the credentials
     */
    public readonly clientEmail!: pulumi.Output<string>;
    /**
     * The Client ID of the credentials
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * A description of the auth method.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the auth method is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The ID of the private key from the credentials
     */
    public readonly privateKeyId!: pulumi.Output<string>;
    /**
     * The GCP Project ID
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            inputs["clientEmail"] = state ? state.clientEmail : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["local"] = state ? state.local : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["privateKeyId"] = state ? state.privateKeyId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            inputs["clientEmail"] = args ? args.clientEmail : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["local"] = args ? args.local : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["privateKeyId"] = args ? args.privateKeyId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The clients email associated with the credentials
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * The Client ID of the credentials
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    readonly credentials?: pulumi.Input<string>;
    /**
     * A description of the auth method.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The ID of the private key from the credentials
     */
    readonly privateKeyId?: pulumi.Input<string>;
    /**
     * The GCP Project ID
     */
    readonly projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The clients email associated with the credentials
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * The Client ID of the credentials
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     */
    readonly credentials?: pulumi.Input<string>;
    /**
     * A description of the auth method.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * The path to mount the auth method — this defaults to 'gcp'.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The ID of the private key from the credentials
     */
    readonly privateKeyId?: pulumi.Input<string>;
    /**
     * The GCP Project ID
     */
    readonly projectId?: pulumi.Input<string>;
}
