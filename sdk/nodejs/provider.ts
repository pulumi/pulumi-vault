// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the vault package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'vault';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    public readonly addAddressToEnv!: pulumi.Output<string | undefined>;
    /**
     * URL of the root of the target Vault server.
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * Path to directory containing CA certificate files to validate the server's certificate.
     */
    public readonly caCertDir!: pulumi.Output<string | undefined>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     */
    public readonly caCertFile!: pulumi.Output<string | undefined>;
    /**
     * The namespace to use. Available only for Vault Enterprise.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Name to use as the SNI host when connecting via TLS.
     */
    public readonly tlsServerName!: pulumi.Output<string | undefined>;
    /**
     * Token to use to authenticate to Vault.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Token name to use for creating the Vault child token.
     */
    public readonly tokenName!: pulumi.Output<string | undefined>;
    /**
     * Override the Vault server version, which is normally determined dynamically from the target Vault server
     */
    public readonly vaultVersionOverride!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["addAddressToEnv"] = args ? args.addAddressToEnv : undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["authLogin"] = pulumi.output(args ? args.authLogin : undefined).apply(JSON.stringify);
            resourceInputs["authLoginAws"] = pulumi.output(args ? args.authLoginAws : undefined).apply(JSON.stringify);
            resourceInputs["authLoginAzure"] = pulumi.output(args ? args.authLoginAzure : undefined).apply(JSON.stringify);
            resourceInputs["authLoginCert"] = pulumi.output(args ? args.authLoginCert : undefined).apply(JSON.stringify);
            resourceInputs["authLoginGcp"] = pulumi.output(args ? args.authLoginGcp : undefined).apply(JSON.stringify);
            resourceInputs["authLoginJwt"] = pulumi.output(args ? args.authLoginJwt : undefined).apply(JSON.stringify);
            resourceInputs["authLoginKerberos"] = pulumi.output(args ? args.authLoginKerberos : undefined).apply(JSON.stringify);
            resourceInputs["authLoginOci"] = pulumi.output(args ? args.authLoginOci : undefined).apply(JSON.stringify);
            resourceInputs["authLoginOidc"] = pulumi.output(args ? args.authLoginOidc : undefined).apply(JSON.stringify);
            resourceInputs["authLoginRadius"] = pulumi.output(args ? args.authLoginRadius : undefined).apply(JSON.stringify);
            resourceInputs["authLoginTokenFile"] = pulumi.output(args ? args.authLoginTokenFile : undefined).apply(JSON.stringify);
            resourceInputs["authLoginUserpass"] = pulumi.output(args ? args.authLoginUserpass : undefined).apply(JSON.stringify);
            resourceInputs["caCertDir"] = args ? args.caCertDir : undefined;
            resourceInputs["caCertFile"] = args ? args.caCertFile : undefined;
            resourceInputs["clientAuth"] = pulumi.output(args ? args.clientAuth : undefined).apply(JSON.stringify);
            resourceInputs["headers"] = pulumi.output(args ? args.headers : undefined).apply(JSON.stringify);
            resourceInputs["maxLeaseTtlSeconds"] = pulumi.output((args ? args.maxLeaseTtlSeconds : undefined) ?? (utilities.getEnvNumber("TERRAFORM_VAULT_MAX_TTL") || 1200)).apply(JSON.stringify);
            resourceInputs["maxRetries"] = pulumi.output((args ? args.maxRetries : undefined) ?? (utilities.getEnvNumber("VAULT_MAX_RETRIES") || 2)).apply(JSON.stringify);
            resourceInputs["maxRetriesCcc"] = pulumi.output(args ? args.maxRetriesCcc : undefined).apply(JSON.stringify);
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["setNamespaceFromToken"] = pulumi.output(args ? args.setNamespaceFromToken : undefined).apply(JSON.stringify);
            resourceInputs["skipChildToken"] = pulumi.output(args ? args.skipChildToken : undefined).apply(JSON.stringify);
            resourceInputs["skipGetVaultVersion"] = pulumi.output(args ? args.skipGetVaultVersion : undefined).apply(JSON.stringify);
            resourceInputs["skipTlsVerify"] = pulumi.output((args ? args.skipTlsVerify : undefined) ?? utilities.getEnvBoolean("VAULT_SKIP_VERIFY")).apply(JSON.stringify);
            resourceInputs["tlsServerName"] = args ? args.tlsServerName : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["tokenName"] = args ? args.tokenName : undefined;
            resourceInputs["vaultVersionOverride"] = args ? args.vaultVersionOverride : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }

    /**
     * This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
     */
    terraformConfig(): pulumi.Output<Provider.TerraformConfigResult> {
        return pulumi.runtime.call("pulumi:providers:vault/terraformConfig", {
            "__self__": this,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    addAddressToEnv?: pulumi.Input<string>;
    /**
     * URL of the root of the target Vault server.
     */
    address?: pulumi.Input<string>;
    /**
     * Login to vault with an existing auth method using auth/<mount>/login
     */
    authLogin?: pulumi.Input<inputs.ProviderAuthLogin>;
    /**
     * Login to vault using the AWS method
     */
    authLoginAws?: pulumi.Input<inputs.ProviderAuthLoginAws>;
    /**
     * Login to vault using the azure method
     */
    authLoginAzure?: pulumi.Input<inputs.ProviderAuthLoginAzure>;
    /**
     * Login to vault using the cert method
     */
    authLoginCert?: pulumi.Input<inputs.ProviderAuthLoginCert>;
    /**
     * Login to vault using the gcp method
     */
    authLoginGcp?: pulumi.Input<inputs.ProviderAuthLoginGcp>;
    /**
     * Login to vault using the jwt method
     */
    authLoginJwt?: pulumi.Input<inputs.ProviderAuthLoginJwt>;
    /**
     * Login to vault using the kerberos method
     */
    authLoginKerberos?: pulumi.Input<inputs.ProviderAuthLoginKerberos>;
    /**
     * Login to vault using the OCI method
     */
    authLoginOci?: pulumi.Input<inputs.ProviderAuthLoginOci>;
    /**
     * Login to vault using the oidc method
     */
    authLoginOidc?: pulumi.Input<inputs.ProviderAuthLoginOidc>;
    /**
     * Login to vault using the radius method
     */
    authLoginRadius?: pulumi.Input<inputs.ProviderAuthLoginRadius>;
    /**
     * Login to vault using
     */
    authLoginTokenFile?: pulumi.Input<inputs.ProviderAuthLoginTokenFile>;
    /**
     * Login to vault using the userpass method
     */
    authLoginUserpass?: pulumi.Input<inputs.ProviderAuthLoginUserpass>;
    /**
     * Path to directory containing CA certificate files to validate the server's certificate.
     */
    caCertDir?: pulumi.Input<string>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     */
    caCertFile?: pulumi.Input<string>;
    /**
     * Client authentication credentials.
     */
    clientAuth?: pulumi.Input<inputs.ProviderClientAuth>;
    /**
     * The headers to send with each Vault request.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.ProviderHeader>[]>;
    /**
     * Maximum TTL for secret leases requested by this provider.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Maximum number of retries when a 5xx error code is encountered.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Maximum number of retries for Client Controlled Consistency related operations
     */
    maxRetriesCcc?: pulumi.Input<number>;
    /**
     * The namespace to use. Available only for Vault Enterprise.
     */
    namespace?: pulumi.Input<string>;
    /**
     * In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
     * token namespace as the root namespace for all resources.
     */
    setNamespaceFromToken?: pulumi.Input<boolean>;
    /**
     * Set this to true to prevent the creation of ephemeral child token used by this provider.
     */
    skipChildToken?: pulumi.Input<boolean>;
    /**
     * Skip the dynamic fetching of the Vault server version.
     */
    skipGetVaultVersion?: pulumi.Input<boolean>;
    /**
     * Set this to true only if the target Vault server is an insecure development instance.
     */
    skipTlsVerify?: pulumi.Input<boolean>;
    /**
     * Name to use as the SNI host when connecting via TLS.
     */
    tlsServerName?: pulumi.Input<string>;
    /**
     * Token to use to authenticate to Vault.
     */
    token?: pulumi.Input<string>;
    /**
     * Token name to use for creating the Vault child token.
     */
    tokenName?: pulumi.Input<string>;
    /**
     * Override the Vault server version, which is normally determined dynamically from the target Vault server
     */
    vaultVersionOverride?: pulumi.Input<string>;
}

export namespace Provider {
    /**
     * The results of the Provider.terraformConfig method.
     */
    export interface TerraformConfigResult {
        readonly result: {[key: string]: any};
    }

}
