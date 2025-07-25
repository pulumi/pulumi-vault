# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
    'GetAccessCredentialsResult',
    'AwaitableGetAccessCredentialsResult',
    'get_access_credentials',
    'get_access_credentials_output',
]

@pulumi.output_type
class GetAccessCredentialsResult:
    """
    A collection of values returned by getAccessCredentials.
    """
    def __init__(__self__, backend=None, current_password=None, id=None, last_password=None, namespace=None, role=None, username=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if current_password and not isinstance(current_password, str):
            raise TypeError("Expected argument 'current_password' to be a str")
        pulumi.set(__self__, "current_password", current_password)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_password and not isinstance(last_password, str):
            raise TypeError("Expected argument 'last_password' to be a str")
        pulumi.set(__self__, "last_password", last_password)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> _builtins.str:
        return pulumi.get(self, "backend")

    @_builtins.property
    @pulumi.getter(name="currentPassword")
    def current_password(self) -> _builtins.str:
        """
        The current set password on the Active Directory service account.
        """
        return pulumi.get(self, "current_password")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="lastPassword")
    def last_password(self) -> _builtins.str:
        """
        The current set password on the Active Directory service account, provided because AD is eventually consistent.
        """
        return pulumi.get(self, "last_password")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter
    def role(self) -> _builtins.str:
        return pulumi.get(self, "role")

    @_builtins.property
    @pulumi.getter
    def username(self) -> _builtins.str:
        """
        The Active Directory service account username.
        """
        return pulumi.get(self, "username")


class AwaitableGetAccessCredentialsResult(GetAccessCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessCredentialsResult(
            backend=self.backend,
            current_password=self.current_password,
            id=self.id,
            last_password=self.last_password,
            namespace=self.namespace,
            role=self.role,
            username=self.username)


def get_access_credentials(backend: Optional[_builtins.str] = None,
                           namespace: Optional[_builtins.str] = None,
                           role: Optional[_builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessCredentialsResult:
    """
    ## Example Usage


    :param _builtins.str backend: The path to the AD secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param _builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param _builtins.str role: The name of the AD secret backend role to read
           credentials from, with no leading or trailing `/`s.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['role'] = role
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:ad/getAccessCredentials:getAccessCredentials', __args__, opts=opts, typ=GetAccessCredentialsResult).value

    return AwaitableGetAccessCredentialsResult(
        backend=pulumi.get(__ret__, 'backend'),
        current_password=pulumi.get(__ret__, 'current_password'),
        id=pulumi.get(__ret__, 'id'),
        last_password=pulumi.get(__ret__, 'last_password'),
        namespace=pulumi.get(__ret__, 'namespace'),
        role=pulumi.get(__ret__, 'role'),
        username=pulumi.get(__ret__, 'username'))
def get_access_credentials_output(backend: Optional[pulumi.Input[_builtins.str]] = None,
                                  namespace: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                  role: Optional[pulumi.Input[_builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAccessCredentialsResult]:
    """
    ## Example Usage


    :param _builtins.str backend: The path to the AD secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param _builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param _builtins.str role: The name of the AD secret backend role to read
           credentials from, with no leading or trailing `/`s.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['role'] = role
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:ad/getAccessCredentials:getAccessCredentials', __args__, opts=opts, typ=GetAccessCredentialsResult)
    return __ret__.apply(lambda __response__: GetAccessCredentialsResult(
        backend=pulumi.get(__response__, 'backend'),
        current_password=pulumi.get(__response__, 'current_password'),
        id=pulumi.get(__response__, 'id'),
        last_password=pulumi.get(__response__, 'last_password'),
        namespace=pulumi.get(__response__, 'namespace'),
        role=pulumi.get(__response__, 'role'),
        username=pulumi.get(__response__, 'username')))
