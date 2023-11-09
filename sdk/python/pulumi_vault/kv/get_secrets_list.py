# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSecretsListResult',
    'AwaitableGetSecretsListResult',
    'get_secrets_list',
    'get_secrets_list_output',
]

@pulumi.output_type
class GetSecretsListResult:
    """
    A collection of values returned by getSecretsList.
    """
    def __init__(__self__, id=None, names=None, namespace=None, path=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        List of all secret names listed under the given path.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")


class AwaitableGetSecretsListResult(GetSecretsListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretsListResult(
            id=self.id,
            names=self.names,
            namespace=self.namespace,
            path=self.path)


def get_secrets_list(namespace: Optional[str] = None,
                     path: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretsListResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_vault as vault

    kvv1 = vault.Mount("kvv1",
        path="kvv1",
        type="kv",
        options={
            "version": "1",
        },
        description="KV Version 1 secret engine mount")
    aws_secret = vault.kv.Secret("awsSecret",
        path=kvv1.path.apply(lambda path: f"{path}/aws-secret"),
        data_json=json.dumps({
            "zip": "zap",
        }))
    azure_secret = vault.kv.Secret("azureSecret",
        path=kvv1.path.apply(lambda path: f"{path}/azure-secret"),
        data_json=json.dumps({
            "foo": "bar",
        }))
    secrets = vault.kv.get_secrets_list_output(path=kvv1.path)
    ```
    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Full KV-V1 path where secrets will be listed.
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kv/getSecretsList:getSecretsList', __args__, opts=opts, typ=GetSecretsListResult).value

    return AwaitableGetSecretsListResult(
        id=pulumi.get(__ret__, 'id'),
        names=pulumi.get(__ret__, 'names'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'))


@_utilities.lift_output_func(get_secrets_list)
def get_secrets_list_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
                            path: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretsListResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_vault as vault

    kvv1 = vault.Mount("kvv1",
        path="kvv1",
        type="kv",
        options={
            "version": "1",
        },
        description="KV Version 1 secret engine mount")
    aws_secret = vault.kv.Secret("awsSecret",
        path=kvv1.path.apply(lambda path: f"{path}/aws-secret"),
        data_json=json.dumps({
            "zip": "zap",
        }))
    azure_secret = vault.kv.Secret("azureSecret",
        path=kvv1.path.apply(lambda path: f"{path}/azure-secret"),
        data_json=json.dumps({
            "foo": "bar",
        }))
    secrets = vault.kv.get_secrets_list_output(path=kvv1.path)
    ```
    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Full KV-V1 path where secrets will be listed.
    """
    ...
