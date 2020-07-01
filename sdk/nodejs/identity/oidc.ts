// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure the [Identity Tokens Backend](https://www.vaultproject.io/docs/secrets/identity/index.html#identity-tokens).
 *
 * The Identity secrets engine is the identity management solution for Vault. It internally maintains
 * the clients who are recognized by Vault.
 *
 * > **NOTE:** Each Vault server may only have one Identity Tokens Backend configuration. Multiple configurations of the resource against the same Vault server will cause a perpetual difference.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const server = new vault.identity.Oidc("server", {
 *     issuer: "https://www.acme.com",
 * });
 * ```
 */
export class Oidc extends pulumi.CustomResource {
    /**
     * Get an existing Oidc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcState, opts?: pulumi.CustomResourceOptions): Oidc {
        return new Oidc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/oidc:Oidc';

    /**
     * Returns true if the given object is an instance of Oidc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Oidc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Oidc.__pulumiType;
    }

    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault's
     * `apiAddr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     */
    public readonly issuer!: pulumi.Output<string>;

    /**
     * Create a Oidc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OidcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcArgs | OidcState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OidcState | undefined;
            inputs["issuer"] = state ? state.issuer : undefined;
        } else {
            const args = argsOrState as OidcArgs | undefined;
            inputs["issuer"] = args ? args.issuer : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Oidc.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Oidc resources.
 */
export interface OidcState {
    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault's
     * `apiAddr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     */
    readonly issuer?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Oidc resource.
 */
export interface OidcArgs {
    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault's
     * `apiAddr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     */
    readonly issuer?: pulumi.Input<string>;
}
