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
 * const creds = vault.azure.getAccessCredentials({
 *     role: "my-role",
 *     validateCreds: true,
 *     numSequentialSuccesses: 8,
 *     numSecondsBetweenTests: 1,
 *     maxCredValidationSeconds: 300,
 * });
 * ```
 * ## Caveats
 *
 * The `validateCreds` option requires read-access to the `backend` config endpoint.
 * If the effective Vault role does not have the required permissions then valid values
 * are required to be set for: `subscriptionId`, `tenantId`, `environment`.
 */
export function getAccessCredentials(args: GetAccessCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessCredentialsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:azure/getAccessCredentials:getAccessCredentials", {
        "backend": args.backend,
        "environment": args.environment,
        "maxCredValidationSeconds": args.maxCredValidationSeconds,
        "namespace": args.namespace,
        "numSecondsBetweenTests": args.numSecondsBetweenTests,
        "numSequentialSuccesses": args.numSequentialSuccesses,
        "role": args.role,
        "subscriptionId": args.subscriptionId,
        "tenantId": args.tenantId,
        "validateCreds": args.validateCreds,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsArgs {
    /**
     * The path to the Azure secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * The Azure environment to use during credential validation.
     * Defaults to the environment configured in the Vault backend.
     * Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
     * *See the caveats section for more information on this field.*
     */
    environment?: string;
    /**
     * If 'validate_creds' is true, 
     * the number of seconds after which to give up validating credentials. Defaults
     * to 300.
     */
    maxCredValidationSeconds?: number;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * If 'validate_creds' is true, 
     * the number of seconds to wait between each test of generated credentials.
     * Defaults to 1.
     */
    numSecondsBetweenTests?: number;
    /**
     * If 'validate_creds' is true, 
     * the number of sequential successes required to validate generated
     * credentials. Defaults to 8.
     */
    numSequentialSuccesses?: number;
    /**
     * The name of the Azure secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: string;
    /**
     * The subscription ID to use during credential
     * validation. Defaults to the subscription ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     */
    subscriptionId?: string;
    /**
     * The tenant ID to use during credential validation.
     * Defaults to the tenant ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     */
    tenantId?: string;
    /**
     * Whether generated credentials should be 
     * validated before being returned. Defaults to `false`, which returns
     * credentials without checking whether they have fully propagated throughout
     * Azure Active Directory. Designating `true` activates testing.
     */
    validateCreds?: boolean;
}

/**
 * A collection of values returned by getAccessCredentials.
 */
export interface GetAccessCredentialsResult {
    readonly backend: string;
    /**
     * The client id for credentials to query the Azure APIs.
     */
    readonly clientId: string;
    /**
     * The client secret for credentials to query the Azure APIs.
     */
    readonly clientSecret: string;
    readonly environment?: string;
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
    readonly maxCredValidationSeconds?: number;
    readonly namespace?: string;
    readonly numSecondsBetweenTests?: number;
    readonly numSequentialSuccesses?: number;
    readonly role: string;
    readonly subscriptionId?: string;
    readonly tenantId?: string;
    readonly validateCreds?: boolean;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const creds = vault.azure.getAccessCredentials({
 *     role: "my-role",
 *     validateCreds: true,
 *     numSequentialSuccesses: 8,
 *     numSecondsBetweenTests: 1,
 *     maxCredValidationSeconds: 300,
 * });
 * ```
 * ## Caveats
 *
 * The `validateCreds` option requires read-access to the `backend` config endpoint.
 * If the effective Vault role does not have the required permissions then valid values
 * are required to be set for: `subscriptionId`, `tenantId`, `environment`.
 */
export function getAccessCredentialsOutput(args: GetAccessCredentialsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessCredentialsResult> {
    return pulumi.output(args).apply((a: any) => getAccessCredentials(a, opts))
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsOutputArgs {
    /**
     * The path to the Azure secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The Azure environment to use during credential validation.
     * Defaults to the environment configured in the Vault backend.
     * Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
     * *See the caveats section for more information on this field.*
     */
    environment?: pulumi.Input<string>;
    /**
     * If 'validate_creds' is true, 
     * the number of seconds after which to give up validating credentials. Defaults
     * to 300.
     */
    maxCredValidationSeconds?: pulumi.Input<number>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * If 'validate_creds' is true, 
     * the number of seconds to wait between each test of generated credentials.
     * Defaults to 1.
     */
    numSecondsBetweenTests?: pulumi.Input<number>;
    /**
     * If 'validate_creds' is true, 
     * the number of sequential successes required to validate generated
     * credentials. Defaults to 8.
     */
    numSequentialSuccesses?: pulumi.Input<number>;
    /**
     * The name of the Azure secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: pulumi.Input<string>;
    /**
     * The subscription ID to use during credential
     * validation. Defaults to the subscription ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     */
    subscriptionId?: pulumi.Input<string>;
    /**
     * The tenant ID to use during credential validation.
     * Defaults to the tenant ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Whether generated credentials should be 
     * validated before being returned. Defaults to `false`, which returns
     * credentials without checking whether they have fully propagated throughout
     * Azure Active Directory. Designating `true` activates testing.
     */
    validateCreds?: pulumi.Input<boolean>;
}
