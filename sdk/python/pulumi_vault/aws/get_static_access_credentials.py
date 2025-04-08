# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetStaticAccessCredentialsResult',
    'AwaitableGetStaticAccessCredentialsResult',
    'get_static_access_credentials',
    'get_static_access_credentials_output',
]

@pulumi.output_type
class GetStaticAccessCredentialsResult:
    """
    A collection of values returned by getStaticAccessCredentials.
    """
    def __init__(__self__, access_key=None, backend=None, id=None, name=None, namespace=None, secret_key=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if secret_key and not isinstance(secret_key, str):
            raise TypeError("Expected argument 'secret_key' to be a str")
        pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> builtins.str:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def backend(self) -> builtins.str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> builtins.str:
        return pulumi.get(self, "secret_key")


class AwaitableGetStaticAccessCredentialsResult(GetStaticAccessCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStaticAccessCredentialsResult(
            access_key=self.access_key,
            backend=self.backend,
            id=self.id,
            name=self.name,
            namespace=self.namespace,
            secret_key=self.secret_key)


def get_static_access_credentials(backend: Optional[builtins.str] = None,
                                  name: Optional[builtins.str] = None,
                                  namespace: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStaticAccessCredentialsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['name'] = name
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:aws/getStaticAccessCredentials:getStaticAccessCredentials', __args__, opts=opts, typ=GetStaticAccessCredentialsResult).value

    return AwaitableGetStaticAccessCredentialsResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        backend=pulumi.get(__ret__, 'backend'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        secret_key=pulumi.get(__ret__, 'secret_key'))
def get_static_access_credentials_output(backend: Optional[pulumi.Input[builtins.str]] = None,
                                         name: Optional[pulumi.Input[builtins.str]] = None,
                                         namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStaticAccessCredentialsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['name'] = name
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:aws/getStaticAccessCredentials:getStaticAccessCredentials', __args__, opts=opts, typ=GetStaticAccessCredentialsResult)
    return __ret__.apply(lambda __response__: GetStaticAccessCredentialsResult(
        access_key=pulumi.get(__response__, 'access_key'),
        backend=pulumi.get(__response__, 'backend'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        secret_key=pulumi.get(__response__, 'secret_key')))
