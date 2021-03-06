// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        public static Task<GetAccessCredentialsResult> InvokeAsync(GetAccessCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessCredentialsResult>("vault:azure/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsArgs(), options.WithVersion());
    }


    public sealed class GetAccessCredentialsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the Azure secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds after which to give up validating credentials. Defaults
        /// to 1,200 (20 minutes).
        /// </summary>
        [Input("maxCredValidationSeconds")]
        public int? MaxCredValidationSeconds { get; set; }

        /// <summary>
        /// If 'validate_creds' is true, 
        /// the number of seconds to wait between each test of generated credentials.
        /// Defaults to 7.
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
        public readonly int? NumSecondsBetweenTests;
        public readonly int? NumSequentialSuccesses;
        public readonly string Role;
        public readonly bool? ValidateCreds;

        [OutputConstructor]
        private GetAccessCredentialsResult(
            string backend,

            string clientId,

            string clientSecret,

            string id,

            int leaseDuration,

            string leaseId,

            bool leaseRenewable,

            string leaseStartTime,

            int? maxCredValidationSeconds,

            int? numSecondsBetweenTests,

            int? numSequentialSuccesses,

            string role,

            bool? validateCreds)
        {
            Backend = backend;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            LeaseStartTime = leaseStartTime;
            MaxCredValidationSeconds = maxCredValidationSeconds;
            NumSecondsBetweenTests = numSecondsBetweenTests;
            NumSequentialSuccesses = numSequentialSuccesses;
            Role = role;
            ValidateCreds = validateCreds;
        }
    }
}
