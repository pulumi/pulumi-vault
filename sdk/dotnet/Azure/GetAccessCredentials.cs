// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure
{
    public static class GetAccessCredentials
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var creds = Vault.Azure.GetAccessCredentials.Invoke(new()
        ///     {
        ///         Role = "my-role",
        ///         ValidateCreds = true,
        ///         NumSequentialSuccesses = 8,
        ///         NumSecondsBetweenTests = 1,
        ///         MaxCredValidationSeconds = 300,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Caveats
        /// 
        /// The `validate_creds` option requires read-access to the `backend` config endpoint.
        /// If the effective Vault role does not have the required permissions then valid values 
        /// are required to be set for: `subscription_id`, `tenant_id`, `environment`.
        /// </summary>
        public static Task<GetAccessCredentialsResult> InvokeAsync(GetAccessCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessCredentialsResult>("vault:azure/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var creds = Vault.Azure.GetAccessCredentials.Invoke(new()
        ///     {
        ///         Role = "my-role",
        ///         ValidateCreds = true,
        ///         NumSequentialSuccesses = 8,
        ///         NumSecondsBetweenTests = 1,
        ///         MaxCredValidationSeconds = 300,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Caveats
        /// 
        /// The `validate_creds` option requires read-access to the `backend` config endpoint.
        /// If the effective Vault role does not have the required permissions then valid values 
        /// are required to be set for: `subscription_id`, `tenant_id`, `environment`.
        /// </summary>
        public static Output<GetAccessCredentialsResult> Invoke(GetAccessCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessCredentialsResult>("vault:azure/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var creds = Vault.Azure.GetAccessCredentials.Invoke(new()
        ///     {
        ///         Role = "my-role",
        ///         ValidateCreds = true,
        ///         NumSequentialSuccesses = 8,
        ///         NumSecondsBetweenTests = 1,
        ///         MaxCredValidationSeconds = 300,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Caveats
        /// 
        /// The `validate_creds` option requires read-access to the `backend` config endpoint.
        /// If the effective Vault role does not have the required permissions then valid values 
        /// are required to be set for: `subscription_id`, `tenant_id`, `environment`.
        /// </summary>
        public static Output<GetAccessCredentialsResult> Invoke(GetAccessCredentialsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessCredentialsResult>("vault:azure/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the Azure secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// The Azure environment to use during credential validation.
        /// Defaults to the environment configured in the Vault backend.
        /// Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("environment")]
        public string? Environment { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds after which to give up validating credentials. Defaults
        /// to 300.
        /// </summary>
        [Input("maxCredValidationSeconds")]
        public int? MaxCredValidationSeconds { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds to wait between each test of generated credentials.
        /// Defaults to 1.
        /// </summary>
        [Input("numSecondsBetweenTests")]
        public int? NumSecondsBetweenTests { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of sequential successes required to validate generated
        /// credentials. Defaults to 8.
        /// </summary>
        [Input("numSequentialSuccesses")]
        public int? NumSequentialSuccesses { get; set; }

        /// <summary>
        /// The name of the Azure secret backend role to read
        /// credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public string Role { get; set; } = null!;

        /// <summary>
        /// The subscription ID to use during credential
        /// validation. Defaults to the subscription ID configured in the Vault `backend`.
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        /// <summary>
        /// The tenant ID to use during credential validation.
        /// Defaults to the tenant ID configured in the Vault `backend`.
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        /// <summary>
        /// Whether generated credentials should be 
        /// validated before being returned. Defaults to `false`, which returns
        /// credentials without checking whether they have fully propagated throughout
        /// Azure Active Directory. Designating `true` activates testing.
        /// </summary>
        [Input("validateCreds")]
        public bool? ValidateCreds { get; set; }

        public GetAccessCredentialsArgs()
        {
        }
        public static new GetAccessCredentialsArgs Empty => new GetAccessCredentialsArgs();
    }

    public sealed class GetAccessCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the Azure secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The Azure environment to use during credential validation.
        /// Defaults to the environment configured in the Vault backend.
        /// Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds after which to give up validating credentials. Defaults
        /// to 300.
        /// </summary>
        [Input("maxCredValidationSeconds")]
        public Input<int>? MaxCredValidationSeconds { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds to wait between each test of generated credentials.
        /// Defaults to 1.
        /// </summary>
        [Input("numSecondsBetweenTests")]
        public Input<int>? NumSecondsBetweenTests { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of sequential successes required to validate generated
        /// credentials. Defaults to 8.
        /// </summary>
        [Input("numSequentialSuccesses")]
        public Input<int>? NumSequentialSuccesses { get; set; }

        /// <summary>
        /// The name of the Azure secret backend role to read
        /// credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The subscription ID to use during credential
        /// validation. Defaults to the subscription ID configured in the Vault `backend`.
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The tenant ID to use during credential validation.
        /// Defaults to the tenant ID configured in the Vault `backend`.
        /// *See the caveats section for more information on this field.*
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Whether generated credentials should be 
        /// validated before being returned. Defaults to `false`, which returns
        /// credentials without checking whether they have fully propagated throughout
        /// Azure Active Directory. Designating `true` activates testing.
        /// </summary>
        [Input("validateCreds")]
        public Input<bool>? ValidateCreds { get; set; }

        public GetAccessCredentialsInvokeArgs()
        {
        }
        public static new GetAccessCredentialsInvokeArgs Empty => new GetAccessCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessCredentialsResult
    {
        public readonly string Backend;
        /// <summary>
        /// The client id for credentials to query the Azure APIs.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The client secret for credentials to query the Azure APIs.
        /// </summary>
        public readonly string ClientSecret;
        public readonly string? Environment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The duration of the secret lease, in seconds relative
        /// to the time the data was requested. Once this time has passed any plan
        /// generated with this data may fail to apply.
        /// </summary>
        public readonly int LeaseDuration;
        /// <summary>
        /// The lease identifier assigned by Vault.
        /// </summary>
        public readonly string LeaseId;
        public readonly bool LeaseRenewable;
        public readonly string LeaseStartTime;
        public readonly int? MaxCredValidationSeconds;
        public readonly string? Namespace;
        public readonly int? NumSecondsBetweenTests;
        public readonly int? NumSequentialSuccesses;
        public readonly string Role;
        public readonly string? SubscriptionId;
        public readonly string? TenantId;
        public readonly bool? ValidateCreds;

        [OutputConstructor]
        private GetAccessCredentialsResult(
            string backend,

            string clientId,

            string clientSecret,

            string? environment,

            string id,

            int leaseDuration,

            string leaseId,

            bool leaseRenewable,

            string leaseStartTime,

            int? maxCredValidationSeconds,

            string? @namespace,

            int? numSecondsBetweenTests,

            int? numSequentialSuccesses,

            string role,

            string? subscriptionId,

            string? tenantId,

            bool? validateCreds)
        {
            Backend = backend;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Environment = environment;
            Id = id;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            LeaseStartTime = leaseStartTime;
            MaxCredValidationSeconds = maxCredValidationSeconds;
            Namespace = @namespace;
            NumSecondsBetweenTests = numSecondsBetweenTests;
            NumSequentialSuccesses = numSequentialSuccesses;
            Role = role;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
            ValidateCreds = validateCreds;
        }
    }
}
