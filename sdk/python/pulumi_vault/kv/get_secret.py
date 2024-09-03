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
    'GetSecretResult',
    'AwaitableGetSecretResult',
    'get_secret',
    'get_secret_output',
]

@pulumi.output_type
class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, data=None, data_json=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, namespace=None, path=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if data_json and not isinstance(data_json, str):
            raise TypeError("Expected argument 'data_json' to be a str")
        pulumi.set(__self__, "data_json", data_json)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lease_duration and not isinstance(lease_duration, int):
            raise TypeError("Expected argument 'lease_duration' to be a int")
        pulumi.set(__self__, "lease_duration", lease_duration)
        if lease_id and not isinstance(lease_id, str):
            raise TypeError("Expected argument 'lease_id' to be a str")
        pulumi.set(__self__, "lease_id", lease_id)
        if lease_renewable and not isinstance(lease_renewable, bool):
            raise TypeError("Expected argument 'lease_renewable' to be a bool")
        pulumi.set(__self__, "lease_renewable", lease_renewable)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def data(self) -> Mapping[str, str]:
        """
        A mapping whose keys are the top-level data keys returned from
        Vault and whose values are the corresponding values. This map can only
        represent string data, so any non-string values returned from Vault are
        serialized as JSON.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> str:
        """
        JSON-encoded string that that is
        read as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> int:
        """
        The duration of the secret lease, in seconds. Once
        this time has passed any plan generated with this data may fail to apply.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> str:
        """
        The lease identifier assigned by Vault, if any.
        """
        return pulumi.get(self, "lease_id")

    @property
    @pulumi.getter(name="leaseRenewable")
    def lease_renewable(self) -> bool:
        """
        True if the duration of this lease can be extended
        through renewal.
        """
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")


class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            data=self.data,
            data_json=self.data_json,
            id=self.id,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            namespace=self.namespace,
            path=self.path)


def get_secret(namespace: Optional[str] = None,
               path: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
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
    secret = vault.kv.Secret("secret",
        path=kvv1.path.apply(lambda path: f"{path}/secret"),
        data_json=json.dumps({
            "zip": "zap",
            "foo": "bar",
        }))
    secret_data = vault.kv.get_secret_output(path=secret.path)
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Full path of the KV-V1 secret.
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kv/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        data=pulumi.get(__ret__, 'data'),
        data_json=pulumi.get(__ret__, 'data_json'),
        id=pulumi.get(__ret__, 'id'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_id=pulumi.get(__ret__, 'lease_id'),
        lease_renewable=pulumi.get(__ret__, 'lease_renewable'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'))


@_utilities.lift_output_func(get_secret)
def get_secret_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
                      path: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretResult]:
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
    secret = vault.kv.Secret("secret",
        path=kvv1.path.apply(lambda path: f"{path}/secret"),
        data_json=json.dumps({
            "zip": "zap",
            "foo": "bar",
        }))
    secret_data = vault.kv.get_secret_output(path=secret.path)
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Full path of the KV-V1 secret.
    """
    ...
