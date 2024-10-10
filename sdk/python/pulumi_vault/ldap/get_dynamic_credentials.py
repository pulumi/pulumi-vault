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
    'GetDynamicCredentialsResult',
    'AwaitableGetDynamicCredentialsResult',
    'get_dynamic_credentials',
    'get_dynamic_credentials_output',
]

@pulumi.output_type
class GetDynamicCredentialsResult:
    """
    A collection of values returned by getDynamicCredentials.
    """
    def __init__(__self__, distinguished_names=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, mount=None, namespace=None, password=None, role_name=None, username=None):
        if distinguished_names and not isinstance(distinguished_names, list):
            raise TypeError("Expected argument 'distinguished_names' to be a list")
        pulumi.set(__self__, "distinguished_names", distinguished_names)
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
        if mount and not isinstance(mount, str):
            raise TypeError("Expected argument 'mount' to be a str")
        pulumi.set(__self__, "mount", mount)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        pulumi.set(__self__, "role_name", role_name)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="distinguishedNames")
    def distinguished_names(self) -> Sequence[str]:
        return pulumi.get(self, "distinguished_names")

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
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> str:
        return pulumi.get(self, "lease_id")

    @property
    @pulumi.getter(name="leaseRenewable")
    def lease_renewable(self) -> bool:
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter
    def mount(self) -> str:
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetDynamicCredentialsResult(GetDynamicCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDynamicCredentialsResult(
            distinguished_names=self.distinguished_names,
            id=self.id,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            mount=self.mount,
            namespace=self.namespace,
            password=self.password,
            role_name=self.role_name,
            username=self.username)


def get_dynamic_credentials(mount: Optional[str] = None,
                            namespace: Optional[str] = None,
                            role_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDynamicCredentialsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['mount'] = mount
    __args__['namespace'] = namespace
    __args__['roleName'] = role_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:ldap/getDynamicCredentials:getDynamicCredentials', __args__, opts=opts, typ=GetDynamicCredentialsResult).value

    return AwaitableGetDynamicCredentialsResult(
        distinguished_names=pulumi.get(__ret__, 'distinguished_names'),
        id=pulumi.get(__ret__, 'id'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_id=pulumi.get(__ret__, 'lease_id'),
        lease_renewable=pulumi.get(__ret__, 'lease_renewable'),
        mount=pulumi.get(__ret__, 'mount'),
        namespace=pulumi.get(__ret__, 'namespace'),
        password=pulumi.get(__ret__, 'password'),
        role_name=pulumi.get(__ret__, 'role_name'),
        username=pulumi.get(__ret__, 'username'))
def get_dynamic_credentials_output(mount: Optional[pulumi.Input[str]] = None,
                                   namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                   role_name: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDynamicCredentialsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['mount'] = mount
    __args__['namespace'] = namespace
    __args__['roleName'] = role_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:ldap/getDynamicCredentials:getDynamicCredentials', __args__, opts=opts, typ=GetDynamicCredentialsResult)
    return __ret__.apply(lambda __response__: GetDynamicCredentialsResult(
        distinguished_names=pulumi.get(__response__, 'distinguished_names'),
        id=pulumi.get(__response__, 'id'),
        lease_duration=pulumi.get(__response__, 'lease_duration'),
        lease_id=pulumi.get(__response__, 'lease_id'),
        lease_renewable=pulumi.get(__response__, 'lease_renewable'),
        mount=pulumi.get(__response__, 'mount'),
        namespace=pulumi.get(__response__, 'namespace'),
        password=pulumi.get(__response__, 'password'),
        role_name=pulumi.get(__response__, 'role_name'),
        username=pulumi.get(__response__, 'username')))
