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
    'GetEncryptResult',
    'AwaitableGetEncryptResult',
    'get_encrypt',
    'get_encrypt_output',
]

@pulumi.output_type
class GetEncryptResult:
    """
    A collection of values returned by getEncrypt.
    """
    def __init__(__self__, backend=None, ciphertext=None, context=None, id=None, key=None, key_version=None, namespace=None, plaintext=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        pulumi.set(__self__, "ciphertext", ciphertext)
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        pulumi.set(__self__, "context", context)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if key_version and not isinstance(key_version, int):
            raise TypeError("Expected argument 'key_version' to be a int")
        pulumi.set(__self__, "key_version", key_version)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        pulumi.set(__self__, "plaintext", plaintext)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> _builtins.str:
        return pulumi.get(self, "backend")

    @_builtins.property
    @pulumi.getter
    def ciphertext(self) -> _builtins.str:
        """
        Encrypted ciphertext returned from Vault
        """
        return pulumi.get(self, "ciphertext")

    @_builtins.property
    @pulumi.getter
    def context(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "context")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def key(self) -> _builtins.str:
        return pulumi.get(self, "key")

    @_builtins.property
    @pulumi.getter(name="keyVersion")
    def key_version(self) -> Optional[_builtins.int]:
        return pulumi.get(self, "key_version")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter
    def plaintext(self) -> _builtins.str:
        return pulumi.get(self, "plaintext")


class AwaitableGetEncryptResult(GetEncryptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEncryptResult(
            backend=self.backend,
            ciphertext=self.ciphertext,
            context=self.context,
            id=self.id,
            key=self.key,
            key_version=self.key_version,
            namespace=self.namespace,
            plaintext=self.plaintext)


def get_encrypt(backend: Optional[_builtins.str] = None,
                context: Optional[_builtins.str] = None,
                key: Optional[_builtins.str] = None,
                key_version: Optional[_builtins.int] = None,
                namespace: Optional[_builtins.str] = None,
                plaintext: Optional[_builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEncryptResult:
    """
    This is a data source which can be used to encrypt plaintext using a Vault Transit key.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['context'] = context
    __args__['key'] = key
    __args__['keyVersion'] = key_version
    __args__['namespace'] = namespace
    __args__['plaintext'] = plaintext
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:transit/getEncrypt:getEncrypt', __args__, opts=opts, typ=GetEncryptResult).value

    return AwaitableGetEncryptResult(
        backend=pulumi.get(__ret__, 'backend'),
        ciphertext=pulumi.get(__ret__, 'ciphertext'),
        context=pulumi.get(__ret__, 'context'),
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        key_version=pulumi.get(__ret__, 'key_version'),
        namespace=pulumi.get(__ret__, 'namespace'),
        plaintext=pulumi.get(__ret__, 'plaintext'))
def get_encrypt_output(backend: Optional[pulumi.Input[_builtins.str]] = None,
                       context: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                       key: Optional[pulumi.Input[_builtins.str]] = None,
                       key_version: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                       namespace: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                       plaintext: Optional[pulumi.Input[_builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEncryptResult]:
    """
    This is a data source which can be used to encrypt plaintext using a Vault Transit key.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['context'] = context
    __args__['key'] = key
    __args__['keyVersion'] = key_version
    __args__['namespace'] = namespace
    __args__['plaintext'] = plaintext
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:transit/getEncrypt:getEncrypt', __args__, opts=opts, typ=GetEncryptResult)
    return __ret__.apply(lambda __response__: GetEncryptResult(
        backend=pulumi.get(__response__, 'backend'),
        ciphertext=pulumi.get(__response__, 'ciphertext'),
        context=pulumi.get(__response__, 'context'),
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        key_version=pulumi.get(__response__, 'key_version'),
        namespace=pulumi.get(__response__, 'namespace'),
        plaintext=pulumi.get(__response__, 'plaintext')))
