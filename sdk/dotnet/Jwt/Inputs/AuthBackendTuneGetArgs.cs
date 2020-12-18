// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Jwt.Inputs
{

    public sealed class AuthBackendTuneGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedResponseHeaders")]
        private InputList<string>? _allowedResponseHeaders;
        public InputList<string> AllowedResponseHeaders
        {
            get => _allowedResponseHeaders ?? (_allowedResponseHeaders = new InputList<string>());
            set => _allowedResponseHeaders = value;
        }

        [Input("auditNonHmacRequestKeys")]
        private InputList<string>? _auditNonHmacRequestKeys;

        /// <summary>
        /// Specifies the list of keys that will 
        /// not be HMAC'd by audit devices in the request data object.
        /// </summary>
        public InputList<string> AuditNonHmacRequestKeys
        {
            get => _auditNonHmacRequestKeys ?? (_auditNonHmacRequestKeys = new InputList<string>());
            set => _auditNonHmacRequestKeys = value;
        }

        [Input("auditNonHmacResponseKeys")]
        private InputList<string>? _auditNonHmacResponseKeys;

        /// <summary>
        /// Specifies the list of keys that will 
        /// not be HMAC'd by audit devices in the response data object.
        /// </summary>
        public InputList<string> AuditNonHmacResponseKeys
        {
            get => _auditNonHmacResponseKeys ?? (_auditNonHmacResponseKeys = new InputList<string>());
            set => _auditNonHmacResponseKeys = value;
        }

        /// <summary>
        /// Specifies the default time-to-live. 
        /// If set, this overrides the global default.
        /// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        /// </summary>
        [Input("defaultLeaseTtl")]
        public Input<string>? DefaultLeaseTtl { get; set; }

        /// <summary>
        /// Specifies whether to show this mount in 
        /// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        /// </summary>
        [Input("listingVisibility")]
        public Input<string>? ListingVisibility { get; set; }

        /// <summary>
        /// Specifies the maximum time-to-live. 
        /// If set, this overrides the global default.
        /// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        /// </summary>
        [Input("maxLeaseTtl")]
        public Input<string>? MaxLeaseTtl { get; set; }

        [Input("passthroughRequestHeaders")]
        private InputList<string>? _passthroughRequestHeaders;

        /// <summary>
        /// List of headers to whitelist and 
        /// pass from the request to the backend.
        /// </summary>
        public InputList<string> PassthroughRequestHeaders
        {
            get => _passthroughRequestHeaders ?? (_passthroughRequestHeaders = new InputList<string>());
            set => _passthroughRequestHeaders = value;
        }

        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        public AuthBackendTuneGetArgs()
        {
        }
    }
}
