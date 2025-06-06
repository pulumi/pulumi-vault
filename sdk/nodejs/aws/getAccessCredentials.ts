// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 */
export function getAccessCredentials(args: GetAccessCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:aws/getAccessCredentials:getAccessCredentials", {
        "backend": args.backend,
        "namespace": args.namespace,
        "region": args.region,
        "role": args.role,
        "roleArn": args.roleArn,
        "ttl": args.ttl,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsArgs {
    /**
     * The path to the AWS secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The region the read credentials belong to.
     */
    region?: string;
    /**
     * The name of the AWS secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: string;
    /**
     * The specific AWS ARN to use
     * from the configured role. If the role does not have multiple ARNs, this does
     * not need to be specified.
     */
    roleArn?: string;
    /**
     * Specifies the TTL for the use of the STS token. This
     * is specified as a string with a duration suffix. Valid only when
     * `credentialType` of the connected `vault.aws.SecretBackendRole` resource is `assumedRole` or `federationToken`
     */
    ttl?: string;
    /**
     * The type of credentials to read. Defaults
     * to `"creds"`, which just returns an AWS Access Key ID and Secret
     * Key. Can also be set to `"sts"`, which will return a security token
     * in addition to the keys.
     */
    type?: string;
}

/**
 * A collection of values returned by getAccessCredentials.
 */
export interface GetAccessCredentialsResult {
    /**
     * The AWS Access Key ID returned by Vault.
     */
    readonly accessKey: string;
    readonly backend: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The duration of the secret lease, in seconds relative
     * to the time the data was requested. Once this time has passed any plan
     * generated with this data may fail to apply.
     */
    readonly leaseDuration: number;
    /**
     * The lease identifier assigned by Vault.
     */
    readonly leaseId: string;
    readonly leaseRenewable: boolean;
    readonly leaseStartTime: string;
    readonly namespace?: string;
    readonly region?: string;
    readonly role: string;
    readonly roleArn?: string;
    /**
     * The AWS Secret Key returned by Vault.
     */
    readonly secretKey: string;
    /**
     * The STS token returned by Vault, if any.
     */
    readonly securityToken: string;
    readonly ttl?: string;
    readonly type?: string;
}
/**
 * ## Example Usage
 */
export function getAccessCredentialsOutput(args: GetAccessCredentialsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccessCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:aws/getAccessCredentials:getAccessCredentials", {
        "backend": args.backend,
        "namespace": args.namespace,
        "region": args.region,
        "role": args.role,
        "roleArn": args.roleArn,
        "ttl": args.ttl,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsOutputArgs {
    /**
     * The path to the AWS secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The region the read credentials belong to.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of the AWS secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: pulumi.Input<string>;
    /**
     * The specific AWS ARN to use
     * from the configured role. If the role does not have multiple ARNs, this does
     * not need to be specified.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Specifies the TTL for the use of the STS token. This
     * is specified as a string with a duration suffix. Valid only when
     * `credentialType` of the connected `vault.aws.SecretBackendRole` resource is `assumedRole` or `federationToken`
     */
    ttl?: pulumi.Input<string>;
    /**
     * The type of credentials to read. Defaults
     * to `"creds"`, which just returns an AWS Access Key ID and Secret
     * Key. Can also be set to `"sts"`, which will return a security token
     * in addition to the keys.
     */
    type?: pulumi.Input<string>;
}
