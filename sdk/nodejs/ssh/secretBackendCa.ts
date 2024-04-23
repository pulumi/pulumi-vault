// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage CA information in an SSH secret backend
 * [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.Mount("example", {type: "ssh"});
 * const foo = new vault.ssh.SecretBackendCa("foo", {backend: example.path});
 * ```
 *
 * ## Import
 *
 * SSH secret backend CAs can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:ssh/secretBackendCa:SecretBackendCa foo ssh
 * ```
 */
export class SecretBackendCa extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendCa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendCaState, opts?: pulumi.CustomResourceOptions): SecretBackendCa {
        return new SecretBackendCa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ssh/secretBackendCa:SecretBackendCa';

    /**
     * Returns true if the given object is an instance of SecretBackendCa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendCa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendCa.__pulumiType;
    }

    /**
     * The path where the SSH secret backend is mounted. Defaults to 'ssh'
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Whether Vault should generate the signing key pair internally. Defaults to true
     */
    public readonly generateSigningKey!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Private key part the SSH CA key pair; required if generate_signing_key is false.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * The public key part the SSH CA key pair; required if generateSigningKey is false.
     */
    public readonly publicKey!: pulumi.Output<string>;

    /**
     * Create a SecretBackendCa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretBackendCaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendCaArgs | SecretBackendCaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendCaState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["generateSigningKey"] = state ? state.generateSigningKey : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
        } else {
            const args = argsOrState as SecretBackendCaArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["generateSigningKey"] = args ? args.generateSigningKey : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackendCa.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendCa resources.
 */
export interface SecretBackendCaState {
    /**
     * The path where the SSH secret backend is mounted. Defaults to 'ssh'
     */
    backend?: pulumi.Input<string>;
    /**
     * Whether Vault should generate the signing key pair internally. Defaults to true
     */
    generateSigningKey?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Private key part the SSH CA key pair; required if generate_signing_key is false.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The public key part the SSH CA key pair; required if generateSigningKey is false.
     */
    publicKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendCa resource.
 */
export interface SecretBackendCaArgs {
    /**
     * The path where the SSH secret backend is mounted. Defaults to 'ssh'
     */
    backend?: pulumi.Input<string>;
    /**
     * Whether Vault should generate the signing key pair internally. Defaults to true
     */
    generateSigningKey?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Private key part the SSH CA key pair; required if generate_signing_key is false.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The public key part the SSH CA key pair; required if generateSigningKey is false.
     */
    publicKey?: pulumi.Input<string>;
}
