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
    'GetResult',
    'AwaitableGetResult',
    'get',
    'get_output',
]

warnings.warn("""vault.namespace/get.get has been deprecated in favor of vault.index/getnamespace.getNamespace""", DeprecationWarning)

@pulumi.output_type
class GetResult:
    """
    A collection of values returned by get.
    """
    def __init__(__self__, custom_metadata=None, id=None, namespace=None, namespace_id=None, path=None, path_fq=None):
        if custom_metadata and not isinstance(custom_metadata, dict):
            raise TypeError("Expected argument 'custom_metadata' to be a dict")
        pulumi.set(__self__, "custom_metadata", custom_metadata)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if path_fq and not isinstance(path_fq, str):
            raise TypeError("Expected argument 'path_fq' to be a str")
        pulumi.set(__self__, "path_fq", path_fq)

    @property
    @pulumi.getter(name="customMetadata")
    def custom_metadata(self) -> Mapping[str, Any]:
        return pulumi.get(self, "custom_metadata")

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
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> str:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pathFq")
    def path_fq(self) -> str:
        return pulumi.get(self, "path_fq")


class AwaitableGetResult(GetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResult(
            custom_metadata=self.custom_metadata,
            id=self.id,
            namespace=self.namespace,
            namespace_id=self.namespace_id,
            path=self.path,
            path_fq=self.path_fq)


def get(namespace: Optional[str] = None,
        path: Optional[str] = None,
        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResult:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get is deprecated: vault.namespace/get.get has been deprecated in favor of vault.index/getnamespace.getNamespace""")
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:namespace/get:get', __args__, opts=opts, typ=GetResult).value

    return AwaitableGetResult(
        custom_metadata=pulumi.get(__ret__, 'custom_metadata'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        namespace_id=pulumi.get(__ret__, 'namespace_id'),
        path=pulumi.get(__ret__, 'path'),
        path_fq=pulumi.get(__ret__, 'path_fq'))


@_utilities.lift_output_func(get)
def get_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
               path: Optional[pulumi.Input[Optional[str]]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResult]:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get is deprecated: vault.namespace/get.get has been deprecated in favor of vault.index/getnamespace.getNamespace""")
    ...
