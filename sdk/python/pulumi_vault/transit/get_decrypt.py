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
    'GetDecryptResult',
    'AwaitableGetDecryptResult',
    'get_decrypt',
    'get_decrypt_output',
]

@pulumi.output_type
class GetDecryptResult:
    """
    A collection of values returned by getDecrypt.
    """
    def __init__(__self__, backend=None, ciphertext=None, context=None, id=None, key=None, namespace=None, plaintext=None):
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
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        pulumi.set(__self__, "plaintext", plaintext)

    @property
    @pulumi.getter
    def backend(self) -> builtins.str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def ciphertext(self) -> builtins.str:
        return pulumi.get(self, "ciphertext")

    @property
    @pulumi.getter
    def context(self) -> Optional[builtins.str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> builtins.str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def plaintext(self) -> builtins.str:
        """
        Decrypted plaintext returned from Vault
        """
        return pulumi.get(self, "plaintext")


class AwaitableGetDecryptResult(GetDecryptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDecryptResult(
            backend=self.backend,
            ciphertext=self.ciphertext,
            context=self.context,
            id=self.id,
            key=self.key,
            namespace=self.namespace,
            plaintext=self.plaintext)


def get_decrypt(backend: Optional[builtins.str] = None,
                ciphertext: Optional[builtins.str] = None,
                context: Optional[builtins.str] = None,
                key: Optional[builtins.str] = None,
                namespace: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDecryptResult:
    """
    This is a data source which can be used to decrypt ciphertext using a Vault Transit key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    test = vault.transit.get_decrypt(backend="transit",
        key="test",
        ciphertext="vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==")
    ```
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['ciphertext'] = ciphertext
    __args__['context'] = context
    __args__['key'] = key
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:transit/getDecrypt:getDecrypt', __args__, opts=opts, typ=GetDecryptResult).value

    return AwaitableGetDecryptResult(
        backend=pulumi.get(__ret__, 'backend'),
        ciphertext=pulumi.get(__ret__, 'ciphertext'),
        context=pulumi.get(__ret__, 'context'),
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        namespace=pulumi.get(__ret__, 'namespace'),
        plaintext=pulumi.get(__ret__, 'plaintext'))
def get_decrypt_output(backend: Optional[pulumi.Input[builtins.str]] = None,
                       ciphertext: Optional[pulumi.Input[builtins.str]] = None,
                       context: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       key: Optional[pulumi.Input[builtins.str]] = None,
                       namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDecryptResult]:
    """
    This is a data source which can be used to decrypt ciphertext using a Vault Transit key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    test = vault.transit.get_decrypt(backend="transit",
        key="test",
        ciphertext="vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==")
    ```
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['ciphertext'] = ciphertext
    __args__['context'] = context
    __args__['key'] = key
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:transit/getDecrypt:getDecrypt', __args__, opts=opts, typ=GetDecryptResult)
    return __ret__.apply(lambda __response__: GetDecryptResult(
        backend=pulumi.get(__response__, 'backend'),
        ciphertext=pulumi.get(__response__, 'ciphertext'),
        context=pulumi.get(__response__, 'context'),
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        namespace=pulumi.get(__response__, 'namespace'),
        plaintext=pulumi.get(__response__, 'plaintext')))
