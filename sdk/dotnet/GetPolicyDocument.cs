// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetPolicyDocument
    {
        /// <summary>
        /// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `vault.Policy` resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var examplePolicyDocument = Output.Create(Vault.GetPolicyDocument.InvokeAsync(new Vault.GetPolicyDocumentArgs
        ///         {
        ///             Rules = 
        ///             {
        ///                 new Vault.Inputs.GetPolicyDocumentRuleArgs
        ///                 {
        ///                     Path = "secret/*",
        ///                     Capabilities = 
        ///                     {
        ///                         "create",
        ///                         "read",
        ///                         "update",
        ///                         "delete",
        ///                         "list",
        ///                     },
        ///                     Description = "allow all on secrets",
        ///                 },
        ///             },
        ///         }));
        ///         var examplePolicy = new Vault.Policy("examplePolicy", new Vault.PolicyArgs
        ///         {
        ///             Policy = examplePolicyDocument.Apply(examplePolicyDocument =&gt; examplePolicyDocument.Hcl),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyDocumentResult> InvokeAsync(GetPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDocumentResult>("vault:index/getPolicyDocument:getPolicyDocument", args ?? new GetPolicyDocumentArgs(), options.WithVersion());
    }


    public sealed class GetPolicyDocumentArgs : Pulumi.InvokeArgs
    {
        [Input("rules")]
        private List<Inputs.GetPolicyDocumentRuleArgs>? _rules;
        public List<Inputs.GetPolicyDocumentRuleArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.GetPolicyDocumentRuleArgs>());
            set => _rules = value;
        }

        public GetPolicyDocumentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyDocumentResult
    {
        /// <summary>
        /// The above arguments serialized as a standard Vault HCL policy document.
        /// </summary>
        public readonly string Hcl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetPolicyDocumentRuleResult> Rules;

        [OutputConstructor]
        private GetPolicyDocumentResult(
            string hcl,

            string id,

            ImmutableArray<Outputs.GetPolicyDocumentRuleResult> rules)
        {
            Hcl = hcl;
            Id = id;
            Rules = rules;
        }
    }
}
