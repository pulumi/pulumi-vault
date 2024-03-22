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
    'GetSecretV2Result',
    'AwaitableGetSecretV2Result',
    'get_secret_v2',
    'get_secret_v2_output',
]

@pulumi.output_type
class GetSecretV2Result:
    """
    A collection of values returned by getSecretV2.
    """
    def __init__(__self__, created_time=None, custom_metadata=None, data=None, data_json=None, deletion_time=None, destroyed=None, id=None, mount=None, name=None, namespace=None, path=None, version=None):
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if custom_metadata and not isinstance(custom_metadata, dict):
            raise TypeError("Expected argument 'custom_metadata' to be a dict")
        pulumi.set(__self__, "custom_metadata", custom_metadata)
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if data_json and not isinstance(data_json, str):
            raise TypeError("Expected argument 'data_json' to be a str")
        pulumi.set(__self__, "data_json", data_json)
        if deletion_time and not isinstance(deletion_time, str):
            raise TypeError("Expected argument 'deletion_time' to be a str")
        pulumi.set(__self__, "deletion_time", deletion_time)
        if destroyed and not isinstance(destroyed, bool):
            raise TypeError("Expected argument 'destroyed' to be a bool")
        pulumi.set(__self__, "destroyed", destroyed)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mount and not isinstance(mount, str):
            raise TypeError("Expected argument 'mount' to be a str")
        pulumi.set(__self__, "mount", mount)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        Time at which secret was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="customMetadata")
    def custom_metadata(self) -> Mapping[str, Any]:
        """
        Custom metadata for the secret.
        """
        return pulumi.get(self, "custom_metadata")

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
        JSON-encoded string that that is
        read as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter(name="deletionTime")
    def deletion_time(self) -> str:
        """
        Deletion time for the secret.
        """
        return pulumi.get(self, "deletion_time")

    @property
    @pulumi.getter
    def destroyed(self) -> bool:
        """
        Indicates whether the secret has been destroyed.
        """
        return pulumi.get(self, "destroyed")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mount(self) -> str:
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        Full path where the KVV2 secret is written.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        Version of the secret.
        """
        return pulumi.get(self, "version")


class AwaitableGetSecretV2Result(GetSecretV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretV2Result(
            created_time=self.created_time,
            custom_metadata=self.custom_metadata,
            data=self.data,
            data_json=self.data_json,
            deletion_time=self.deletion_time,
            destroyed=self.destroyed,
            id=self.id,
            mount=self.mount,
            name=self.name,
            namespace=self.namespace,
            path=self.path,
            version=self.version)


def get_secret_v2(mount: Optional[str] = None,
                  name: Optional[str] = None,
                  namespace: Optional[str] = None,
                  version: Optional[int] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretV2Result:
    """
    ## Example Usage


    :param str mount: Path where KV-V2 engine is mounted.
    :param str name: Full name of the secret. For a nested secret
           the name is the nested path excluding the mount and data
           prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
           the name is `foo/bar/baz`.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param int version: Version of the secret to retrieve.
    """
    __args__ = dict()
    __args__['mount'] = mount
    __args__['name'] = name
    __args__['namespace'] = namespace
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kv/getSecretV2:getSecretV2', __args__, opts=opts, typ=GetSecretV2Result).value

    return AwaitableGetSecretV2Result(
        created_time=pulumi.get(__ret__, 'created_time'),
        custom_metadata=pulumi.get(__ret__, 'custom_metadata'),
        data=pulumi.get(__ret__, 'data'),
        data_json=pulumi.get(__ret__, 'data_json'),
        deletion_time=pulumi.get(__ret__, 'deletion_time'),
        destroyed=pulumi.get(__ret__, 'destroyed'),
        id=pulumi.get(__ret__, 'id'),
        mount=pulumi.get(__ret__, 'mount'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_secret_v2)
def get_secret_v2_output(mount: Optional[pulumi.Input[str]] = None,
                         name: Optional[pulumi.Input[str]] = None,
                         namespace: Optional[pulumi.Input[Optional[str]]] = None,
                         version: Optional[pulumi.Input[Optional[int]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretV2Result]:
    """
    ## Example Usage


    :param str mount: Path where KV-V2 engine is mounted.
    :param str name: Full name of the secret. For a nested secret
           the name is the nested path excluding the mount and data
           prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
           the name is `foo/bar/baz`.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param int version: Version of the secret to retrieve.
    """
    ...
