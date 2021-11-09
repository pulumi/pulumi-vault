# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    def __init__(__self__, data=None, data_json=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, lease_start_time=None, path=None, version=None):
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
        if lease_start_time and not isinstance(lease_start_time, str):
            raise TypeError("Expected argument 'lease_start_time' to be a str")
        pulumi.set(__self__, "lease_start_time", lease_start_time)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def data(self) -> Mapping[str, Any]:
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
        A string containing the full data payload retrieved from
        Vault, serialized in JSON format.
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
        The duration of the secret lease, in seconds relative
        to the time the data was requested. Once this time has passed any plan
        generated with this data may fail to apply.
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
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter(name="leaseStartTime")
    def lease_start_time(self) -> str:
        return pulumi.get(self, "lease_start_time")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


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
            lease_start_time=self.lease_start_time,
            path=self.path,
            version=self.version)


def get_secret(path: Optional[str] = None,
               version: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
    """
    Use this data source to access information about an existing resource.

    :param str path: The full logical path from which to request data.
           To read data from the "generic" secret backend mounted in Vault by
           default, this should be prefixed with `secret/`. Reading from other backends
           with this data source is possible; consult each backend's documentation
           to see which endpoints support the `GET` method.
    :param int version: The version of the secret to read. This is used by the
           Vault KV secrets engine - version 2 to indicate which version of the secret
           to read.
    """
    __args__ = dict()
    __args__['path'] = path
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:generic/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        data=__ret__.data,
        data_json=__ret__.data_json,
        id=__ret__.id,
        lease_duration=__ret__.lease_duration,
        lease_id=__ret__.lease_id,
        lease_renewable=__ret__.lease_renewable,
        lease_start_time=__ret__.lease_start_time,
        path=__ret__.path,
        version=__ret__.version)


@_utilities.lift_output_func(get_secret)
def get_secret_output(path: Optional[pulumi.Input[str]] = None,
                      version: Optional[pulumi.Input[Optional[int]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretResult]:
    """
    Use this data source to access information about an existing resource.

    :param str path: The full logical path from which to request data.
           To read data from the "generic" secret backend mounted in Vault by
           default, this should be prefixed with `secret/`. Reading from other backends
           with this data source is possible; consult each backend's documentation
           to see which endpoints support the `GET` method.
    :param int version: The version of the secret to read. This is used by the
           Vault KV secrets engine - version 2 to indicate which version of the secret
           to read.
    """
    ...
