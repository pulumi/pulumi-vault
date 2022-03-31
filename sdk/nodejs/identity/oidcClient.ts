// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages OIDC Clients in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
 * for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const testOidcAssignment = new vault.identity.OidcAssignment("testOidcAssignment", {
 *     entityIds: ["ascbascas-2231a-sdfaa"],
 *     groupIds: ["sajkdsad-32414-sfsada"],
 * });
 * const testOidcClient = new vault.identity.OidcClient("testOidcClient", {
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     assignments: [testOidcAssignment.name],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * ```
 *
 * ## Import
 *
 * OIDC Clients can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:identity/oidcClient:OidcClient test my-app
 * ```
 */
export class OidcClient extends pulumi.CustomResource {
    /**
     * Get an existing OidcClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcClientState, opts?: pulumi.CustomResourceOptions): OidcClient {
        return new OidcClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/oidcClient:OidcClient';

    /**
     * Returns true if the given object is an instance of OidcClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcClient.__pulumiType;
    }

    /**
     * The time-to-live for access tokens obtained by the client.
     */
    public readonly accessTokenTtl!: pulumi.Output<number>;
    /**
     * A list of assignment resources associated with the client.
     */
    public readonly assignments!: pulumi.Output<string[] | undefined>;
    /**
     * The Client ID from Vault.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * The Client Secret from Vault.
     */
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    /**
     * The client type based on its ability to maintain confidentiality of credentials.
     * The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
     */
    public readonly clientType!: pulumi.Output<string>;
    /**
     * The time-to-live for ID tokens obtained by the client. 
     * The value should be less than the `verificationTtl` on the key.
     */
    public readonly idTokenTtl!: pulumi.Output<number>;
    /**
     * A reference to a named key resource in Vault.
     * This cannot be modified after creation. If not provided, the `default`
     * key is used.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The name of the client.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Redirection URI values used by the client. 
     * One of these values must exactly match the `redirectUri` parameter value
     * used in each authentication request.
     */
    public readonly redirectUris!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OidcClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OidcClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcClientArgs | OidcClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OidcClientState | undefined;
            resourceInputs["accessTokenTtl"] = state ? state.accessTokenTtl : undefined;
            resourceInputs["assignments"] = state ? state.assignments : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["clientType"] = state ? state.clientType : undefined;
            resourceInputs["idTokenTtl"] = state ? state.idTokenTtl : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redirectUris"] = state ? state.redirectUris : undefined;
        } else {
            const args = argsOrState as OidcClientArgs | undefined;
            resourceInputs["accessTokenTtl"] = args ? args.accessTokenTtl : undefined;
            resourceInputs["assignments"] = args ? args.assignments : undefined;
            resourceInputs["clientType"] = args ? args.clientType : undefined;
            resourceInputs["idTokenTtl"] = args ? args.idTokenTtl : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redirectUris"] = args ? args.redirectUris : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["clientSecret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OidcClient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OidcClient resources.
 */
export interface OidcClientState {
    /**
     * The time-to-live for access tokens obtained by the client.
     */
    accessTokenTtl?: pulumi.Input<number>;
    /**
     * A list of assignment resources associated with the client.
     */
    assignments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Client ID from Vault.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The Client Secret from Vault.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The client type based on its ability to maintain confidentiality of credentials.
     * The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
     */
    clientType?: pulumi.Input<string>;
    /**
     * The time-to-live for ID tokens obtained by the client. 
     * The value should be less than the `verificationTtl` on the key.
     */
    idTokenTtl?: pulumi.Input<number>;
    /**
     * A reference to a named key resource in Vault.
     * This cannot be modified after creation. If not provided, the `default`
     * key is used.
     */
    key?: pulumi.Input<string>;
    /**
     * The name of the client.
     */
    name?: pulumi.Input<string>;
    /**
     * Redirection URI values used by the client. 
     * One of these values must exactly match the `redirectUri` parameter value
     * used in each authentication request.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OidcClient resource.
 */
export interface OidcClientArgs {
    /**
     * The time-to-live for access tokens obtained by the client.
     */
    accessTokenTtl?: pulumi.Input<number>;
    /**
     * A list of assignment resources associated with the client.
     */
    assignments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The client type based on its ability to maintain confidentiality of credentials.
     * The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
     */
    clientType?: pulumi.Input<string>;
    /**
     * The time-to-live for ID tokens obtained by the client. 
     * The value should be less than the `verificationTtl` on the key.
     */
    idTokenTtl?: pulumi.Input<number>;
    /**
     * A reference to a named key resource in Vault.
     * This cannot be modified after creation. If not provided, the `default`
     * key is used.
     */
    key?: pulumi.Input<string>;
    /**
     * The name of the client.
     */
    name?: pulumi.Input<string>;
    /**
     * Redirection URI values used by the client. 
     * One of these values must exactly match the `redirectUri` parameter value
     * used in each authentication request.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}
