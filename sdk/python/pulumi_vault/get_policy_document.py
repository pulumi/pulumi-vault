# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetPolicyDocumentResult',
    'AwaitableGetPolicyDocumentResult',
    'get_policy_document',
    'get_policy_document_output',
]

@pulumi.output_type
class GetPolicyDocumentResult:
    """
    A collection of values returned by getPolicyDocument.
    """
    def __init__(__self__, hcl=None, id=None, namespace=None, rules=None):
        if hcl and not isinstance(hcl, str):
            raise TypeError("Expected argument 'hcl' to be a str")
        pulumi.set(__self__, "hcl", hcl)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def hcl(self) -> str:
        """
        The above arguments serialized as a standard Vault HCL policy document.
        """
        return pulumi.get(self, "hcl")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetPolicyDocumentRuleResult']:
        return pulumi.get(self, "rules")


class AwaitableGetPolicyDocumentResult(GetPolicyDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyDocumentResult(
            hcl=self.hcl,
            id=self.id,
            namespace=self.namespace,
            rules=self.rules)


def get_policy_document(namespace: Optional[str] = None,
                        rules: Optional[Sequence[Union['GetPolicyDocumentRuleArgs', 'GetPolicyDocumentRuleArgsDict']]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyDocumentResult:
    """
    This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `Policy` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    example = vault.get_policy_document(rules=[{
        "path": "secret/*",
        "capabilities": [
            "create",
            "read",
            "update",
            "delete",
            "list",
        ],
        "description": "allow all on secrets",
    }])
    example_policy = vault.Policy("example",
        name="example_policy",
        policy=example.hcl)
    ```
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['rules'] = rules
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:index/getPolicyDocument:getPolicyDocument', __args__, opts=opts, typ=GetPolicyDocumentResult).value

    return AwaitableGetPolicyDocumentResult(
        hcl=pulumi.get(__ret__, 'hcl'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        rules=pulumi.get(__ret__, 'rules'))


@_utilities.lift_output_func(get_policy_document)
def get_policy_document_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
                               rules: Optional[pulumi.Input[Optional[Sequence[Union['GetPolicyDocumentRuleArgs', 'GetPolicyDocumentRuleArgsDict']]]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyDocumentResult]:
    """
    This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `Policy` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    example = vault.get_policy_document(rules=[{
        "path": "secret/*",
        "capabilities": [
            "create",
            "read",
            "update",
            "delete",
            "list",
        ],
        "description": "allow all on secrets",
    }])
    example_policy = vault.Policy("example",
        name="example_policy",
        policy=example.hcl)
    ```
    """
    ...
