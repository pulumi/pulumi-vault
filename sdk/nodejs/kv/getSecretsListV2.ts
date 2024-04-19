// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kvv2 = new vault.Mount("kvv2", {
 *     path: "kvv2",
 *     type: "kv",
 *     options: {
 *         version: "2",
 *     },
 *     description: "KV Version 2 secret engine mount",
 * });
 * const awsSecret = new vault.kv.SecretV2("awsSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         zip: "zap",
 *     }),
 * });
 * const azureSecret = new vault.kv.SecretV2("azureSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         foo: "bar",
 *     }),
 * });
 * const nestedSecret = new vault.kv.SecretV2("nestedSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         password: "test",
 *     }),
 * });
 * const secrets = vault.kv.getSecretsListV2Output({
 *     mount: kvv2.path,
 * });
 * const nestedSecrets = kvv2.path.apply(path => vault.kv.getSecretsListV2Output({
 *     mount: path,
 *     name: vault_kv_secret_v2.test_2.name,
 * }));
 * ```
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `read` capability on the given path.
 */
export function getSecretsListV2(args: GetSecretsListV2Args, opts?: pulumi.InvokeOptions): Promise<GetSecretsListV2Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:kv/getSecretsListV2:getSecretsListV2", {
        "mount": args.mount,
        "name": args.name,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretsListV2.
 */
export interface GetSecretsListV2Args {
    /**
     * Path where KV-V2 engine is mounted.
     */
    mount: string;
    /**
     * Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     */
    name?: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
}

/**
 * A collection of values returned by getSecretsListV2.
 */
export interface GetSecretsListV2Result {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mount: string;
    readonly name?: string;
    /**
     * List of all secret names listed under the given path.
     */
    readonly names: string[];
    readonly namespace?: string;
    /**
     * Full path where the KV-V2 secrets are listed.
     */
    readonly path: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kvv2 = new vault.Mount("kvv2", {
 *     path: "kvv2",
 *     type: "kv",
 *     options: {
 *         version: "2",
 *     },
 *     description: "KV Version 2 secret engine mount",
 * });
 * const awsSecret = new vault.kv.SecretV2("awsSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         zip: "zap",
 *     }),
 * });
 * const azureSecret = new vault.kv.SecretV2("azureSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         foo: "bar",
 *     }),
 * });
 * const nestedSecret = new vault.kv.SecretV2("nestedSecret", {
 *     mount: kvv2.path,
 *     dataJson: JSON.stringify({
 *         password: "test",
 *     }),
 * });
 * const secrets = vault.kv.getSecretsListV2Output({
 *     mount: kvv2.path,
 * });
 * const nestedSecrets = kvv2.path.apply(path => vault.kv.getSecretsListV2Output({
 *     mount: path,
 *     name: vault_kv_secret_v2.test_2.name,
 * }));
 * ```
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `read` capability on the given path.
 */
export function getSecretsListV2Output(args: GetSecretsListV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretsListV2Result> {
    return pulumi.output(args).apply((a: any) => getSecretsListV2(a, opts))
}

/**
 * A collection of arguments for invoking getSecretsListV2.
 */
export interface GetSecretsListV2OutputArgs {
    /**
     * Path where KV-V2 engine is mounted.
     */
    mount: pulumi.Input<string>;
    /**
     * Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
