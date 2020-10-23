# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetEncryptResult',
    'AwaitableGetEncryptResult',
    'get_encrypt',
]

@pulumi.output_type
class GetEncryptResult:
    """
    A collection of values returned by getEncrypt.
    """
    def __init__(__self__, backend=None, ciphertext=None, context=None, id=None, key=None, key_version=None, plaintext=None):
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
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        pulumi.set(__self__, "plaintext", plaintext)

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def ciphertext(self) -> str:
        """
        Encrypted ciphertext returned from Vault
        """
        return pulumi.get(self, "ciphertext")

    @property
    @pulumi.getter
    def context(self) -> Optional[str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyVersion")
    def key_version(self) -> Optional[int]:
        return pulumi.get(self, "key_version")

    @property
    @pulumi.getter
    def plaintext(self) -> str:
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
            plaintext=self.plaintext)


def get_encrypt(backend: Optional[str] = None,
                context: Optional[str] = None,
                key: Optional[str] = None,
                key_version: Optional[int] = None,
                plaintext: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEncryptResult:
    """
    This is a data source which can be used to encrypt plaintext using a Vault Transit key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    test_mount = vault.Mount("testMount",
        description="This is an example mount",
        path="transit",
        type="transit")
    test_secret_backend_key = vault.transit.SecretBackendKey("testSecretBackendKey", backend=test_mount.path)
    test_encrypt = pulumi.Output.all(test_mount.path, test_secret_backend_key.name).apply(lambda path, name: vault.transit.get_encrypt(backend=path,
        key=name,
        plaintext="foobar"))
    ```


    :param str backend: The path the transit secret backend is mounted at, with no leading or trailing `/`.
    :param str context: Context for key derivation. This is required if key derivation is enabled for this key.
    :param str key: Specifies the name of the transit key to encrypt against.
    :param int key_version: The version of the key to use for encryption. If not set, uses the latest version. Must be greater than or equal to the key's `min_encryption_version`, if set.
    :param str plaintext: Plaintext to be encoded.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['context'] = context
    __args__['key'] = key
    __args__['keyVersion'] = key_version
    __args__['plaintext'] = plaintext
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:transit/getEncrypt:getEncrypt', __args__, opts=opts, typ=GetEncryptResult).value

    return AwaitableGetEncryptResult(
        backend=__ret__.backend,
        ciphertext=__ret__.ciphertext,
        context=__ret__.context,
        id=__ret__.id,
        key=__ret__.key,
        key_version=__ret__.key_version,
        plaintext=__ret__.plaintext)
