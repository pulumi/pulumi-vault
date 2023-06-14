// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    public static class GetAccessCredentials
    {
        public static Task<GetAccessCredentialsResult> InvokeAsync(GetAccessCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessCredentialsResult>("vault:aws/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsArgs(), options.WithDefaults());

        public static Output<GetAccessCredentialsResult> Invoke(GetAccessCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessCredentialsResult>("vault:aws/getAccessCredentials:getAccessCredentials", args ?? new GetAccessCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the AWS secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// The region the read credentials belong to.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The name of the AWS secret backend role to read
        /// credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public string Role { get; set; } = null!;

        /// <summary>
        /// The specific AWS ARN to use
        /// from the configured role. If the role does not have multiple ARNs, this does
        /// not need to be specified.
        /// </summary>
        [Input("roleArn")]
        public string? RoleArn { get; set; }

        /// <summary>
        /// Specifies the TTL for the use of the STS token. This
        /// is specified as a string with a duration suffix. Valid only when
        /// `credential_type` of the connected `vault.aws.SecretBackendRole` resource is `assumed_role` or `federation_token`
        /// </summary>
        [Input("ttl")]
        public string? Ttl { get; set; }

        /// <summary>
        /// The type of credentials to read. Defaults
        /// to `"creds"`, which just returns an AWS Access Key ID and Secret
        /// Key. Can also be set to `"sts"`, which will return a security token
        /// in addition to the keys.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetAccessCredentialsArgs()
        {
        }
        public static new GetAccessCredentialsArgs Empty => new GetAccessCredentialsArgs();
    }

    public sealed class GetAccessCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the AWS secret backend to
        /// read credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The region the read credentials belong to.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the AWS secret backend role to read
        /// credentials from, with no leading or trailing `/`s.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The specific AWS ARN to use
        /// from the configured role. If the role does not have multiple ARNs, this does
        /// not need to be specified.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Specifies the TTL for the use of the STS token. This
        /// is specified as a string with a duration suffix. Valid only when
        /// `credential_type` of the connected `vault.aws.SecretBackendRole` resource is `assumed_role` or `federation_token`
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The type of credentials to read. Defaults
        /// to `"creds"`, which just returns an AWS Access Key ID and Secret
        /// Key. Can also be set to `"sts"`, which will return a security token
        /// in addition to the keys.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetAccessCredentialsInvokeArgs()
        {
        }
        public static new GetAccessCredentialsInvokeArgs Empty => new GetAccessCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessCredentialsResult
    {
        /// <summary>
        /// The AWS Access Key ID returned by Vault.
        /// </summary>
        public readonly string AccessKey;
        public readonly string Backend;
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
        public readonly string? Namespace;
        public readonly string? Region;
        public readonly string Role;
        public readonly string? RoleArn;
        /// <summary>
        /// The AWS Secret Key returned by Vault.
        /// </summary>
        public readonly string SecretKey;
        /// <summary>
        /// The STS token returned by Vault, if any.
        /// </summary>
        public readonly string SecurityToken;
        public readonly string? Ttl;
        public readonly string? Type;

        [OutputConstructor]
        private GetAccessCredentialsResult(
            string accessKey,

            string backend,

            string id,

            int leaseDuration,

            string leaseId,

            bool leaseRenewable,

            string leaseStartTime,

            string? @namespace,

            string? region,

            string role,

            string? roleArn,

            string secretKey,

            string securityToken,

            string? ttl,

            string? type)
        {
            AccessKey = accessKey;
            Backend = backend;
            Id = id;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            LeaseStartTime = leaseStartTime;
            Namespace = @namespace;
            Region = region;
            Role = role;
            RoleArn = roleArn;
            SecretKey = secretKey;
            SecurityToken = securityToken;
            Ttl = ttl;
            Type = type;
        }
    }
}
