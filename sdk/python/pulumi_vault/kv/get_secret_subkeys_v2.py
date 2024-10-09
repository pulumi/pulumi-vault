# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetSecretSubkeysV2Result',
    'AwaitableGetSecretSubkeysV2Result',
    'get_secret_subkeys_v2',
    'get_secret_subkeys_v2_output',
]

@pulumi.output_type
class GetSecretSubkeysV2Result:
    """
    A collection of values returned by getSecretSubkeysV2.
    """
    def __init__(__self__, data=None, data_json=None, depth=None, id=None, mount=None, name=None, namespace=None, path=None, version=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if data_json and not isinstance(data_json, str):
            raise TypeError("Expected argument 'data_json' to be a str")
        pulumi.set(__self__, "data_json", data_json)
        if depth and not isinstance(depth, int):
            raise TypeError("Expected argument 'depth' to be a int")
        pulumi.set(__self__, "depth", depth)
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
    @pulumi.getter
    def data(self) -> Mapping[str, str]:
        """
        Subkeys for the KV-V2 secret stored as a serialized map of strings.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> str:
        """
        Subkeys for the KV-V2 secret read from Vault.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter
    def depth(self) -> Optional[int]:
        return pulumi.get(self, "depth")

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
        Full path where the KV-V2 secrets are listed.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


class AwaitableGetSecretSubkeysV2Result(GetSecretSubkeysV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretSubkeysV2Result(
            data=self.data,
            data_json=self.data_json,
            depth=self.depth,
            id=self.id,
            mount=self.mount,
            name=self.name,
            namespace=self.namespace,
            path=self.path,
            version=self.version)


def get_secret_subkeys_v2(depth: Optional[int] = None,
                          mount: Optional[str] = None,
                          name: Optional[str] = None,
                          namespace: Optional[str] = None,
                          version: Optional[int] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretSubkeysV2Result:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_vault as vault

    kvv2 = vault.Mount("kvv2",
        path="kvv2",
        type="kv",
        options={
            "version": "2",
        },
        description="KV Version 2 secret engine mount")
    aws_secret = vault.kv.SecretV2("aws_secret",
        mount=kvv2.path,
        name="aws_secret",
        data_json=json.dumps({
            "zip": "zap",
            "foo": "bar",
        }))
    test = vault.kv.get_secret_subkeys_v2_output(mount=kvv2.path,
        name=aws_secret.name)
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param int depth: Specifies the deepest nesting level to provide in the output.
           If non-zero, keys that reside at the specified depth value will be
           artificially treated as leaves and will thus be `null` even if further
           underlying sub-keys exist.
    :param str mount: Path where KV-V2 engine is mounted.
    :param str name: Full name of the secret. For a nested secret
           the name is the nested path excluding the mount and data
           prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
           the name is `foo/bar/baz`.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param int version: Specifies the version to return. If not 
           set the latest version is returned.
    """
    __args__ = dict()
    __args__['depth'] = depth
    __args__['mount'] = mount
    __args__['name'] = name
    __args__['namespace'] = namespace
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kv/getSecretSubkeysV2:getSecretSubkeysV2', __args__, opts=opts, typ=GetSecretSubkeysV2Result).value

    return AwaitableGetSecretSubkeysV2Result(
        data=pulumi.get(__ret__, 'data'),
        data_json=pulumi.get(__ret__, 'data_json'),
        depth=pulumi.get(__ret__, 'depth'),
        id=pulumi.get(__ret__, 'id'),
        mount=pulumi.get(__ret__, 'mount'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'),
        version=pulumi.get(__ret__, 'version'))
def get_secret_subkeys_v2_output(depth: Optional[pulumi.Input[Optional[int]]] = None,
                                 mount: Optional[pulumi.Input[str]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                 version: Optional[pulumi.Input[Optional[int]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretSubkeysV2Result]:
    """
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_vault as vault

    kvv2 = vault.Mount("kvv2",
        path="kvv2",
        type="kv",
        options={
            "version": "2",
        },
        description="KV Version 2 secret engine mount")
    aws_secret = vault.kv.SecretV2("aws_secret",
        mount=kvv2.path,
        name="aws_secret",
        data_json=json.dumps({
            "zip": "zap",
            "foo": "bar",
        }))
    test = vault.kv.get_secret_subkeys_v2_output(mount=kvv2.path,
        name=aws_secret.name)
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `read` capability on the given path.


    :param int depth: Specifies the deepest nesting level to provide in the output.
           If non-zero, keys that reside at the specified depth value will be
           artificially treated as leaves and will thus be `null` even if further
           underlying sub-keys exist.
    :param str mount: Path where KV-V2 engine is mounted.
    :param str name: Full name of the secret. For a nested secret
           the name is the nested path excluding the mount and data
           prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
           the name is `foo/bar/baz`.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param int version: Specifies the version to return. If not 
           set the latest version is returned.
    """
    __args__ = dict()
    __args__['depth'] = depth
    __args__['mount'] = mount
    __args__['name'] = name
    __args__['namespace'] = namespace
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:kv/getSecretSubkeysV2:getSecretSubkeysV2', __args__, opts=opts, typ=GetSecretSubkeysV2Result)
    return __ret__.apply(lambda __response__: GetSecretSubkeysV2Result(
        data=pulumi.get(__response__, 'data'),
        data_json=pulumi.get(__response__, 'data_json'),
        depth=pulumi.get(__response__, 'depth'),
        id=pulumi.get(__response__, 'id'),
        mount=pulumi.get(__response__, 'mount'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        path=pulumi.get(__response__, 'path'),
        version=pulumi.get(__response__, 'version')))
