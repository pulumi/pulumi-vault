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
from . import _utilities

__all__ = [
    'GetAuthBackendResult',
    'AwaitableGetAuthBackendResult',
    'get_auth_backend',
    'get_auth_backend_output',
]

@pulumi.output_type
class GetAuthBackendResult:
    """
    A collection of values returned by getAuthBackend.
    """
    def __init__(__self__, accessor=None, default_lease_ttl_seconds=None, description=None, id=None, listing_visibility=None, local=None, max_lease_ttl_seconds=None, namespace=None, path=None, type=None):
        if accessor and not isinstance(accessor, str):
            raise TypeError("Expected argument 'accessor' to be a str")
        pulumi.set(__self__, "accessor", accessor)
        if default_lease_ttl_seconds and not isinstance(default_lease_ttl_seconds, int):
            raise TypeError("Expected argument 'default_lease_ttl_seconds' to be a int")
        pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listing_visibility and not isinstance(listing_visibility, str):
            raise TypeError("Expected argument 'listing_visibility' to be a str")
        pulumi.set(__self__, "listing_visibility", listing_visibility)
        if local and not isinstance(local, bool):
            raise TypeError("Expected argument 'local' to be a bool")
        pulumi.set(__self__, "local", local)
        if max_lease_ttl_seconds and not isinstance(max_lease_ttl_seconds, int):
            raise TypeError("Expected argument 'max_lease_ttl_seconds' to be a int")
        pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def accessor(self) -> str:
        """
        The accessor for this auth method.
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> int:
        """
        The default lease duration in seconds.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the auth method.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> str:
        """
        Specifies whether to show this mount in the UI-specific listing endpoint.
        """
        return pulumi.get(self, "listing_visibility")

    @property
    @pulumi.getter
    def local(self) -> bool:
        """
        Specifies if the auth method is local only.
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> int:
        """
        The maximum lease duration in seconds.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The name of the auth method type.
        """
        return pulumi.get(self, "type")


class AwaitableGetAuthBackendResult(GetAuthBackendResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendResult(
            accessor=self.accessor,
            default_lease_ttl_seconds=self.default_lease_ttl_seconds,
            description=self.description,
            id=self.id,
            listing_visibility=self.listing_visibility,
            local=self.local,
            max_lease_ttl_seconds=self.max_lease_ttl_seconds,
            namespace=self.namespace,
            path=self.path,
            type=self.type)


def get_auth_backend(namespace: Optional[str] = None,
                     path: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthBackendResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    example = vault.get_auth_backend(path="userpass")
    ```


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param str path: The auth backend mount point.
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:index/getAuthBackend:getAuthBackend', __args__, opts=opts, typ=GetAuthBackendResult).value

    return AwaitableGetAuthBackendResult(
        accessor=pulumi.get(__ret__, 'accessor'),
        default_lease_ttl_seconds=pulumi.get(__ret__, 'default_lease_ttl_seconds'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        listing_visibility=pulumi.get(__ret__, 'listing_visibility'),
        local=pulumi.get(__ret__, 'local'),
        max_lease_ttl_seconds=pulumi.get(__ret__, 'max_lease_ttl_seconds'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'),
        type=pulumi.get(__ret__, 'type'))
def get_auth_backend_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
                            path: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuthBackendResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    example = vault.get_auth_backend(path="userpass")
    ```


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param str path: The auth backend mount point.
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:index/getAuthBackend:getAuthBackend', __args__, opts=opts, typ=GetAuthBackendResult)
    return __ret__.apply(lambda __response__: GetAuthBackendResult(
        accessor=pulumi.get(__response__, 'accessor'),
        default_lease_ttl_seconds=pulumi.get(__response__, 'default_lease_ttl_seconds'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        listing_visibility=pulumi.get(__response__, 'listing_visibility'),
        local=pulumi.get(__response__, 'local'),
        max_lease_ttl_seconds=pulumi.get(__response__, 'max_lease_ttl_seconds'),
        namespace=pulumi.get(__response__, 'namespace'),
        path=pulumi.get(__response__, 'path'),
        type=pulumi.get(__response__, 'type')))
